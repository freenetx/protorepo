// package: freenet.api
// file: proto/freenet/freenext.proto

import * as jspb from "google-protobuf";

export class SyncMessage extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  clearForwardIpsList(): void;
  getForwardIpsList(): Array<string>;
  setForwardIpsList(value: Array<string>): void;
  addForwardIps(value: string, index?: number): string;

  getKnownClientsMap(): jspb.Map<string, string>;
  clearKnownClientsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SyncMessage.AsObject;
  static toObject(includeInstance: boolean, msg: SyncMessage): SyncMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SyncMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SyncMessage;
  static deserializeBinaryFromReader(message: SyncMessage, reader: jspb.BinaryReader): SyncMessage;
}

export namespace SyncMessage {
  export type AsObject = {
    id: string,
    forwardIpsList: Array<string>,
    knownClientsMap: Array<[string, string]>,
  }
}

export class MessageAck extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageAck.AsObject;
  static toObject(includeInstance: boolean, msg: MessageAck): MessageAck.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MessageAck, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageAck;
  static deserializeBinaryFromReader(message: MessageAck, reader: jspb.BinaryReader): MessageAck;
}

export namespace MessageAck {
  export type AsObject = {
    id: string,
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

export class ClientInitMessage extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  clearLanIpsList(): void;
  getLanIpsList(): Array<string>;
  setLanIpsList(value: Array<string>): void;
  addLanIps(value: string, index?: number): string;

  clearVpnIpsList(): void;
  getVpnIpsList(): Array<string>;
  setVpnIpsList(value: Array<string>): void;
  addVpnIps(value: string, index?: number): string;

  getClientAddress(): string;
  setClientAddress(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientInitMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ClientInitMessage): ClientInitMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientInitMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientInitMessage;
  static deserializeBinaryFromReader(message: ClientInitMessage, reader: jspb.BinaryReader): ClientInitMessage;
}

export namespace ClientInitMessage {
  export type AsObject = {
    id: string,
    lanIpsList: Array<string>,
    vpnIpsList: Array<string>,
    clientAddress: string,
  }
}

export class NeedClientInitMessage extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NeedClientInitMessage.AsObject;
  static toObject(includeInstance: boolean, msg: NeedClientInitMessage): NeedClientInitMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: NeedClientInitMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NeedClientInitMessage;
  static deserializeBinaryFromReader(message: NeedClientInitMessage, reader: jspb.BinaryReader): NeedClientInitMessage;
}

export namespace NeedClientInitMessage {
  export type AsObject = {
  }
}

export class ClientMessage extends jspb.Message {
  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

  hasClientInitMessage(): boolean;
  clearClientInitMessage(): void;
  getClientInitMessage(): ClientInitMessage | undefined;
  setClientInitMessage(value?: ClientInitMessage): void;

  hasMessageAck(): boolean;
  clearMessageAck(): void;
  getMessageAck(): MessageAck | undefined;
  setMessageAck(value?: MessageAck): void;

  getRequestCase(): ClientMessage.RequestCase;
  getResponseCase(): ClientMessage.ResponseCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ClientMessage): ClientMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientMessage;
  static deserializeBinaryFromReader(message: ClientMessage, reader: jspb.BinaryReader): ClientMessage;
}

export namespace ClientMessage {
  export type AsObject = {
    data?: Data.AsObject,
    clientInitMessage?: ClientInitMessage.AsObject,
    messageAck?: MessageAck.AsObject,
  }

  export enum RequestCase {
    REQUEST_NOT_SET = 0,
    DATA = 1,
    CLIENT_INIT_MESSAGE = 2,
  }

  export enum ResponseCase {
    RESPONSE_NOT_SET = 0,
    MESSAGE_ACK = 100,
  }
}

export class ServerMessage extends jspb.Message {
  hasData(): boolean;
  clearData(): void;
  getData(): Data | undefined;
  setData(value?: Data): void;

  hasSyncMessage(): boolean;
  clearSyncMessage(): void;
  getSyncMessage(): SyncMessage | undefined;
  setSyncMessage(value?: SyncMessage): void;

  hasNeedClientInitMessage(): boolean;
  clearNeedClientInitMessage(): void;
  getNeedClientInitMessage(): NeedClientInitMessage | undefined;
  setNeedClientInitMessage(value?: NeedClientInitMessage): void;

  hasMessageAck(): boolean;
  clearMessageAck(): void;
  getMessageAck(): MessageAck | undefined;
  setMessageAck(value?: MessageAck): void;

  getRequestCase(): ServerMessage.RequestCase;
  getResponseCase(): ServerMessage.ResponseCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerMessage.AsObject;
  static toObject(includeInstance: boolean, msg: ServerMessage): ServerMessage.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServerMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerMessage;
  static deserializeBinaryFromReader(message: ServerMessage, reader: jspb.BinaryReader): ServerMessage;
}

export namespace ServerMessage {
  export type AsObject = {
    data?: Data.AsObject,
    syncMessage?: SyncMessage.AsObject,
    needClientInitMessage?: NeedClientInitMessage.AsObject,
    messageAck?: MessageAck.AsObject,
  }

  export enum RequestCase {
    REQUEST_NOT_SET = 0,
    DATA = 1,
    SYNC_MESSAGE = 2,
    NEED_CLIENT_INIT_MESSAGE = 3,
  }

  export enum ResponseCase {
    RESPONSE_NOT_SET = 0,
    MESSAGE_ACK = 100,
  }
}

