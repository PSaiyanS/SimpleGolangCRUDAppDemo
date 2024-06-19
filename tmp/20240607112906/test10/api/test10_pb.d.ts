import * as jspb from 'google-protobuf'



export class PingRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): PingRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PingRequest): PingRequest.AsObject;
  static serializeBinaryToWriter(message: PingRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingRequest;
  static deserializeBinaryFromReader(message: PingRequest, reader: jspb.BinaryReader): PingRequest;
}

export namespace PingRequest {
  export type AsObject = {
    message: string,
  }
}

export class PingRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): PingRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PingRespond.AsObject;
  static toObject(includeInstance: boolean, msg: PingRespond): PingRespond.AsObject;
  static serializeBinaryToWriter(message: PingRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PingRespond;
  static deserializeBinaryFromReader(message: PingRespond, reader: jspb.BinaryReader): PingRespond;
}

export namespace PingRespond {
  export type AsObject = {
    message: string,
  }
}

