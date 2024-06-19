// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api/database.proto

package database

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Database_AuthMethod int32

const (
	Database_AUTH_METHOD_UNSPECIFIED       Database_AuthMethod = 0
	Database_AUTH_METHOD_NONE              Database_AuthMethod = 1
	Database_AUTH_METHOD_USERNAME_PASSWORD Database_AuthMethod = 2
	Database_AUTH_METHOD_AWS_IAM           Database_AuthMethod = 4
)

// Enum value maps for Database_AuthMethod.
var (
	Database_AuthMethod_name = map[int32]string{
		0: "AUTH_METHOD_UNSPECIFIED",
		1: "AUTH_METHOD_NONE",
		2: "AUTH_METHOD_USERNAME_PASSWORD",
		4: "AUTH_METHOD_AWS_IAM",
	}
	Database_AuthMethod_value = map[string]int32{
		"AUTH_METHOD_UNSPECIFIED":       0,
		"AUTH_METHOD_NONE":              1,
		"AUTH_METHOD_USERNAME_PASSWORD": 2,
		"AUTH_METHOD_AWS_IAM":           4,
	}
)

func (x Database_AuthMethod) Enum() *Database_AuthMethod {
	p := new(Database_AuthMethod)
	*p = x
	return p
}

func (x Database_AuthMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Database_AuthMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_api_database_proto_enumTypes[0].Descriptor()
}

func (Database_AuthMethod) Type() protoreflect.EnumType {
	return &file_api_database_proto_enumTypes[0]
}

func (x Database_AuthMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Database_AuthMethod.Descriptor instead.
func (Database_AuthMethod) EnumDescriptor() ([]byte, []int) {
	return file_api_database_proto_rawDescGZIP(), []int{0, 0}
}

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host            string              `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port            uint32              `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Name            string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // database name
	TracingEnabled  bool                `protobuf:"varint,4,opt,name=tracing_enabled,json=tracingEnabled,proto3" json:"tracing_enabled,omitempty"`
	Debug           bool                `protobuf:"varint,5,opt,name=debug,proto3" json:"debug,omitempty"`
	MaxIdleConns    uint32              `protobuf:"varint,6,opt,name=max_idle_conns,json=maxIdleConns,proto3" json:"max_idle_conns,omitempty"`
	MaxOpenConns    uint32              `protobuf:"varint,7,opt,name=max_open_conns,json=maxOpenConns,proto3" json:"max_open_conns,omitempty"`
	ConnMaxLifeTime uint32              `protobuf:"varint,8,opt,name=conn_max_life_time,json=connMaxLifeTime,proto3" json:"conn_max_life_time,omitempty"` // minutes
	ConnMaxIdleTime uint32              `protobuf:"varint,9,opt,name=conn_max_idle_time,json=connMaxIdleTime,proto3" json:"conn_max_idle_time,omitempty"` // minutes
	AuthMethod      Database_AuthMethod `protobuf:"varint,10,opt,name=auth_method,json=authMethod,proto3,enum=greyhole.database.Database_AuthMethod" json:"auth_method,omitempty"`
	// AUTH_METHOD_USERNAME_PASSWORD
	Username string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	// AUTH_METHOD_AWS_IAM
	AwsRegion string `protobuf:"bytes,13,opt,name=aws_region,json=awsRegion,proto3" json:"aws_region,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_api_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_api_database_proto_rawDescGZIP(), []int{0}
}

func (x *Database) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Database) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetTracingEnabled() bool {
	if x != nil {
		return x.TracingEnabled
	}
	return false
}

func (x *Database) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Database) GetMaxIdleConns() uint32 {
	if x != nil {
		return x.MaxIdleConns
	}
	return 0
}

func (x *Database) GetMaxOpenConns() uint32 {
	if x != nil {
		return x.MaxOpenConns
	}
	return 0
}

func (x *Database) GetConnMaxLifeTime() uint32 {
	if x != nil {
		return x.ConnMaxLifeTime
	}
	return 0
}

func (x *Database) GetConnMaxIdleTime() uint32 {
	if x != nil {
		return x.ConnMaxIdleTime
	}
	return 0
}

func (x *Database) GetAuthMethod() Database_AuthMethod {
	if x != nil {
		return x.AuthMethod
	}
	return Database_AUTH_METHOD_UNSPECIFIED
}

func (x *Database) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Database) GetAwsRegion() string {
	if x != nil {
		return x.AwsRegion
	}
	return ""
}

var File_api_database_proto protoreflect.FileDescriptor

var file_api_database_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x72, 0x65, 0x79, 0x68, 0x6f, 0x6c, 0x65, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xee, 0x04, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18,
	0xff, 0xff, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d,
	0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6d,
	0x61, 0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e,
	0x73, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69,
	0x66, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x6e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x4d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x79, 0x68, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x77, 0x73, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x4d, 0x45,
	0x54, 0x48, 0x4f, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x55, 0x54, 0x48,
	0x5f, 0x4d, 0x45, 0x54, 0x48, 0x4f, 0x44, 0x5f, 0x41, 0x57, 0x53, 0x5f, 0x49, 0x41, 0x4d, 0x10,
	0x04, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x75, 0x67, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_database_proto_rawDescOnce sync.Once
	file_api_database_proto_rawDescData = file_api_database_proto_rawDesc
)

func file_api_database_proto_rawDescGZIP() []byte {
	file_api_database_proto_rawDescOnce.Do(func() {
		file_api_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_database_proto_rawDescData)
	})
	return file_api_database_proto_rawDescData
}

var file_api_database_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_database_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_database_proto_goTypes = []interface{}{
	(Database_AuthMethod)(0), // 0: greyhole.database.Database.AuthMethod
	(*Database)(nil),         // 1: greyhole.database.Database
}
var file_api_database_proto_depIdxs = []int32{
	0, // 0: greyhole.database.Database.auth_method:type_name -> greyhole.database.Database.AuthMethod
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_database_proto_init() }
func file_api_database_proto_init() {
	if File_api_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_database_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_database_proto_goTypes,
		DependencyIndexes: file_api_database_proto_depIdxs,
		EnumInfos:         file_api_database_proto_enumTypes,
		MessageInfos:      file_api_database_proto_msgTypes,
	}.Build()
	File_api_database_proto = out.File
	file_api_database_proto_rawDesc = nil
	file_api_database_proto_goTypes = nil
	file_api_database_proto_depIdxs = nil
}