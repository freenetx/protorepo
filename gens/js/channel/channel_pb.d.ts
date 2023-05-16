// package: file.api
// file: proto/channel/channel.proto

import * as jspb from "google-protobuf";
import * as share_share_pb from "../../share/share_pb";

export class Create extends jspb.Message {
  getToken(): Uint8Array | string;
  getToken_asU8(): Uint8Array;
  getToken_asB64(): string;
  setToken(value: Uint8Array | string): void;

  getDestination(): Uint8Array | string;
  getDestination_asU8(): Uint8Array;
  getDestination_asB64(): string;
  setDestination(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Create.AsObject;
  static toObject(includeInstance: boolean, msg: Create): Create.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Create, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Create;
  static deserializeBinaryFromReader(message: Create, reader: jspb.BinaryReader): Create;
}

export namespace Create {
  export type AsObject = {
    token: Uint8Array | string,
    destination: Uint8Array | string,
  }
}

export class Data extends jspb.Message {
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Data.AsObject;
  static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Data;
  static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
}

export namespace Data {
  export type AsObject = {
    data: Uint8Array | string,
  }
}

export class Close extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Close.AsObject;
  static toObject(includeInstance: boolean, msg: Close): Close.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Close, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Close;
  static deserializeBinaryFromReader(message: Close, reader: jspb.BinaryReader): Close;
}

export namespace Close {
  export type AsObject = {
  }
}

export class ChannelRequest extends jspb.Message {
  hasCreate(): boolean;
  clearCreate(): void;
  getCreate(): Create | undefined;
  setCreate(value?: Create): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

  hasClose(): boolean;
  clearClose(): void;
  getClose(): Close | undefined;
  setClose(value?: Close): void;

  getReqCase(): ChannelRequest.ReqCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChannelRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChannelRequest): ChannelRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ChannelRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChannelRequest;
  static deserializeBinaryFromReader(message: ChannelRequest, reader: jspb.BinaryReader): ChannelRequest;
}

export namespace ChannelRequest {
  export type AsObject = {
    create?: Create.AsObject,
    data?: Data.AsObject,
    close?: Close.AsObject,
  }

  export enum ReqCase {
    REQ_NOT_SET = 0,
    CREATE = 1,
    DATA = 2,
    CLOSE = 3,
  }
}

export class Created extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Created.AsObject;
  static toObject(includeInstance: boolean, msg: Created): Created.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Created, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Created;
  static deserializeBinaryFromReader(message: Created, reader: jspb.BinaryReader): Created;
}

export namespace Created {
  export type AsObject = {
  }
}

export class Closed extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Closed.AsObject;
  static toObject(includeInstance: boolean, msg: Closed): Closed.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Closed, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Closed;
  static deserializeBinaryFromReader(message: Closed, reader: jspb.BinaryReader): Closed;
}

export namespace Closed {
  export type AsObject = {
  }
}

export class ReceiveResponse extends jspb.Message {
  hasCreated(): boolean;
  clearCreated(): void;
  getCreated(): Created | undefined;
  setCreated(value?: Created): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

  hasClosed(): boolean;
  clearClosed(): void;
  getClosed(): Closed | undefined;
  setClosed(value?: Closed): void;

  getRespCase(): ReceiveResponse.RespCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReceiveResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ReceiveResponse): ReceiveResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReceiveResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReceiveResponse;
  static deserializeBinaryFromReader(message: ReceiveResponse, reader: jspb.BinaryReader): ReceiveResponse;
}

export namespace ReceiveResponse {
  export type AsObject = {
    created?: Created.AsObject,
    data?: Data.AsObject,
    closed?: Closed.AsObject,
  }

  export enum RespCase {
    RESP_NOT_SET = 0,
    CREATED = 1,
    DATA = 2,
    CLOSED = 3,
  }
}

