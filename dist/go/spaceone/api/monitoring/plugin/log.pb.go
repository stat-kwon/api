// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: spaceone/api/monitoring/plugin/log.proto

package plugin

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

type Sort struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Desc          bool                   `protobuf:"varint,2,opt,name=desc,proto3" json:"desc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sort) Reset() {
	*x = Sort{}
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Sort) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type LogRequest struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Options    *_struct.Struct        `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	SecretData *_struct.Struct        `protobuf:"bytes,2,opt,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty"`
	// +optional
	Schema string          `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Query  *_struct.Struct `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Keyword string `protobuf:"bytes,5,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Start   string `protobuf:"bytes,10,opt,name=start,proto3" json:"start,omitempty"`
	End     string `protobuf:"bytes,11,opt,name=end,proto3" json:"end,omitempty"`
	// +optional
	Sort *Sort `protobuf:"bytes,12,opt,name=sort,proto3" json:"sort,omitempty"`
	// +optional
	Limit         int32 `protobuf:"varint,13,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogRequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *LogRequest) GetSecretData() *_struct.Struct {
	if x != nil {
		return x.SecretData
	}
	return nil
}

func (x *LogRequest) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *LogRequest) GetQuery() *_struct.Struct {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *LogRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *LogRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *LogRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

func (x *LogRequest) GetSort() *Sort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *LogRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type LogsDataInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*_struct.Struct      `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsDataInfo) Reset() {
	*x = LogsDataInfo{}
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsDataInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsDataInfo) ProtoMessage() {}

func (x *LogsDataInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_monitoring_plugin_log_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsDataInfo.ProtoReflect.Descriptor instead.
func (*LogsDataInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogsDataInfo) GetResults() []*_struct.Struct {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_spaceone_api_monitoring_plugin_log_proto protoreflect.FileDescriptor

const file_spaceone_api_monitoring_plugin_log_proto_rawDesc = "" +
	"\n" +
	"(spaceone/api/monitoring/plugin/log.proto\x12\x1espaceone.api.monitoring.plugin\x1a\x1cgoogle/protobuf/struct.proto\",\n" +
	"\x04Sort\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x12\n" +
	"\x04desc\x18\x02 \x01(\bR\x04desc\"\xd2\x02\n" +
	"\n" +
	"LogRequest\x121\n" +
	"\aoptions\x18\x01 \x01(\v2\x17.google.protobuf.StructR\aoptions\x128\n" +
	"\vsecret_data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\n" +
	"secretData\x12\x16\n" +
	"\x06schema\x18\x03 \x01(\tR\x06schema\x12-\n" +
	"\x05query\x18\x04 \x01(\v2\x17.google.protobuf.StructR\x05query\x12\x18\n" +
	"\akeyword\x18\x05 \x01(\tR\akeyword\x12\x14\n" +
	"\x05start\x18\n" +
	" \x01(\tR\x05start\x12\x10\n" +
	"\x03end\x18\v \x01(\tR\x03end\x128\n" +
	"\x04sort\x18\f \x01(\v2$.spaceone.api.monitoring.plugin.SortR\x04sort\x12\x14\n" +
	"\x05limit\x18\r \x01(\x05R\x05limit\"A\n" +
	"\fLogsDataInfo\x121\n" +
	"\aresults\x18\x01 \x03(\v2\x17.google.protobuf.StructR\aresults2k\n" +
	"\x03Log\x12d\n" +
	"\x04list\x12*.spaceone.api.monitoring.plugin.LogRequest\x1a,.spaceone.api.monitoring.plugin.LogsDataInfo\"\x000\x01BEZCgithub.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/pluginb\x06proto3"

var (
	file_spaceone_api_monitoring_plugin_log_proto_rawDescOnce sync.Once
	file_spaceone_api_monitoring_plugin_log_proto_rawDescData []byte
)

func file_spaceone_api_monitoring_plugin_log_proto_rawDescGZIP() []byte {
	file_spaceone_api_monitoring_plugin_log_proto_rawDescOnce.Do(func() {
		file_spaceone_api_monitoring_plugin_log_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spaceone_api_monitoring_plugin_log_proto_rawDesc), len(file_spaceone_api_monitoring_plugin_log_proto_rawDesc)))
	})
	return file_spaceone_api_monitoring_plugin_log_proto_rawDescData
}

var file_spaceone_api_monitoring_plugin_log_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_monitoring_plugin_log_proto_goTypes = []any{
	(*Sort)(nil),           // 0: spaceone.api.monitoring.plugin.Sort
	(*LogRequest)(nil),     // 1: spaceone.api.monitoring.plugin.LogRequest
	(*LogsDataInfo)(nil),   // 2: spaceone.api.monitoring.plugin.LogsDataInfo
	(*_struct.Struct)(nil), // 3: google.protobuf.Struct
}
var file_spaceone_api_monitoring_plugin_log_proto_depIdxs = []int32{
	3, // 0: spaceone.api.monitoring.plugin.LogRequest.options:type_name -> google.protobuf.Struct
	3, // 1: spaceone.api.monitoring.plugin.LogRequest.secret_data:type_name -> google.protobuf.Struct
	3, // 2: spaceone.api.monitoring.plugin.LogRequest.query:type_name -> google.protobuf.Struct
	0, // 3: spaceone.api.monitoring.plugin.LogRequest.sort:type_name -> spaceone.api.monitoring.plugin.Sort
	3, // 4: spaceone.api.monitoring.plugin.LogsDataInfo.results:type_name -> google.protobuf.Struct
	1, // 5: spaceone.api.monitoring.plugin.Log.list:input_type -> spaceone.api.monitoring.plugin.LogRequest
	2, // 6: spaceone.api.monitoring.plugin.Log.list:output_type -> spaceone.api.monitoring.plugin.LogsDataInfo
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_spaceone_api_monitoring_plugin_log_proto_init() }
func file_spaceone_api_monitoring_plugin_log_proto_init() {
	if File_spaceone_api_monitoring_plugin_log_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spaceone_api_monitoring_plugin_log_proto_rawDesc), len(file_spaceone_api_monitoring_plugin_log_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_monitoring_plugin_log_proto_goTypes,
		DependencyIndexes: file_spaceone_api_monitoring_plugin_log_proto_depIdxs,
		MessageInfos:      file_spaceone_api_monitoring_plugin_log_proto_msgTypes,
	}.Build()
	File_spaceone_api_monitoring_plugin_log_proto = out.File
	file_spaceone_api_monitoring_plugin_log_proto_goTypes = nil
	file_spaceone_api_monitoring_plugin_log_proto_depIdxs = nil
}
