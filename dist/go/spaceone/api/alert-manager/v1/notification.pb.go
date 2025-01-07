// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v3.6.1
// source: spaceone/api/alert_manager/v1/notification.proto

package v1

import (
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

type NotificationInfo struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	NotificationId    string                 `protobuf:"bytes,1,opt,name=notification_id,json=notificationId,proto3" json:"notification_id,omitempty"`
	ServiceChannels   *_struct.Struct        `protobuf:"bytes,2,opt,name=service_channels,json=serviceChannels,proto3" json:"service_channels,omitempty"`
	UserGroupChannels *_struct.Struct        `protobuf:"bytes,3,opt,name=user_group_channels,json=userGroupChannels,proto3" json:"user_group_channels,omitempty"`
	UserChannels      *_struct.Struct        `protobuf:"bytes,4,opt,name=user_channels,json=userChannels,proto3" json:"user_channels,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NotificationInfo) Reset() {
	*x = NotificationInfo{}
	mi := &file_spaceone_api_alert_manager_v1_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationInfo) ProtoMessage() {}

func (x *NotificationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_alert_manager_v1_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationInfo.ProtoReflect.Descriptor instead.
func (*NotificationInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_alert_manager_v1_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationInfo) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *NotificationInfo) GetServiceChannels() *_struct.Struct {
	if x != nil {
		return x.ServiceChannels
	}
	return nil
}

func (x *NotificationInfo) GetUserGroupChannels() *_struct.Struct {
	if x != nil {
		return x.UserGroupChannels
	}
	return nil
}

func (x *NotificationInfo) GetUserChannels() *_struct.Struct {
	if x != nil {
		return x.UserChannels
	}
	return nil
}

var File_spaceone_api_alert_manager_v1_notification_proto protoreflect.FileDescriptor

var file_spaceone_api_alert_manager_v1_notification_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1d, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x02, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x42, 0x0a,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x47, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x65,
	0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_alert_manager_v1_notification_proto_rawDescOnce sync.Once
	file_spaceone_api_alert_manager_v1_notification_proto_rawDescData = file_spaceone_api_alert_manager_v1_notification_proto_rawDesc
)

func file_spaceone_api_alert_manager_v1_notification_proto_rawDescGZIP() []byte {
	file_spaceone_api_alert_manager_v1_notification_proto_rawDescOnce.Do(func() {
		file_spaceone_api_alert_manager_v1_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_alert_manager_v1_notification_proto_rawDescData)
	})
	return file_spaceone_api_alert_manager_v1_notification_proto_rawDescData
}

var file_spaceone_api_alert_manager_v1_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spaceone_api_alert_manager_v1_notification_proto_goTypes = []any{
	(*NotificationInfo)(nil), // 0: spaceone.api.alert_manager.v1.NotificationInfo
	(*_struct.Struct)(nil),   // 1: google.protobuf.Struct
}
var file_spaceone_api_alert_manager_v1_notification_proto_depIdxs = []int32{
	1, // 0: spaceone.api.alert_manager.v1.NotificationInfo.service_channels:type_name -> google.protobuf.Struct
	1, // 1: spaceone.api.alert_manager.v1.NotificationInfo.user_group_channels:type_name -> google.protobuf.Struct
	1, // 2: spaceone.api.alert_manager.v1.NotificationInfo.user_channels:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_alert_manager_v1_notification_proto_init() }
func file_spaceone_api_alert_manager_v1_notification_proto_init() {
	if File_spaceone_api_alert_manager_v1_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_alert_manager_v1_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spaceone_api_alert_manager_v1_notification_proto_goTypes,
		DependencyIndexes: file_spaceone_api_alert_manager_v1_notification_proto_depIdxs,
		MessageInfos:      file_spaceone_api_alert_manager_v1_notification_proto_msgTypes,
	}.Build()
	File_spaceone_api_alert_manager_v1_notification_proto = out.File
	file_spaceone_api_alert_manager_v1_notification_proto_rawDesc = nil
	file_spaceone_api_alert_manager_v1_notification_proto_goTypes = nil
	file_spaceone_api_alert_manager_v1_notification_proto_depIdxs = nil
}
