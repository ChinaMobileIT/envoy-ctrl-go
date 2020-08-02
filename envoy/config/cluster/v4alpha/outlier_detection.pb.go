// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/config/cluster/v4alpha/outlier_detection.proto

package envoy_config_cluster_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OutlierDetection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consecutive_5Xx                        *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=consecutive_5xx,json=consecutive5xx,proto3" json:"consecutive_5xx,omitempty"`
	Interval                               *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	BaseEjectionTime                       *duration.Duration    `protobuf:"bytes,3,opt,name=base_ejection_time,json=baseEjectionTime,proto3" json:"base_ejection_time,omitempty"`
	MaxEjectionPercent                     *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_ejection_percent,json=maxEjectionPercent,proto3" json:"max_ejection_percent,omitempty"`
	EnforcingConsecutive_5Xx               *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=enforcing_consecutive_5xx,json=enforcingConsecutive5xx,proto3" json:"enforcing_consecutive_5xx,omitempty"`
	EnforcingSuccessRate                   *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=enforcing_success_rate,json=enforcingSuccessRate,proto3" json:"enforcing_success_rate,omitempty"`
	SuccessRateMinimumHosts                *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=success_rate_minimum_hosts,json=successRateMinimumHosts,proto3" json:"success_rate_minimum_hosts,omitempty"`
	SuccessRateRequestVolume               *wrappers.UInt32Value `protobuf:"bytes,8,opt,name=success_rate_request_volume,json=successRateRequestVolume,proto3" json:"success_rate_request_volume,omitempty"`
	SuccessRateStdevFactor                 *wrappers.UInt32Value `protobuf:"bytes,9,opt,name=success_rate_stdev_factor,json=successRateStdevFactor,proto3" json:"success_rate_stdev_factor,omitempty"`
	ConsecutiveGatewayFailure              *wrappers.UInt32Value `protobuf:"bytes,10,opt,name=consecutive_gateway_failure,json=consecutiveGatewayFailure,proto3" json:"consecutive_gateway_failure,omitempty"`
	EnforcingConsecutiveGatewayFailure     *wrappers.UInt32Value `protobuf:"bytes,11,opt,name=enforcing_consecutive_gateway_failure,json=enforcingConsecutiveGatewayFailure,proto3" json:"enforcing_consecutive_gateway_failure,omitempty"`
	SplitExternalLocalOriginErrors         bool                  `protobuf:"varint,12,opt,name=split_external_local_origin_errors,json=splitExternalLocalOriginErrors,proto3" json:"split_external_local_origin_errors,omitempty"`
	ConsecutiveLocalOriginFailure          *wrappers.UInt32Value `protobuf:"bytes,13,opt,name=consecutive_local_origin_failure,json=consecutiveLocalOriginFailure,proto3" json:"consecutive_local_origin_failure,omitempty"`
	EnforcingConsecutiveLocalOriginFailure *wrappers.UInt32Value `protobuf:"bytes,14,opt,name=enforcing_consecutive_local_origin_failure,json=enforcingConsecutiveLocalOriginFailure,proto3" json:"enforcing_consecutive_local_origin_failure,omitempty"`
	EnforcingLocalOriginSuccessRate        *wrappers.UInt32Value `protobuf:"bytes,15,opt,name=enforcing_local_origin_success_rate,json=enforcingLocalOriginSuccessRate,proto3" json:"enforcing_local_origin_success_rate,omitempty"`
	FailurePercentageThreshold             *wrappers.UInt32Value `protobuf:"bytes,16,opt,name=failure_percentage_threshold,json=failurePercentageThreshold,proto3" json:"failure_percentage_threshold,omitempty"`
	EnforcingFailurePercentage             *wrappers.UInt32Value `protobuf:"bytes,17,opt,name=enforcing_failure_percentage,json=enforcingFailurePercentage,proto3" json:"enforcing_failure_percentage,omitempty"`
	EnforcingFailurePercentageLocalOrigin  *wrappers.UInt32Value `protobuf:"bytes,18,opt,name=enforcing_failure_percentage_local_origin,json=enforcingFailurePercentageLocalOrigin,proto3" json:"enforcing_failure_percentage_local_origin,omitempty"`
	FailurePercentageMinimumHosts          *wrappers.UInt32Value `protobuf:"bytes,19,opt,name=failure_percentage_minimum_hosts,json=failurePercentageMinimumHosts,proto3" json:"failure_percentage_minimum_hosts,omitempty"`
	FailurePercentageRequestVolume         *wrappers.UInt32Value `protobuf:"bytes,20,opt,name=failure_percentage_request_volume,json=failurePercentageRequestVolume,proto3" json:"failure_percentage_request_volume,omitempty"`
}

