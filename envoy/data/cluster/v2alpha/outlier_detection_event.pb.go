// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/data/cluster/v2alpha/outlier_detection_event.proto

package envoy_data_cluster_v2alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type OutlierEjectionType int32

const (
	OutlierEjectionType_CONSECUTIVE_5XX                  OutlierEjectionType = 0
	OutlierEjectionType_CONSECUTIVE_GATEWAY_FAILURE      OutlierEjectionType = 1
	OutlierEjectionType_SUCCESS_RATE                     OutlierEjectionType = 2
	OutlierEjectionType_CONSECUTIVE_LOCAL_ORIGIN_FAILURE OutlierEjectionType = 3
	OutlierEjectionType_SUCCESS_RATE_LOCAL_ORIGIN        OutlierEjectionType = 4
	OutlierEjectionType_FAILURE_PERCENTAGE               OutlierEjectionType = 5
	OutlierEjectionType_FAILURE_PERCENTAGE_LOCAL_ORIGIN  OutlierEjectionType = 6
)

// Enum value maps for OutlierEjectionType.
var (
	OutlierEjectionType_name = map[int32]string{
		0: "CONSECUTIVE_5XX",
		1: "CONSECUTIVE_GATEWAY_FAILURE",
		2: "SUCCESS_RATE",
		3: "CONSECUTIVE_LOCAL_ORIGIN_FAILURE",
		4: "SUCCESS_RATE_LOCAL_ORIGIN",
		5: "FAILURE_PERCENTAGE",
		6: "FAILURE_PERCENTAGE_LOCAL_ORIGIN",
	}
	OutlierEjectionType_value = map[string]int32{
		"CONSECUTIVE_5XX":                  0,
		"CONSECUTIVE_GATEWAY_FAILURE":      1,
		"SUCCESS_RATE":                     2,
		"CONSECUTIVE_LOCAL_ORIGIN_FAILURE": 3,
		"SUCCESS_RATE_LOCAL_ORIGIN":        4,
		"FAILURE_PERCENTAGE":               5,
		"FAILURE_PERCENTAGE_LOCAL_ORIGIN":  6,
	}
)

func (x OutlierEjectionType) Enum() *OutlierEjectionType {
	p := new(OutlierEjectionType)
	*p = x
	return p
}

func (x OutlierEjectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OutlierEjectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes[0].Descriptor()
}

func (OutlierEjectionType) Type() protoreflect.EnumType {
	return &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes[0]
}

