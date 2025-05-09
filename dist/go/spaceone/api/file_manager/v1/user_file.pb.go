// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/file_manager/v1/user_file.proto

package v1

import (
	v2 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type UserFileReference struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceType  string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId    string                 `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFileReference) Reset() {
	*x = UserFileReference{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFileReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileReference) ProtoMessage() {}

func (x *UserFileReference) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileReference.ProtoReflect.Descriptor instead.
func (*UserFileReference) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{0}
}

func (x *UserFileReference) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *UserFileReference) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CreateUserFileRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	Reference *UserFileReference `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	UserId        string `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserFileRequest) Reset() {
	*x = CreateUserFileRequest{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserFileRequest) ProtoMessage() {}

func (x *CreateUserFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserFileRequest.ProtoReflect.Descriptor instead.
func (*CreateUserFileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserFileRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CreateUserFileRequest) GetReference() *UserFileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *CreateUserFileRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *CreateUserFileRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateUserFileRequest struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	FileId string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,2,opt,name=tags,proto3" json:"tags,omitempty"`
	// +optional
	Reference     *UserFileReference `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserFileRequest) Reset() {
	*x = UpdateUserFileRequest{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserFileRequest) ProtoMessage() {}

func (x *UpdateUserFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserFileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UpdateUserFileRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UpdateUserFileRequest) GetReference() *UserFileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

type UserFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFileRequest) Reset() {
	*x = UserFileRequest{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileRequest) ProtoMessage() {}

func (x *UserFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileRequest.ProtoReflect.Descriptor instead.
func (*UserFileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{3}
}

func (x *UserFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type UserFileSearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// +optional
	Query *v2.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	FileId string `protobuf:"bytes,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// / +optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	ResourceType string `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// +optional
	ResourceId string `protobuf:"bytes,5,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// +optional
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	// +optional
	UserId        string `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFileSearchQuery) Reset() {
	*x = UserFileSearchQuery{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFileSearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileSearchQuery) ProtoMessage() {}

func (x *UserFileSearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileSearchQuery.ProtoReflect.Descriptor instead.
func (*UserFileSearchQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{4}
}

func (x *UserFileSearchQuery) GetQuery() *v2.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UserFileSearchQuery) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UserFileSearchQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserFileSearchQuery) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *UserFileSearchQuery) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UserFileSearchQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *UserFileSearchQuery) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserFileInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileId        string                 `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DownloadUrl   string                 `protobuf:"bytes,3,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	Tags          *_struct.Struct        `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	Reference     *UserFileReference     `protobuf:"bytes,5,opt,name=reference,proto3" json:"reference,omitempty"`
	DomainId      string                 `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	UserId        string                 `protobuf:"bytes,22,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,31,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFileInfo) Reset() {
	*x = UserFileInfo{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileInfo) ProtoMessage() {}

func (x *UserFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileInfo.ProtoReflect.Descriptor instead.
func (*UserFileInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{5}
}

func (x *UserFileInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UserFileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserFileInfo) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *UserFileInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *UserFileInfo) GetReference() *UserFileReference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *UserFileInfo) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

func (x *UserFileInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserFileInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UserFilesInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*UserFileInfo        `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount    int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFilesInfo) Reset() {
	*x = UserFilesInfo{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFilesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFilesInfo) ProtoMessage() {}

func (x *UserFilesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFilesInfo.ProtoReflect.Descriptor instead.
func (*UserFilesInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{6}
}

func (x *UserFilesInfo) GetResults() []*UserFileInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *UserFilesInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type UserFileStatQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *v2.StatisticsQuery    `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserFileStatQuery) Reset() {
	*x = UserFileStatQuery{}
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFileStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFileStatQuery) ProtoMessage() {}

