# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/opsflow/v1/task.proto
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
from spaceone.api.opsflow.v1 import comment_pb2 as spaceone_dot_api_dot_opsflow_dot_v1_dot_comment__pb2
from spaceone.api.file_manager.v1 import file_pb2 as spaceone_dot_api_dot_file__manager_dot_v1_dot_file__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\"spaceone/api/opsflow/v1/task.proto\x12\x17spaceone.api.opsflow.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a%spaceone/api/opsflow/v1/comment.proto\x1a\'spaceone/api/file_manager/v1/file.proto\"\xa9\x02\n\x11TaskCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tstatus_id\x18\x02 \x01(\t\x12\x37\n\x08priority\x18\x03 \x01(\x0e\x32%.spaceone.api.opsflow.v1.TaskPriority\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\r\n\x05\x66iles\x18\x05 \x03(\t\x12\x33\n\x08mentions\x18\x06 \x01(\x0b\x32!.spaceone.api.opsflow.v1.Mentions\x12\x10\n\x08\x61ssignee\x18\x07 \x01(\t\x12%\n\x04\x64\x61ta\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nproject_id\x18\x15 \x01(\t\x12\x14\n\x0ctask_type_id\x18\x16 \x01(\t\"\x91\x02\n\x11TaskUpdateRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\x08priority\x18\x03 \x01(\x0e\x32%.spaceone.api.opsflow.v1.TaskPriority\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\r\n\x05\x66iles\x18\x05 \x03(\t\x12\x33\n\x08mentions\x18\x06 \x01(\x0b\x32!.spaceone.api.opsflow.v1.Mentions\x12\x10\n\x08\x61ssignee\x18\x07 \x01(\t\x12%\n\x04\x64\x61ta\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x12\n\nproject_id\x18\x15 \x01(\t\"O\n\x17TaskChangeStatusRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x11\n\tstatus_id\x18\x02 \x01(\t\x12\x10\n\x08\x61ssignee\x18\x03 \x01(\t\"\x1e\n\x0bTaskRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\t\"\x92\x02\n\x0fTaskSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07task_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tstatus_id\x18\x04 \x01(\t\x12\x13\n\x0bstatus_type\x18\x05 \x01(\t\x12\x37\n\x08priority\x18\x06 \x01(\x0e\x32%.spaceone.api.opsflow.v1.TaskPriority\x12\x14\n\x0cworkspace_id\x18\x0b \x01(\t\x12\x12\n\nproject_id\x18\x0c \x01(\t\x12\x13\n\x0b\x63\x61tegory_id\x18\r \x01(\t\x12\x14\n\x0ctask_type_id\x18\x0e \x01(\t\"\xc6\x03\n\x08TaskInfo\x12\x0f\n\x07task_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06status\x18\x03 \x01(\t\x12\x13\n\x0bstatus_type\x18\x04 \x01(\t\x12\x37\n\x08priority\x18\x05 \x01(\x0e\x32%.spaceone.api.opsflow.v1.TaskPriority\x12\x13\n\x0b\x64\x65scription\x18\x06 \x01(\t\x12%\n\x04\x64\x61ta\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x35\n\x05\x66iles\x18\x08 \x03(\x0b\x32&.spaceone.api.file_manager.v1.FileInfo\x12\x10\n\x08\x61ssignee\x18\t \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nproject_id\x18\x17 \x01(\t\x12\x13\n\x0b\x63\x61tegory_id\x18\x18 \x01(\t\x12\x14\n\x0ctask_type_id\x18\x19 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nstarted_at\x18  \x01(\t\x12\x12\n\nupdated_at\x18! \x01(\t\x12\x14\n\x0c\x63ompleted_at\x18\" \x01(\t\"T\n\tTasksInfo\x12\x32\n\x07results\x18\x01 \x03(\x0b\x32!.spaceone.api.opsflow.v1.TaskInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"E\n\rTaskStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery*E\n\x0cTaskPriority\x12\x16\n\x12TASK_PRIORITY_NONE\x10\x00\x12\x08\n\x04HIGH\x10\x01\x12\n\n\x06MEDIUM\x10\x02\x12\x07\n\x03LOW\x10\x03\x32\xd2\x06\n\x04Task\x12{\n\x06\x63reate\x12*.spaceone.api.opsflow.v1.TaskCreateRequest\x1a!.spaceone.api.opsflow.v1.TaskInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/opsflow/v1/task/create:\x01*\x12{\n\x06update\x12*.spaceone.api.opsflow.v1.TaskUpdateRequest\x1a!.spaceone.api.opsflow.v1.TaskInfo\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/opsflow/v1/task/update:\x01*\x12\x8f\x01\n\rchange_status\x12\x30.spaceone.api.opsflow.v1.TaskChangeStatusRequest\x1a!.spaceone.api.opsflow.v1.TaskInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/opsflow/v1/task/change_status:\x01*\x12j\n\x06\x64\x65lete\x12$.spaceone.api.opsflow.v1.TaskRequest\x1a\x16.google.protobuf.Empty\"\"\x82\xd3\xe4\x93\x02\x1c\"\x17/opsflow/v1/task/delete:\x01*\x12o\n\x03get\x12$.spaceone.api.opsflow.v1.TaskRequest\x1a!.spaceone.api.opsflow.v1.TaskInfo\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/opsflow/v1/task/get:\x01*\x12v\n\x04list\x12(.spaceone.api.opsflow.v1.TaskSearchQuery\x1a\".spaceone.api.opsflow.v1.TasksInfo\" \x82\xd3\xe4\x93\x02\x1a\"\x15/opsflow/v1/task/list:\x01*\x12i\n\x04stat\x12&.spaceone.api.opsflow.v1.TaskStatQuery\x1a\x17.google.protobuf.Struct\" \x82\xd3\xe4\x93\x02\x1a\"\x15/opsflow/v1/task/stat:\x01*B>Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.opsflow.v1.task_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z<github.com/cloudforet-io/api/dist/go/spaceone/api/opsflow/v1'
  _globals['_TASK'].methods_by_name['create']._loaded_options = None
  _globals['_TASK'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\034\"\027/opsflow/v1/task/create:\001*'
  _globals['_TASK'].methods_by_name['update']._loaded_options = None
  _globals['_TASK'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\034\"\027/opsflow/v1/task/update:\001*'
  _globals['_TASK'].methods_by_name['change_status']._loaded_options = None
  _globals['_TASK'].methods_by_name['change_status']._serialized_options = b'\202\323\344\223\002#\"\036/opsflow/v1/task/change_status:\001*'
  _globals['_TASK'].methods_by_name['delete']._loaded_options = None
  _globals['_TASK'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\034\"\027/opsflow/v1/task/delete:\001*'
  _globals['_TASK'].methods_by_name['get']._loaded_options = None
  _globals['_TASK'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\031\"\024/opsflow/v1/task/get:\001*'
  _globals['_TASK'].methods_by_name['list']._loaded_options = None
  _globals['_TASK'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\032\"\025/opsflow/v1/task/list:\001*'
  _globals['_TASK'].methods_by_name['stat']._loaded_options = None
  _globals['_TASK'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\032\"\025/opsflow/v1/task/stat:\001*'
  _globals['_TASKPRIORITY']._serialized_start=1846
  _globals['_TASKPRIORITY']._serialized_end=1915
  _globals['_TASKCREATEREQUEST']._serialized_start=267
  _globals['_TASKCREATEREQUEST']._serialized_end=564
  _globals['_TASKUPDATEREQUEST']._serialized_start=567
  _globals['_TASKUPDATEREQUEST']._serialized_end=840
  _globals['_TASKCHANGESTATUSREQUEST']._serialized_start=842
  _globals['_TASKCHANGESTATUSREQUEST']._serialized_end=921
  _globals['_TASKREQUEST']._serialized_start=923
  _globals['_TASKREQUEST']._serialized_end=953
  _globals['_TASKSEARCHQUERY']._serialized_start=956
  _globals['_TASKSEARCHQUERY']._serialized_end=1230
  _globals['_TASKINFO']._serialized_start=1233
  _globals['_TASKINFO']._serialized_end=1687
  _globals['_TASKSINFO']._serialized_start=1689
  _globals['_TASKSINFO']._serialized_end=1773
  _globals['_TASKSTATQUERY']._serialized_start=1775
  _globals['_TASKSTATQUERY']._serialized_end=1844
  _globals['_TASK']._serialized_start=1918
  _globals['_TASK']._serialized_end=2768
# @@protoc_insertion_point(module_scope)
