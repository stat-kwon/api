// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/api_key.proto

package v2

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type APIKeyInfo_State int32

const (
	APIKeyInfo_NONE_STATE APIKeyInfo_State = 0
	APIKeyInfo_ENABLED    APIKeyInfo_State = 1
	APIKeyInfo_DISABLED   APIKeyInfo_State = 2
	APIKeyInfo_EXPIRED    APIKeyInfo_State = 3
)

// Enum value maps for APIKeyInfo_State.
var (
	APIKeyInfo_State_name = map[int32]string{
		0: "NONE_STATE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	APIKeyInfo_State_value = map[string]int32{
		"NONE_STATE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x APIKeyInfo_State) Enum() *APIKeyInfo_State {
	p := new(APIKeyInfo_State)
	*p = x
	return p
}

func (x APIKeyInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIKeyInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_api_key_proto_enumTypes[0].Descriptor()
}

func (APIKeyInfo_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_api_key_proto_enumTypes[0]
}

func (x APIKeyInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIKeyInfo_State.Descriptor instead.
func (APIKeyInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{3, 0}
}

type APIKeySearchQuery_State int32

const (
	APIKeySearchQuery_NONE_STATE APIKeySearchQuery_State = 0
	APIKeySearchQuery_ENABLED    APIKeySearchQuery_State = 1
	APIKeySearchQuery_DISABLED   APIKeySearchQuery_State = 2
	APIKeySearchQuery_EXPIRED    APIKeySearchQuery_State = 3
)

// Enum value maps for APIKeySearchQuery_State.
var (
	APIKeySearchQuery_State_name = map[int32]string{
		0: "NONE_STATE",
		1: "ENABLED",
		2: "DISABLED",
		3: "EXPIRED",
	}
	APIKeySearchQuery_State_value = map[string]int32{
		"NONE_STATE": 0,
		"ENABLED":    1,
		"DISABLED":   2,
		"EXPIRED":    3,
	}
)

func (x APIKeySearchQuery_State) Enum() *APIKeySearchQuery_State {
	p := new(APIKeySearchQuery_State)
	*p = x
	return p
}

func (x APIKeySearchQuery_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APIKeySearchQuery_State) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v2_api_key_proto_enumTypes[1].Descriptor()
}

func (APIKeySearchQuery_State) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v2_api_key_proto_enumTypes[1]
}

func (x APIKeySearchQuery_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APIKeySearchQuery_State.Descriptor instead.
func (APIKeySearchQuery_State) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{4, 0}
}

type CreateAPIKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	ExpiredAt string `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
	DomainId  string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CreateAPIKeyRequest) Reset() {
	*x = CreateAPIKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAPIKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAPIKeyRequest) ProtoMessage() {}

func (x *CreateAPIKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAPIKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateAPIKeyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAPIKeyRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateAPIKeyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAPIKeyRequest) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

func (x *CreateAPIKeyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type UpdateAPIKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKeyId string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	// +optional
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *UpdateAPIKeyRequest) Reset() {
	*x = UpdateAPIKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAPIKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAPIKeyRequest) ProtoMessage() {}

func (x *UpdateAPIKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAPIKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateAPIKeyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAPIKeyRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *UpdateAPIKeyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAPIKeyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type APIKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKeyId string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *APIKeyRequest) Reset() {
	*x = APIKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeyRequest) ProtoMessage() {}

func (x *APIKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeyRequest.ProtoReflect.Descriptor instead.
func (*APIKeyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{2}
}

func (x *APIKeyRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *APIKeyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type APIKeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKeyId       string           `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	ApiKey         string           `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	Name           string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	State          APIKeyInfo_State `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v2.APIKeyInfo_State" json:"state,omitempty"`
	DomainId       string           `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	UserId         string           `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt      string           `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	LastAccessedAt string           `protobuf:"bytes,32,opt,name=last_accessed_at,json=lastAccessedAt,proto3" json:"last_accessed_at,omitempty"`
	ExpiredAt      string           `protobuf:"bytes,33,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *APIKeyInfo) Reset() {
	*x = APIKeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeyInfo) ProtoMessage() {}

func (x *APIKeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeyInfo.ProtoReflect.Descriptor instead.
func (*APIKeyInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{3}
}

func (x *APIKeyInfo) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *APIKeyInfo) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *APIKeyInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *APIKeyInfo) GetState() APIKeyInfo_State {
	if x != nil {
		return x.State
	}
	return APIKeyInfo_NONE_STATE
}

