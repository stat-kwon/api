# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/service_channel.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n3spaceone/api/alert_manager/v1/service_channel.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"G\n\x13\x43hannelScheduleInfo\x12\x14\n\x0cis_scheduled\x18\x01 \x01(\x08\x12\r\n\x05start\x18\x02 \x01(\x05\x12\x0b\n\x03\x65nd\x18\x03 \x01(\x05\"\xb0\x05\n\x0f\x43hannelSchedule\x12Y\n\rSCHEDULE_TYPE\x18\x01 \x01(\x0e\x32\x42.spaceone.api.alert_manager.v1.ChannelSchedule.ChannelScheduleType\x12\x10\n\x08TIMEZONE\x18\x02 \x01(\t\x12?\n\x03MON\x18\x03 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03TUE\x18\x04 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03WED\x18\x05 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03THU\x18\x06 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03\x46RI\x18\x07 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03SAT\x18\x08 \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\x12?\n\x03SUN\x18\t \x01(\x0b\x32\x32.spaceone.api.alert_manager.v1.ChannelScheduleInfo\"i\n\x13\x43hannelScheduleType\x12\x1e\n\x1a\x43HANNEL_SCHEDULE_TYPE_NONE\x10\x00\x12\x0b\n\x07\x41LL_DAY\x10\x01\x12\x0c\n\x08WEEK_DAY\x10\x02\x12\x0b\n\x07WEEKEND\x10\x03\x12\n\n\x06\x43USTOM\x10\x04\"\xe4\x01\n\x1bServiceChannelCreateRequest\x12\x13\n\x0bprotocol_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12@\n\x08schedule\x18\x03 \x01(\x0b\x32..spaceone.api.alert_manager.v1.ChannelSchedule\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nservice_id\x18\x15 \x01(\t\"\xdd\x01\n)ServiceChannelCreateForwardChannelRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12@\n\x08schedule\x18\x02 \x01(\x0b\x32..spaceone.api.alert_manager.v1.ChannelSchedule\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nservice_id\x18\x15 \x01(\t\"\xcf\x01\n\x1bServiceChannelUpdateRequest\x12\x12\n\nchannel_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12@\n\x08schedule\x18\x03 \x01(\x0b\x32..spaceone.api.alert_manager.v1.ChannelSchedule\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\"+\n\x15ServiceChannelRequest\x12\x12\n\nchannel_id\x18\x01 \x01(\t\"\xfd\x03\n\x19ServiceChannelSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x12\n\nchannel_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12T\n\x05state\x18\x04 \x01(\x0e\x32\x45.spaceone.api.alert_manager.v1.ServiceChannelSearchQuery.ChannelState\x12\x61\n\x0c\x63hannel_type\x18\x05 \x01(\x0e\x32K.spaceone.api.alert_manager.v1.ServiceChannelSearchQuery.ServiceChannelType\x12\x14\n\x0cworkspace_id\x18\x15 \x01(\t\x12\x12\n\nservice_id\x18\x16 \x01(\t\x12\x13\n\x0bprotocol_id\x18\x17 \x01(\t\x12\x11\n\tplugin_id\x18\x18 \x01(\t\"D\n\x12ServiceChannelType\x12\x15\n\x11\x43HANNEL_TYPE_NONE\x10\x00\x12\n\n\x06\x44IRECT\x10\x01\x12\x0b\n\x07\x46ORWARD\x10\x02\"A\n\x0c\x43hannelState\x12\x16\n\x12\x43HANNEL_STATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"O\n\x17ServiceChannelStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\x86\x05\n\x12ServiceChannelInfo\x12\x12\n\nchannel_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12M\n\x05state\x18\x03 \x01(\x0e\x32>.spaceone.api.alert_manager.v1.ServiceChannelInfo.ChannelState\x12Z\n\x0c\x63hannel_type\x18\x04 \x01(\x0e\x32\x44.spaceone.api.alert_manager.v1.ServiceChannelInfo.ServiceChannelType\x12@\n\x08schedule\x18\x05 \x01(\x0b\x32..spaceone.api.alert_manager.v1.ChannelSchedule\x12%\n\x04\x64\x61ta\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x0c \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nservice_id\x18\x17 \x01(\t\x12\x13\n\x0bprotocol_id\x18\x18 \x01(\t\x12\x11\n\tsecret_id\x18\x19 \x01(\t\x12\x11\n\tplugin_id\x18\x1a \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"D\n\x12ServiceChannelType\x12\x15\n\x11\x43HANNEL_TYPE_NONE\x10\x00\x12\n\n\x06\x44IRECT\x10\x01\x12\x0b\n\x07\x46ORWARD\x10\x02\"A\n\x0c\x43hannelState\x12\x16\n\x12\x43HANNEL_STATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"n\n\x13ServiceChannelsInfo\x12\x42\n\x07results\x18\x01 \x03(\x0b\x32\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x87\x0c\n\x0eServiceChannel\x12\xac\x01\n\x06\x63reate\x12:.spaceone.api.alert_manager.v1.ServiceChannelCreateRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/service-channel/create:\x01*\x12\xda\x01\n\x16\x63reate_forward_channel\x12H.spaceone.api.alert_manager.v1.ServiceChannelCreateForwardChannelRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"C\x82\xd3\xe4\x93\x02=\"8/alert-manager/v1/service-channel/create-forward-channel:\x01*\x12\xac\x01\n\x06update\x12:.spaceone.api.alert_manager.v1.ServiceChannelUpdateRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/service-channel/update:\x01*\x12\xa6\x01\n\x06\x65nable\x12\x34.spaceone.api.alert_manager.v1.ServiceChannelRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/service-channel/enable:\x01*\x12\xa8\x01\n\x07\x64isable\x12\x34.spaceone.api.alert_manager.v1.ServiceChannelRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"4\x82\xd3\xe4\x93\x02.\")/alert-manager/v1/service-channel/disable:\x01*\x12\x8b\x01\n\x06\x64\x65lete\x12\x34.spaceone.api.alert_manager.v1.ServiceChannelRequest\x1a\x16.google.protobuf.Empty\"3\x82\xd3\xe4\x93\x02-\"(/alert-manager/v1/service-channel/delete:\x01*\x12\xa0\x01\n\x03get\x12\x34.spaceone.api.alert_manager.v1.ServiceChannelRequest\x1a\x31.spaceone.api.alert_manager.v1.ServiceChannelInfo\"0\x82\xd3\xe4\x93\x02*\"%/alert-manager/v1/service-channel/get:\x01*\x12\xa7\x01\n\x04list\x12\x38.spaceone.api.alert_manager.v1.ServiceChannelSearchQuery\x1a\x32.spaceone.api.alert_manager.v1.ServiceChannelsInfo\"1\x82\xd3\xe4\x93\x02+\"&/alert-manager/v1/service-channel/list:\x01*\x12\x8a\x01\n\x04stat\x12\x36.spaceone.api.alert_manager.v1.ServiceChannelStatQuery\x1a\x17.google.protobuf.Struct\"1\x82\xd3\xe4\x93\x02+\"&/alert-manager/v1/service-channel/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.service_channel_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_SERVICECHANNEL'].methods_by_name['create']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/service-channel/create:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['create_forward_channel']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['create_forward_channel']._serialized_options = b'\202\323\344\223\002=\"8/alert-manager/v1/service-channel/create-forward-channel:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['update']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/service-channel/update:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['enable']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['enable']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/service-channel/enable:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['disable']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['disable']._serialized_options = b'\202\323\344\223\002.\")/alert-manager/v1/service-channel/disable:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['delete']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002-\"(/alert-manager/v1/service-channel/delete:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['get']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002*\"%/alert-manager/v1/service-channel/get:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['list']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002+\"&/alert-manager/v1/service-channel/list:\001*'
  _globals['_SERVICECHANNEL'].methods_by_name['stat']._loaded_options = None
  _globals['_SERVICECHANNEL'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002+\"&/alert-manager/v1/service-channel/stat:\001*'
  _globals['_CHANNELSCHEDULEINFO']._serialized_start=209
  _globals['_CHANNELSCHEDULEINFO']._serialized_end=280
  _globals['_CHANNELSCHEDULE']._serialized_start=283
  _globals['_CHANNELSCHEDULE']._serialized_end=971
  _globals['_CHANNELSCHEDULE_CHANNELSCHEDULETYPE']._serialized_start=866
  _globals['_CHANNELSCHEDULE_CHANNELSCHEDULETYPE']._serialized_end=971
  _globals['_SERVICECHANNELCREATEREQUEST']._serialized_start=974
  _globals['_SERVICECHANNELCREATEREQUEST']._serialized_end=1202
  _globals['_SERVICECHANNELCREATEFORWARDCHANNELREQUEST']._serialized_start=1205
  _globals['_SERVICECHANNELCREATEFORWARDCHANNELREQUEST']._serialized_end=1426
  _globals['_SERVICECHANNELUPDATEREQUEST']._serialized_start=1429
  _globals['_SERVICECHANNELUPDATEREQUEST']._serialized_end=1636
  _globals['_SERVICECHANNELREQUEST']._serialized_start=1638
  _globals['_SERVICECHANNELREQUEST']._serialized_end=1681
  _globals['_SERVICECHANNELSEARCHQUERY']._serialized_start=1684
  _globals['_SERVICECHANNELSEARCHQUERY']._serialized_end=2193
  _globals['_SERVICECHANNELSEARCHQUERY_SERVICECHANNELTYPE']._serialized_start=2058
  _globals['_SERVICECHANNELSEARCHQUERY_SERVICECHANNELTYPE']._serialized_end=2126
  _globals['_SERVICECHANNELSEARCHQUERY_CHANNELSTATE']._serialized_start=2128
  _globals['_SERVICECHANNELSEARCHQUERY_CHANNELSTATE']._serialized_end=2193
  _globals['_SERVICECHANNELSTATQUERY']._serialized_start=2195
  _globals['_SERVICECHANNELSTATQUERY']._serialized_end=2274
  _globals['_SERVICECHANNELINFO']._serialized_start=2277
  _globals['_SERVICECHANNELINFO']._serialized_end=2923
  _globals['_SERVICECHANNELINFO_SERVICECHANNELTYPE']._serialized_start=2058
  _globals['_SERVICECHANNELINFO_SERVICECHANNELTYPE']._serialized_end=2126
  _globals['_SERVICECHANNELINFO_CHANNELSTATE']._serialized_start=2128
  _globals['_SERVICECHANNELINFO_CHANNELSTATE']._serialized_end=2193
  _globals['_SERVICECHANNELSINFO']._serialized_start=2925
  _globals['_SERVICECHANNELSINFO']._serialized_end=3035
  _globals['_SERVICECHANNEL']._serialized_start=3038
  _globals['_SERVICECHANNEL']._serialized_end=4581
# @@protoc_insertion_point(module_scope)
