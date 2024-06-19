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

  getEmail(): string;
  setEmail(value: string): User;

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
    email: string,
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

  getUser(): User | undefined;
  setUser(value?: User): UserUpdateRespond;
  hasUser(): boolean;
  clearUser(): UserUpdateRespond;

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
    user?: User.AsObject,
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

export class Room extends jspb.Message {
  getId(): number;
  setId(value: number): Room;

  getName(): string;
  setName(value: string): Room;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Room.AsObject;
  static toObject(includeInstance: boolean, msg: Room): Room.AsObject;
  static serializeBinaryToWriter(message: Room, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Room;
  static deserializeBinaryFromReader(message: Room, reader: jspb.BinaryReader): Room;
}

export namespace Room {
  export type AsObject = {
    id: number,
    name: string,
  }
}

export class RoomCreateRequest extends jspb.Message {
  getRoom(): Room | undefined;
  setRoom(value?: Room): RoomCreateRequest;
  hasRoom(): boolean;
  clearRoom(): RoomCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RoomCreateRequest): RoomCreateRequest.AsObject;
  static serializeBinaryToWriter(message: RoomCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomCreateRequest;
  static deserializeBinaryFromReader(message: RoomCreateRequest, reader: jspb.BinaryReader): RoomCreateRequest;
}

export namespace RoomCreateRequest {
  export type AsObject = {
    room?: Room.AsObject,
  }
}

export class RoomCreateRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): RoomCreateRespond;

  getRoom(): Room | undefined;
  setRoom(value?: Room): RoomCreateRespond;
  hasRoom(): boolean;
  clearRoom(): RoomCreateRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomCreateRespond.AsObject;
  static toObject(includeInstance: boolean, msg: RoomCreateRespond): RoomCreateRespond.AsObject;
  static serializeBinaryToWriter(message: RoomCreateRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomCreateRespond;
  static deserializeBinaryFromReader(message: RoomCreateRespond, reader: jspb.BinaryReader): RoomCreateRespond;
}

export namespace RoomCreateRespond {
  export type AsObject = {
    message: string,
    room?: Room.AsObject,
  }
}

export class RoomListRequest extends jspb.Message {
  getOffset(): number;
  setOffset(value: number): RoomListRequest;

  getLimit(): number;
  setLimit(value: number): RoomListRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomListRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RoomListRequest): RoomListRequest.AsObject;
  static serializeBinaryToWriter(message: RoomListRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomListRequest;
  static deserializeBinaryFromReader(message: RoomListRequest, reader: jspb.BinaryReader): RoomListRequest;
}

export namespace RoomListRequest {
  export type AsObject = {
    offset: number,
    limit: number,
  }
}

export class RoomListRespond extends jspb.Message {
  getRoomsList(): Array<Room>;
  setRoomsList(value: Array<Room>): RoomListRespond;
  clearRoomsList(): RoomListRespond;
  addRooms(value?: Room, index?: number): Room;

  getTotalCount(): number;
  setTotalCount(value: number): RoomListRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomListRespond.AsObject;
  static toObject(includeInstance: boolean, msg: RoomListRespond): RoomListRespond.AsObject;
  static serializeBinaryToWriter(message: RoomListRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomListRespond;
  static deserializeBinaryFromReader(message: RoomListRespond, reader: jspb.BinaryReader): RoomListRespond;
}

export namespace RoomListRespond {
  export type AsObject = {
    roomsList: Array<Room.AsObject>,
    totalCount: number,
  }
}

export class RoomUpdateRequest extends jspb.Message {
  getRoom(): Room | undefined;
  setRoom(value?: Room): RoomUpdateRequest;
  hasRoom(): boolean;
  clearRoom(): RoomUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RoomUpdateRequest): RoomUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: RoomUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomUpdateRequest;
  static deserializeBinaryFromReader(message: RoomUpdateRequest, reader: jspb.BinaryReader): RoomUpdateRequest;
}

export namespace RoomUpdateRequest {
  export type AsObject = {
    room?: Room.AsObject,
  }
}

export class RoomUpdateRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): RoomUpdateRespond;

  getRoom(): Room | undefined;
  setRoom(value?: Room): RoomUpdateRespond;
  hasRoom(): boolean;
  clearRoom(): RoomUpdateRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomUpdateRespond.AsObject;
  static toObject(includeInstance: boolean, msg: RoomUpdateRespond): RoomUpdateRespond.AsObject;
  static serializeBinaryToWriter(message: RoomUpdateRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomUpdateRespond;
  static deserializeBinaryFromReader(message: RoomUpdateRespond, reader: jspb.BinaryReader): RoomUpdateRespond;
}

export namespace RoomUpdateRespond {
  export type AsObject = {
    message: string,
    room?: Room.AsObject,
  }
}

export class RoomDeleteRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RoomDeleteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RoomDeleteRequest): RoomDeleteRequest.AsObject;
  static serializeBinaryToWriter(message: RoomDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomDeleteRequest;
  static deserializeBinaryFromReader(message: RoomDeleteRequest, reader: jspb.BinaryReader): RoomDeleteRequest;
}

export namespace RoomDeleteRequest {
  export type AsObject = {
    id: number,
  }
}

export class RoomDeleteRespond extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): RoomDeleteRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomDeleteRespond.AsObject;
  static toObject(includeInstance: boolean, msg: RoomDeleteRespond): RoomDeleteRespond.AsObject;
  static serializeBinaryToWriter(message: RoomDeleteRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomDeleteRespond;
  static deserializeBinaryFromReader(message: RoomDeleteRespond, reader: jspb.BinaryReader): RoomDeleteRespond;
}

export namespace RoomDeleteRespond {
  export type AsObject = {
    message: string,
  }
}

export class RoomGetByIdRequest extends jspb.Message {
  getId(): number;
  setId(value: number): RoomGetByIdRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomGetByIdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RoomGetByIdRequest): RoomGetByIdRequest.AsObject;
  static serializeBinaryToWriter(message: RoomGetByIdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomGetByIdRequest;
  static deserializeBinaryFromReader(message: RoomGetByIdRequest, reader: jspb.BinaryReader): RoomGetByIdRequest;
}

export namespace RoomGetByIdRequest {
  export type AsObject = {
    id: number,
  }
}

export class RoomGetByIdRespond extends jspb.Message {
  getRoom(): Room | undefined;
  setRoom(value?: Room): RoomGetByIdRespond;
  hasRoom(): boolean;
  clearRoom(): RoomGetByIdRespond;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RoomGetByIdRespond.AsObject;
  static toObject(includeInstance: boolean, msg: RoomGetByIdRespond): RoomGetByIdRespond.AsObject;
  static serializeBinaryToWriter(message: RoomGetByIdRespond, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RoomGetByIdRespond;
  static deserializeBinaryFromReader(message: RoomGetByIdRespond, reader: jspb.BinaryReader): RoomGetByIdRespond;
}

export namespace RoomGetByIdRespond {
  export type AsObject = {
    room?: Room.AsObject,
  }
}

export class Message extends jspb.Message {
  getId(): number;
  setId(value: number): Message;

  getContent(): string;
  setContent(value: string): Message;

  getUserId(): string;
  setUserId(value: string): Message;

  getRoomId(): string;
  setRoomId(value: string): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    id: number,
    content: string,
    userId: string,
    roomId: string,
  }
}

