/*
 * Copyright (C) 2015 Red Hat, Inc.
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

package config

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"

	"github.com/skydive-project/skydive/common"
)

var cfg *viper.Viper

func init() {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	cfg = viper.New()
	cfg.SetDefault("host_id", host)
	cfg.SetDefault("agent.listen", "127.0.0.1:8081")
	cfg.SetDefault("ovs.ovsdb", "unix:///var/run/openvswitch/db.sock")
	cfg.SetDefault("graph.backend", "memory")
	cfg.SetDefault("graph.gremlin", "ws://127.0.0.1:8182")
	cfg.SetDefault("sflow.port_min", 6345)
	cfg.SetDefault("sflow.port_max", 6355)
	cfg.SetDefault("analyzer.listen", "127.0.0.1:8082")
	cfg.SetDefault("analyzer.flowtable_expire", 600)
	cfg.SetDefault("analyzer.flowtable_update", 60)
	cfg.SetDefault("analyzer.flowtable_agent_ratio", 0.5)
	cfg.SetDefault("storage.elasticsearch.host", "127.0.0.1:9200")
	cfg.SetDefault("storage.elasticsearch.maxconns", 10)
	cfg.SetDefault("storage.elasticsearch.retry", 60)
	cfg.SetDefault("storage.elasticsearch.bulk_maxdocs", 0)
	cfg.SetDefault("ws_pong_timeout", 5)
	cfg.SetDefault("docker.url", "unix:///var/run/docker.sock")
	cfg.SetDefault("netns.run_path", "/var/run/netns")
	cfg.SetDefault("etcd.data_dir", "/tmp/skydive-etcd")
	cfg.SetDefault("etcd.embedded", true)
	cfg.SetDefault("etcd.port", 2379)
	cfg.SetDefault("auth.type", "noauth")
	cfg.SetDefault("auth.keystone.tenant", "admin")
	cfg.SetDefault("storage.orientdb.addr", "http://localhost:2480")
	cfg.SetDefault("storage.orientdb.database", "Skydive")
	cfg.SetDefault("storage.orientdb.username", "root")
	cfg.SetDefault("storage.orientdb.password", "root")
	cfg.SetDefault("openstack.endpoint_type", "public")
	cfg.SetDefault("agent.topology.probes", []string{"netlink", "netns"})
	cfg.SetDefault("analyzer.topology.probes", []string{})
	cfg.SetDefault("opencontrail.mpls_udp_port", 51234)

	replacer := strings.NewReplacer(".", "_", "-", "_")
	cfg.SetEnvPrefix("SKYDIVE")
	cfg.SetEnvKeyReplacer(replacer)
	cfg.AutomaticEnv()
}

func checkStrictPositiveInt(key string) error {
	if value := cfg.GetInt(key); value <= 0 {
		return fmt.Errorf("invalid value for %s (%d)", key, value)
	}

	return nil
}

func checkStrictRangeFloat(key string, min, max float64) error {
	if value := cfg.GetFloat64(key); value <= min || value > max {
		return fmt.Errorf("invalid value for %s (%f)", key, value)
	}

	return nil
}

func checkConfig() error {
	if err := checkStrictRangeFloat("analyzer.flowtable_agent_ratio", 0.0, 1.0); err != nil {
		if cfg.GetFloat64("analyzer.flowtable_agent_ratio") != 0.0 {
			return err
		}
	}

	if err := checkStrictPositiveInt("analyzer.flowtable_expire"); err != nil {
		return err
	}

	if err := checkStrictPositiveInt("analyzer.flowtable_update"); err != nil {
		return err
	}

	return nil
}

func checkViperSupportedExts(ext string) bool {
	for _, e := range viper.SupportedExts {
		if e == ext {
			return true
		}
	}
	return false
}

func InitConfig(backend string, path string) error {
	if path == "" {
		return fmt.Errorf("Empty configuration path")
	}

	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	if ext == "" || !checkViperSupportedExts(ext) {
		ext = "yaml"
	}
	cfg.SetConfigType(ext)

	switch backend {
	case "file":
		configFile, err := os.Open(path)
		if err != nil {
			return err
		}
		if err := cfg.ReadConfig(configFile); err != nil {
			return err
		}
	case "etcd":
		u, err := url.Parse(path)
		if err != nil {
			return err
		}
		if err := cfg.AddRemoteProvider("etcd", fmt.Sprintf("%s://%s", u.Scheme, u.Host), u.Path); err != nil {
			return err
		}
		if err := cfg.ReadRemoteConfig(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid backend: %s", backend)
	}

	return checkConfig()
}

func GetConfig() *viper.Viper {
	return cfg
}

func SetDefault(key string, value interface{}) {
	cfg.SetDefault(key, value)
}

func GetAnalyzerServiceAddresses() ([]common.ServiceAddress, error) {
	var addresses []common.ServiceAddress
	for _, a := range GetConfig().GetStringSlice("analyzers") {
		sa, err := common.ServiceAddressFromString(a)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, sa)
	}

	// shuffle in order to load balance connections
	for i := range addresses {
		j := rand.Intn(i + 1)
		addresses[i], addresses[j] = addresses[j], addresses[i]
	}

	return addresses, nil
}

func GetOneAnalyzerServiceAddress() (common.ServiceAddress, error) {
	addresses, err := GetAnalyzerServiceAddresses()
	if err != nil || len(addresses) == 0 {
		return common.ServiceAddress{}, err
	}
	return addresses[rand.Intn(len(addresses))], nil
}

func GetEtcdServerAddrs() []string {
	etcdServers := GetConfig().GetStringSlice("etcd.servers")
	if len(etcdServers) > 0 {
		return etcdServers
	}
	if addresse, err := GetOneAnalyzerServiceAddress(); err == nil {
		return []string{"http://" + addresse.Addr + ":2379"}
	}
	return []string{}
}
