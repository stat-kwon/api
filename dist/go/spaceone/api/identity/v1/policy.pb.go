// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v1/policy.proto

package v1

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

type CreatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// +optional
	Tags     *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId string          `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CreatePolicyRequest) Reset() {
	*x = CreatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolicyRequest) ProtoMessage() {}

func (x *CreatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolicyRequest.ProtoReflect.Descriptor instead.
func (*CreatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePolicyRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *CreatePolicyRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreatePolicyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type UpdatePolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Permissions []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// +optional
	Tags     *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId string          `protobuf:"bytes,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *UpdatePolicyRequest) Reset() {
	*x = UpdatePolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePolicyRequest) ProtoMessage() {}

func (x *UpdatePolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePolicyRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *UpdatePolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePolicyRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UpdatePolicyRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdatePolicyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type PolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *PolicyRequest) Reset() {
	*x = PolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyRequest) ProtoMessage() {}

func (x *PolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyRequest.ProtoReflect.Descriptor instead.
func (*PolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{2}
}

func (x *PolicyRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *PolicyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type GetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	Only []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{3}
}

func (x *GetPolicyRequest) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *GetPolicyRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetPolicyRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

type PolicyQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	PolicyId string `protobuf:"bytes,2,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *PolicyQuery) Reset() {
	*x = PolicyQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyQuery) ProtoMessage() {}

func (x *PolicyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyQuery.ProtoReflect.Descriptor instead.
func (*PolicyQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{4}
}

func (x *PolicyQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PolicyQuery) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *PolicyQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type PolicyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyId    string          `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	Name        string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions []string        `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	DomainId    string          `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Tags        *_struct.Struct `protobuf:"bytes,5,opt,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt   string          `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PolicyInfo) Reset() {
	*x = PolicyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyInfo) ProtoMessage() {}

func (x *PolicyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyInfo.ProtoReflect.Descriptor instead.
func (*PolicyInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{5}
}

func (x *PolicyInfo) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *PolicyInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PolicyInfo) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *PolicyInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *PolicyInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PolicyInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type PoliciesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*PolicyInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *PoliciesInfo) Reset() {
	*x = PoliciesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoliciesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoliciesInfo) ProtoMessage() {}

func (x *PoliciesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoliciesInfo.ProtoReflect.Descriptor instead.
func (*PoliciesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{6}
}

func (x *PoliciesInfo) GetResults() []*PolicyInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *PoliciesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type PolicyStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *PolicyStatQuery) Reset() {
	*x = PolicyStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolicyStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolicyStatQuery) ProtoMessage() {}

func (x *PolicyStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_policy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolicyStatQuery.ProtoReflect.Descriptor instead.
func (*PolicyStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_policy_proto_rawDescGZIP(), []int{7}
}

func (x *PolicyStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PolicyStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_identity_v1_policy_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v1_policy_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0d, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x6f, 0x0a, 0x0c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x0f, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x32, 0xf4, 0x05, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x84, 0x01, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x70, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22,
	0x1a, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x7b, 0x0a, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22,
	0x17, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x7a, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x71, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x29, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74,
	0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v1_policy_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v1_policy_proto_rawDescData = file_spaceone_api_identity_v1_policy_proto_rawDesc
)

func file_spaceone_api_identity_v1_policy_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v1_policy_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v1_policy_proto_rawDescData)
	})
	return file_spaceone_api_identity_v1_policy_proto_rawDescData
}