func (x OutlierEjectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OutlierEjectionType.Descriptor instead.
func (OutlierEjectionType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{0}
}

type Action int32

const (
	Action_EJECT   Action = 0
	Action_UNEJECT Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "EJECT",
		1: "UNEJECT",
	}
	Action_value = map[string]int32{
		"EJECT":   0,
		"UNEJECT": 1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes[1].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes[1]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{1}
}

type OutlierDetectionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                OutlierEjectionType   `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.data.cluster.v2alpha.OutlierEjectionType" json:"type,omitempty"`
	Timestamp           *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SecsSinceLastAction *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=secs_since_last_action,json=secsSinceLastAction,proto3" json:"secs_since_last_action,omitempty"`
	ClusterName         string                `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	UpstreamUrl         string                `protobuf:"bytes,5,opt,name=upstream_url,json=upstreamUrl,proto3" json:"upstream_url,omitempty"`
	Action              Action                `protobuf:"varint,6,opt,name=action,proto3,enum=envoy.data.cluster.v2alpha.Action" json:"action,omitempty"`
	NumEjections        uint32                `protobuf:"varint,7,opt,name=num_ejections,json=numEjections,proto3" json:"num_ejections,omitempty"`
	Enforced            bool                  `protobuf:"varint,8,opt,name=enforced,proto3" json:"enforced,omitempty"`
	// Types that are assignable to Event:
	//	*OutlierDetectionEvent_EjectSuccessRateEvent
	//	*OutlierDetectionEvent_EjectConsecutiveEvent
	//	*OutlierDetectionEvent_EjectFailurePercentageEvent
	Event isOutlierDetectionEvent_Event `protobuf_oneof:"event"`
}

func (x *OutlierDetectionEvent) Reset() {
	*x = OutlierDetectionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierDetectionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierDetectionEvent) ProtoMessage() {}

func (x *OutlierDetectionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierDetectionEvent.ProtoReflect.Descriptor instead.
func (*OutlierDetectionEvent) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{0}
}

func (x *OutlierDetectionEvent) GetType() OutlierEjectionType {
	if x != nil {
		return x.Type
	}
	return OutlierEjectionType_CONSECUTIVE_5XX
}

func (x *OutlierDetectionEvent) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *OutlierDetectionEvent) GetSecsSinceLastAction() *wrappers.UInt64Value {
	if x != nil {
		return x.SecsSinceLastAction
	}
	return nil
}

func (x *OutlierDetectionEvent) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *OutlierDetectionEvent) GetUpstreamUrl() string {
	if x != nil {
		return x.UpstreamUrl
	}
	return ""
}

func (x *OutlierDetectionEvent) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_EJECT
}

func (x *OutlierDetectionEvent) GetNumEjections() uint32 {
	if x != nil {
		return x.NumEjections
	}
	return 0
}

func (x *OutlierDetectionEvent) GetEnforced() bool {
	if x != nil {
		return x.Enforced
	}
	return false
}

func (m *OutlierDetectionEvent) GetEvent() isOutlierDetectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectSuccessRateEvent() *OutlierEjectSuccessRate {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		return x.EjectSuccessRateEvent
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectConsecutiveEvent() *OutlierEjectConsecutive {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		return x.EjectConsecutiveEvent
	}
	return nil
}

func (x *OutlierDetectionEvent) GetEjectFailurePercentageEvent() *OutlierEjectFailurePercentage {
	if x, ok := x.GetEvent().(*OutlierDetectionEvent_EjectFailurePercentageEvent); ok {
		return x.EjectFailurePercentageEvent
	}
	return nil
}

type isOutlierDetectionEvent_Event interface {
	isOutlierDetectionEvent_Event()
}

type OutlierDetectionEvent_EjectSuccessRateEvent struct {
	EjectSuccessRateEvent *OutlierEjectSuccessRate `protobuf:"bytes,9,opt,name=eject_success_rate_event,json=ejectSuccessRateEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectConsecutiveEvent struct {
	EjectConsecutiveEvent *OutlierEjectConsecutive `protobuf:"bytes,10,opt,name=eject_consecutive_event,json=ejectConsecutiveEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectFailurePercentageEvent struct {
	EjectFailurePercentageEvent *OutlierEjectFailurePercentage `protobuf:"bytes,11,opt,name=eject_failure_percentage_event,json=ejectFailurePercentageEvent,proto3,oneof"`
}

func (*OutlierDetectionEvent_EjectSuccessRateEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectConsecutiveEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectFailurePercentageEvent) isOutlierDetectionEvent_Event() {}

type OutlierEjectSuccessRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostSuccessRate                     uint32 `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	ClusterAverageSuccessRate           uint32 `protobuf:"varint,2,opt,name=cluster_average_success_rate,json=clusterAverageSuccessRate,proto3" json:"cluster_average_success_rate,omitempty"`
	ClusterSuccessRateEjectionThreshold uint32 `protobuf:"varint,3,opt,name=cluster_success_rate_ejection_threshold,json=clusterSuccessRateEjectionThreshold,proto3" json:"cluster_success_rate_ejection_threshold,omitempty"`
}

func (x *OutlierEjectSuccessRate) Reset() {
	*x = OutlierEjectSuccessRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectSuccessRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectSuccessRate) ProtoMessage() {}

func (x *OutlierEjectSuccessRate) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectSuccessRate.ProtoReflect.Descriptor instead.
func (*OutlierEjectSuccessRate) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{1}
}

func (x *OutlierEjectSuccessRate) GetHostSuccessRate() uint32 {
	if x != nil {
		return x.HostSuccessRate
	}
	return 0
}

func (x *OutlierEjectSuccessRate) GetClusterAverageSuccessRate() uint32 {
	if x != nil {
		return x.ClusterAverageSuccessRate
	}
	return 0
}

func (x *OutlierEjectSuccessRate) GetClusterSuccessRateEjectionThreshold() uint32 {
	if x != nil {
		return x.ClusterSuccessRateEjectionThreshold
	}
	return 0
}

type OutlierEjectConsecutive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OutlierEjectConsecutive) Reset() {
	*x = OutlierEjectConsecutive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectConsecutive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectConsecutive) ProtoMessage() {}

func (x *OutlierEjectConsecutive) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectConsecutive.ProtoReflect.Descriptor instead.
func (*OutlierEjectConsecutive) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{2}
}