func (x *APIKeyInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *APIKeyInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *APIKeyInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *APIKeyInfo) GetLastAccessedAt() string {
	if x != nil {
		return x.LastAccessedAt
	}
	return ""
}

func (x *APIKeyInfo) GetExpiredAt() string {
	if x != nil {
		return x.ExpiredAt
	}
	return ""
}

type APIKeySearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	ApiKeyId string `protobuf:"bytes,2,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	// +optional
	State    APIKeySearchQuery_State `protobuf:"varint,3,opt,name=state,proto3,enum=spaceone.api.identity.v2.APIKeySearchQuery_State" json:"state,omitempty"`
	DomainId string                  `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	UserId string `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *APIKeySearchQuery) Reset() {
	*x = APIKeySearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeySearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeySearchQuery) ProtoMessage() {}

func (x *APIKeySearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeySearchQuery.ProtoReflect.Descriptor instead.
func (*APIKeySearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{4}
}

func (x *APIKeySearchQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *APIKeySearchQuery) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *APIKeySearchQuery) GetState() APIKeySearchQuery_State {
	if x != nil {
		return x.State
	}
	return APIKeySearchQuery_NONE_STATE
}

func (x *APIKeySearchQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *APIKeySearchQuery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type APIKeysInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*APIKeyInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *APIKeysInfo) Reset() {
	*x = APIKeysInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeysInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeysInfo) ProtoMessage() {}

func (x *APIKeysInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeysInfo.ProtoReflect.Descriptor instead.
func (*APIKeysInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{5}
}

func (x *APIKeysInfo) GetResults() []*APIKeyInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *APIKeysInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type APIKeyStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *APIKeyStatQuery) Reset() {
	*x = APIKeyStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKeyStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKeyStatQuery) ProtoMessage() {}

func (x *APIKeyStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_api_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKeyStatQuery.ProtoReflect.Descriptor instead.
func (*APIKeyStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP(), []int{6}
}

func (x *APIKeyStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *APIKeyStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_identity_v2_api_key_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_api_key_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x64, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0xf8, 0x02, 0x0a, 0x0a, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03, 0x22, 0xa4, 0x02, 0x0a, 0x11, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x49,
	0x64, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x31, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b,
	0x65, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x4e,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10,
	0x03, 0x22, 0x6e, 0x0a, 0x0b, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x6b, 0x0a, 0x0f, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0x80,
	0x08, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x85, 0x01, 0x0a, 0x06, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x85, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b,
	0x65, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x7f, 0x0a, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41,
	0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d,
	0x6b, 0x65, 0x79, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x07, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a,
	0x22, 0x1c, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x71,
	0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x79, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50, 0x49,
	0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a,
	0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x80, 0x01, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32,
	0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x50,
	0x49, 0x4b, 0x65, 0x79, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x70, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x6b, 0x65, 0x79, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_api_key_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_api_key_proto_rawDescData = file_spaceone_api_identity_v2_api_key_proto_rawDesc
)

func file_spaceone_api_identity_v2_api_key_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_api_key_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_api_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_api_key_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_api_key_proto_rawDescData
}

var file_spaceone_api_identity_v2_api_key_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_spaceone_api_identity_v2_api_key_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_spaceone_api_identity_v2_api_key_proto_goTypes = []interface{}{
	(APIKeyInfo_State)(0),        // 0: spaceone.api.identity.v2.APIKeyInfo.State
	(APIKeySearchQuery_State)(0), // 1: spaceone.api.identity.v2.APIKeySearchQuery.State
	(*CreateAPIKeyRequest)(nil),  // 2: spaceone.api.identity.v2.CreateAPIKeyRequest
	(*UpdateAPIKeyRequest)(nil),  // 3: spaceone.api.identity.v2.UpdateAPIKeyRequest
	(*APIKeyRequest)(nil),        // 4: spaceone.api.identity.v2.APIKeyRequest
	(*APIKeyInfo)(nil),           // 5: spaceone.api.identity.v2.APIKeyInfo
	(*APIKeySearchQuery)(nil),    // 6: spaceone.api.identity.v2.APIKeySearchQuery
	(*APIKeysInfo)(nil),          // 7: spaceone.api.identity.v2.APIKeysInfo
	(*APIKeyStatQuery)(nil),      // 8: spaceone.api.identity.v2.APIKeyStatQuery
	(*v1.Query)(nil),             // 9: spaceone.api.core.v1.Query
	(*v1.StatisticsQuery)(nil),   // 10: spaceone.api.core.v1.StatisticsQuery
	(*empty.Empty)(nil),          // 11: google.protobuf.Empty
	(*_struct.Struct)(nil),       // 12: google.protobuf.Struct
}
var file_spaceone_api_identity_v2_api_key_proto_depIdxs = []int32{
	0,  // 0: spaceone.api.identity.v2.APIKeyInfo.state:type_name -> spaceone.api.identity.v2.APIKeyInfo.State
	9,  // 1: spaceone.api.identity.v2.APIKeySearchQuery.query:type_name -> spaceone.api.core.v1.Query
	1,  // 2: spaceone.api.identity.v2.APIKeySearchQuery.state:type_name -> spaceone.api.identity.v2.APIKeySearchQuery.State
	5,  // 3: spaceone.api.identity.v2.APIKeysInfo.results:type_name -> spaceone.api.identity.v2.APIKeyInfo
	10, // 4: spaceone.api.identity.v2.APIKeyStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	2,  // 5: spaceone.api.identity.v2.APIKey.create:input_type -> spaceone.api.identity.v2.CreateAPIKeyRequest
	3,  // 6: spaceone.api.identity.v2.APIKey.update:input_type -> spaceone.api.identity.v2.UpdateAPIKeyRequest
	4,  // 7: spaceone.api.identity.v2.APIKey.enable:input_type -> spaceone.api.identity.v2.APIKeyRequest
	4,  // 8: spaceone.api.identity.v2.APIKey.disable:input_type -> spaceone.api.identity.v2.APIKeyRequest
	4,  // 9: spaceone.api.identity.v2.APIKey.delete:input_type -> spaceone.api.identity.v2.APIKeyRequest
	4,  // 10: spaceone.api.identity.v2.APIKey.get:input_type -> spaceone.api.identity.v2.APIKeyRequest
	6,  // 11: spaceone.api.identity.v2.APIKey.list:input_type -> spaceone.api.identity.v2.APIKeySearchQuery
	8,  // 12: spaceone.api.identity.v2.APIKey.stat:input_type -> spaceone.api.identity.v2.APIKeyStatQuery
	5,  // 13: spaceone.api.identity.v2.APIKey.create:output_type -> spaceone.api.identity.v2.APIKeyInfo
	5,  // 14: spaceone.api.identity.v2.APIKey.update:output_type -> spaceone.api.identity.v2.APIKeyInfo
	5,  // 15: spaceone.api.identity.v2.APIKey.enable:output_type -> spaceone.api.identity.v2.APIKeyInfo
	5,  // 16: spaceone.api.identity.v2.APIKey.disable:output_type -> spaceone.api.identity.v2.APIKeyInfo
	11, // 17: spaceone.api.identity.v2.APIKey.delete:output_type -> google.protobuf.Empty
	5,  // 18: spaceone.api.identity.v2.APIKey.get:output_type -> spaceone.api.identity.v2.APIKeyInfo
	7,  // 19: spaceone.api.identity.v2.APIKey.list:output_type -> spaceone.api.identity.v2.APIKeysInfo
	12, // 20: spaceone.api.identity.v2.APIKey.stat:output_type -> google.protobuf.Struct
	13, // [13:21] is the sub-list for method output_type
	5,  // [5:13] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_api_key_proto_init() }
func file_spaceone_api_identity_v2_api_key_proto_init() {
	if File_spaceone_api_identity_v2_api_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAPIKeyRequest); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAPIKeyRequest); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeyRequest); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeyInfo); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeySearchQuery); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeysInfo); i {
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
		file_spaceone_api_identity_v2_api_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKeyStatQuery); i {
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
			RawDescriptor: file_spaceone_api_identity_v2_api_key_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_api_key_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_api_key_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v2_api_key_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v2_api_key_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_api_key_proto = out.File
	file_spaceone_api_identity_v2_api_key_proto_rawDesc = nil
	file_spaceone_api_identity_v2_api_key_proto_goTypes = nil
	file_spaceone_api_identity_v2_api_key_proto_depIdxs = nil
}
