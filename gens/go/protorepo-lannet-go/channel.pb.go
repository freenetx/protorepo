// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: lannet/channel.proto

package lannetpb

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

type HelloServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key              string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	MyIps            []string `protobuf:"bytes,2,rep,name=my_ips,json=myIps,proto3" json:"my_ips,omitempty"`
	MyEncodedAddress string   `protobuf:"bytes,3,opt,name=my_encoded_address,json=myEncodedAddress,proto3" json:"my_encoded_address,omitempty"`
}

func (x *HelloServer) Reset() {
	*x = HelloServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloServer) ProtoMessage() {}

func (x *HelloServer) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloServer.ProtoReflect.Descriptor instead.
func (*HelloServer) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{0}
}

func (x *HelloServer) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HelloServer) GetMyIps() []string {
	if x != nil {
		return x.MyIps
	}
	return nil
}

func (x *HelloServer) GetMyEncodedAddress() string {
	if x != nil {
		return x.MyEncodedAddress
	}
	return ""
}

type CenterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//
	//	*CenterRequest_Hello
	//	*CenterRequest_Data
	Req isCenterRequest_Req `protobuf_oneof:"Req"`
}

func (x *CenterRequest) Reset() {
	*x = CenterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CenterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CenterRequest) ProtoMessage() {}

func (x *CenterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CenterRequest.ProtoReflect.Descriptor instead.
func (*CenterRequest) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{1}
}

func (m *CenterRequest) GetReq() isCenterRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *CenterRequest) GetHello() *HelloServer {
	if x, ok := x.GetReq().(*CenterRequest_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *CenterRequest) GetData() []byte {
	if x, ok := x.GetReq().(*CenterRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isCenterRequest_Req interface {
	isCenterRequest_Req()
}

type CenterRequest_Hello struct {
	Hello *HelloServer `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type CenterRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*CenterRequest_Hello) isCenterRequest_Req() {}

func (*CenterRequest_Data) isCenterRequest_Req() {}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Strings []string `protobuf:"bytes,1,rep,name=strings,proto3" json:"strings,omitempty"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{2}
}

func (x *StringList) GetStrings() []string {
	if x != nil {
		return x.Strings
	}
	return nil
}

type HelloClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients   map[string]string      `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClientIps map[string]*StringList `protobuf:"bytes,2,rep,name=client_ips,json=clientIps,proto3" json:"client_ips,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Vpn       string                 `protobuf:"bytes,3,opt,name=vpn,proto3" json:"vpn,omitempty"`
}

func (x *HelloClient) Reset() {
	*x = HelloClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloClient) ProtoMessage() {}

func (x *HelloClient) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloClient.ProtoReflect.Descriptor instead.
func (*HelloClient) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{3}
}

func (x *HelloClient) GetClients() map[string]string {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *HelloClient) GetClientIps() map[string]*StringList {
	if x != nil {
		return x.ClientIps
	}
	return nil
}

func (x *HelloClient) GetVpn() string {
	if x != nil {
		return x.Vpn
	}
	return ""
}

type CenterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//
	//	*CenterResponse_Hello
	//	*CenterResponse_Data
	Resp isCenterResponse_Resp `protobuf_oneof:"Resp"`
}

func (x *CenterResponse) Reset() {
	*x = CenterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CenterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CenterResponse) ProtoMessage() {}

func (x *CenterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CenterResponse.ProtoReflect.Descriptor instead.
func (*CenterResponse) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{4}
}

