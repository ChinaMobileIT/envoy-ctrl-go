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
package example

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"

	cluster "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/cluster/v3"
	core "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/core/v3"
	endpoint "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/endpoint/v3"
	listener "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/listener/v3"
	route "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/route/v3"
	hcm "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/extensions/filters/network/http_connection_manager/v3"
	socketsTlsV3 "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/extensions/transport_sockets/tls/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/envoy"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/protobuf"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/cache/types"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/cache/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/resource/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/wellknown"
)

func makeCluster(clusterName string, upstreamHost string, upstreamPort uint32) *cluster.Cluster {
	context := &socketsTlsV3.UpstreamTlsContext{
		Sni: upstreamHost,
	}
	return &cluster.Cluster{
		Name:                 clusterName,
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		ConnectTimeout:       ptypes.DurationProto(250 * time.Millisecond),
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName, upstreamHost, upstreamPort),
		TransportSocket:      envoy.UpstreamTLSTransportSocket(context),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
	}
}

func makeEndpoint(clusterName string, upstreamHost string, upstreamPort uint32) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  upstreamHost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: upstreamPort,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func makeRoute(routeName string, clusterName, upstreamHost string) *route.RouteConfiguration {
	return &route.RouteConfiguration{
		Name: routeName,
		VirtualHosts: []*route.VirtualHost{{
			Name:    "local_service_" + upstreamHost,
			Domains: []string{"*"},
			Routes: []*route.Route{{
				Match: &route.RouteMatch{
					PathSpecifier: &route.RouteMatch_Prefix{
						Prefix: "/",
					},
				},
				Action: &route.Route_Route{
					Route: &route.RouteAction{
						ClusterSpecifier: &route.RouteAction_Cluster{
							Cluster: clusterName,
						},
						HostRewriteSpecifier: &route.RouteAction_HostRewriteLiteral{
							HostRewriteLiteral: upstreamHost,
						},
					},
				},
			}},
		}},
	}
}

func makeHTTPListener(listenerName string, route string, port uint32) *listener.Listener {
	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: "ingress_http",
		RouteSpecifier: &hcm.HttpConnectionManager_Rds{
			Rds: &hcm.Rds{
				ConfigSource:    makeConfigSource(),
				RouteConfigName: route,
			},
		},
		HttpFilters: []*hcm.HttpFilter{{
			Name: wellknown.Router,
		}},
	}

	pbst := protobuf.MustMarshalAny(manager)

	return &listener.Listener{
		Name: listenerName,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}
}

func makeConfigSource() *core.ConfigSource {
	source := &core.ConfigSource{}
	source.ResourceApiVersion = resource.DefaultAPIVersion
	source.ConfigSourceSpecifier = &core.ConfigSource_ApiConfigSource{
		ApiConfigSource: &core.ApiConfigSource{
			TransportApiVersion:       resource.DefaultAPIVersion,
			ApiType:                   core.ApiConfigSource_GRPC,
			SetNodeOnFirstMessageOnly: true,
			GrpcServices: []*core.GrpcService{{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{ClusterName: "xds_cluster"},
				},
			}},
		},
	}
	return source
}

var endpoints []types.Resource
var clusters []types.Resource
var routes []types.Resource
var listeners []types.Resource
var runtimes []types.Resource
var secrets []types.Resource

func GenerateSnapshot(upstreamHost string, upstreamPort, listenerPort uint32) cache.Snapshot {
	version := addNewGuide(upstreamHost, upstreamPort, listenerPort)
	snapshot := cache.NewSnapshot(version, endpoints, clusters, routes, listeners, runtimes, secrets)
	if err := snapshot.Consistent(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	return snapshot
}

func addNewGuide(upstreamHost string, upstreamPort, listenerPort uint32) (version string) {
	clusterName := "cluster_" + upstreamHost
	routeName := "route_" + upstreamHost
	listenerName := "listener_" + upstreamHost
	clusters = append(clusters, makeCluster(clusterName, upstreamHost, upstreamPort))
	routes = append(routes, makeRoute(routeName, clusterName, upstreamHost))
	listeners = append(listeners, makeHTTPListener(listenerName, routeName, listenerPort))
	fmt.Println(upstreamHost, upstreamPort, listenerPort)
	return fmt.Sprintf("%d", time.Now().Nanosecond())
}

func Reset() cache.Snapshot {
	endpoints = []types.Resource{}
	clusters = []types.Resource{}
	routes = []types.Resource{}
	listeners = []types.Resource{}
	runtimes = []types.Resource{}
	secrets = []types.Resource{}
	snapshot := cache.NewSnapshot("0", endpoints, clusters, routes, listeners, runtimes, secrets)
	if err := snapshot.Consistent(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	return snapshot
}
