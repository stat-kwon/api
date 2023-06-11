// A CloudServiceStats is statistics data created from from cloud service query sets.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: spaceone/api/inventory/v1/cloud_service_stats.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
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
//	   "query": <SearchQuery>,
//	   "query_set_id": "query-set-abcd1234",
//	   "key": "Disk Size
//	   "provider": "aws",
//	   "cloud_service_group": "EC2",
//	   "cloud_service_type": "Instance",
//	   "region_code": "us-east-1",
//	   "account": "aws-account-id",
//	   "project_id": "project-abcd1234"
//	}
type CloudServiceStatsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	QuerySetId string `protobuf:"bytes,2,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	// +optional
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// +optional
	Provider string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	// +optional
	CloudServiceGroup string `protobuf:"bytes,5,opt,name=cloud_service_group,json=cloudServiceGroup,proto3" json:"cloud_service_group,omitempty"`
	// +optional
	CloudServiceType string `protobuf:"bytes,6,opt,name=cloud_service_type,json=cloudServiceType,proto3" json:"cloud_service_type,omitempty"`
	// +optional
	RegionCode string `protobuf:"bytes,7,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// +optional
	Account string `protobuf:"bytes,8,opt,name=account,proto3" json:"account,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,11,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DomainId  string `protobuf:"bytes,12,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CloudServiceStatsQuery) Reset() {
	*x = CloudServiceStatsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudServiceStatsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsQuery) ProtoMessage() {}

func (x *CloudServiceStatsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{0}
}

