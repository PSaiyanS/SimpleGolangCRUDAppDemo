/**
 * @fileoverview gRPC-Web generated client stub for test10
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as test10_api_test10_pb from '@marketplace/test10/test10_pb';


export class Test10Client {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorPing = new grpcWeb.MethodDescriptor(
    '/test10.Test10/Ping',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.PingRequest,
    test10_api_test10_pb.PingRespond,
    (request: test10_api_test10_pb.PingRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.PingRespond.deserializeBinary
  );

  ping(
    request: test10_api_test10_pb.PingRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.PingRespond>;

  ping(
    request: test10_api_test10_pb.PingRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.PingRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.PingRespond>;

  ping(
    request: test10_api_test10_pb.PingRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.PingRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/Ping',
        request,
        metadata || {},
        this.methodDescriptorPing,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/Ping',
    request,
    metadata || {},
    this.methodDescriptorPing);
  }

  methodDescriptorUserCreate = new grpcWeb.MethodDescriptor(
    '/test10.Test10/UserCreate',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserCreateRequest,
    test10_api_test10_pb.UserCreateRespond,
    (request: test10_api_test10_pb.UserCreateRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.UserCreateRespond.deserializeBinary
  );

  userCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.UserCreateRespond>;

  userCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserCreateRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.UserCreateRespond>;

  userCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserCreateRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/UserCreate',
        request,
        metadata || {},
        this.methodDescriptorUserCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/UserCreate',
    request,
    metadata || {},
    this.methodDescriptorUserCreate);
  }

  methodDescriptorUserList = new grpcWeb.MethodDescriptor(
    '/test10.Test10/UserList',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserListRequest,
    test10_api_test10_pb.UserListRespond,
    (request: test10_api_test10_pb.UserListRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.UserListRespond.deserializeBinary
  );

  userList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.UserListRespond>;

  userList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserListRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.UserListRespond>;

  userList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserListRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/UserList',
        request,
        metadata || {},
        this.methodDescriptorUserList,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/UserList',
    request,
    metadata || {},
    this.methodDescriptorUserList);
  }

  methodDescriptorUserUpdate = new grpcWeb.MethodDescriptor(
    '/test10.Test10/UserUpdate',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserUpdateRequest,
    test10_api_test10_pb.UserUpdateRespond,
    (request: test10_api_test10_pb.UserUpdateRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.UserUpdateRespond.deserializeBinary
  );

  userUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.UserUpdateRespond>;

  userUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserUpdateRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.UserUpdateRespond>;

  userUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserUpdateRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/UserUpdate',
        request,
        metadata || {},
        this.methodDescriptorUserUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/UserUpdate',
    request,
    metadata || {},
    this.methodDescriptorUserUpdate);
  }

  methodDescriptorUserDelete = new grpcWeb.MethodDescriptor(
    '/test10.Test10/UserDelete',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserDeleteRequest,
    test10_api_test10_pb.UserDeleteRespond,
    (request: test10_api_test10_pb.UserDeleteRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.UserDeleteRespond.deserializeBinary
  );

  userDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.UserDeleteRespond>;

  userDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserDeleteRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.UserDeleteRespond>;

  userDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserDeleteRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/UserDelete',
        request,
        metadata || {},
        this.methodDescriptorUserDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/UserDelete',
    request,
    metadata || {},
    this.methodDescriptorUserDelete);
  }

  methodDescriptorUserGetById = new grpcWeb.MethodDescriptor(
    '/test10.Test10/UserGetById',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserGetByIdRequest,
    test10_api_test10_pb.UserGetByIdRespond,
    (request: test10_api_test10_pb.UserGetByIdRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.UserGetByIdRespond.deserializeBinary
  );

  userGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.UserGetByIdRespond>;

  userGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserGetByIdRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.UserGetByIdRespond>;

  userGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.UserGetByIdRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/UserGetById',
        request,
        metadata || {},
        this.methodDescriptorUserGetById,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/UserGetById',
    request,
    metadata || {},
    this.methodDescriptorUserGetById);
  }

  methodDescriptorRoomCreate = new grpcWeb.MethodDescriptor(
    '/test10.Test10/RoomCreate',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserCreateRequest,
    test10_api_test10_pb.RoomCreateRespond,
    (request: test10_api_test10_pb.UserCreateRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.RoomCreateRespond.deserializeBinary
  );

  roomCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.RoomCreateRespond>;

  roomCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomCreateRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.RoomCreateRespond>;

  roomCreate(
    request: test10_api_test10_pb.UserCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomCreateRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/RoomCreate',
        request,
        metadata || {},
        this.methodDescriptorRoomCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/RoomCreate',
    request,
    metadata || {},
    this.methodDescriptorRoomCreate);
  }

  methodDescriptorRoomList = new grpcWeb.MethodDescriptor(
    '/test10.Test10/RoomList',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserListRequest,
    test10_api_test10_pb.RoomListRespond,
    (request: test10_api_test10_pb.UserListRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.RoomListRespond.deserializeBinary
  );

  roomList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.RoomListRespond>;

  roomList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomListRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.RoomListRespond>;

  roomList(
    request: test10_api_test10_pb.UserListRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomListRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/RoomList',
        request,
        metadata || {},
        this.methodDescriptorRoomList,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/RoomList',
    request,
    metadata || {},
    this.methodDescriptorRoomList);
  }

  methodDescriptorRoomUpdate = new grpcWeb.MethodDescriptor(
    '/test10.Test10/RoomUpdate',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserUpdateRequest,
    test10_api_test10_pb.RoomUpdateRespond,
    (request: test10_api_test10_pb.UserUpdateRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.RoomUpdateRespond.deserializeBinary
  );

  roomUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.RoomUpdateRespond>;

  roomUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomUpdateRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.RoomUpdateRespond>;

  roomUpdate(
    request: test10_api_test10_pb.UserUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomUpdateRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/RoomUpdate',
        request,
        metadata || {},
        this.methodDescriptorRoomUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/RoomUpdate',
    request,
    metadata || {},
    this.methodDescriptorRoomUpdate);
  }

  methodDescriptorRoomDelete = new grpcWeb.MethodDescriptor(
    '/test10.Test10/RoomDelete',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserDeleteRequest,
    test10_api_test10_pb.RoomDeleteRespond,
    (request: test10_api_test10_pb.UserDeleteRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.RoomDeleteRespond.deserializeBinary
  );

  roomDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.RoomDeleteRespond>;

  roomDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomDeleteRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.RoomDeleteRespond>;

  roomDelete(
    request: test10_api_test10_pb.UserDeleteRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomDeleteRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/RoomDelete',
        request,
        metadata || {},
        this.methodDescriptorRoomDelete,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/RoomDelete',
    request,
    metadata || {},
    this.methodDescriptorRoomDelete);
  }

  methodDescriptorRoomGetById = new grpcWeb.MethodDescriptor(
    '/test10.Test10/RoomGetById',
    grpcWeb.MethodType.UNARY,
    test10_api_test10_pb.UserGetByIdRequest,
    test10_api_test10_pb.RoomGetByIdRespond,
    (request: test10_api_test10_pb.UserGetByIdRequest) => {
      return request.serializeBinary();
    },
    test10_api_test10_pb.RoomGetByIdRespond.deserializeBinary
  );

  roomGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null): Promise<test10_api_test10_pb.RoomGetByIdRespond>;

  roomGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomGetByIdRespond) => void): grpcWeb.ClientReadableStream<test10_api_test10_pb.RoomGetByIdRespond>;

  roomGetById(
    request: test10_api_test10_pb.UserGetByIdRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: test10_api_test10_pb.RoomGetByIdRespond) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/test10.Test10/RoomGetById',
        request,
        metadata || {},
        this.methodDescriptorRoomGetById,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/test10.Test10/RoomGetById',
    request,
    metadata || {},
    this.methodDescriptorRoomGetById);
  }

}

