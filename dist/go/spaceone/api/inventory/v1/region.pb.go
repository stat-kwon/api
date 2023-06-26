// A Region is a resource storing regional information from each cloud service provider. Regional data stored by the resource includes the latitude and longitude of the region.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/inventory/v1/region.proto

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

//	{
//	   "name": "Asia Pacific (Seoul)",
//	   "region_code": "ap-northeast-2",
//	   "provider": "aws",
//	   "tags": {
//	       "continent": "asis_pacific",
//	       "longitude": "73.013805",
//	       "latitude": "19.147428"
//	   }
//	}
type CreateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	Tags     *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId string          `protobuf:"bytes,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CreateRegionRequest) Reset() {
	*x = CreateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegionRequest) ProtoMessage() {}

func (x *CreateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegionRequest.ProtoReflect.Descriptor instead.
func (*CreateRegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRegionRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CreateRegionRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateRegionRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateRegionRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "region_id": "region-e41deed3c939",
//	   "name": "Region home",
//	   "tags": {
//	       "latitude": "6.34545",
//	       "continent": "asia_pacific",
//	       "longitude": "5.6433213"
//	   }
//	}
type UpdateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags     *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId string          `protobuf:"bytes,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *UpdateRegionRequest) Reset() {
	*x = UpdateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionRequest) ProtoMessage() {}

func (x *UpdateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRegionRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *UpdateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRegionRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateRegionRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "region_id": "region-e41deed3c939"
//	}
type RegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *RegionRequest) Reset() {
	*x = RegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionRequest) ProtoMessage() {}

func (x *RegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionRequest.ProtoReflect.Descriptor instead.
func (*RegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{2}
}

func (x *RegionRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "region_id": "region-f803eb00b567"
//	}
type GetRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId string `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	DomainId string `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	Only []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
}

func (x *GetRegionRequest) Reset() {
	*x = GetRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionRequest) ProtoMessage() {}

func (x *GetRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionRequest.ProtoReflect.Descriptor instead.
func (*GetRegionRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{3}
}

func (x *GetRegionRequest) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *GetRegionRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *GetRegionRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

//	{
//	   "query": {
//	       "filter": [
//	           {
//	               "key": "name",
//	               "value": "Asia Pacific",
//	               "operator": "contain"
//	           }
//	       ]
//	   }
//	}
type RegionQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	RegionId string `protobuf:"bytes,2,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	RegionKey string `protobuf:"bytes,4,opt,name=region_key,json=regionKey,proto3" json:"region_key,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
	DomainId string `protobuf:"bytes,7,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *RegionQuery) Reset() {
	*x = RegionQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionQuery) ProtoMessage() {}

func (x *RegionQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionQuery.ProtoReflect.Descriptor instead.
func (*RegionQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{4}
}

func (x *RegionQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RegionQuery) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionQuery) GetRegionKey() string {
	if x != nil {
		return x.RegionKey
	}
	return ""
}

func (x *RegionQuery) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RegionQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *RegionQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "region_id": "region-e41deed3c939",
//	   "name": "Asia Pacific (Seoul)",
//	   "region_key": "aws.ap-northeast-2",
//	   "region_code": "ap-northeast-2",
//	   "provider": "aws",
//	   "tags": {
//	       "continent": "asia_pacific",
//	       "longitude": "73.013805",
//	       "latitude": "19.147428"
//	   },
//	   "domain_id": "domain-x1b3c34v432",
//	   "created_at": "2021-11-18T13:07:31.382Z",
//	   "updated_at": "2022-06-17T00:07:35.469Z"
//	}
type RegionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionId   string          `protobuf:"bytes,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Name       string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegionKey  string          `protobuf:"bytes,3,opt,name=region_key,json=regionKey,proto3" json:"region_key,omitempty"`
	RegionCode string          `protobuf:"bytes,4,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Provider   string          `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	Tags       *_struct.Struct `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	DomainId   string          `protobuf:"bytes,7,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt  string          `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string          `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *RegionInfo) Reset() {
	*x = RegionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionInfo) ProtoMessage() {}

func (x *RegionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionInfo.ProtoReflect.Descriptor instead.
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{5}
}

func (x *RegionInfo) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *RegionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionInfo) GetRegionKey() string {
	if x != nil {
		return x.RegionKey
	}
	return ""
}

