// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: exchange/exchange.proto

package exchangev1

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber string `protobuf:"bytes,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	ApiKey       string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *RegisterRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentUuid string `protobuf:"bytes,1,opt,name=agent_uuid,json=agentUuid,proto3" json:"agent_uuid,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetAgentUuid() string {
	if x != nil {
		return x.AgentUuid
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentUuid string `protobuf:"bytes,1,opt,name=agent_uuid,json=agentUuid,proto3" json:"agent_uuid,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetAgentUuid() string {
	if x != nil {
		return x.AgentUuid
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Btemp   *float32 `protobuf:"fixed32,2,opt,name=btemp,proto3,oneof" json:"btemp,omitempty"`
	Cputemp *float32 `protobuf:"fixed32,3,opt,name=cputemp,proto3,oneof" json:"cputemp,omitempty"`
	Health  *bool    `protobuf:"varint,4,opt,name=health,proto3,oneof" json:"health,omitempty"`
	Online  *bool    `protobuf:"varint,5,opt,name=online,proto3,oneof" json:"online,omitempty"`
}

func (x *UpdateDataRequest) Reset() {
	*x = UpdateDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDataRequest) ProtoMessage() {}

func (x *UpdateDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateDataRequest) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDataRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdateDataRequest) GetBtemp() float32 {
	if x != nil && x.Btemp != nil {
		return *x.Btemp
	}
	return 0
}

func (x *UpdateDataRequest) GetCputemp() float32 {
	if x != nil && x.Cputemp != nil {
		return *x.Cputemp
	}
	return 0
}

func (x *UpdateDataRequest) GetHealth() bool {
	if x != nil && x.Health != nil {
		return *x.Health
	}
	return false
}

func (x *UpdateDataRequest) GetOnline() bool {
	if x != nil && x.Online != nil {
		return *x.Online
	}
	return false
}

type UpdateDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Errmessage string `protobuf:"bytes,3,opt,name=errmessage,proto3" json:"errmessage,omitempty"`
}

func (x *UpdateDataResponse) Reset() {
	*x = UpdateDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_exchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDataResponse) ProtoMessage() {}

func (x *UpdateDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_exchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDataResponse.ProtoReflect.Descriptor instead.
func (*UpdateDataResponse) Descriptor() ([]byte, []int) {
	return file_exchange_exchange_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDataResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UpdateDataResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateDataResponse) GetErrmessage() string {
	if x != nil {
		return x.Errmessage
	}
	return ""
}

var File_exchange_exchange_proto protoreflect.FileDescriptor

var file_exchange_exchange_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x61,
	0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc7, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x62, 0x74, 0x65, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x05, 0x62, 0x74, 0x65, 0x6d, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x74, 0x65, 0x6d, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x07, 0x63, 0x70, 0x75, 0x74, 0x65, 0x6d, 0x70, 0x88, 0x01,
	0x01, 0x12, 0x1b, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x02, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x03,
	0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x62, 0x74, 0x65, 0x6d, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x70, 0x75, 0x74, 0x65, 0x6d,
	0x70, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x60, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x72, 0x72, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x08, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x16, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a,
	0x6e, 0x6e, 0x6b, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x3b,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_exchange_exchange_proto_rawDescOnce sync.Once
	file_exchange_exchange_proto_rawDescData = file_exchange_exchange_proto_rawDesc
)

func file_exchange_exchange_proto_rawDescGZIP() []byte {
	file_exchange_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_exchange_proto_rawDescData)
	})
	return file_exchange_exchange_proto_rawDescData
}

var file_exchange_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_exchange_exchange_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),    // 0: exchange.RegisterRequest
	(*RegisterResponse)(nil),   // 1: exchange.RegisterResponse
	(*LoginRequest)(nil),       // 2: exchange.LoginRequest
	(*LoginResponse)(nil),      // 3: exchange.loginResponse
	(*UpdateDataRequest)(nil),  // 4: exchange.UpdateDataRequest
	(*UpdateDataResponse)(nil), // 5: exchange.UpdateDataResponse
}
var file_exchange_exchange_proto_depIdxs = []int32{
	0, // 0: exchange.Exchange.Register:input_type -> exchange.RegisterRequest
	2, // 1: exchange.Exchange.Login:input_type -> exchange.LoginRequest
	4, // 2: exchange.Exchange.UpdateData:input_type -> exchange.UpdateDataRequest
	1, // 3: exchange.Exchange.Register:output_type -> exchange.RegisterResponse
	3, // 4: exchange.Exchange.Login:output_type -> exchange.loginResponse
	5, // 5: exchange.Exchange.UpdateData:output_type -> exchange.UpdateDataResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exchange_exchange_proto_init() }
func file_exchange_exchange_proto_init() {
	if File_exchange_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_exchange_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_exchange_exchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_exchange_exchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_exchange_exchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDataRequest); i {
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
		file_exchange_exchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDataResponse); i {
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
	file_exchange_exchange_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_exchange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_exchange_proto_depIdxs,
		MessageInfos:      file_exchange_exchange_proto_msgTypes,
	}.Build()
	File_exchange_exchange_proto = out.File
	file_exchange_exchange_proto_rawDesc = nil
	file_exchange_exchange_proto_goTypes = nil
	file_exchange_exchange_proto_depIdxs = nil
}
