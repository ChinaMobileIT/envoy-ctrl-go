// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package main

import (
	"context"
	"flag"

	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/cache"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/web"
	serverv3 "github.com/ChinaMobileIT/envoy-ctrl-go/pkg/server/v3"
	testv3 "github.com/ChinaMobileIT/envoy-ctrl-go/pkg/test/v3"
)

var (
	l      example.Logger
	port   uint
	nodeID string
)

func init() {
	l = example.Logger{}

	flag.BoolVar(&l.Debug, "debug", true, "Enable xDS server debug logging")

	// The port that this xDS server listens on
	flag.UintVar(&port, "port", 18000, "xDS management server port")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")
}

func main() {
	flag.Parse()

	cache := cache.NewCache(l)

	// 启动web服务器，提供简化的操作接口
	go web.Start()

	// Run the xDS server
	ctx := context.Background()
	cb := &testv3.Callbacks{Debug: l.Debug}
	srv := serverv3.NewServer(ctx, cache, cb)
	example.RunServer(ctx, srv, port)
}