var file_spaceone_api_identity_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_identity_v1_policy_proto_goTypes = []interface{}{
	(*CreatePolicyRequest)(nil), // 0: spaceone.api.identity.v1.CreatePolicyRequest
	(*UpdatePolicyRequest)(nil), // 1: spaceone.api.identity.v1.UpdatePolicyRequest
	(*PolicyRequest)(nil),       // 2: spaceone.api.identity.v1.PolicyRequest
	(*GetPolicyRequest)(nil),    // 3: spaceone.api.identity.v1.GetPolicyRequest
	(*PolicyQuery)(nil),         // 4: spaceone.api.identity.v1.PolicyQuery
	(*PolicyInfo)(nil),          // 5: spaceone.api.identity.v1.PolicyInfo
	(*PoliciesInfo)(nil),        // 6: spaceone.api.identity.v1.PoliciesInfo
	(*PolicyStatQuery)(nil),     // 7: spaceone.api.identity.v1.PolicyStatQuery
	(*_struct.Struct)(nil),      // 8: google.protobuf.Struct
	(*v1.Query)(nil),            // 9: spaceone.api.core.v1.Query
	(*v1.StatisticsQuery)(nil),  // 10: spaceone.api.core.v1.StatisticsQuery
	(*empty.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_spaceone_api_identity_v1_policy_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.identity.v1.CreatePolicyRequest.tags:type_name -> google.protobuf.Struct
	8,  // 1: spaceone.api.identity.v1.UpdatePolicyRequest.tags:type_name -> google.protobuf.Struct
	9,  // 2: spaceone.api.identity.v1.PolicyQuery.query:type_name -> spaceone.api.core.v1.Query
	8,  // 3: spaceone.api.identity.v1.PolicyInfo.tags:type_name -> google.protobuf.Struct
	5,  // 4: spaceone.api.identity.v1.PoliciesInfo.results:type_name -> spaceone.api.identity.v1.PolicyInfo
	10, // 5: spaceone.api.identity.v1.PolicyStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	0,  // 6: spaceone.api.identity.v1.Policy.create:input_type -> spaceone.api.identity.v1.CreatePolicyRequest
	1,  // 7: spaceone.api.identity.v1.Policy.update:input_type -> spaceone.api.identity.v1.UpdatePolicyRequest
	2,  // 8: spaceone.api.identity.v1.Policy.delete:input_type -> spaceone.api.identity.v1.PolicyRequest
	3,  // 9: spaceone.api.identity.v1.Policy.get:input_type -> spaceone.api.identity.v1.GetPolicyRequest
	4,  // 10: spaceone.api.identity.v1.Policy.list:input_type -> spaceone.api.identity.v1.PolicyQuery
	7,  // 11: spaceone.api.identity.v1.Policy.stat:input_type -> spaceone.api.identity.v1.PolicyStatQuery
	5,  // 12: spaceone.api.identity.v1.Policy.create:output_type -> spaceone.api.identity.v1.PolicyInfo
	5,  // 13: spaceone.api.identity.v1.Policy.update:output_type -> spaceone.api.identity.v1.PolicyInfo
	11, // 14: spaceone.api.identity.v1.Policy.delete:output_type -> google.protobuf.Empty
	5,  // 15: spaceone.api.identity.v1.Policy.get:output_type -> spaceone.api.identity.v1.PolicyInfo
	6,  // 16: spaceone.api.identity.v1.Policy.list:output_type -> spaceone.api.identity.v1.PoliciesInfo
	8,  // 17: spaceone.api.identity.v1.Policy.stat:output_type -> google.protobuf.Struct
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v1_policy_proto_init() }
func file_spaceone_api_identity_v1_policy_proto_init() {
	if File_spaceone_api_identity_v1_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v1_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolicyRequest); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePolicyRequest); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyRequest); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyRequest); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyQuery); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyInfo); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoliciesInfo); i {
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
		file_spaceone_api_identity_v1_policy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolicyStatQuery); i {
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
			RawDescriptor: file_spaceone_api_identity_v1_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v1_policy_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v1_policy_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v1_policy_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v1_policy_proto = out.File
	file_spaceone_api_identity_v1_policy_proto_rawDesc = nil
	file_spaceone_api_identity_v1_policy_proto_goTypes = nil
	file_spaceone_api_identity_v1_policy_proto_depIdxs = nil
}