func (x *UserFileStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_file_manager_v1_user_file_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFileStatQuery.ProtoReflect.Descriptor instead.
func (*UserFileStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP(), []int{7}
}

func (x *UserFileStatQuery) GetQuery() *v2.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

var File_spaceone_api_file_manager_v1_user_file_proto protoreflect.FileDescriptor

const file_spaceone_api_file_manager_v1_user_file_proto_rawDesc = "" +
	"\n" +
	",spaceone/api/file_manager/v1/user_file.proto\x12\x1cspaceone.api.file_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"Y\n" +
	"\x11UserFileReference\x12#\n" +
	"\rresource_type\x18\x01 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x02 \x01(\tR\n" +
	"resourceId\"\xdd\x01\n" +
	"\x15CreateUserFileRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12+\n" +
	"\x04tags\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12M\n" +
	"\treference\x18\x03 \x01(\v2/.spaceone.api.file_manager.v1.UserFileReferenceR\treference\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x17\n" +
	"\auser_id\x18\x16 \x01(\tR\x06userId\"\xac\x01\n" +
	"\x15UpdateUserFileRequest\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12+\n" +
	"\x04tags\x18\x02 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12M\n" +
	"\treference\x18\x03 \x01(\v2/.spaceone.api.file_manager.v1.UserFileReferenceR\treference\"*\n" +
	"\x0fUserFileRequest\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\"\xf1\x01\n" +
	"\x13UserFileSearchQuery\x121\n" +
	"\x05query\x18\x01 \x01(\v2\x1b.spaceone.api.core.v2.QueryR\x05query\x12\x17\n" +
	"\afile_id\x18\x02 \x01(\tR\x06fileId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12#\n" +
	"\rresource_type\x18\x04 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x05 \x01(\tR\n" +
	"resourceId\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x17\n" +
	"\auser_id\x18\x16 \x01(\tR\x06userId\"\xaf\x02\n" +
	"\fUserFileInfo\x12\x17\n" +
	"\afile_id\x18\x01 \x01(\tR\x06fileId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12!\n" +
	"\fdownload_url\x18\x03 \x01(\tR\vdownloadUrl\x12+\n" +
	"\x04tags\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x04tags\x12M\n" +
	"\treference\x18\x05 \x01(\v2/.spaceone.api.file_manager.v1.UserFileReferenceR\treference\x12\x1b\n" +
	"\tdomain_id\x18\x15 \x01(\tR\bdomainId\x12\x17\n" +
	"\auser_id\x18\x16 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"created_at\x18\x1f \x01(\tR\tcreatedAt\"v\n" +
	"\rUserFilesInfo\x12D\n" +
	"\aresults\x18\x01 \x03(\v2*.spaceone.api.file_manager.v1.UserFileInfoR\aresults\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"P\n" +
	"\x11UserFileStatQuery\x12;\n" +
	"\x05query\x18\x01 \x01(\v2%.spaceone.api.core.v2.StatisticsQueryR\x05query2\xc4\x05\n" +
	"\bUserFile\x12\x97\x01\n" +
	"\x06update\x123.spaceone.api.file_manager.v1.UpdateUserFileRequest\x1a*.spaceone.api.file_manager.v1.UserFileInfo\",\x82\xd3\xe4\x93\x02&:\x01*\"!/file-manager/v1/user_file/update\x12}\n" +
	"\x06delete\x12-.spaceone.api.file_manager.v1.UserFileRequest\x1a\x16.google.protobuf.Empty\",\x82\xd3\xe4\x93\x02&:\x01*\"!/file-manager/v1/user_file/delete\x12\x8b\x01\n" +
	"\x03get\x12-.spaceone.api.file_manager.v1.UserFileRequest\x1a*.spaceone.api.file_manager.v1.UserFileInfo\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/file-manager/v1/user_file/get\x12\x92\x01\n" +
	"\x04list\x121.spaceone.api.file_manager.v1.UserFileSearchQuery\x1a+.spaceone.api.file_manager.v1.UserFilesInfo\"*\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/file-manager/v1/user_file/list\x12|\n" +
	"\x04stat\x12/.spaceone.api.file_manager.v1.UserFileStatQuery\x1a\x17.google.protobuf.Struct\"*\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/file-manager/v1/user_file/statBCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/file_manager/v1b\x06proto3"

var (
	file_spaceone_api_file_manager_v1_user_file_proto_rawDescOnce sync.Once
	file_spaceone_api_file_manager_v1_user_file_proto_rawDescData []byte
)

func file_spaceone_api_file_manager_v1_user_file_proto_rawDescGZIP() []byte {
	file_spaceone_api_file_manager_v1_user_file_proto_rawDescOnce.Do(func() {
		file_spaceone_api_file_manager_v1_user_file_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_file_manager_v1_user_file_proto_rawDesc), len(file_spaceone_api_file_manager_v1_user_file_proto_rawDesc)))
	})
	return file_spaceone_api_file_manager_v1_user_file_proto_rawDescData
}

