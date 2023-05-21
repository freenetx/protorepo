// package: channel.api
// file: proto/channel/channel.proto

import * as proto_channel_channel_pb from "../../proto/channel/channel_pb";
import * as share_share_pb from "../../share/share_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ChannelServiceTest = {
  readonly methodName: string;
  readonly service: typeof ChannelService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof share_share_pb.Empty;
  readonly responseType: typeof share_share_pb.Empty;
};

type ChannelServiceChannel = {
  readonly methodName: string;
  readonly service: typeof ChannelService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof proto_channel_channel_pb.ChannelRequest;
  readonly responseType: typeof proto_channel_channel_pb.ReceiveResponse;
};

export class ChannelService {
  static readonly serviceName: string;
  static readonly Test: ChannelServiceTest;
  static readonly Channel: ChannelServiceChannel;
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

export class ChannelServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  test(
    requestMessage: share_share_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: share_share_pb.Empty|null) => void
  ): UnaryResponse;
  test(
    requestMessage: share_share_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: share_share_pb.Empty|null) => void
  ): UnaryResponse;
  channel(metadata?: grpc.Metadata): BidirectionalStream<proto_channel_channel_pb.ChannelRequest, proto_channel_channel_pb.ReceiveResponse>;
}