func (x *OutlierDetection) Reset() {
	*x = OutlierDetection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_cluster_v4alpha_outlier_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetection) ProtoMessage() {}

func (x *OutlierDetection) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_cluster_v4alpha_outlier_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetection.ProtoReflect.Descriptor instead.
func (*OutlierDetection) Descriptor() ([]byte, []int) {
	return file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescGZIP(), []int{0}
}

func (x *OutlierDetection) GetConsecutive_5Xx() *wrappers.UInt32Value {
	if x != nil {
		return x.Consecutive_5Xx
	}
	return nil
}

func (x *OutlierDetection) GetInterval() *duration.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *OutlierDetection) GetBaseEjectionTime() *duration.Duration {
	if x != nil {
		return x.BaseEjectionTime
	}
	return nil
}

func (x *OutlierDetection) GetMaxEjectionPercent() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxEjectionPercent
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutive_5Xx() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutive_5Xx
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingSuccessRate() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingSuccessRate
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateMinimumHosts() *wrappers.UInt32Value {
	if x != nil {
		return x.SuccessRateMinimumHosts
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateRequestVolume() *wrappers.UInt32Value {
	if x != nil {
		return x.SuccessRateRequestVolume
	}
	return nil
}

func (x *OutlierDetection) GetSuccessRateStdevFactor() *wrappers.UInt32Value {
	if x != nil {
		return x.SuccessRateStdevFactor
	}
	return nil
}

func (x *OutlierDetection) GetConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if x != nil {
		return x.ConsecutiveGatewayFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutiveGatewayFailure() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutiveGatewayFailure
	}
	return nil
}

func (x *OutlierDetection) GetSplitExternalLocalOriginErrors() bool {
	if x != nil {
		return x.SplitExternalLocalOriginErrors
	}
	return false
}

func (x *OutlierDetection) GetConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if x != nil {
		return x.ConsecutiveLocalOriginFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingConsecutiveLocalOriginFailure() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingConsecutiveLocalOriginFailure
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingLocalOriginSuccessRate() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingLocalOriginSuccessRate
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageThreshold() *wrappers.UInt32Value {
	if x != nil {
		return x.FailurePercentageThreshold
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingFailurePercentage() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingFailurePercentage
	}
	return nil
}

func (x *OutlierDetection) GetEnforcingFailurePercentageLocalOrigin() *wrappers.UInt32Value {
	if x != nil {
		return x.EnforcingFailurePercentageLocalOrigin
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageMinimumHosts() *wrappers.UInt32Value {
	if x != nil {
		return x.FailurePercentageMinimumHosts
	}
	return nil
}

func (x *OutlierDetection) GetFailurePercentageRequestVolume() *wrappers.UInt32Value {
	if x != nil {
		return x.FailurePercentageRequestVolume
	}
	return nil
}

var File_envoy_config_cluster_v4alpha_outlier_detection_proto protoreflect.FileDescriptor

var file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDesc = []byte{
	0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6f,
	0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xef, 0x0f, 0x0a, 0x10, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x35, 0x78, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x35, 0x78, 0x78, 0x12, 0x3f, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02,
	0x2a, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x51, 0x0a, 0x12,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x10, 0x62,
	0x61, 0x73, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x57, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x18, 0x64, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x19, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x35, 0x78, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x18, 0x64, 0x52, 0x17, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x35, 0x78, 0x78, 0x12, 0x5b, 0x0a, 0x16, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x18, 0x64, 0x52, 0x14, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x1a, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x17, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x5b, 0x0a, 0x1b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x18, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x57, 0x0a, 0x19, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x74, 0x64, 0x65, 0x76, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x16, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x64, 0x65, 0x76, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5c, 0x0a, 0x1b, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x78, 0x0a, 0x25, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x22, 0x65,
	0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x76, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x12, 0x4a, 0x0a, 0x22, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x73,
	0x70, 0x6c, 0x69, 0x74, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x65, 0x0a,
	0x20, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x2a, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64,
	0x52, 0x26, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x73, 0x0a, 0x23, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x1f, 0x65, 0x6e,
	0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x67, 0x0a,
	0x1c, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x1a, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x67, 0x0a, 0x1c, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x18, 0x64, 0x52, 0x1a, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12,
	0x7f, 0x0a, 0x29, 0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x25, 0x65, 0x6e, 0x66, 0x6f, 0x72,
	0x63, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x65, 0x0a, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72,
	0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x6e, 0x69, 0x6d,
	0x75, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x67, 0x0a, 0x21, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x3a, 0x2f, 0x9a, 0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x4d, 0x0a, 0x2a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x15, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescOnce sync.Once
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescData = file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDesc
)

func file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescGZIP() []byte {
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescOnce.Do(func() {
		file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescData)
	})
	return file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDescData
}

var file_envoy_config_cluster_v4alpha_outlier_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_cluster_v4alpha_outlier_detection_proto_goTypes = []interface{}{
	(*OutlierDetection)(nil),     // 0: envoy.config.cluster.v4alpha.OutlierDetection
	(*wrappers.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*duration.Duration)(nil),    // 2: google.protobuf.Duration
}
var file_envoy_config_cluster_v4alpha_outlier_detection_proto_depIdxs = []int32{
	1,  // 0: envoy.config.cluster.v4alpha.OutlierDetection.consecutive_5xx:type_name -> google.protobuf.UInt32Value
	2,  // 1: envoy.config.cluster.v4alpha.OutlierDetection.interval:type_name -> google.protobuf.Duration
	2,  // 2: envoy.config.cluster.v4alpha.OutlierDetection.base_ejection_time:type_name -> google.protobuf.Duration
	1,  // 3: envoy.config.cluster.v4alpha.OutlierDetection.max_ejection_percent:type_name -> google.protobuf.UInt32Value
	1,  // 4: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_consecutive_5xx:type_name -> google.protobuf.UInt32Value
	1,  // 5: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_success_rate:type_name -> google.protobuf.UInt32Value
	1,  // 6: envoy.config.cluster.v4alpha.OutlierDetection.success_rate_minimum_hosts:type_name -> google.protobuf.UInt32Value
	1,  // 7: envoy.config.cluster.v4alpha.OutlierDetection.success_rate_request_volume:type_name -> google.protobuf.UInt32Value
	1,  // 8: envoy.config.cluster.v4alpha.OutlierDetection.success_rate_stdev_factor:type_name -> google.protobuf.UInt32Value
	1,  // 9: envoy.config.cluster.v4alpha.OutlierDetection.consecutive_gateway_failure:type_name -> google.protobuf.UInt32Value
	1,  // 10: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_consecutive_gateway_failure:type_name -> google.protobuf.UInt32Value
	1,  // 11: envoy.config.cluster.v4alpha.OutlierDetection.consecutive_local_origin_failure:type_name -> google.protobuf.UInt32Value
	1,  // 12: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_consecutive_local_origin_failure:type_name -> google.protobuf.UInt32Value
	1,  // 13: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_local_origin_success_rate:type_name -> google.protobuf.UInt32Value
	1,  // 14: envoy.config.cluster.v4alpha.OutlierDetection.failure_percentage_threshold:type_name -> google.protobuf.UInt32Value
	1,  // 15: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_failure_percentage:type_name -> google.protobuf.UInt32Value
	1,  // 16: envoy.config.cluster.v4alpha.OutlierDetection.enforcing_failure_percentage_local_origin:type_name -> google.protobuf.UInt32Value
	1,  // 17: envoy.config.cluster.v4alpha.OutlierDetection.failure_percentage_minimum_hosts:type_name -> google.protobuf.UInt32Value
	1,  // 18: envoy.config.cluster.v4alpha.OutlierDetection.failure_percentage_request_volume:type_name -> google.protobuf.UInt32Value
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_envoy_config_cluster_v4alpha_outlier_detection_proto_init() }
func file_envoy_config_cluster_v4alpha_outlier_detection_proto_init() {
	if File_envoy_config_cluster_v4alpha_outlier_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_cluster_v4alpha_outlier_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_cluster_v4alpha_outlier_detection_proto_goTypes,
		DependencyIndexes: file_envoy_config_cluster_v4alpha_outlier_detection_proto_depIdxs,
		MessageInfos:      file_envoy_config_cluster_v4alpha_outlier_detection_proto_msgTypes,
	}.Build()
	File_envoy_config_cluster_v4alpha_outlier_detection_proto = out.File
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_rawDesc = nil
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_goTypes = nil
	file_envoy_config_cluster_v4alpha_outlier_detection_proto_depIdxs = nil
}
