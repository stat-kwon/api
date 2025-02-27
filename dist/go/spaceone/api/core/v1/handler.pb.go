// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/core/v1/handler.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthorizationRequest_Scope int32

const (
	AuthorizationRequest_NONE              AuthorizationRequest_Scope = 0
	AuthorizationRequest_SYSTEM            AuthorizationRequest_Scope = 1
	AuthorizationRequest_DOMAIN            AuthorizationRequest_Scope = 2
	AuthorizationRequest_PROJECT           AuthorizationRequest_Scope = 3
	AuthorizationRequest_USER              AuthorizationRequest_Scope = 4
	AuthorizationRequest_PUBLIC            AuthorizationRequest_Scope = 5
	AuthorizationRequest_PUBLIC_OR_DOMAIN  AuthorizationRequest_Scope = 6
	AuthorizationRequest_DOMAIN_OR_PROJECT AuthorizationRequest_Scope = 7
)

// Enum value maps for AuthorizationRequest_Scope.
var (
	AuthorizationRequest_Scope_name = map[int32]string{
		0: "NONE",
		1: "SYSTEM",
		2: "DOMAIN",
		3: "PROJECT",
		4: "USER",
		5: "PUBLIC",
		6: "PUBLIC_OR_DOMAIN",
		7: "DOMAIN_OR_PROJECT",
	}
	AuthorizationRequest_Scope_value = map[string]int32{
		"NONE":              0,
		"SYSTEM":            1,
		"DOMAIN":            2,
		"PROJECT":           3,
		"USER":              4,
		"PUBLIC":            5,
		"PUBLIC_OR_DOMAIN":  6,
		"DOMAIN_OR_PROJECT": 7,
	}
)

func (x AuthorizationRequest_Scope) Enum() *AuthorizationRequest_Scope {
	p := new(AuthorizationRequest_Scope)
	*p = x
	return p
}

func (x AuthorizationRequest_Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthorizationRequest_Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_core_v1_handler_proto_enumTypes[0].Descriptor()
}

func (AuthorizationRequest_Scope) Type() protoreflect.EnumType {
	return &file_spaceone_api_core_v1_handler_proto_enumTypes[0]
}

func (x AuthorizationRequest_Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthorizationRequest_Scope.Descriptor instead.
func (AuthorizationRequest_Scope) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{0, 0}
}

type AuthorizationRequest struct {
	state                 protoimpl.MessageState     `protogen:"open.v1"`
	Service               string                     `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Resource              string                     `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Verb                  string                     `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	Scope                 AuthorizationRequest_Scope `protobuf:"varint,4,opt,name=scope,proto3,enum=spaceone.api.core.v1.AuthorizationRequest_Scope" json:"scope,omitempty"`
	DomainId              string                     `protobuf:"bytes,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	ProjectId             string                     `protobuf:"bytes,6,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectGroupId        string                     `protobuf:"bytes,7,opt,name=project_group_id,json=projectGroupId,proto3" json:"project_group_id,omitempty"`
	UserId                string                     `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RequireProjectId      bool                       `protobuf:"varint,9,opt,name=require_project_id,json=requireProjectId,proto3" json:"require_project_id,omitempty"`
	RequireProjectGroupId bool                       `protobuf:"varint,10,opt,name=require_project_group_id,json=requireProjectGroupId,proto3" json:"require_project_group_id,omitempty"`
	RequireUserId         bool                       `protobuf:"varint,11,opt,name=require_user_id,json=requireUserId,proto3" json:"require_user_id,omitempty"`
	RequireDomainId       bool                       `protobuf:"varint,12,opt,name=require_domain_id,json=requireDomainId,proto3" json:"require_domain_id,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *AuthorizationRequest) Reset() {
	*x = AuthorizationRequest{}
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationRequest) ProtoMessage() {}

func (x *AuthorizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationRequest.ProtoReflect.Descriptor instead.
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizationRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *AuthorizationRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AuthorizationRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *AuthorizationRequest) GetScope() AuthorizationRequest_Scope {
	if x != nil {
		return x.Scope
	}
	return AuthorizationRequest_NONE
}

func (x *AuthorizationRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AuthorizationRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *AuthorizationRequest) GetProjectGroupId() string {
	if x != nil {
		return x.ProjectGroupId
	}
	return ""
}

func (x *AuthorizationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthorizationRequest) GetRequireProjectId() bool {
	if x != nil {
		return x.RequireProjectId
	}
	return false
}

func (x *AuthorizationRequest) GetRequireProjectGroupId() bool {
	if x != nil {
		return x.RequireProjectGroupId
	}
	return false
}

func (x *AuthorizationRequest) GetRequireUserId() bool {
	if x != nil {
		return x.RequireUserId
	}
	return false
}

