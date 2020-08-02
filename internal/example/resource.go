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
	"time"

	"github.com/golang/protobuf/ptypes"

	cluster "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/cluster/v3"
	core "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/core/v3"
	endpoint "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/endpoint/v3"
	listener "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/listener/v3"
	route "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/route/v3"
	hcm "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/extensions/transport_sockets/tls/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/protobuf"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/cache/types"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/cache/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/resource/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/pkg/wellknown"
)

// UpstreamTLSTransportSocket returns a custom transport socket using the UpstreamTlsContext provided.
func UpstreamTLSTransportSocket(tls *envoy_extensions_transport_sockets_tls_v3.UpstreamTlsContext) *core.TransportSocket {
	config := protobuf.MustMarshalAny(tls)
	return &core.TransportSocket{
		Name: "envoy.transport_sockets.tls",
		ConfigType: &core.TransportSocket_TypedConfig{
			TypedConfig: config,
		},
	}
}

//
// // DownstreamTLSTransportSocket returns a custom transport socket using the DownstreamTlsContext provided.
// func DownstreamTLSTransportSocket(tls *envoy_api_v2_auth.DownstreamTlsContext) *envoy_api_v2_core.TransportSocket {
// 	return &envoy_api_v2_core.TransportSocket{
// 		Name: "envoy.transport_sockets.tls",
// 		ConfigType: &envoy_api_v2_core.TransportSocket_TypedConfig{
// 			TypedConfig: protobuf.MustMarshalAny(tls),
// 		},
// 	}
// }
func makeCluster(clusterName string, upstreamHost string, upstreamPort uint32) *cluster.Cluster {
	context := &envoy_extensions_transport_sockets_tls_v3.UpstreamTlsContext{
		Sni: upstreamHost,
	}
	return &cluster.Cluster{
		Name:                 clusterName,
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		ConnectTimeout:       ptypes.DurationProto(250 * time.Millisecond),
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName, upstreamHost, upstreamPort),
		TransportSocket:      UpstreamTLSTransportSocket(context),
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

func GenerateSnapshot(version, UpstreamHost string, UpstreamPort, ListenerPort uint32) cache.Snapshot {

	var endpoints []types.Resource
	var clusters []types.Resource
	var routes []types.Resource
	var listeners []types.Resource
	var runtimes []types.Resource
	var secrets []types.Resource

	clusters, routes, listeners = addNewGuide(clusters, routes, listeners, UpstreamHost, UpstreamPort, ListenerPort)

	return cache.NewSnapshot(
		version,
		endpoints,
		clusters,
		routes,
		listeners,
		runtimes,
		secrets,
	)
}

func addNewGuide(clusters, routes, listeners []types.Resource, UpstreamHost string, UpstreamPort, ListenerPort uint32) ([]types.Resource, []types.Resource, []types.Resource) {
	ClusterName := "cluseter_" + UpstreamHost
	RouteName := "route_" + UpstreamHost
	ListenerName := "listener_" + UpstreamHost
	clusters = append(clusters, makeCluster(ClusterName, UpstreamHost, UpstreamPort))
	routes = append(routes, makeRoute(RouteName, ClusterName, UpstreamHost))
	listeners = append(listeners, makeHTTPListener(ListenerName, RouteName, ListenerPort))
	fmt.Println(UpstreamHost, UpstreamPort, ListenerPort)
	return clusters, routes, listeners
}
