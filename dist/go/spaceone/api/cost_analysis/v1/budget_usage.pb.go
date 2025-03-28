// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.6.1
// source: spaceone/api/cost_analysis/v1/budget_usage.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type BudgetUsageInfo_ResourceGroup int32

const (
	BudgetUsageInfo_RESOURCE_GROUP_NONE BudgetUsageInfo_ResourceGroup = 0
	BudgetUsageInfo_WORKSPACE           BudgetUsageInfo_ResourceGroup = 1
	BudgetUsageInfo_PROJECT             BudgetUsageInfo_ResourceGroup = 2
)

// Enum value maps for BudgetUsageInfo_ResourceGroup.
var (
	BudgetUsageInfo_ResourceGroup_name = map[int32]string{
		0: "RESOURCE_GROUP_NONE",
		1: "WORKSPACE",
		2: "PROJECT",
	}
	BudgetUsageInfo_ResourceGroup_value = map[string]int32{
		"RESOURCE_GROUP_NONE": 0,
		"WORKSPACE":           1,
		"PROJECT":             2,
	}
)

func (x BudgetUsageInfo_ResourceGroup) Enum() *BudgetUsageInfo_ResourceGroup {
	p := new(BudgetUsageInfo_ResourceGroup)
	*p = x
	return p
}

func (x BudgetUsageInfo_ResourceGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BudgetUsageInfo_ResourceGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_enumTypes[0].Descriptor()
}

func (BudgetUsageInfo_ResourceGroup) Type() protoreflect.EnumType {
	return &file_spaceone_api_cost_analysis_v1_budget_usage_proto_enumTypes[0]
}

func (x BudgetUsageInfo_ResourceGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BudgetUsageInfo_ResourceGroup.Descriptor instead.
func (BudgetUsageInfo_ResourceGroup) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{1, 0}
}

//	{
//	   "query": {}
//	}
type BudgetUsageQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// +optional
	WorkspaceId string `protobuf:"bytes,21,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// +optional
	ProjectId string `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// +optional
	ServiceAccountId string `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// +optional
	BudgetId string `protobuf:"bytes,25,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	// +optional
	DataSourceId  string `protobuf:"bytes,26,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetUsageQuery) Reset() {
	*x = BudgetUsageQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetUsageQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageQuery) ProtoMessage() {}

func (x *BudgetUsageQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{0}
}

func (x *BudgetUsageQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BudgetUsageQuery) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BudgetUsageQuery) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *BudgetUsageQuery) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BudgetUsageQuery) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *BudgetUsageQuery) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

//	{
//	   "budget_id": "budget-abb377eb9e8b",
//	   "name": "Cloudforet-Budget3",
//	   "date": "2025-01",
//	   "cost": 7671.164,
//	   "limit": 10000.0,
//	   "currency": "USD",
//	   "data_source_id": "data-source-1b2b3b4b5b6b",
//	   "service_account_id": "sa-1b2b3b4b5b6b",
//	   "project_id": "project-1b2b3b4b5b6b",
//	   "workspace_id": "workspace-1b2b3b4b5b6b",
//	   "domain_id": "domain-58010aa2e451",
//	   "updated_at": "2025-01-19T04:26:08.099Z"
//	}
type BudgetUsageInfo struct {
	state            protoimpl.MessageState        `protogen:"open.v1"`
	BudgetId         string                        `protobuf:"bytes,1,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	Name             string                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Date             string                        `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Cost             float32                       `protobuf:"fixed32,5,opt,name=cost,proto3" json:"cost,omitempty"`
	Limit            float32                       `protobuf:"fixed32,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Currency         string                        `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	ResourceGroup    BudgetUsageInfo_ResourceGroup `protobuf:"varint,9,opt,name=resource_group,json=resourceGroup,proto3,enum=spaceone.api.cost_analysis.v1.BudgetUsageInfo_ResourceGroup" json:"resource_group,omitempty"`
	DomainId         string                        `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	WorkspaceId      string                        `protobuf:"bytes,22,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ProjectId        string                        `protobuf:"bytes,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceAccountId string                        `protobuf:"bytes,24,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	DataSourceId     string                        `protobuf:"bytes,25,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	UpdatedAt        string                        `protobuf:"bytes,31,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *BudgetUsageInfo) Reset() {
	*x = BudgetUsageInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetUsageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageInfo) ProtoMessage() {}

