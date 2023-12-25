// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: app/config/client.proto

package config

import (
	common "github.com/ydb-platform/fq-connector-go/api/common"
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

// Connector client configuration
type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Connector service instance network address we connect to
	Endpoint *common.TEndpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// TLS credentials for Connector
	Tls *ClientTLSConfig `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Data source instance we read data from
	DataSourceInstance *common.TDataSourceInstance `protobuf:"bytes,3,opt,name=data_source_instance,json=dataSourceInstance,proto3" json:"data_source_instance,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_app_config_client_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConfig) GetEndpoint() *common.TEndpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *ClientConfig) GetTls() *ClientTLSConfig {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ClientConfig) GetDataSourceInstance() *common.TDataSourceInstance {
	if x != nil {
		return x.DataSourceInstance
	}
	return nil
}

type ClientTLSConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CA certificate path
	Ca string `protobuf:"bytes,1,opt,name=ca,proto3" json:"ca,omitempty"`
}

func (x *ClientTLSConfig) Reset() {
	*x = ClientTLSConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_config_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTLSConfig) ProtoMessage() {}

func (x *ClientTLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_config_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTLSConfig.ProtoReflect.Descriptor instead.
func (*ClientTLSConfig) Descriptor() ([]byte, []int) {
	return file_app_config_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientTLSConfig) GetCa() string {
	if x != nil {
		return x.Ca
	}
	return ""
}

var File_app_config_client_proto protoreflect.FileDescriptor

var file_app_config_client_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x4e, 0x59, 0x71, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x48, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x2f, 0x79, 0x71, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45,
	0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e,
	0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e,
	0x54, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x03, 0x74, 0x6c,
	0x73, 0x12, 0x5b, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x12, 0x64, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x21,
	0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x4c, 0x53, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63,
	0x61, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x64, 0x62, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x66, 0x71, 0x2d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_config_client_proto_rawDescOnce sync.Once
	file_app_config_client_proto_rawDescData = file_app_config_client_proto_rawDesc
)

func file_app_config_client_proto_rawDescGZIP() []byte {
	file_app_config_client_proto_rawDescOnce.Do(func() {
		file_app_config_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_config_client_proto_rawDescData)
	})
	return file_app_config_client_proto_rawDescData
}

var file_app_config_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_config_client_proto_goTypes = []interface{}{
	(*ClientConfig)(nil),               // 0: NYql.Connector.App.Config.ClientConfig
	(*ClientTLSConfig)(nil),            // 1: NYql.Connector.App.Config.ClientTLSConfig
	(*common.TEndpoint)(nil),           // 2: NYql.NConnector.NApi.TEndpoint
	(*common.TDataSourceInstance)(nil), // 3: NYql.NConnector.NApi.TDataSourceInstance
}
var file_app_config_client_proto_depIdxs = []int32{
	2, // 0: NYql.Connector.App.Config.ClientConfig.endpoint:type_name -> NYql.NConnector.NApi.TEndpoint
	1, // 1: NYql.Connector.App.Config.ClientConfig.tls:type_name -> NYql.Connector.App.Config.ClientTLSConfig
	3, // 2: NYql.Connector.App.Config.ClientConfig.data_source_instance:type_name -> NYql.NConnector.NApi.TDataSourceInstance
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_app_config_client_proto_init() }
func file_app_config_client_proto_init() {
	if File_app_config_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_config_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_app_config_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTLSConfig); i {
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
			RawDescriptor: file_app_config_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_config_client_proto_goTypes,
		DependencyIndexes: file_app_config_client_proto_depIdxs,
		MessageInfos:      file_app_config_client_proto_msgTypes,
	}.Build()
	File_app_config_client_proto = out.File
	file_app_config_client_proto_rawDesc = nil
	file_app_config_client_proto_goTypes = nil
	file_app_config_client_proto_depIdxs = nil
}
