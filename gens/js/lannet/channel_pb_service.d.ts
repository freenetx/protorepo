// package: lannet.api
// file: lannet/channel.proto

import * as lannet_channel_pb from "../lannet/channel_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CenterServiceChannel = {
  readonly methodName: string;
  readonly service: typeof CenterService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof lannet_channel_pb.CenterRequest;
  readonly responseType: typeof lannet_channel_pb.CenterResponse;
};

export class CenterService {
  static readonly serviceName: string;
  static readonly Channel: CenterServiceChannel;
}

type DirectNetServiceChannel = {
  readonly methodName: string;
  readonly service: typeof DirectNetService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof lannet_channel_pb.DirectNetRequest;
  readonly responseType: typeof lannet_channel_pb.DirectNetResponse;
};

export class DirectNetService {
  static readonly serviceName: string;
  static readonly Channel: DirectNetServiceChannel;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class CenterServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  channel(metadata?: grpc.Metadata): BidirectionalStream<lannet_channel_pb.CenterRequest, lannet_channel_pb.CenterResponse>;
}

export class DirectNetServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  channel(metadata?: grpc.Metadata): BidirectionalStream<lannet_channel_pb.DirectNetRequest, lannet_channel_pb.DirectNetResponse>;
}