func (x *CloudServiceStatsQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsQuery) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetCloudServiceGroup() string {
	if x != nil {
		return x.CloudServiceGroup
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetCloudServiceType() string {
	if x != nil {
		return x.CloudServiceType
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CloudServiceStatsQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

//	{
//	   "query_set_id": "query-set-abcd1234",
//	   "key": "Disk Size",
//	   "value": "1040",
//	   "unit": "GB",
//	   "provider": "aws",
//	   "cloud_service_group": "EC2",
//	   "cloud_service_type": "Instance",
//	   "region_code": "us-east-1",
//	   "account": "aws-account-id",
//	   "additional_info": {
//	       "instance_type": "t2.micro"
//	   },
//	   "project_id": "project-abcd1234",
//	   "domain_id": "domain-58010aa2e451",
//	   "created_at": "2022-06-22T01:38:16.301Z"
//	}
type CloudServiceStatInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuerySetId        string          `protobuf:"bytes,1,opt,name=query_set_id,json=querySetId,proto3" json:"query_set_id,omitempty"`
	Key               string          `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value             float32         `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	Unit              string          `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
	Provider          string          `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
	CloudServiceGroup string          `protobuf:"bytes,6,opt,name=cloud_service_group,json=cloudServiceGroup,proto3" json:"cloud_service_group,omitempty"`
	CloudServiceType  string          `protobuf:"bytes,7,opt,name=cloud_service_type,json=cloudServiceType,proto3" json:"cloud_service_type,omitempty"`
	RegionCode        string          `protobuf:"bytes,8,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	Account           string          `protobuf:"bytes,9,opt,name=account,proto3" json:"account,omitempty"`
	AdditionalInfo    *_struct.Struct `protobuf:"bytes,11,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	ProjectId         string          `protobuf:"bytes,21,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DomainId          string          `protobuf:"bytes,22,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt         string          `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CloudServiceStatInfo) Reset() {
	*x = CloudServiceStatInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudServiceStatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatInfo) ProtoMessage() {}

func (x *CloudServiceStatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatInfo.ProtoReflect.Descriptor instead.
func (*CloudServiceStatInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{1}
}

func (x *CloudServiceStatInfo) GetQuerySetId() string {
	if x != nil {
		return x.QuerySetId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CloudServiceStatInfo) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CloudServiceStatInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CloudServiceStatInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCloudServiceGroup() string {
	if x != nil {
		return x.CloudServiceGroup
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCloudServiceType() string {
	if x != nil {
		return x.CloudServiceType
	}
	return ""
}

func (x *CloudServiceStatInfo) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *CloudServiceStatInfo) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CloudServiceStatInfo) GetAdditionalInfo() *_struct.Struct {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

func (x *CloudServiceStatInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CloudServiceStatInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

//	{
//	   "results": [
//	       {
//	           "query_set_id": "query-set-abcd1234",
//	           "key": "Disk Size",
//	           "value": "1040",
//	           "unit": "GB",
//	           "provider": "aws",
//	           "cloud_service_group": "EC2",
//	           "cloud_service_type": "Instance",
//	           "region_code": "us-east-1",
//	           "account": "aws-account-id",
//	           "additional_info": {
//	               "instance_type": "t2.micro"
//	           },
//	           "project_id": "project-abcd1234",
//	           "domain_id": "domain-58010aa2e451",
//	           "created_at": "2022-06-22T01:38:16.301Z"
//	       },
//	       {...}
//	   ],
//	   "total_count": 2
//	}
type CloudServiceStatsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*CloudServiceStatInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32                   `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *CloudServiceStatsInfo) Reset() {
	*x = CloudServiceStatsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudServiceStatsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsInfo) ProtoMessage() {}

func (x *CloudServiceStatsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsInfo.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{2}
}

func (x *CloudServiceStatsInfo) GetResults() []*CloudServiceStatInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CloudServiceStatsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type CloudServiceStatsAnalyzeQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string                     `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CloudServiceStatsAnalyzeQuery) Reset() {
	*x = CloudServiceStatsAnalyzeQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudServiceStatsAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsAnalyzeQuery) ProtoMessage() {}

func (x *CloudServiceStatsAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{3}
}

func (x *CloudServiceStatsAnalyzeQuery) GetQuery() *v1.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsAnalyzeQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type CloudServiceStatsStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	DomainId string              `protobuf:"bytes,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CloudServiceStatsStatQuery) Reset() {
	*x = CloudServiceStatsStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudServiceStatsStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudServiceStatsStatQuery) ProtoMessage() {}

func (x *CloudServiceStatsStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudServiceStatsStatQuery.ProtoReflect.Descriptor instead.
func (*CloudServiceStatsStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP(), []int{4}
}

func (x *CloudServiceStatsStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CloudServiceStatsStatQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_inventory_v1_cloud_service_stats_proto protoreflect.FileDescriptor

var file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc = []byte{
	0x0a, 0x33, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x02, 0x0a, 0x16, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0c,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0xc6, 0x03, 0x0a, 0x14, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x40, 0x0a,
	0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x80, 0x01, 0x0a, 0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x1a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0xd5, 0x03, 0x0a, 0x11,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x9e, 0x01, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x30, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x92, 0x01, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x38,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74,
	0x12, 0x35, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescOnce sync.Once
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData = file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc
)

func file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescGZIP() []byte {
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescOnce.Do(func() {
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData)
	})
	return file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDescData
}

var file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes = []interface{}{
	(*CloudServiceStatsQuery)(nil),        // 0: spaceone.api.inventory.v1.CloudServiceStatsQuery
	(*CloudServiceStatInfo)(nil),          // 1: spaceone.api.inventory.v1.CloudServiceStatInfo
	(*CloudServiceStatsInfo)(nil),         // 2: spaceone.api.inventory.v1.CloudServiceStatsInfo
	(*CloudServiceStatsAnalyzeQuery)(nil), // 3: spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery
	(*CloudServiceStatsStatQuery)(nil),    // 4: spaceone.api.inventory.v1.CloudServiceStatsStatQuery
	(*v1.Query)(nil),                      // 5: spaceone.api.core.v1.Query
	(*_struct.Struct)(nil),                // 6: google.protobuf.Struct
	(*v1.TimeSeriesAnalyzeQuery)(nil),     // 7: spaceone.api.core.v1.TimeSeriesAnalyzeQuery
	(*v1.StatisticsQuery)(nil),            // 8: spaceone.api.core.v1.StatisticsQuery
}
var file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs = []int32{
	5, // 0: spaceone.api.inventory.v1.CloudServiceStatsQuery.query:type_name -> spaceone.api.core.v1.Query
	6, // 1: spaceone.api.inventory.v1.CloudServiceStatInfo.additional_info:type_name -> google.protobuf.Struct
	1, // 2: spaceone.api.inventory.v1.CloudServiceStatsInfo.results:type_name -> spaceone.api.inventory.v1.CloudServiceStatInfo
	7, // 3: spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery.query:type_name -> spaceone.api.core.v1.TimeSeriesAnalyzeQuery
	8, // 4: spaceone.api.inventory.v1.CloudServiceStatsStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	0, // 5: spaceone.api.inventory.v1.CloudServiceStats.list:input_type -> spaceone.api.inventory.v1.CloudServiceStatsQuery
	3, // 6: spaceone.api.inventory.v1.CloudServiceStats.analyze:input_type -> spaceone.api.inventory.v1.CloudServiceStatsAnalyzeQuery
	4, // 7: spaceone.api.inventory.v1.CloudServiceStats.stat:input_type -> spaceone.api.inventory.v1.CloudServiceStatsStatQuery
	2, // 8: spaceone.api.inventory.v1.CloudServiceStats.list:output_type -> spaceone.api.inventory.v1.CloudServiceStatsInfo
	6, // 9: spaceone.api.inventory.v1.CloudServiceStats.analyze:output_type -> google.protobuf.Struct
	6, // 10: spaceone.api.inventory.v1.CloudServiceStats.stat:output_type -> google.protobuf.Struct
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_inventory_v1_cloud_service_stats_proto_init() }
func file_spaceone_api_inventory_v1_cloud_service_stats_proto_init() {
	if File_spaceone_api_inventory_v1_cloud_service_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudServiceStatsQuery); i {
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
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudServiceStatInfo); i {
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
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudServiceStatsInfo); i {
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
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudServiceStatsAnalyzeQuery); i {
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
		file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudServiceStatsStatQuery); i {
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
			RawDescriptor: file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes,
		DependencyIndexes: file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs,
		MessageInfos:      file_spaceone_api_inventory_v1_cloud_service_stats_proto_msgTypes,
	}.Build()
	File_spaceone_api_inventory_v1_cloud_service_stats_proto = out.File
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_rawDesc = nil
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_goTypes = nil
	file_spaceone_api_inventory_v1_cloud_service_stats_proto_depIdxs = nil
}
