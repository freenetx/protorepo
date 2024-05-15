// package: lannet.api
// file: lannet/channel.proto

import * as jspb from "google-protobuf";

export class HelloServer extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  clearMyIpsList(): void;
  getMyIpsList(): Array<string>;
  setMyIpsList(value: Array<string>): void;
  addMyIps(value: string, index?: number): string;

  getMyEncodedAddress(): string;
  setMyEncodedAddress(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloServer.AsObject;
  static toObject(includeInstance: boolean, msg: HelloServer): HelloServer.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloServer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloServer;
  static deserializeBinaryFromReader(message: HelloServer, reader: jspb.BinaryReader): HelloServer;
}

export namespace HelloServer {
  export type AsObject = {
    key: string,
    myIpsList: Array<string>,
    myEncodedAddress: string,
  }
}

export class CenterRequest extends jspb.Message {
  hasHello(): boolean;
  clearHello(): void;
  getHello(): HelloServer | undefined;
  setHello(value?: HelloServer): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getReqCase(): CenterRequest.ReqCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CenterRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CenterRequest): CenterRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CenterRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CenterRequest;
  static deserializeBinaryFromReader(message: CenterRequest, reader: jspb.BinaryReader): CenterRequest;
}

export namespace CenterRequest {
  export type AsObject = {
    hello?: HelloServer.AsObject,
    data: Uint8Array | string,
  }

  export enum ReqCase {
    REQ_NOT_SET = 0,
    HELLO = 1,
    DATA = 2,
  }
}

export class StringList extends jspb.Message {
  clearStringsList(): void;
  getStringsList(): Array<string>;
  setStringsList(value: Array<string>): void;
  addStrings(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StringList.AsObject;
  static toObject(includeInstance: boolean, msg: StringList): StringList.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StringList, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StringList;
  static deserializeBinaryFromReader(message: StringList, reader: jspb.BinaryReader): StringList;
}

export namespace StringList {
  export type AsObject = {
    stringsList: Array<string>,
  }
}

export class HelloClient extends jspb.Message {
  getClientsMap(): jspb.Map<string, string>;
  clearClientsMap(): void;
  getClientIpsMap(): jspb.Map<string, StringList>;
  clearClientIpsMap(): void;
  getVpn(): string;
  setVpn(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HelloClient.AsObject;
  static toObject(includeInstance: boolean, msg: HelloClient): HelloClient.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HelloClient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HelloClient;
  static deserializeBinaryFromReader(message: HelloClient, reader: jspb.BinaryReader): HelloClient;
}

export namespace HelloClient {
  export type AsObject = {
    clientsMap: Array<[string, string]>,
    clientIpsMap: Array<[string, StringList.AsObject]>,
    vpn: string,
  }
}

export class CenterResponse extends jspb.Message {
  hasHello(): boolean;
  clearHello(): void;
  getHello(): HelloClient | undefined;
  setHello(value?: HelloClient): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getRespCase(): CenterResponse.RespCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CenterResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CenterResponse): CenterResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CenterResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CenterResponse;
  static deserializeBinaryFromReader(message: CenterResponse, reader: jspb.BinaryReader): CenterResponse;
}

export namespace CenterResponse {
  export type AsObject = {
    hello?: HelloClient.AsObject,
    data: Uint8Array | string,
  }

  export enum RespCase {
    RESP_NOT_SET = 0,
    HELLO = 1,
    DATA = 2,
  }
}

export class DirectNetHelloServer extends jspb.Message {
  getKey(): string;
  setKey(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirectNetHelloServer.AsObject;
  static toObject(includeInstance: boolean, msg: DirectNetHelloServer): DirectNetHelloServer.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirectNetHelloServer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirectNetHelloServer;
  static deserializeBinaryFromReader(message: DirectNetHelloServer, reader: jspb.BinaryReader): DirectNetHelloServer;
}

export namespace DirectNetHelloServer {
  export type AsObject = {
    key: string,
  }
}

export class DirectNetRequest extends jspb.Message {
  hasHello(): boolean;
  clearHello(): void;
  getHello(): DirectNetHelloServer | undefined;
  setHello(value?: DirectNetHelloServer): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getReqCase(): DirectNetRequest.ReqCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirectNetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DirectNetRequest): DirectNetRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirectNetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirectNetRequest;
  static deserializeBinaryFromReader(message: DirectNetRequest, reader: jspb.BinaryReader): DirectNetRequest;
}

export namespace DirectNetRequest {
  export type AsObject = {
    hello?: DirectNetHelloServer.AsObject,
    data: Uint8Array | string,
  }

  export enum ReqCase {
    REQ_NOT_SET = 0,
    HELLO = 1,
    DATA = 2,
  }
}

export class DirectNetHelloClient extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirectNetHelloClient.AsObject;
  static toObject(includeInstance: boolean, msg: DirectNetHelloClient): DirectNetHelloClient.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirectNetHelloClient, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirectNetHelloClient;
  static deserializeBinaryFromReader(message: DirectNetHelloClient, reader: jspb.BinaryReader): DirectNetHelloClient;
}

export namespace DirectNetHelloClient {
  export type AsObject = {
  }
}

export class DirectNetResponse extends jspb.Message {
  hasHello(): boolean;
  clearHello(): void;
  getHello(): DirectNetHelloClient | undefined;
  setHello(value?: DirectNetHelloClient): void;

  hasData(): boolean;
  clearData(): void;
  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  getRespCase(): DirectNetResponse.RespCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DirectNetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DirectNetResponse): DirectNetResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: DirectNetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DirectNetResponse;
  static deserializeBinaryFromReader(message: DirectNetResponse, reader: jspb.BinaryReader): DirectNetResponse;
}

export namespace DirectNetResponse {
  export type AsObject = {
    hello?: DirectNetHelloClient.AsObject,
    data: Uint8Array | string,
  }

  export enum RespCase {
    RESP_NOT_SET = 0,
    HELLO = 1,
    DATA = 2,
  }
}

