# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import app_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2


class AppStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.App/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CreateAppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.identity.v2.App/update',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.UpdateAppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.generate_api_key = channel.unary_unary(
                '/spaceone.api.identity.v2.App/generate_api_key',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.GenerateAPIKeyAppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.enable = channel.unary_unary(
                '/spaceone.api.identity.v2.App/enable',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.disable = channel.unary_unary(
                '/spaceone.api.identity.v2.App/disable',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v2.App/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.App/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
                )
        self.check = channel.unary_unary(
                '/spaceone.api.identity.v2.App/check',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppCheckRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CheckAppInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.App/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.App/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class AppServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def generate_api_key(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def enable(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def disable(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def check(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AppServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CreateAppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.UpdateAppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'generate_api_key': grpc.unary_unary_rpc_method_handler(
                    servicer.generate_api_key,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.GenerateAPIKeyAppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'enable': grpc.unary_unary_rpc_method_handler(
                    servicer.enable,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'disable': grpc.unary_unary_rpc_method_handler(
                    servicer.disable,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.SerializeToString,
            ),
            'check': grpc.unary_unary_rpc_method_handler(
                    servicer.check,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppCheckRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CheckAppInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.App', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class App(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CreateAppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/update',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.UpdateAppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def generate_api_key(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/generate_api_key',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.GenerateAPIKeyAppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def enable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/enable',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def disable(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/disable',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/delete',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def check(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/check',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppCheckRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.CheckAppInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppsInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.identity.v2.App/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_app__pb2.AppStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)