# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/config/v1/domain_config.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n*spaceone/api/config/v1/domain_config.proto\x12\x16spaceone.api.config.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"t\n\x16SetDomainConfigRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\"#\n\x13\x44omainConfigRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"S\n\x17\x44omainConfigSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xa9\x01\n\x10\x44omainConfigInfo\x12\x0c\n\x04name\x18\x01 \x01(\t\x12%\n\x04\x64\x61ta\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"c\n\x11\x44omainConfigsInfo\x12\x39\n\x07results\x18\x01 \x03(\x0b\x32(.spaceone.api.config.v1.DomainConfigInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15\x44omainConfigStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xcd\x06\n\x0c\x44omainConfig\x12\x8e\x01\n\x06\x63reate\x12..spaceone.api.config.v1.SetDomainConfigRequest\x1a(.spaceone.api.config.v1.DomainConfigInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/domain-config/create:\x01*\x12\x8e\x01\n\x06update\x12..spaceone.api.config.v1.SetDomainConfigRequest\x1a(.spaceone.api.config.v1.DomainConfigInfo\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/domain-config/update:\x01*\x12\x88\x01\n\x03set\x12..spaceone.api.config.v1.SetDomainConfigRequest\x1a(.spaceone.api.config.v1.DomainConfigInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/config/v1/domain-config/set:\x01*\x12y\n\x06\x64\x65lete\x12+.spaceone.api.config.v1.DomainConfigRequest\x1a\x16.google.protobuf.Empty\"*\x82\xd3\xe4\x93\x02$\"\x1f/config/v1/domain-config/delete:\x01*\x12\x85\x01\n\x03get\x12+.spaceone.api.config.v1.DomainConfigRequest\x1a(.spaceone.api.config.v1.DomainConfigInfo\"\'\x82\xd3\xe4\x93\x02!\"\x1c/config/v1/domain-config/get:\x01*\x12\x8c\x01\n\x04list\x12/.spaceone.api.config.v1.DomainConfigSearchQuery\x1a).spaceone.api.config.v1.DomainConfigsInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/config/v1/domain-config/list:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/config/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.config.v1.domain_config_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/config/v1'
  _globals['_DOMAINCONFIG'].methods_by_name['create']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/domain-config/create:\001*'
  _globals['_DOMAINCONFIG'].methods_by_name['update']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/domain-config/update:\001*'
  _globals['_DOMAINCONFIG'].methods_by_name['set']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['set']._serialized_options = b'\202\323\344\223\002!\"\034/config/v1/domain-config/set:\001*'
  _globals['_DOMAINCONFIG'].methods_by_name['delete']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002$\"\037/config/v1/domain-config/delete:\001*'
  _globals['_DOMAINCONFIG'].methods_by_name['get']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002!\"\034/config/v1/domain-config/get:\001*'
  _globals['_DOMAINCONFIG'].methods_by_name['list']._loaded_options = None
  _globals['_DOMAINCONFIG'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\"\"\035/config/v1/domain-config/list:\001*'
  _globals['_SETDOMAINCONFIGREQUEST']._serialized_start=193
  _globals['_SETDOMAINCONFIGREQUEST']._serialized_end=309
  _globals['_DOMAINCONFIGREQUEST']._serialized_start=311
  _globals['_DOMAINCONFIGREQUEST']._serialized_end=346
  _globals['_DOMAINCONFIGSEARCHQUERY']._serialized_start=348
  _globals['_DOMAINCONFIGSEARCHQUERY']._serialized_end=431
  _globals['_DOMAINCONFIGINFO']._serialized_start=434
  _globals['_DOMAINCONFIGINFO']._serialized_end=603
  _globals['_DOMAINCONFIGSINFO']._serialized_start=605
  _globals['_DOMAINCONFIGSINFO']._serialized_end=704
  _globals['_DOMAINCONFIGSTATQUERY']._serialized_start=706
  _globals['_DOMAINCONFIGSTATQUERY']._serialized_end=783
  _globals['_DOMAINCONFIG']._serialized_start=786
  _globals['_DOMAINCONFIG']._serialized_end=1631
# @@protoc_insertion_point(module_scope)
