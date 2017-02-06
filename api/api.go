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

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/abbot/go-http-auth"
	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"

	shttp "github.com/skydive-project/skydive/http"
	"github.com/skydive-project/skydive/logging"
	"github.com/skydive-project/skydive/validator"
	"github.com/skydive-project/skydive/version"
)

type ApiServer struct {
	HTTPServer  *shttp.Server
	EtcdKeyAPI  etcd.KeysAPI
	ServiceName string
	handlers    map[string]ApiHandler
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func writeError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}

func (a *ApiServer) RegisterApiHandler(handler ApiHandler) error {
	name := handler.Name()
	title := strings.Title(name)

	routes := []shttp.Route{
		{
			title + "Index",
			"GET",
			"/api/" + name,
			func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)

				resources := handler.Index()
				for _, resource := range resources {
					handler.Decorate(resource)
				}

				if err := json.NewEncoder(w).Encode(resources); err != nil {
					logging.GetLogger().Criticalf("Failed to display %s: %s", name, err.Error())
				}
			},
		},
		{
			title + "Show",
			"GET",
			shttp.PathPrefix(fmt.Sprintf("/api/%s/", name)),
			func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
				id := r.URL.Path[len(fmt.Sprintf("/api/%s/", name)):]
				if id == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				resource, ok := handler.Get(id)
				if !ok {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				handler.Decorate(resource)
				if err := json.NewEncoder(w).Encode(resource); err != nil {
					logging.GetLogger().Criticalf("Failed to display %s: %s", name, err.Error())
				}
			},
		},
		{
			title + "Insert",
			"POST",
			"/api/" + name,
			func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
				resource := handler.New()

				// keep the original ID
				id := resource.ID()

				if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
					writeError(w, http.StatusBadRequest, err)
					return
				}

				resource.SetID(id)

				if err := validator.Validate(resource); err != nil {
					writeError(w, http.StatusBadRequest, err)
					return
				}

				if err := handler.Create(resource); err != nil {
					writeError(w, http.StatusBadRequest, err)
					return
				}

				data, err := json.Marshal(&resource)
				if err != nil {
					writeError(w, http.StatusBadRequest, err)
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
				if _, err := w.Write(data); err != nil {
					logging.GetLogger().Criticalf("Failed to create %s: %s", name, err.Error())
				}
			},
		},
		{
			title + "Delete",
			"DELETE",
			shttp.PathPrefix(fmt.Sprintf("/api/%s/", name)),
			func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
				id := r.URL.Path[len(fmt.Sprintf("/api/%s/", name)):]
				if id == "" {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				if err := handler.Delete(id); err != nil {
					writeError(w, http.StatusBadRequest, err)
					return
				}

				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)
			},
		},
	}

	a.HTTPServer.RegisterRoutes(routes)

	if _, err := a.EtcdKeyAPI.Set(context.Background(), "/"+name, "", &etcd.SetOptions{Dir: true}); err != nil {
		if _, err = a.EtcdKeyAPI.Get(context.Background(), "/"+name, nil); err != nil {
			return err
		}
	}

	a.handlers[handler.Name()] = handler

	return nil
}

func (a *ApiServer) addAPIRootRoute() {
	s := struct {
		Service string
		Version string
	}{
		Service: a.ServiceName,
		Version: version.Version,
	}

	routes := []shttp.Route{
		{
			"Skydive API",
			"GET",
			"/api",
			func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusOK)

				if err := json.NewEncoder(w).Encode(s); err != nil {
					logging.GetLogger().Criticalf("Failed to display /api: %s", err.Error())
				}
			},
		}}

	a.HTTPServer.RegisterRoutes(routes)
}

func (a *ApiServer) GetHandler(s string) ApiHandler {
	return a.handlers[s]
}

func NewApi(server *shttp.Server, kapi etcd.KeysAPI, serviceName string) (*ApiServer, error) {
	apiServer := &ApiServer{
		HTTPServer:  server,
		EtcdKeyAPI:  kapi,
		ServiceName: serviceName,
		handlers:    make(map[string]ApiHandler),
	}

	apiServer.addAPIRootRoute()

	return apiServer, nil
}

func NewCrudClientFromConfig(authOptions *shttp.AuthenticationOpts) (*shttp.CrudClient, error) {
	return shttp.NewCrudClientFromConfig(authOptions, "api")
}

func NewRestClientFromConfig(authOptions *shttp.AuthenticationOpts) (*shttp.RestClient, error) {
	return shttp.NewRestClientFromConfig(authOptions)
}