func (x *BudgetUsageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageInfo.ProtoReflect.Descriptor instead.
func (*BudgetUsageInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{1}
}

func (x *BudgetUsageInfo) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BudgetUsageInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BudgetUsageInfo) GetCost() float32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *BudgetUsageInfo) GetLimit() float32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *BudgetUsageInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *BudgetUsageInfo) GetResourceGroup() BudgetUsageInfo_ResourceGroup {
	if x != nil {
		return x.ResourceGroup
	}
	return BudgetUsageInfo_RESOURCE_GROUP_NONE
}

func (x *BudgetUsageInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *BudgetUsageInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *BudgetUsageInfo) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *BudgetUsageInfo) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *BudgetUsageInfo) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *BudgetUsageInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

//	{
//	       "results": [
//	           {
//	               "budget_id": "budget-abb377eb9e8b",
//	               "name": "Cloudforet-Budget3",
//	               "date": "2025-01",
//	               "cost": 7671.164,
//	               "limit": 10000.0,
//	               "currency": "USD",
//	               "service_account_id" : "sa-1b2b3b4b5b6b",
//	               "workspace_id": "workspace-1b2b3b4b5b6b",
//	               "project_id": "project-1b2b3b4b5b6b",
//	               "data_source_id": "data-source-1b2b3b4b5b6b",
//	               "domain_id": "domain-58010aa2e451",
//	               "updated_at": "2022-07-19T04:26:08.099Z"
//	           },
//	           {
//	               "budget_id": "budget-abb377eb9e8b",
//	               "name": "Cloudforet-Budget3",
//	               "date": "2022-02",
//	               "cost": 5931.771,
//	               "limit": 11000.0,
//	               "currency": "USD",
//	               "service_account_id" : "sa-1b2b3b4b5b6b",
//	               "project_id": "project-1b2b3b4b5b6b",
//	               "workspace_id": "workspace-1b2b3b4b5b6b",
//	               "data_source_id": "data-source-1b2b3b4b5b6b",
//	               "domain_id": "domain-58010aa2e451",
//	               "updated_at": "2022-07-19T04:26:08.105Z"
//	           }
//	       ],
//	       "total_count": 12
//	}
type BudgetUsagesInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*BudgetUsageInfo     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetUsagesInfo) Reset() {
	*x = BudgetUsagesInfo{}
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetUsagesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsagesInfo) ProtoMessage() {}

func (x *BudgetUsagesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsagesInfo.ProtoReflect.Descriptor instead.
func (*BudgetUsagesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetUsagesInfo) GetResults() []*BudgetUsageInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *BudgetUsagesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

//	{
//	  "query": {
//	      "group_by": [
//	          {
//	              "key": "budget_id",
//	              "name": "Budget"
//	          }
//	      ],
//	      "filter": {},
//	      "fields": {
//	          "_total_usage": {
//	              "key": "cost",
//	              "operator": "sum"
//	          }
//	      }
//	  }
//	}
type BudgetUsageAnalyzeQuery struct {
	state protoimpl.MessageState     `protogen:"open.v1"`
	Query *v2.TimeSeriesAnalyzeQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	BudgetId string `protobuf:"bytes,2,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	// +optional
	DataSourceId  string `protobuf:"bytes,3,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetUsageAnalyzeQuery) Reset() {
	*x = BudgetUsageAnalyzeQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetUsageAnalyzeQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageAnalyzeQuery) ProtoMessage() {}

func (x *BudgetUsageAnalyzeQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageAnalyzeQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageAnalyzeQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{3}
}

func (x *BudgetUsageAnalyzeQuery) GetQuery() *v2.TimeSeriesAnalyzeQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageAnalyzeQuery) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageAnalyzeQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

type BudgetUsageStatQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Query *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	BudgetId string `protobuf:"bytes,2,opt,name=budget_id,json=budgetId,proto3" json:"budget_id,omitempty"`
	// +optional
	DataSourceId  string `protobuf:"bytes,3,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BudgetUsageStatQuery) Reset() {
	*x = BudgetUsageStatQuery{}
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BudgetUsageStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUsageStatQuery) ProtoMessage() {}

func (x *BudgetUsageStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUsageStatQuery.ProtoReflect.Descriptor instead.
func (*BudgetUsageStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP(), []int{4}
}

func (x *BudgetUsageStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *BudgetUsageStatQuery) GetBudgetId() string {
	if x != nil {
		return x.BudgetId
	}
	return ""
}

func (x *BudgetUsageStatQuery) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

var File_spaceone_api_cost_analysis_v1_budget_usage_proto protoreflect.FileDescriptor

var file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x10, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x99, 0x04, 0x0a, 0x0f, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x63, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x44, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x57, 0x4f, 0x52, 0x4b, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x02, 0x22, 0x7d,
	0x0a, 0x10, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa0, 0x01,
	0x0a, 0x17, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x96, 0x01, 0x0a, 0x14, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x32, 0xbf, 0x03, 0x0a, 0x0b, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22,
	0x23, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x8d, 0x01, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65,
	0x12, 0x36, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x61, 0x6c,
	0x79, 0x7a, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x63,
	0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x7a, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x33, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x2d, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescOnce sync.Once
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData []byte
)

func file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescGZIP() []byte {
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescOnce.Do(func() {
		file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc)))
	})
	return file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDescData
}

var file_spaceone_api_cost_analysis_v1_budget_usage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes = []any{
	(BudgetUsageInfo_ResourceGroup)(0), // 0: spaceone.api.cost_analysis.v1.BudgetUsageInfo.ResourceGroup
	(*BudgetUsageQuery)(nil),           // 1: spaceone.api.cost_analysis.v1.BudgetUsageQuery
	(*BudgetUsageInfo)(nil),            // 2: spaceone.api.cost_analysis.v1.BudgetUsageInfo
	(*BudgetUsagesInfo)(nil),           // 3: spaceone.api.cost_analysis.v1.BudgetUsagesInfo
	(*BudgetUsageAnalyzeQuery)(nil),    // 4: spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery
	(*BudgetUsageStatQuery)(nil),       // 5: spaceone.api.cost_analysis.v1.BudgetUsageStatQuery
	(*v2.Query)(nil),                   // 6: spaceone.api.core.v2.Query
	(*v2.TimeSeriesAnalyzeQuery)(nil),  // 7: spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	(*v2.StatisticsQuery)(nil),         // 8: spaceone.api.core.v2.StatisticsQuery
	(*_struct.Struct)(nil),             // 9: google.protobuf.Struct
}
var file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs = []int32{
	6, // 0: spaceone.api.cost_analysis.v1.BudgetUsageQuery.query:type_name -> spaceone.api.core.v2.Query
	0, // 1: spaceone.api.cost_analysis.v1.BudgetUsageInfo.resource_group:type_name -> spaceone.api.cost_analysis.v1.BudgetUsageInfo.ResourceGroup
	2, // 2: spaceone.api.cost_analysis.v1.BudgetUsagesInfo.results:type_name -> spaceone.api.cost_analysis.v1.BudgetUsageInfo
	7, // 3: spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery.query:type_name -> spaceone.api.core.v2.TimeSeriesAnalyzeQuery
	8, // 4: spaceone.api.cost_analysis.v1.BudgetUsageStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	1, // 5: spaceone.api.cost_analysis.v1.BudgetUsage.list:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageQuery
	4, // 6: spaceone.api.cost_analysis.v1.BudgetUsage.analyze:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageAnalyzeQuery
	5, // 7: spaceone.api.cost_analysis.v1.BudgetUsage.stat:input_type -> spaceone.api.cost_analysis.v1.BudgetUsageStatQuery
	3, // 8: spaceone.api.cost_analysis.v1.BudgetUsage.list:output_type -> spaceone.api.cost_analysis.v1.BudgetUsagesInfo
	9, // 9: spaceone.api.cost_analysis.v1.BudgetUsage.analyze:output_type -> google.protobuf.Struct
	9, // 10: spaceone.api.cost_analysis.v1.BudgetUsage.stat:output_type -> google.protobuf.Struct
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_cost_analysis_v1_budget_usage_proto_init() }
func file_spaceone_api_cost_analysis_v1_budget_usage_proto_init() {
	if File_spaceone_api_cost_analysis_v1_budget_usage_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc), len(file_spaceone_api_cost_analysis_v1_budget_usage_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes,
		DependencyIndexes: file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs,
		EnumInfos:         file_spaceone_api_cost_analysis_v1_budget_usage_proto_enumTypes,
		MessageInfos:      file_spaceone_api_cost_analysis_v1_budget_usage_proto_msgTypes,
	}.Build()
	File_spaceone_api_cost_analysis_v1_budget_usage_proto = out.File
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_goTypes = nil
	file_spaceone_api_cost_analysis_v1_budget_usage_proto_depIdxs = nil
}
