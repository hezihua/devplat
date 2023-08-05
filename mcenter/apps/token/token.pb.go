// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: apps/token/pb/token.proto

package token

import (
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

type GRANT_TYPE int32

const (
	// 用户密码授权
	GRANT_TYPE_PASSWORD GRANT_TYPE = 0
	// LDAP授权
	GRANT_TYPE_LDAP GRANT_TYPE = 1
	// 飞书授权
	GRANT_TYPE_FEISHU GRANT_TYPE = 8
	// 飞书授权
	GRANT_TYPE_DINGDING GRANT_TYPE = 9
)

// Enum value maps for GRANT_TYPE.
var (
	GRANT_TYPE_name = map[int32]string{
		0: "PASSWORD",
		1: "LDAP",
		8: "FEISHU",
		9: "DINGDING",
	}
	GRANT_TYPE_value = map[string]int32{
		"PASSWORD": 0,
		"LDAP":     1,
		"FEISHU":   8,
		"DINGDING": 9,
	}
)

func (x GRANT_TYPE) Enum() *GRANT_TYPE {
	p := new(GRANT_TYPE)
	*p = x
	return p
}

func (x GRANT_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GRANT_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_token_proto_enumTypes[0].Descriptor()
}

func (GRANT_TYPE) Type() protoreflect.EnumType {
	return &file_apps_token_pb_token_proto_enumTypes[0]
}

func (x GRANT_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GRANT_TYPE.Descriptor instead.
func (GRANT_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{0}
}

// 令牌类型
type TOKEN_TYPE int32

const (
	// Bearer Token: xxxx
	TOKEN_TYPE_BEARER TOKEN_TYPE = 0
	// 基于Mac的Token
	TOKEN_TYPE_MAC TOKEN_TYPE = 1
	// Json Web Token
	TOKEN_TYPE_JWT TOKEN_TYPE = 2
)

// Enum value maps for TOKEN_TYPE.
var (
	TOKEN_TYPE_name = map[int32]string{
		0: "BEARER",
		1: "MAC",
		2: "JWT",
	}
	TOKEN_TYPE_value = map[string]int32{
		"BEARER": 0,
		"MAC":    1,
		"JWT":    2,
	}
)

func (x TOKEN_TYPE) Enum() *TOKEN_TYPE {
	p := new(TOKEN_TYPE)
	*p = x
	return p
}

func (x TOKEN_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TOKEN_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_token_proto_enumTypes[1].Descriptor()
}

func (TOKEN_TYPE) Type() protoreflect.EnumType {
	return &file_apps_token_pb_token_proto_enumTypes[1]
}

func (x TOKEN_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TOKEN_TYPE.Descriptor instead.
func (TOKEN_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{1}
}

// 冻结类型
type BLOCK_TYPE int32

const (
	// 刷新Token过期, 回话中断
	BLOCK_TYPE_REFRESH_TOKEN_EXPIRED BLOCK_TYPE = 0
	// 异地登陆
	BLOCK_TYPE_OTHER_PLACE_LOGGED_IN BLOCK_TYPE = 1
	// 异常Ip登陆
	BLOCK_TYPE_OTHER_IP_LOGGED_IN BLOCK_TYPE = 2
)

// Enum value maps for BLOCK_TYPE.
var (
	BLOCK_TYPE_name = map[int32]string{
		0: "REFRESH_TOKEN_EXPIRED",
		1: "OTHER_PLACE_LOGGED_IN",
		2: "OTHER_IP_LOGGED_IN",
	}
	BLOCK_TYPE_value = map[string]int32{
		"REFRESH_TOKEN_EXPIRED": 0,
		"OTHER_PLACE_LOGGED_IN": 1,
		"OTHER_IP_LOGGED_IN":    2,
	}
)

func (x BLOCK_TYPE) Enum() *BLOCK_TYPE {
	p := new(BLOCK_TYPE)
	*p = x
	return p
}

func (x BLOCK_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BLOCK_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_token_pb_token_proto_enumTypes[2].Descriptor()
}

func (BLOCK_TYPE) Type() protoreflect.EnumType {
	return &file_apps_token_pb_token_proto_enumTypes[2]
}

func (x BLOCK_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BLOCK_TYPE.Descriptor instead.
func (BLOCK_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{2}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bson:"_id" json:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"_id"`
	// @gotags: bson:"refresh_token" json:"refresh_token"
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token" bson:"refresh_token"`
	// 颁发时间
	// @gotags: bson:"issue_at" json:"issue_at"
	IssueAt int64 `protobuf:"varint,3,opt,name=issue_at,json=issueAt,proto3" json:"issue_at" bson:"issue_at"`
	// 访问令牌过期时间
	// @gotags: bson:"access_expired_at" json:"access_expired_at"
	AccessExpiredAt int64 `protobuf:"varint,4,opt,name=access_expired_at,json=accessExpiredAt,proto3" json:"access_expired_at" bson:"access_expired_at"`
	// 刷新令牌过期时间
	// @gotags: bson:"refresh_expired_at" json:"refresh_expired_at"
	RefreshExpiredAt int64 `protobuf:"varint,5,opt,name=refresh_expired_at,json=refreshExpiredAt,proto3" json:"refresh_expired_at" bson:"refresh_expired_at"`
	// 用户名
	// @gotags: bson:"username" json:"username"
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username" bson:"username"`
	// 用户Id
	// @gotags: bson:"user_id" json:"user_id"
	UserId string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id" bson:"user_id"`
	// 令牌类型
	// @gotags: bson:"type" json:"type"
	Type TOKEN_TYPE `protobuf:"varint,12,opt,name=type,proto3,enum=devcloud.mcenter.token.TOKEN_TYPE" json:"type" bson:"type"`
	// 令牌状态
	// @gotags: bson:"status" json:"status,omitempty"
	Status *Status `protobuf:"bytes,13,opt,name=status,proto3" json:"status,omitempty" bson:"status"`
	// 其他信息
	// @gotags: bson:"meta" json:"meta,omitempty"
	Meta map[string]string `protobuf:"bytes,14,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"meta"`
	// 授权类型
	// @gotags: bson:"grant_type" json:"grant_type"
	GrantType GRANT_TYPE `protobuf:"varint,11,opt,name=grant_type,json=grantType,proto3,enum=devcloud.mcenter.token.GRANT_TYPE" json:"grant_type" bson:"grant_type"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetIssueAt() int64 {
	if x != nil {
		return x.IssueAt
	}
	return 0
}

func (x *Token) GetAccessExpiredAt() int64 {
	if x != nil {
		return x.AccessExpiredAt
	}
	return 0
}

func (x *Token) GetRefreshExpiredAt() int64 {
	if x != nil {
		return x.RefreshExpiredAt
	}
	return 0
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Token) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Token) GetType() TOKEN_TYPE {
	if x != nil {
		return x.Type
	}
	return TOKEN_TYPE_BEARER
}

func (x *Token) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Token) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Token) GetGrantType() GRANT_TYPE {
	if x != nil {
		return x.GrantType
	}
	return GRANT_TYPE_PASSWORD
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否冻结
	// @gotags: bson:"is_block" json:"is_block"
	IsBlock bool `protobuf:"varint,1,opt,name=is_block,json=isBlock,proto3" json:"is_block" bson:"is_block"`
	// 冻结类型
	// @gotags: bson:"block_type" json:"block_type"
	BlockType BLOCK_TYPE `protobuf:"varint,2,opt,name=block_type,json=blockType,proto3,enum=devcloud.mcenter.token.BLOCK_TYPE" json:"block_type" bson:"block_type"`
	// 冻结时间
	// @gotags: bson:"block_at" json:"block_at"
	BlockAt int64 `protobuf:"varint,3,opt,name=block_at,json=blockAt,proto3" json:"block_at" bson:"block_at"`
	// 冻结原因
	// @gotags: bson:"block_reason" json:"block_reason"
	BlockReason string `protobuf:"bytes,4,opt,name=block_reason,json=blockReason,proto3" json:"block_reason" bson:"block_reason"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_token_pb_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_token_pb_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_token_pb_token_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetIsBlock() bool {
	if x != nil {
		return x.IsBlock
	}
	return false
}

func (x *Status) GetBlockType() BLOCK_TYPE {
	if x != nil {
		return x.BlockType
	}
	return BLOCK_TYPE_REFRESH_TOKEN_EXPIRED
}

func (x *Status) GetBlockAt() int64 {
	if x != nil {
		return x.BlockAt
	}
	return 0
}

func (x *Status) GetBlockReason() string {
	if x != nil {
		return x.BlockReason
	}
	return ""
}

var File_apps_token_pb_token_proto protoreflect.FileDescriptor

var file_apps_token_pb_token_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x65, 0x76,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xa2, 0x04, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x41, 0x74,
	0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x12,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x0a,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a,
	0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x41,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a,
	0x3e, 0x0a, 0x0a, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0c, 0x0a,
	0x08, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x44, 0x41, 0x50, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x49, 0x53, 0x48, 0x55, 0x10,
	0x08, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x4e, 0x47, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x2a,
	0x2a, 0x0a, 0x0a, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0a, 0x0a,
	0x06, 0x42, 0x45, 0x41, 0x52, 0x45, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x43,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x10, 0x02, 0x2a, 0x5a, 0x0a, 0x0a, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x46,
	0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x50, 0x4c,
	0x41, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x49, 0x50, 0x5f, 0x4c, 0x4f, 0x47, 0x47,
	0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x02, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x7a, 0x69, 0x68, 0x75, 0x61, 0x2f, 0x64, 0x65,
	0x76, 0x70, 0x6c, 0x61, 0x74, 0x2f, 0x6d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_token_pb_token_proto_rawDescOnce sync.Once
	file_apps_token_pb_token_proto_rawDescData = file_apps_token_pb_token_proto_rawDesc
)

func file_apps_token_pb_token_proto_rawDescGZIP() []byte {
	file_apps_token_pb_token_proto_rawDescOnce.Do(func() {
		file_apps_token_pb_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_token_pb_token_proto_rawDescData)
	})
	return file_apps_token_pb_token_proto_rawDescData
}

var file_apps_token_pb_token_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_apps_token_pb_token_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_token_pb_token_proto_goTypes = []interface{}{
	(GRANT_TYPE)(0), // 0: devcloud.mcenter.token.GRANT_TYPE
	(TOKEN_TYPE)(0), // 1: devcloud.mcenter.token.TOKEN_TYPE
	(BLOCK_TYPE)(0), // 2: devcloud.mcenter.token.BLOCK_TYPE
	(*Token)(nil),   // 3: devcloud.mcenter.token.Token
	(*Status)(nil),  // 4: devcloud.mcenter.token.Status
	nil,             // 5: devcloud.mcenter.token.Token.MetaEntry
}
var file_apps_token_pb_token_proto_depIdxs = []int32{
	1, // 0: devcloud.mcenter.token.Token.type:type_name -> devcloud.mcenter.token.TOKEN_TYPE
	4, // 1: devcloud.mcenter.token.Token.status:type_name -> devcloud.mcenter.token.Status
	5, // 2: devcloud.mcenter.token.Token.meta:type_name -> devcloud.mcenter.token.Token.MetaEntry
	0, // 3: devcloud.mcenter.token.Token.grant_type:type_name -> devcloud.mcenter.token.GRANT_TYPE
	2, // 4: devcloud.mcenter.token.Status.block_type:type_name -> devcloud.mcenter.token.BLOCK_TYPE
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_apps_token_pb_token_proto_init() }
func file_apps_token_pb_token_proto_init() {
	if File_apps_token_pb_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_token_pb_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_apps_token_pb_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_apps_token_pb_token_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_token_pb_token_proto_goTypes,
		DependencyIndexes: file_apps_token_pb_token_proto_depIdxs,
		EnumInfos:         file_apps_token_pb_token_proto_enumTypes,
		MessageInfos:      file_apps_token_pb_token_proto_msgTypes,
	}.Build()
	File_apps_token_pb_token_proto = out.File
	file_apps_token_pb_token_proto_rawDesc = nil
	file_apps_token_pb_token_proto_goTypes = nil
	file_apps_token_pb_token_proto_depIdxs = nil
}
