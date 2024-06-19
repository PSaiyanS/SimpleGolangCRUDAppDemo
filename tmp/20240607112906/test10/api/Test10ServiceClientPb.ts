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

}

