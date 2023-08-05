// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/service/pb/service.proto

package service

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

type ServiceSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户总数
	// @gotags: json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	// 用户列表
	// @gotags: json:"items"
	Items []*Service `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *ServiceSet) Reset() {
	*x = ServiceSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceSet) ProtoMessage() {}

func (x *ServiceSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceSet.ProtoReflect.Descriptor instead.
func (*ServiceSet) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ServiceSet) GetItems() []*Service {
	if x != nil {
		return x.Items
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateServiceRequest `protobuf:"bytes,15,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// @gotags: bson:"credential" json:"credentail"
	Credential *Crendential `protobuf:"bytes,2,opt,name=credential,proto3" json:"credentail" bson:"credential"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Service) GetSpec() *CreateServiceRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Service) GetCredential() *Crendential {
	if x != nil {
		return x.Credential
	}
	return nil
}

type Crendential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"client_id" json: "client_id"
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" bson:"client_id"`
	// @gotags: bson:"client_secret" json: "client_secret"
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty" bson:"client_secret"`
	// @gotags: bson:"create_at" json: "create_at"
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty" bson:"create_at"`
}

func (x *Crendential) Reset() {
	*x = Crendential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crendential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crendential) ProtoMessage() {}

func (x *Crendential) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crendential.ProtoReflect.Descriptor instead.
func (*Crendential) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{2}
}

func (x *Crendential) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Crendential) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Crendential) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

type CreateServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"domain" json: "domain" validate:"required"
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty" bson:"domain" validate:"required"`
	// @gotags: bson:"namespace" json: "namespace" validate:"required"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" bson:"namespace" validate:"required"`
	// @gotags: bson:"name" json: "name" validate:"required"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" bson:"name" validate:"required"`
}

func (x *CreateServiceRequest) Reset() {
	*x = CreateServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_service_pb_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRequest) ProtoMessage() {}

func (x *CreateServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_service_pb_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return file_apps_service_pb_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateServiceRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateServiceRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_apps_service_pb_service_proto protoreflect.FileDescriptor

var file_apps_service_pb_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5b, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc5,
	0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x76, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12,
	0x45, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x6c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x60, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x7a, 0x69, 0x68, 0x75, 0x61, 0x2f, 0x64, 0x65, 0x76,
	0x70, 0x6c, 0x61, 0x74, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_service_pb_service_proto_rawDescOnce sync.Once
	file_apps_service_pb_service_proto_rawDescData = file_apps_service_pb_service_proto_rawDesc
)

func file_apps_service_pb_service_proto_rawDescGZIP() []byte {
	file_apps_service_pb_service_proto_rawDescOnce.Do(func() {
		file_apps_service_pb_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_service_pb_service_proto_rawDescData)
	})
	return file_apps_service_pb_service_proto_rawDescData
}

var file_apps_service_pb_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_service_pb_service_proto_goTypes = []interface{}{
	(*ServiceSet)(nil),           // 0: devcloud.mcenter.service.ServiceSet
	(*Service)(nil),              // 1: devcloud.mcenter.service.Service
	(*Crendential)(nil),          // 2: devcloud.mcenter.service.Crendential
	(*CreateServiceRequest)(nil), // 3: devcloud.mcenter.service.CreateServiceRequest
	(*meta.Meta)(nil),            // 4: devcloud.mcenter.meta.Meta
}
var file_apps_service_pb_service_proto_depIdxs = []int32{
	1, // 0: devcloud.mcenter.service.ServiceSet.items:type_name -> devcloud.mcenter.service.Service
	4, // 1: devcloud.mcenter.service.Service.meta:type_name -> devcloud.mcenter.meta.Meta
	3, // 2: devcloud.mcenter.service.Service.spec:type_name -> devcloud.mcenter.service.CreateServiceRequest
	2, // 3: devcloud.mcenter.service.Service.credential:type_name -> devcloud.mcenter.service.Crendential
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_service_pb_service_proto_init() }
func file_apps_service_pb_service_proto_init() {
	if File_apps_service_pb_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_service_pb_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceSet); i {
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
		file_apps_service_pb_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_apps_service_pb_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crendential); i {
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
		file_apps_service_pb_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServiceRequest); i {
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
			RawDescriptor: file_apps_service_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_service_pb_service_proto_goTypes,
		DependencyIndexes: file_apps_service_pb_service_proto_depIdxs,
		MessageInfos:      file_apps_service_pb_service_proto_msgTypes,
	}.Build()
	File_apps_service_pb_service_proto = out.File
	file_apps_service_pb_service_proto_rawDesc = nil
	file_apps_service_pb_service_proto_goTypes = nil
	file_apps_service_pb_service_proto_depIdxs = nil
}