type OutlierEjectFailurePercentage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostSuccessRate uint32 `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
}

func (x *OutlierEjectFailurePercentage) Reset() {
	*x = OutlierEjectFailurePercentage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutlierEjectFailurePercentage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutlierEjectFailurePercentage) ProtoMessage() {}

func (x *OutlierEjectFailurePercentage) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutlierEjectFailurePercentage.ProtoReflect.Descriptor instead.
func (*OutlierEjectFailurePercentage) Descriptor() ([]byte, []int) {
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP(), []int{3}
}

func (x *OutlierEjectFailurePercentage) GetHostSuccessRate() uint32 {
	if x != nil {
		return x.HostSuccessRate
	}
	return 0
}

var File_envoy_data_cluster_v2alpha_outlier_detection_event_proto protoreflect.FileDescriptor

var file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x06, 0x0a, 0x15, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x51, 0x0a, 0x16, 0x73, 0x65, 0x63, 0x73, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x13, 0x73, 0x65, 0x63, 0x73, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x0b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x12, 0x44, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x45,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x64, 0x12, 0x6e, 0x0a, 0x18, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x15, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x17, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x48, 0x00, 0x52, 0x15, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x1e, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65,
	0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x1b, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x22, 0xf7, 0x01, 0x0a, 0x17, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72,
	0x45, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x33, 0x0a, 0x11, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x18, 0x64, 0x52, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x1c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x18, 0x64, 0x52, 0x19, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x5d, 0x0a, 0x27, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x23, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x22, 0x19,
	0x0a, 0x17, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x22, 0x54, 0x0a, 0x1d, 0x4f, 0x75, 0x74,
	0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x18, 0x64, 0x52, 0x0f,
	0x68, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x2a,
	0xdf, 0x01, 0x0a, 0x13, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x45, 0x6a, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x53, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x35, 0x58, 0x58, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x4f, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45,
	0x57, 0x41, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x24, 0x0a, 0x20, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x43, 0x55, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x4c,
	0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x5f, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47,
	0x49, 0x4e, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f,
	0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41, 0x47, 0x45, 0x10, 0x05, 0x12, 0x23, 0x0a, 0x1f,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41,
	0x47, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x10,
	0x06, 0x2a, 0x20, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x45, 0x4a, 0x45, 0x43,
	0x54, 0x10, 0x01, 0x42, 0x6d, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42,
	0x1a, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe,
	0x8f, 0x05, 0x17, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescOnce sync.Once
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescData = file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDesc
)

func file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescGZIP() []byte {
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescOnce.Do(func() {
		file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescData)
	})
	return file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDescData
}

var file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_goTypes = []interface{}{
	(OutlierEjectionType)(0),              // 0: envoy.data.cluster.v2alpha.OutlierEjectionType
	(Action)(0),                           // 1: envoy.data.cluster.v2alpha.Action
	(*OutlierDetectionEvent)(nil),         // 2: envoy.data.cluster.v2alpha.OutlierDetectionEvent
	(*OutlierEjectSuccessRate)(nil),       // 3: envoy.data.cluster.v2alpha.OutlierEjectSuccessRate
	(*OutlierEjectConsecutive)(nil),       // 4: envoy.data.cluster.v2alpha.OutlierEjectConsecutive
	(*OutlierEjectFailurePercentage)(nil), // 5: envoy.data.cluster.v2alpha.OutlierEjectFailurePercentage
	(*timestamp.Timestamp)(nil),           // 6: google.protobuf.Timestamp
	(*wrappers.UInt64Value)(nil),          // 7: google.protobuf.UInt64Value
}
var file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_depIdxs = []int32{
	0, // 0: envoy.data.cluster.v2alpha.OutlierDetectionEvent.type:type_name -> envoy.data.cluster.v2alpha.OutlierEjectionType
	6, // 1: envoy.data.cluster.v2alpha.OutlierDetectionEvent.timestamp:type_name -> google.protobuf.Timestamp
	7, // 2: envoy.data.cluster.v2alpha.OutlierDetectionEvent.secs_since_last_action:type_name -> google.protobuf.UInt64Value
	1, // 3: envoy.data.cluster.v2alpha.OutlierDetectionEvent.action:type_name -> envoy.data.cluster.v2alpha.Action
	3, // 4: envoy.data.cluster.v2alpha.OutlierDetectionEvent.eject_success_rate_event:type_name -> envoy.data.cluster.v2alpha.OutlierEjectSuccessRate
	4, // 5: envoy.data.cluster.v2alpha.OutlierDetectionEvent.eject_consecutive_event:type_name -> envoy.data.cluster.v2alpha.OutlierEjectConsecutive
	5, // 6: envoy.data.cluster.v2alpha.OutlierDetectionEvent.eject_failure_percentage_event:type_name -> envoy.data.cluster.v2alpha.OutlierEjectFailurePercentage
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_init() }
func file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_init() {
	if File_envoy_data_cluster_v2alpha_outlier_detection_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierDetectionEvent); i {
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
		file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectSuccessRate); i {
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
		file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectConsecutive); i {
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
		file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutlierEjectFailurePercentage); i {
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
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*OutlierDetectionEvent_EjectSuccessRateEvent)(nil),
		(*OutlierDetectionEvent_EjectConsecutiveEvent)(nil),
		(*OutlierDetectionEvent_EjectFailurePercentageEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_goTypes,
		DependencyIndexes: file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_depIdxs,
		EnumInfos:         file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_enumTypes,
		MessageInfos:      file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_msgTypes,
	}.Build()
	File_envoy_data_cluster_v2alpha_outlier_detection_event_proto = out.File
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_rawDesc = nil
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_goTypes = nil
	file_envoy_data_cluster_v2alpha_outlier_detection_event_proto_depIdxs = nil
}
