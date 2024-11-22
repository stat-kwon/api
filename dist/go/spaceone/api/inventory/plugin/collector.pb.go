// A Collector is a plugin collecting data of external infrastructure resources.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: spaceone/api/inventory/plugin/collector.proto

package plugin

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type ResourceInfo_State int32

const (
	ResourceInfo_NONE       ResourceInfo_State = 0
	ResourceInfo_CREATED    ResourceInfo_State = 1
	ResourceInfo_PENDING    ResourceInfo_State = 2
	ResourceInfo_INPROGRESS ResourceInfo_State = 3
	ResourceInfo_SUCCESS    ResourceInfo_State = 4
	ResourceInfo_FAILURE    ResourceInfo_State = 5
	ResourceInfo_TIMEOUT    ResourceInfo_State = 6
	ResourceInfo_IDLE       ResourceInfo_State = 7
	ResourceInfo_SKIP       ResourceInfo_State = 8
)

// Enum value maps for ResourceInfo_State.
var (
	ResourceInfo_State_name = map[int32]string{
		0: "NONE",
		1: "CREATED",
		2: "PENDING",
		3: "INPROGRESS",
		4: "SUCCESS",
		5: "FAILURE",
		6: "TIMEOUT",
		7: "IDLE",
		8: "SKIP",
	}
	ResourceInfo_State_value = map[string]int32{
		"NONE":       0,
		"CREATED":    1,
		"PENDING":    2,
		"INPROGRESS": 3,
		"SUCCESS":    4,
		"FAILURE":    5,
		"TIMEOUT":    6,
		"IDLE":       7,
		"SKIP":       8,
	}
)

func (x ResourceInfo_State) Enum() *ResourceInfo_State {
	p := new(ResourceInfo_State)
	*p = x
	return p
}

func (x ResourceInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_inventory_plugin_collector_proto_enumTypes[0].Descriptor()
}

func (ResourceInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_inventory_plugin_collector_proto_enumTypes[0]
}

func (x ResourceInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceInfo_State.Descriptor instead.
func (ResourceInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{3, 0}
}

// {
//
// }
type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{0}
}

func (x *InitRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

// {
//
// }
type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options    *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *VerifyRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

// {
//
// }
type CollectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options     *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData  *_struct.Struct `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	Filter      *_struct.Struct `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	TaskOptions *_struct.Struct `protobuf:"bytes,4,opt,name=task_options,json=taskOptions,proto3" json:"task_options,omitempty"`
}

func (x *CollectRequest) Reset() {
	*x = CollectRequest{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectRequest) ProtoMessage() {}

func (x *CollectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectRequest.ProtoReflect.Descriptor instead.
func (*CollectRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{2}
}

func (x *CollectRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *CollectRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *CollectRequest) GetFilter() *_struct.Struct {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *CollectRequest) GetTaskOptions() *_struct.Struct {
	if x != nil {
		return x.TaskOptions
	}
	return nil
}

type ResourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State        ResourceInfo_State `protobuf:"varint,1,opt,name=state,proto3,enum=spaceone.api.inventory.plugin.ResourceInfo_State" json:"state,omitempty"`
	Message      string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ResourceType string             `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	MatchRules   *_struct.Struct    `protobuf:"bytes,4,opt,name=match_rules,json=matchRules,proto3" json:"match_rules,omitempty"`
	Resource     *_struct.Struct    `protobuf:"bytes,6,opt,name=resource,proto3" json:"resource,omitempty"`
	Options      *_struct.Struct    `protobuf:"bytes,7,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ResourceInfo) Reset() {
	*x = ResourceInfo{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfo) ProtoMessage() {}

func (x *ResourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfo.ProtoReflect.Descriptor instead.
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{3}
}

func (x *ResourceInfo) GetState() ResourceInfo_State {
	if x != nil {
		return x.State
	}
	return ResourceInfo_NONE
}

func (x *ResourceInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResourceInfo) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceInfo) GetMatchRules() *_struct.Struct {
	if x != nil {
		return x.MatchRules
	}
	return nil
}

func (x *ResourceInfo) GetResource() *_struct.Struct {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

type CollectorVerifyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *_struct.Struct `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CollectorVerifyInfo) Reset() {
	*x = CollectorVerifyInfo{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectorVerifyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectorVerifyInfo) ProtoMessage() {}

func (x *CollectorVerifyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectorVerifyInfo.ProtoReflect.Descriptor instead.
func (*CollectorVerifyInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{4}
}

func (x *CollectorVerifyInfo) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

// {
//
// }
type PluginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *_struct.Struct `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_plugin_collector_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP(), []int{5}
}

func (x *PluginInfo) GetMetadata() *_struct.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_spaceone_api_inventory_plugin_collector_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_plugin_collector_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0b, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7c, 0x0a, 0x0d, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0xea, 0x01, 0x0a, 0x0e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x38, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb0, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x76, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x49, 0x4d, 0x45,
	0x4f, 0x55, 0x54, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x07, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x08, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x41, 0x0a, 0x0a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa9, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x5f, 0x0a, 0x04, 0x69, 0x6e, 0x69, 0x74, 0x12, 0x2a, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12,
	0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_plugin_collector_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_plugin_collector_proto_rawDescData = file_spaceone_api_inventory_plugin_collector_proto_rawDesc
)