func (m *CenterResponse) GetResp() isCenterResponse_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *CenterResponse) GetHello() *HelloClient {
	if x, ok := x.GetResp().(*CenterResponse_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *CenterResponse) GetData() []byte {
	if x, ok := x.GetResp().(*CenterResponse_Data); ok {
		return x.Data
	}
	return nil
}

type isCenterResponse_Resp interface {
	isCenterResponse_Resp()
}

type CenterResponse_Hello struct {
	Hello *HelloClient `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type CenterResponse_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*CenterResponse_Hello) isCenterResponse_Resp() {}

func (*CenterResponse_Data) isCenterResponse_Resp() {}

type DirectNetHelloServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DirectNetHelloServer) Reset() {
	*x = DirectNetHelloServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectNetHelloServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectNetHelloServer) ProtoMessage() {}

func (x *DirectNetHelloServer) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectNetHelloServer.ProtoReflect.Descriptor instead.
func (*DirectNetHelloServer) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{5}
}

func (x *DirectNetHelloServer) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DirectNetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Req:
	//
	//	*DirectNetRequest_Hello
	//	*DirectNetRequest_Data
	Req isDirectNetRequest_Req `protobuf_oneof:"Req"`
}

func (x *DirectNetRequest) Reset() {
	*x = DirectNetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectNetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectNetRequest) ProtoMessage() {}

func (x *DirectNetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectNetRequest.ProtoReflect.Descriptor instead.
func (*DirectNetRequest) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{6}
}

func (m *DirectNetRequest) GetReq() isDirectNetRequest_Req {
	if m != nil {
		return m.Req
	}
	return nil
}

func (x *DirectNetRequest) GetHello() *DirectNetHelloServer {
	if x, ok := x.GetReq().(*DirectNetRequest_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *DirectNetRequest) GetData() []byte {
	if x, ok := x.GetReq().(*DirectNetRequest_Data); ok {
		return x.Data
	}
	return nil
}

type isDirectNetRequest_Req interface {
	isDirectNetRequest_Req()
}

type DirectNetRequest_Hello struct {
	Hello *DirectNetHelloServer `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type DirectNetRequest_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*DirectNetRequest_Hello) isDirectNetRequest_Req() {}

func (*DirectNetRequest_Data) isDirectNetRequest_Req() {}

type DirectNetHelloClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DirectNetHelloClient) Reset() {
	*x = DirectNetHelloClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectNetHelloClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectNetHelloClient) ProtoMessage() {}

func (x *DirectNetHelloClient) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectNetHelloClient.ProtoReflect.Descriptor instead.
func (*DirectNetHelloClient) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{7}
}

type DirectNetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Resp:
	//
	//	*DirectNetResponse_Hello
	//	*DirectNetResponse_Data
	Resp isDirectNetResponse_Resp `protobuf_oneof:"Resp"`
}

func (x *DirectNetResponse) Reset() {
	*x = DirectNetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lannet_channel_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectNetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectNetResponse) ProtoMessage() {}

func (x *DirectNetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lannet_channel_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectNetResponse.ProtoReflect.Descriptor instead.
func (*DirectNetResponse) Descriptor() ([]byte, []int) {
	return file_lannet_channel_proto_rawDescGZIP(), []int{8}
}

func (m *DirectNetResponse) GetResp() isDirectNetResponse_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *DirectNetResponse) GetHello() *DirectNetHelloClient {
	if x, ok := x.GetResp().(*DirectNetResponse_Hello); ok {
		return x.Hello
	}
	return nil
}

func (x *DirectNetResponse) GetData() []byte {
	if x, ok := x.GetResp().(*DirectNetResponse_Data); ok {
		return x.Data
	}
	return nil
}

type isDirectNetResponse_Resp interface {
	isDirectNetResponse_Resp()
}

type DirectNetResponse_Hello struct {
	Hello *DirectNetHelloClient `protobuf:"bytes,1,opt,name=hello,proto3,oneof"`
}

type DirectNetResponse_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

func (*DirectNetResponse_Hello) isDirectNetResponse_Resp() {}

func (*DirectNetResponse_Data) isDirectNetResponse_Resp() {}

var File_lannet_channel_proto protoreflect.FileDescriptor

var file_lannet_channel_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x22, 0x64, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x79, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x79, 0x49, 0x70, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x79,
	0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x79, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5d, 0x0a, 0x0d, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x22, 0x26, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x22,
	0xb8, 0x02, 0x0a, 0x0b, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x3e, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x45, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x70, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x70, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x70, 0x6e, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x54, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5f, 0x0a, 0x0e, 0x43, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6c, 0x61,
	0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x22, 0x28, 0x0a, 0x14, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x69, 0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71,
	0x22, 0x16, 0x0a, 0x14, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x11, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x4e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6c,
	0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x4e, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a,
	0x04, 0x52, 0x65, 0x73, 0x70, 0x32, 0x57, 0x0a, 0x0d, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6c,
	0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x60,
	0x0a, 0x10, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x2e,
	0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x4e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6c, 0x61,
	0x6e, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x70,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72,
	0x65, 0x70, 0x6f, 0x2d, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x74, 0x2d, 0x67, 0x6f, 0x3b, 0x6c, 0x61,
	0x6e, 0x6e, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lannet_channel_proto_rawDescOnce sync.Once
	file_lannet_channel_proto_rawDescData = file_lannet_channel_proto_rawDesc
)

