/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package etcd

import (
	"sync"
	"sync/atomic"
	"time"

	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"

	"github.com/skydive-project/skydive/common"
	"github.com/skydive-project/skydive/config"
	"github.com/skydive-project/skydive/logging"
)

const (
	timeout = time.Second * 30
)

type EtcdLeaderElectionListener interface {
	OnLeader()
	OnFollower()
}

type EtcdLeaderElector struct {
	sync.RWMutex
	EtcdKeyAPI etcd.KeysAPI
	Host       string
	path       string
	listeners  []EtcdLeaderElectionListener
	cancel     context.CancelFunc
	leader     bool
	state      int64
	wg         sync.WaitGroup
}

func (le *EtcdLeaderElector) holdLock(quit chan bool) {
	defer close(quit)

	tick := time.NewTicker(timeout / 2)
	defer tick.Stop()

	setOptions := &etcd.SetOptions{
		TTL:       timeout,
		PrevExist: etcd.PrevExist,
		PrevValue: le.Host,
	}

	ch := tick.C

	for {
		select {
		case <-ch:
			if _, err := le.EtcdKeyAPI.Set(context.Background(), le.path, le.Host, setOptions); err != nil {
				return
			}
		case <-quit:
			return
		}
	}
}

func (le *EtcdLeaderElector) IsLeader() bool {
	le.RLock()
	defer le.RUnlock()

	return le.leader
}

func (le *EtcdLeaderElector) start() {
	// delete previous Lock
	le.EtcdKeyAPI.Delete(context.Background(), le.path, &etcd.DeleteOptions{PrevValue: le.Host})

	quit := make(chan bool)

	// try to get the lock
	setOptions := &etcd.SetOptions{
		TTL:       timeout,
		PrevExist: etcd.PrevNoExist,
	}

	if _, err := le.EtcdKeyAPI.Set(context.Background(), le.path, le.Host, setOptions); err == nil {
		logging.GetLogger().Infof("starting as the leader: %s", le.Host)

		le.Lock()
		le.leader = true
		le.Unlock()

		go le.holdLock(quit)

		for _, listener := range le.listeners {
			listener.OnLeader()
		}
	} else {
		logging.GetLogger().Infof("starting as a follower: %s", le.Host)
		for _, listener := range le.listeners {
			listener.OnFollower()
		}
	}

	// now watch for changes
	watcher := le.EtcdKeyAPI.Watcher(le.path, &etcd.WatcherOptions{})

	ctx, cancel := context.WithCancel(context.Background())
	le.cancel = cancel

	le.wg.Add(1)
	defer le.wg.Done()

	atomic.StoreInt64(&le.state, common.RunningState)
	for atomic.LoadInt64(&le.state) == common.RunningState {

		resp, err := watcher.Next(ctx)
		if err != nil {
			logging.GetLogger().Errorf("Error while watching etcd: %s", err.Error())

			time.Sleep(1 * time.Second)
			continue
		}

		switch resp.Action {
		case "expire", "delete", "compareAndDelete":
			_, err = le.EtcdKeyAPI.Set(context.Background(), le.path, le.Host, setOptions)
			if err == nil && !le.leader {
				le.Lock()
				le.leader = true
				le.Unlock()

				go le.holdLock(quit)

				logging.GetLogger().Infof("I'm now the leader: %s", le.Host)
				for _, listener := range le.listeners {
					listener.OnLeader()
				}
			}
		case "create", "update":
			le.RLock()
			leader := le.leader
			le.RUnlock()

			if !leader {
				logging.GetLogger().Infof("The leader is now: %s", resp.Node.Value)
				for _, listener := range le.listeners {
					listener.OnFollower()
				}
			}
		}
	}

	// unlock before leaving so that another can take the lead
	le.EtcdKeyAPI.Delete(context.Background(), le.path, &etcd.DeleteOptions{PrevValue: le.Host})
}

func (le *EtcdLeaderElector) Start() {
	go le.start()
}

func (le *EtcdLeaderElector) Stop() {
	if atomic.CompareAndSwapInt64(&le.state, common.RunningState, common.StoppingState) {
		le.cancel()
		le.wg.Wait()
	}
}

func (le *EtcdLeaderElector) AddEventListener(listener EtcdLeaderElectionListener) {
	le.listeners = append(le.listeners, listener)
}

func NewEtcdLeaderElector(host string, serviceType common.ServiceType, key string, etcdClient *EtcdClient) *EtcdLeaderElector {
	return &EtcdLeaderElector{
		EtcdKeyAPI: etcdClient.KeysApi,
		Host:       host,
		path:       "/leader-" + serviceType.String() + "-" + key,
		leader:     false,
	}
}

func NewEtcdLeaderElectorFromConfig(serviceType common.ServiceType, key string, etcdClient *EtcdClient) *EtcdLeaderElector {
	host := config.GetConfig().GetString("host_id")
	return NewEtcdLeaderElector(host, serviceType, key, etcdClient)
}