func file_spaceone_api_inventory_plugin_collector_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_plugin_collector_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_plugin_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_plugin_collector_proto_rawDescData)
	})
	return file_spaceone_api_inventory_plugin_collector_proto_rawDescData
}

var file_spaceone_api_inventory_plugin_collector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_inventory_plugin_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spaceone_api_inventory_plugin_collector_proto_goTypes = []any{
	(ResourceInfo_State)(0),     // 0: spaceone.api.inventory.plugin.ResourceInfo.State
	(*InitRequest)(nil),         // 1: spaceone.api.inventory.plugin.InitRequest
	(*VerifyRequest)(nil),       // 2: spaceone.api.inventory.plugin.VerifyRequest
	(*CollectRequest)(nil),      // 3: spaceone.api.inventory.plugin.CollectRequest
	(*ResourceInfo)(nil),        // 4: spaceone.api.inventory.plugin.ResourceInfo
	(*CollectorVerifyInfo)(nil), // 5: spaceone.api.inventory.plugin.CollectorVerifyInfo
	(*PluginInfo)(nil),          // 6: spaceone.api.inventory.plugin.PluginInfo
	(*_struct.Struct)(nil),      // 7: google.protobuf.Struct
	(*empty.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_spaceone_api_inventory_plugin_collector_proto_depIdxs = []int32{
	7,  // 0: spaceone.api.inventory.plugin.InitRequest.options:type_name -> google.protobuf.Struct
	7,  // 1: spaceone.api.inventory.plugin.VerifyRequest.options:type_name -> google.protobuf.Struct
	7,  // 2: spaceone.api.inventory.plugin.VerifyRequest.secret_data:type_name -> google.protobuf.Struct
	7,  // 3: spaceone.api.inventory.plugin.CollectRequest.options:type_name -> google.protobuf.Struct
	7,  // 4: spaceone.api.inventory.plugin.CollectRequest.secret_data:type_name -> google.protobuf.Struct
	7,  // 5: spaceone.api.inventory.plugin.CollectRequest.filter:type_name -> google.protobuf.Struct
	7,  // 6: spaceone.api.inventory.plugin.CollectRequest.task_options:type_name -> google.protobuf.Struct
	0,  // 7: spaceone.api.inventory.plugin.ResourceInfo.state:type_name -> spaceone.api.inventory.plugin.ResourceInfo.State
	7,  // 8: spaceone.api.inventory.plugin.ResourceInfo.match_rules:type_name -> google.protobuf.Struct
	7,  // 9: spaceone.api.inventory.plugin.ResourceInfo.resource:type_name -> google.protobuf.Struct
	7,  // 10: spaceone.api.inventory.plugin.ResourceInfo.options:type_name -> google.protobuf.Struct
	7,  // 11: spaceone.api.inventory.plugin.CollectorVerifyInfo.options:type_name -> google.protobuf.Struct
	7,  // 12: spaceone.api.inventory.plugin.PluginInfo.metadata:type_name -> google.protobuf.Struct
	1,  // 13: spaceone.api.inventory.plugin.Collector.init:input_type -> spaceone.api.inventory.plugin.InitRequest
	2,  // 14: spaceone.api.inventory.plugin.Collector.verify:input_type -> spaceone.api.inventory.plugin.VerifyRequest
	3,  // 15: spaceone.api.inventory.plugin.Collector.collect:input_type -> spaceone.api.inventory.plugin.CollectRequest
	6,  // 16: spaceone.api.inventory.plugin.Collector.init:output_type -> spaceone.api.inventory.plugin.PluginInfo
	8,  // 17: spaceone.api.inventory.plugin.Collector.verify:output_type -> google.protobuf.Empty
	4,  // 18: spaceone.api.inventory.plugin.Collector.collect:output_type -> spaceone.api.inventory.plugin.ResourceInfo
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_plugin_collector_proto_init() }
func file_spaceone_api_inventory_plugin_collector_proto_init() {
	if File_spaceone_api_inventory_plugin_collector_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_inventory_plugin_collector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_plugin_collector_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_plugin_collector_proto_depIdxs,
		EnumInfos:         file_spaceone_api_inventory_plugin_collector_proto_enumTypes,
		MessageInfos:      file_spaceone_api_inventory_plugin_collector_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_plugin_collector_proto = out.File
	file_spaceone_api_inventory_plugin_collector_proto_rawDesc = nil
	file_spaceone_api_inventory_plugin_collector_proto_goTypes = nil
	file_spaceone_api_inventory_plugin_collector_proto_depIdxs = nil
}