func file_lannet_channel_proto_rawDescGZIP() []byte {
	file_lannet_channel_proto_rawDescOnce.Do(func() {
		file_lannet_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_lannet_channel_proto_rawDescData)
	})
	return file_lannet_channel_proto_rawDescData
}

var file_lannet_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_lannet_channel_proto_goTypes = []interface{}{
	(*HelloServer)(nil),          // 0: lannet.api.HelloServer
	(*CenterRequest)(nil),        // 1: lannet.api.CenterRequest
	(*StringList)(nil),           // 2: lannet.api.StringList
	(*HelloClient)(nil),          // 3: lannet.api.HelloClient
	(*CenterResponse)(nil),       // 4: lannet.api.CenterResponse
	(*DirectNetHelloServer)(nil), // 5: lannet.api.DirectNetHelloServer
	(*DirectNetRequest)(nil),     // 6: lannet.api.DirectNetRequest
	(*DirectNetHelloClient)(nil), // 7: lannet.api.DirectNetHelloClient
	(*DirectNetResponse)(nil),    // 8: lannet.api.DirectNetResponse
	nil,                          // 9: lannet.api.HelloClient.ClientsEntry
	nil,                          // 10: lannet.api.HelloClient.ClientIpsEntry
}
var file_lannet_channel_proto_depIdxs = []int32{
	0,  // 0: lannet.api.CenterRequest.hello:type_name -> lannet.api.HelloServer
	9,  // 1: lannet.api.HelloClient.clients:type_name -> lannet.api.HelloClient.ClientsEntry
	10, // 2: lannet.api.HelloClient.client_ips:type_name -> lannet.api.HelloClient.ClientIpsEntry
	3,  // 3: lannet.api.CenterResponse.hello:type_name -> lannet.api.HelloClient
	5,  // 4: lannet.api.DirectNetRequest.hello:type_name -> lannet.api.DirectNetHelloServer
	7,  // 5: lannet.api.DirectNetResponse.hello:type_name -> lannet.api.DirectNetHelloClient
	2,  // 6: lannet.api.HelloClient.ClientIpsEntry.value:type_name -> lannet.api.StringList
	1,  // 7: lannet.api.CenterService.Channel:input_type -> lannet.api.CenterRequest
	6,  // 8: lannet.api.DirectNetService.Channel:input_type -> lannet.api.DirectNetRequest
	4,  // 9: lannet.api.CenterService.Channel:output_type -> lannet.api.CenterResponse
	8,  // 10: lannet.api.DirectNetService.Channel:output_type -> lannet.api.DirectNetResponse
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_lannet_channel_proto_init() }
func file_lannet_channel_proto_init() {
	if File_lannet_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lannet_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloServer); i {
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
		file_lannet_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CenterRequest); i {
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
		file_lannet_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringList); i {
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
		file_lannet_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloClient); i {
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
		file_lannet_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CenterResponse); i {
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
		file_lannet_channel_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectNetHelloServer); i {
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
		file_lannet_channel_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectNetRequest); i {
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
		file_lannet_channel_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectNetHelloClient); i {
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
		file_lannet_channel_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectNetResponse); i {
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
	file_lannet_channel_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CenterRequest_Hello)(nil),
		(*CenterRequest_Data)(nil),
	}
	file_lannet_channel_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*CenterResponse_Hello)(nil),
		(*CenterResponse_Data)(nil),
	}
	file_lannet_channel_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*DirectNetRequest_Hello)(nil),
		(*DirectNetRequest_Data)(nil),
	}
	file_lannet_channel_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*DirectNetResponse_Hello)(nil),
		(*DirectNetResponse_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lannet_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_lannet_channel_proto_goTypes,
		DependencyIndexes: file_lannet_channel_proto_depIdxs,
		MessageInfos:      file_lannet_channel_proto_msgTypes,
	}.Build()
	File_lannet_channel_proto = out.File
	file_lannet_channel_proto_rawDesc = nil
	file_lannet_channel_proto_goTypes = nil
	file_lannet_channel_proto_depIdxs = nil
}