var file_spaceone_api_file_manager_v1_user_file_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_file_manager_v1_user_file_proto_goTypes = []any{
	(*UserFileReference)(nil),     // 0: spaceone.api.file_manager.v1.UserFileReference
	(*CreateUserFileRequest)(nil), // 1: spaceone.api.file_manager.v1.CreateUserFileRequest
	(*UpdateUserFileRequest)(nil), // 2: spaceone.api.file_manager.v1.UpdateUserFileRequest
	(*UserFileRequest)(nil),       // 3: spaceone.api.file_manager.v1.UserFileRequest
	(*UserFileSearchQuery)(nil),   // 4: spaceone.api.file_manager.v1.UserFileSearchQuery
	(*UserFileInfo)(nil),          // 5: spaceone.api.file_manager.v1.UserFileInfo
	(*UserFilesInfo)(nil),         // 6: spaceone.api.file_manager.v1.UserFilesInfo
	(*UserFileStatQuery)(nil),     // 7: spaceone.api.file_manager.v1.UserFileStatQuery
	(*_struct.Struct)(nil),        // 8: google.protobuf.Struct
	(*v2.Query)(nil),              // 9: spaceone.api.core.v2.Query
	(*v2.StatisticsQuery)(nil),    // 10: spaceone.api.core.v2.StatisticsQuery
	(*empty.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_spaceone_api_file_manager_v1_user_file_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.file_manager.v1.CreateUserFileRequest.tags:type_name -> google.protobuf.Struct
	0,  // 1: spaceone.api.file_manager.v1.CreateUserFileRequest.reference:type_name -> spaceone.api.file_manager.v1.UserFileReference
	8,  // 2: spaceone.api.file_manager.v1.UpdateUserFileRequest.tags:type_name -> google.protobuf.Struct
	0,  // 3: spaceone.api.file_manager.v1.UpdateUserFileRequest.reference:type_name -> spaceone.api.file_manager.v1.UserFileReference
	9,  // 4: spaceone.api.file_manager.v1.UserFileSearchQuery.query:type_name -> spaceone.api.core.v2.Query
	8,  // 5: spaceone.api.file_manager.v1.UserFileInfo.tags:type_name -> google.protobuf.Struct
	0,  // 6: spaceone.api.file_manager.v1.UserFileInfo.reference:type_name -> spaceone.api.file_manager.v1.UserFileReference
	5,  // 7: spaceone.api.file_manager.v1.UserFilesInfo.results:type_name -> spaceone.api.file_manager.v1.UserFileInfo
	10, // 8: spaceone.api.file_manager.v1.UserFileStatQuery.query:type_name -> spaceone.api.core.v2.StatisticsQuery
	2,  // 9: spaceone.api.file_manager.v1.UserFile.update:input_type -> spaceone.api.file_manager.v1.UpdateUserFileRequest
	3,  // 10: spaceone.api.file_manager.v1.UserFile.delete:input_type -> spaceone.api.file_manager.v1.UserFileRequest
	3,  // 11: spaceone.api.file_manager.v1.UserFile.get:input_type -> spaceone.api.file_manager.v1.UserFileRequest
	4,  // 12: spaceone.api.file_manager.v1.UserFile.list:input_type -> spaceone.api.file_manager.v1.UserFileSearchQuery
	7,  // 13: spaceone.api.file_manager.v1.UserFile.stat:input_type -> spaceone.api.file_manager.v1.UserFileStatQuery
	5,  // 14: spaceone.api.file_manager.v1.UserFile.update:output_type -> spaceone.api.file_manager.v1.UserFileInfo
	11, // 15: spaceone.api.file_manager.v1.UserFile.delete:output_type -> google.protobuf.Empty
	5,  // 16: spaceone.api.file_manager.v1.UserFile.get:output_type -> spaceone.api.file_manager.v1.UserFileInfo
	6,  // 17: spaceone.api.file_manager.v1.UserFile.list:output_type -> spaceone.api.file_manager.v1.UserFilesInfo
	8,  // 18: spaceone.api.file_manager.v1.UserFile.stat:output_type -> google.protobuf.Struct
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_spaceone_api_file_manager_v1_user_file_proto_init() }
func file_spaceone_api_file_manager_v1_user_file_proto_init() {
	if File_spaceone_api_file_manager_v1_user_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_file_manager_v1_user_file_proto_rawDesc), len(file_spaceone_api_file_manager_v1_user_file_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_file_manager_v1_user_file_proto_goTypes,
		DependencyIndexes: file_spaceone_api_file_manager_v1_user_file_proto_depIdxs,
		MessageInfos:      file_spaceone_api_file_manager_v1_user_file_proto_msgTypes,
	}.Build()
	File_spaceone_api_file_manager_v1_user_file_proto = out.File
	file_spaceone_api_file_manager_v1_user_file_proto_goTypes = nil
	file_spaceone_api_file_manager_v1_user_file_proto_depIdxs = nil
}