func (x *RegionInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *RegionInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *RegionInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RegionInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *RegionInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RegionInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "region_id": "region-e41deed3c939",
//	           "name": "Asia Pacific (Mumbai)",
//	           "region_key": "aws.ap-south-1",
//	           "region_code": "ap-south-1",
//	           "provider": "aws",
//	           "tags": {
//	               "continent": "asia_pacific",
//	               "longitude": "73.013805",
//	               "latitude": "19.147428"
//	           },
//	           "domain_id": "domain-x1b3c34v432",
//	           "created_at": "2021-11-18T13:07:31.382Z",
//	           "updated_at": "2022-06-17T00:07:35.469Z"
//	       },
//	       {
//	           "region_id": "region-f803eb00b567",
//	           "name": "Asia Pacific (Seoul)",
//	           "region_key": "aws.ap-northeast-2",
//	           "region_code": "ap-northeast-2",
//	           "provider": "aws",
//	           "tags": {
//	               "latitude": "6.34545",
//	               "continent": "asia_pacific",
//	               "longitude": "5.6433213"
//	           },
//	           "domain_id": "domain-x1b3c34v432",
//	           "created_at": "2022-03-21T09:08:31.961Z",
//	           "updated_at": "2022-06-17T00:07:35.749Z"
//	       }
//	   ],
//	   "total_count": 2
//	}
type RegionsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*RegionInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32         `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *RegionsInfo) Reset() {
	*x = RegionsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionsInfo) ProtoMessage() {}

func (x *RegionsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionsInfo.ProtoReflect.Descriptor instead.
func (*RegionsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{6}
}

func (x *RegionsInfo) GetResults() []*RegionInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *RegionsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type RegionStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *RegionStatQuery) Reset() {
	*x = RegionStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionStatQuery) ProtoMessage() {}

func (x *RegionStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_region_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionStatQuery.ProtoReflect.Descriptor instead.
func (*RegionStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_region_proto_rawDescGZIP(), []int{7}
}

func (x *RegionStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *RegionStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_inventory_v1_region_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v1_region_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0x90, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f, 0x6e, 0x6c,
	0x79, 0x22, 0xea, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xa1,
	0x02, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x6f, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x32, 0x81, 0x06, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x87, 0x01, 0x0a, 0x06,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a,
	0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x72, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x7e, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f,
	0x67, 0x65, 0x74, 0x12, 0x7c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x12, 0x71, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x69, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_v1_region_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v1_region_proto_rawDescData = file_spaceone_api_inventory_v1_region_proto_rawDesc
)

func file_spaceone_api_inventory_v1_region_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v1_region_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v1_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_v1_region_proto_rawDescData)
	})
	return file_spaceone_api_inventory_v1_region_proto_rawDescData
}

var file_spaceone_api_inventory_v1_region_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_inventory_v1_region_proto_goTypes = []interface{}{
	(*CreateRegionRequest)(nil), // 0: spaceone.api.inventory.v1.CreateRegionRequest
	(*UpdateRegionRequest)(nil), // 1: spaceone.api.inventory.v1.UpdateRegionRequest
	(*RegionRequest)(nil),       // 2: spaceone.api.inventory.v1.RegionRequest
	(*GetRegionRequest)(nil),    // 3: spaceone.api.inventory.v1.GetRegionRequest
	(*RegionQuery)(nil),         // 4: spaceone.api.inventory.v1.RegionQuery
	(*RegionInfo)(nil),          // 5: spaceone.api.inventory.v1.RegionInfo
	(*RegionsInfo)(nil),         // 6: spaceone.api.inventory.v1.RegionsInfo
	(*RegionStatQuery)(nil),     // 7: spaceone.api.inventory.v1.RegionStatQuery
	(*_struct.Struct)(nil),      // 8: google.protobuf.Struct
	(*v1.Query)(nil),            // 9: spaceone.api.core.v1.Query
	(*v1.StatisticsQuery)(nil),  // 10: spaceone.api.core.v1.StatisticsQuery
	(*empty.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_spaceone_api_inventory_v1_region_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.inventory.v1.CreateRegionRequest.tags:type_name -> google.protobuf.Struct
	8,  // 1: spaceone.api.inventory.v1.UpdateRegionRequest.tags:type_name -> google.protobuf.Struct
	9,  // 2: spaceone.api.inventory.v1.RegionQuery.query:type_name -> spaceone.api.core.v1.Query
	8,  // 3: spaceone.api.inventory.v1.RegionInfo.tags:type_name -> google.protobuf.Struct
	5,  // 4: spaceone.api.inventory.v1.RegionsInfo.results:type_name -> spaceone.api.inventory.v1.RegionInfo
	10, // 5: spaceone.api.inventory.v1.RegionStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	0,  // 6: spaceone.api.inventory.v1.Region.create:input_type -> spaceone.api.inventory.v1.CreateRegionRequest
	1,  // 7: spaceone.api.inventory.v1.Region.update:input_type -> spaceone.api.inventory.v1.UpdateRegionRequest
	2,  // 8: spaceone.api.inventory.v1.Region.delete:input_type -> spaceone.api.inventory.v1.RegionRequest
	3,  // 9: spaceone.api.inventory.v1.Region.get:input_type -> spaceone.api.inventory.v1.GetRegionRequest
	4,  // 10: spaceone.api.inventory.v1.Region.list:input_type -> spaceone.api.inventory.v1.RegionQuery
	7,  // 11: spaceone.api.inventory.v1.Region.stat:input_type -> spaceone.api.inventory.v1.RegionStatQuery
	5,  // 12: spaceone.api.inventory.v1.Region.create:output_type -> spaceone.api.inventory.v1.RegionInfo
	5,  // 13: spaceone.api.inventory.v1.Region.update:output_type -> spaceone.api.inventory.v1.RegionInfo
	11, // 14: spaceone.api.inventory.v1.Region.delete:output_type -> google.protobuf.Empty
	5,  // 15: spaceone.api.inventory.v1.Region.get:output_type -> spaceone.api.inventory.v1.RegionInfo
	6,  // 16: spaceone.api.inventory.v1.Region.list:output_type -> spaceone.api.inventory.v1.RegionsInfo
	8,  // 17: spaceone.api.inventory.v1.Region.stat:output_type -> google.protobuf.Struct
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v1_region_proto_init() }
func file_spaceone_api_inventory_v1_region_proto_init() {
	if File_spaceone_api_inventory_v1_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_inventory_v1_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegionRequest); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegionRequest); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionRequest); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegionRequest); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionQuery); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionInfo); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionsInfo); i {
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
		file_spaceone_api_inventory_v1_region_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionStatQuery); i {
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
			RawDescriptor: file_spaceone_api_inventory_v1_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v1_region_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v1_region_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v1_region_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v1_region_proto = out.File
	file_spaceone_api_inventory_v1_region_proto_rawDesc = nil
	file_spaceone_api_inventory_v1_region_proto_goTypes = nil
	file_spaceone_api_inventory_v1_region_proto_depIdxs = nil
}
