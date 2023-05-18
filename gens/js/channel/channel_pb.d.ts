// package: file.api
// file: proto/channel/channel.proto

import * as jspb from "google-protobuf";
import * as share_share_pb from "../../share/share_pb";

export class SetKey extends jspb.Message {
  getSelfKey(): string;
  setSelfKey(value: string): void;

  getDestinationKey(): string;
  setDestinationKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetKey.AsObject;
  static toObject(includeInstance: boolean, msg: SetKey): SetKey.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SetKey, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetKey;
  static deserializeBinaryFromReader(message: SetKey, reader: jspb.BinaryReader): SetKey;
}

export namespace SetKey {
  export type AsObject = {
    selfKey: string,
    destinationKey: string,
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

export class ChannelRequest extends jspb.Message {
  hasSetKey(): boolean;
  clearSetKey(): void;
  getSetKey(): SetKey | undefined;
  setSetKey(value?: SetKey): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

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
    setKey?: SetKey.AsObject,
    data?: Data.AsObject,
  }

  export enum ReqCase {
    REQ_NOT_SET = 0,
    SET_KEY = 1,
    DATA = 2,
  }
}

export class KeyAccept extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): KeyAccept.AsObject;
  static toObject(includeInstance: boolean, msg: KeyAccept): KeyAccept.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: KeyAccept, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): KeyAccept;
  static deserializeBinaryFromReader(message: KeyAccept, reader: jspb.BinaryReader): KeyAccept;
}

export namespace KeyAccept {
  export type AsObject = {
  }
}

export class ReceiveResponse extends jspb.Message {
  hasKeyAccept(): boolean;
  clearKeyAccept(): void;
  getKeyAccept(): KeyAccept | undefined;
  setKeyAccept(value?: KeyAccept): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

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
    keyAccept?: KeyAccept.AsObject,
    data?: Data.AsObject,
  }

  export enum RespCase {
    RESP_NOT_SET = 0,
    KEY_ACCEPT = 1,
    DATA = 2,
  }
}

