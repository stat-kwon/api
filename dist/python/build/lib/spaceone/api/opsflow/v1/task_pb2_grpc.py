# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.opsflow.v1 import task_pb2 as spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in spaceone/api/opsflow/v1/task_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class TaskStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/create',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskCreateRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/update',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.update_description = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/update_description',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateDescriptionRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.change_assignee = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/change_assignee',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeAssigneeRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.change_status = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/change_status',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeStatusRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/delete',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/get',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/list',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TasksInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.opsflow.v1.Task/stat',
                request_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class TaskServicer(object):
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

    def update_description(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def change_assignee(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def change_status(self, request, context):
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


def add_TaskServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskCreateRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'update_description': grpc.unary_unary_rpc_method_handler(
                    servicer.update_description,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateDescriptionRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'change_assignee': grpc.unary_unary_rpc_method_handler(
                    servicer.change_assignee,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeAssigneeRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'change_status': grpc.unary_unary_rpc_method_handler(
                    servicer.change_status,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeStatusRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TasksInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.opsflow.v1.Task', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.opsflow.v1.Task', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Task(object):
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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/create',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskCreateRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/update',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def update_description(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/update_description',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskUpdateDescriptionRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def change_assignee(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/change_assignee',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeAssigneeRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def change_status(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/change_status',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskChangeStatusRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/delete',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/get',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskRequest.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/list',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskSearchQuery.SerializeToString,
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TasksInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

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
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.opsflow.v1.Task/stat',
            spaceone_dot_api_dot_opsflow_dot_v1_dot_task__pb2.TaskStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
