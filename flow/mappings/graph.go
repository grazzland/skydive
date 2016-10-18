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

package mappings

import (
	"github.com/skydive-project/skydive/flow"
	"github.com/skydive-project/skydive/flow/packet"
	"github.com/skydive-project/skydive/logging"
	"github.com/skydive-project/skydive/topology/graph"
)

type GraphFlowEnhancer struct {
	Graph *graph.Graph
}

func (gfe *GraphFlowEnhancer) getNodeUUIDMac(mac string) string {
	if packet.IsBroadcastMac(mac) || packet.IsMulticastMac(mac) {
		return "*"
	}

	gfe.Graph.RLock()
	defer gfe.Graph.RUnlock()

	intfs := gfe.Graph.LookupNodes(graph.Metadata{"MAC": mac})
	if len(intfs) > 1 {
		logging.GetLogger().Infof("GraphFlowEnhancer found more than one interface for the mac: %s", mac)
	} else if len(intfs) == 1 {
		return string(intfs[0].ID)
	}
	return ""
}

func (gfe *GraphFlowEnhancer) getNodeUUIDIP(ip string) string {
	if packet.IsMulticastIP(ip) {
		return "*"
	}

	gfe.Graph.RLock()
	defer gfe.Graph.RUnlock()

	intfs := gfe.Graph.LookupNodes(graph.Metadata{"IP": ip})
	if len(intfs) > 1 {
		logging.GetLogger().Infof("GraphFlowEnhancer found more than one interface for the ip: %s", ip)
	} else if len(intfs) == 1 {
		return string(intfs[0].ID)
	}
	return ""
}

func (gfe *GraphFlowEnhancer) Enhance(f *flow.Flow) {
	if f.ANodeUUID == "" || f.BNodeUUID == "" {
		if f.Link == nil && f.Network == nil {
			return
		}
	}
	if f.ANodeUUID == "" {
		if f.Link != nil {
			f.ANodeUUID = gfe.getNodeUUIDMac(f.Link.A)
		} else {
			f.ANodeUUID = gfe.getNodeUUIDIP(f.Network.A)
		}
	}
	if f.BNodeUUID == "" {
		if f.Link != nil {
			f.BNodeUUID = gfe.getNodeUUIDMac(f.Link.B)
		} else {
			f.BNodeUUID = gfe.getNodeUUIDIP(f.Network.B)
		}
	}
}

func NewGraphFlowEnhancer(g *graph.Graph) *GraphFlowEnhancer {
	return &GraphFlowEnhancer{
		Graph: g,
	}
}
