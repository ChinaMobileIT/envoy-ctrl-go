package envoy

import (
	core "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/config/core/v3"
	socketsTlsV3 "github.com/ChinaMobileIT/envoy-ctrl-go/envoy/extensions/transport_sockets/tls/v3"
	"github.com/ChinaMobileIT/envoy-ctrl-go/internal/example/protobuf"
)

// UpstreamTLSTransportSocket returns a custom transport socket using the UpstreamTlsContext provided.
func UpstreamTLSTransportSocket(tls *socketsTlsV3.UpstreamTlsContext) *core.TransportSocket {
	return &core.TransportSocket{
		Name: "envoy.transport_sockets.tls",
		ConfigType: &core.TransportSocket_TypedConfig{
			TypedConfig: protobuf.MustMarshalAny(tls),
		},
	}
}

// DownstreamTLSTransportSocket returns a custom transport socket using the DownstreamTlsContext provided.
func DownstreamTLSTransportSocket(tls *socketsTlsV3.DownstreamTlsContext) *core.TransportSocket {
	return &core.TransportSocket{
		Name: "envoy.transport_sockets.tls",
		ConfigType: &core.TransportSocket_TypedConfig{
			TypedConfig: protobuf.MustMarshalAny(tls),
		},
	}
}
