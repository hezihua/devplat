// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/endpoint/pb/endpoint.proto

package endpoint

import (
	meta "github.com/hezihua/devplat/mcenter/common/meta"
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

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateEndpointRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Endpoint) GetSpec() *CreateEndpointRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type CreateEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务id
	// @gotags: bson:"service_id" json:"service_id"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" bson:"service_id"`
	// @gotags: bson:"method" json:"method"
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method" bson:"method"`
	// @gotags: bson:"path" json:"path"
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path" bson:"path"`
	// 函数名称
	// @gotags: bson:"operation" json:"operation"
	Operation string `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation" bson:"operation"`
	// 是否开启认证
	// @gotags: bson:"auth" json:"auth"
	Auth bool `protobuf:"varint,5,opt,name=auth,proto3" json:"auth" bson:"auth"`
}

func (x *CreateEndpointRequest) Reset() {
	*x = CreateEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEndpointRequest) ProtoMessage() {}

func (x *CreateEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_endpoint_pb_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEndpointRequest.ProtoReflect.Descriptor instead.
func (*CreateEndpointRequest) Descriptor() ([]byte, []int) {
	return file_apps_endpoint_pb_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEndpointRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *CreateEndpointRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CreateEndpointRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateEndpointRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *CreateEndpointRequest) GetAuth() bool {
	if x != nil {
		return x.Auth
	}
	return false
}

var File_apps_endpoint_pb_endpoint_proto protoreflect.FileDescriptor

var file_apps_endpoint_pb_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x94, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65,
	0x7a, 0x69, 0x68, 0x75, 0x61, 0x2f, 0x64, 0x65, 0x76, 0x70, 0x6c, 0x61, 0x74, 0x2f, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_endpoint_pb_endpoint_proto_rawDescOnce sync.Once
	file_apps_endpoint_pb_endpoint_proto_rawDescData = file_apps_endpoint_pb_endpoint_proto_rawDesc
)

func file_apps_endpoint_pb_endpoint_proto_rawDescGZIP() []byte {
	file_apps_endpoint_pb_endpoint_proto_rawDescOnce.Do(func() {
		file_apps_endpoint_pb_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_endpoint_pb_endpoint_proto_rawDescData)
	})
	return file_apps_endpoint_pb_endpoint_proto_rawDescData
}

var file_apps_endpoint_pb_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_endpoint_pb_endpoint_proto_goTypes = []interface{}{
	(*Endpoint)(nil),              // 0: devcloud.mcenter.endpoint.Endpoint
	(*CreateEndpointRequest)(nil), // 1: devcloud.mcenter.endpoint.CreateEndpointRequest
	(*meta.Meta)(nil),             // 2: devcloud.mcenter.meta.Meta
}
var file_apps_endpoint_pb_endpoint_proto_depIdxs = []int32{
	2, // 0: devcloud.mcenter.endpoint.Endpoint.meta:type_name -> devcloud.mcenter.meta.Meta
	1, // 1: devcloud.mcenter.endpoint.Endpoint.spec:type_name -> devcloud.mcenter.endpoint.CreateEndpointRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_endpoint_pb_endpoint_proto_init() }
func file_apps_endpoint_pb_endpoint_proto_init() {
	if File_apps_endpoint_pb_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_endpoint_pb_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_apps_endpoint_pb_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEndpointRequest); i {
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
			RawDescriptor: file_apps_endpoint_pb_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_endpoint_pb_endpoint_proto_goTypes,
		DependencyIndexes: file_apps_endpoint_pb_endpoint_proto_depIdxs,
		MessageInfos:      file_apps_endpoint_pb_endpoint_proto_msgTypes,
	}.Build()
	File_apps_endpoint_pb_endpoint_proto = out.File
	file_apps_endpoint_pb_endpoint_proto_rawDesc = nil
	file_apps_endpoint_pb_endpoint_proto_goTypes = nil
	file_apps_endpoint_pb_endpoint_proto_depIdxs = nil
}