func (x *AuthorizationRequest) GetRequireDomainId() bool {
	if x != nil {
		return x.RequireDomainId
	}
	return false
}

type AuthorizationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoleType      string                 `protobuf:"bytes,1,opt,name=role_type,json=roleType,proto3" json:"role_type,omitempty"`
	Projects      []string               `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	ProjectGroups []string               `protobuf:"bytes,3,rep,name=project_groups,json=projectGroups,proto3" json:"project_groups,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorizationResponse) Reset() {
	*x = AuthorizationResponse{}
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationResponse) ProtoMessage() {}

func (x *AuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizationResponse) GetRoleType() string {
	if x != nil {
		return x.RoleType
	}
	return ""
}

func (x *AuthorizationResponse) GetProjects() []string {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *AuthorizationResponse) GetProjectGroups() []string {
	if x != nil {
		return x.ProjectGroups
	}
	return nil
}

type AuthenticationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DomainId      string                 `protobuf:"bytes,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	PublicKey     string                 `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{3}
}

func (x *AuthenticationResponse) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *AuthenticationResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type EventRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       string                 `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Resource      string                 `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Verb          string                 `protobuf:"bytes,3,opt,name=verb,proto3" json:"verb,omitempty"`
	Status        string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Message       *_struct.Struct        `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_core_v1_handler_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_core_v1_handler_proto_rawDescGZIP(), []int{4}
}

func (x *EventRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EventRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *EventRequest) GetVerb() string {
	if x != nil {
		return x.Verb
	}
	return ""
}

func (x *EventRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EventRequest) GetMessage() *_struct.Struct {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_spaceone_api_core_v1_handler_proto protoreflect.FileDescriptor

var file_spaceone_api_core_v1_handler_proto_rawDesc = string([]byte{
	0x0a, 0x22, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x04, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x46, 0x0a, 0x05, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x79, 0x0a,
	0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x4a,
	0x45, 0x43, 0x54, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x43, 0x5f, 0x4f, 0x52, 0x5f, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x10,
	0x06, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x4f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x50,
	0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x07, 0x22, 0x77, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x22, 0x34, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x16, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0xa3, 0x01,
	0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x65, 0x72, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x76, 0x65, 0x72, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_spaceone_api_core_v1_handler_proto_rawDescOnce sync.Once
	file_spaceone_api_core_v1_handler_proto_rawDescData []byte
)

func file_spaceone_api_core_v1_handler_proto_rawDescGZIP() []byte {
	file_spaceone_api_core_v1_handler_proto_rawDescOnce.Do(func() {
		file_spaceone_api_core_v1_handler_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v1_handler_proto_rawDesc), len(file_spaceone_api_core_v1_handler_proto_rawDesc)))
	})
	return file_spaceone_api_core_v1_handler_proto_rawDescData
}

var file_spaceone_api_core_v1_handler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_core_v1_handler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_core_v1_handler_proto_goTypes = []any{
	(AuthorizationRequest_Scope)(0), // 0: spaceone.api.core.v1.AuthorizationRequest.Scope
	(*AuthorizationRequest)(nil),    // 1: spaceone.api.core.v1.AuthorizationRequest
	(*AuthorizationResponse)(nil),   // 2: spaceone.api.core.v1.AuthorizationResponse
	(*AuthenticationRequest)(nil),   // 3: spaceone.api.core.v1.AuthenticationRequest
	(*AuthenticationResponse)(nil),  // 4: spaceone.api.core.v1.AuthenticationResponse
	(*EventRequest)(nil),            // 5: spaceone.api.core.v1.EventRequest
	(*_struct.Struct)(nil),          // 6: google.protobuf.Struct
}
var file_spaceone_api_core_v1_handler_proto_depIdxs = []int32{
	0, // 0: spaceone.api.core.v1.AuthorizationRequest.scope:type_name -> spaceone.api.core.v1.AuthorizationRequest.Scope
	6, // 1: spaceone.api.core.v1.EventRequest.message:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spaceone_api_core_v1_handler_proto_init() }
func file_spaceone_api_core_v1_handler_proto_init() {
	if File_spaceone_api_core_v1_handler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_core_v1_handler_proto_rawDesc), len(file_spaceone_api_core_v1_handler_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaceone_api_core_v1_handler_proto_goTypes,
		DependencyIndexes: file_spaceone_api_core_v1_handler_proto_depIdxs,
		EnumInfos:         file_spaceone_api_core_v1_handler_proto_enumTypes,
		MessageInfos:      file_spaceone_api_core_v1_handler_proto_msgTypes,
	}.Build()
	File_spaceone_api_core_v1_handler_proto = out.File
	file_spaceone_api_core_v1_handler_proto_goTypes = nil
	file_spaceone_api_core_v1_handler_proto_depIdxs = nil
}
