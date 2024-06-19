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

export class User extends jspb.Message {
  getId(): number;
  setId(value: number): User;

  getUsername(): string;
  setUsername(value: string): User;

  getPassword(): string;
  setPassword(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: number,
    username: string,
    password: string,
  }
}

export class UserCreateRequest extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): UserCreateRequest;
  hasUser(): boolean;
  clearUser(): UserCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserCreateRequest): UserCreateRequest.AsObject;
  static serializeBinaryToWriter(message: UserCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCreateRequest;
  static deserializeBinaryFromReader(message: UserCreateRequest, reader: jspb.BinaryReader): UserCreateRequest;
}

export namespace UserCreateRequest {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export class UserCreateRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UserCreateRespond;

  getUser(): User | undefined;
  setUser(value?: User): UserCreateRespond;
  hasUser(): boolean;
  clearUser(): UserCreateRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCreateRespond.AsObject;
  static toObject(includeInstance: boolean, msg: UserCreateRespond): UserCreateRespond.AsObject;
  static serializeBinaryToWriter(message: UserCreateRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCreateRespond;
  static deserializeBinaryFromReader(message: UserCreateRespond, reader: jspb.BinaryReader): UserCreateRespond;
}

export namespace UserCreateRespond {
  export type AsObject = {
    message: string,
    user?: User.AsObject,
  }
}

export class UserListRequest extends jspb.Message {
  getOffset(): number;
  setOffset(value: number): UserListRequest;

  getLimit(): number;
  setLimit(value: number): UserListRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserListRequest): UserListRequest.AsObject;
  static serializeBinaryToWriter(message: UserListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserListRequest;
  static deserializeBinaryFromReader(message: UserListRequest, reader: jspb.BinaryReader): UserListRequest;
}

export namespace UserListRequest {
  export type AsObject = {
    offset: number,
    limit: number,
  }
}

export class UserListRespond extends jspb.Message {
  getUsersList(): Array<User>;
  setUsersList(value: Array<User>): UserListRespond;
  clearUsersList(): UserListRespond;
  addUsers(value?: User, index?: number): User;

  getTotalCount(): number;
  setTotalCount(value: number): UserListRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserListRespond.AsObject;
  static toObject(includeInstance: boolean, msg: UserListRespond): UserListRespond.AsObject;
  static serializeBinaryToWriter(message: UserListRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserListRespond;
  static deserializeBinaryFromReader(message: UserListRespond, reader: jspb.BinaryReader): UserListRespond;
}

export namespace UserListRespond {
  export type AsObject = {
    usersList: Array<User.AsObject>,
    totalCount: number,
  }
}

export class UserUpdateRequest extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): UserUpdateRequest;
  hasUser(): boolean;
  clearUser(): UserUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserUpdateRequest): UserUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UserUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserUpdateRequest;
  static deserializeBinaryFromReader(message: UserUpdateRequest, reader: jspb.BinaryReader): UserUpdateRequest;
}

export namespace UserUpdateRequest {
  export type AsObject = {
    user?: User.AsObject,
  }
}

export class UserUpdateRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UserUpdateRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserUpdateRespond.AsObject;
  static toObject(includeInstance: boolean, msg: UserUpdateRespond): UserUpdateRespond.AsObject;
  static serializeBinaryToWriter(message: UserUpdateRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserUpdateRespond;
  static deserializeBinaryFromReader(message: UserUpdateRespond, reader: jspb.BinaryReader): UserUpdateRespond;
}

export namespace UserUpdateRespond {
  export type AsObject = {
    message: string,
  }
}

export class UserDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): UserDeleteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserDeleteRequest): UserDeleteRequest.AsObject;
  static serializeBinaryToWriter(message: UserDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserDeleteRequest;
  static deserializeBinaryFromReader(message: UserDeleteRequest, reader: jspb.BinaryReader): UserDeleteRequest;
}

export namespace UserDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class UserDeleteRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UserDeleteRespond;

  getUser(): User | undefined;
  setUser(value?: User): UserDeleteRespond;
  hasUser(): boolean;
  clearUser(): UserDeleteRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserDeleteRespond.AsObject;
  static toObject(includeInstance: boolean, msg: UserDeleteRespond): UserDeleteRespond.AsObject;
  static serializeBinaryToWriter(message: UserDeleteRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserDeleteRespond;
  static deserializeBinaryFromReader(message: UserDeleteRespond, reader: jspb.BinaryReader): UserDeleteRespond;
}

export namespace UserDeleteRespond {
  export type AsObject = {
    message: string,
    user?: User.AsObject,
  }
}

export class UserGetByIdRequest extends jspb.Message {
  getId(): number;
  setId(value: number): UserGetByIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetByIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetByIdRequest): UserGetByIdRequest.AsObject;
  static serializeBinaryToWriter(message: UserGetByIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetByIdRequest;
  static deserializeBinaryFromReader(message: UserGetByIdRequest, reader: jspb.BinaryReader): UserGetByIdRequest;
}

export namespace UserGetByIdRequest {
  export type AsObject = {
    id: number,
  }
}

export class UserGetByIdRespond extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): UserGetByIdRespond;
  hasUser(): boolean;
  clearUser(): UserGetByIdRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetByIdRespond.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetByIdRespond): UserGetByIdRespond.AsObject;
  static serializeBinaryToWriter(message: UserGetByIdRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetByIdRespond;
  static deserializeBinaryFromReader(message: UserGetByIdRespond, reader: jspb.BinaryReader): UserGetByIdRespond;
}

export namespace UserGetByIdRespond {
  export type AsObject = {
    user?: User.AsObject,
  }
}

