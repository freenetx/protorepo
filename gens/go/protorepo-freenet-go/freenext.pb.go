// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.5
// source: proto/freenet/freenext.proto

package freenetpb

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

type SyncMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ForwardIps   []string          `protobuf:"bytes,2,rep,name=forward_ips,json=forwardIps,proto3" json:"forward_ips,omitempty"`
	KnownClients map[string]string `protobuf:"bytes,3,rep,name=known_clients,json=knownClients,proto3" json:"known_clients,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SyncMessage) Reset() {
	*x = SyncMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncMessage) ProtoMessage() {}

func (x *SyncMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncMessage.ProtoReflect.Descriptor instead.
func (*SyncMessage) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{0}
}

func (x *SyncMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SyncMessage) GetForwardIps() []string {
	if x != nil {
		return x.ForwardIps
	}
	return nil
}

func (x *SyncMessage) GetKnownClients() map[string]string {
	if x != nil {
		return x.KnownClients
	}
	return nil
}

type MessageAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageAck) Reset() {
	*x = MessageAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAck) ProtoMessage() {}

func (x *MessageAck) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAck.ProtoReflect.Descriptor instead.
func (*MessageAck) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{1}
}

func (x *MessageAck) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{2}
}

func (x *Data) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClientInitMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LanIps        []string `protobuf:"bytes,2,rep,name=lan_ips,json=lanIps,proto3" json:"lan_ips,omitempty"`
	VpnIps        []string `protobuf:"bytes,3,rep,name=vpn_ips,json=vpnIps,proto3" json:"vpn_ips,omitempty"`
	ClientAddress string   `protobuf:"bytes,4,opt,name=client_address,json=clientAddress,proto3" json:"client_address,omitempty"`
}

func (x *ClientInitMessage) Reset() {
	*x = ClientInitMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInitMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInitMessage) ProtoMessage() {}

func (x *ClientInitMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInitMessage.ProtoReflect.Descriptor instead.
func (*ClientInitMessage) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{3}
}

func (x *ClientInitMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientInitMessage) GetLanIps() []string {
	if x != nil {
		return x.LanIps
	}
	return nil
}

func (x *ClientInitMessage) GetVpnIps() []string {
	if x != nil {
		return x.VpnIps
	}
	return nil
}

func (x *ClientInitMessage) GetClientAddress() string {
	if x != nil {
		return x.ClientAddress
	}
	return ""
}

type NeedClientInitMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NeedClientInitMessage) Reset() {
	*x = NeedClientInitMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeedClientInitMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeedClientInitMessage) ProtoMessage() {}

func (x *NeedClientInitMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeedClientInitMessage.ProtoReflect.Descriptor instead.
func (*NeedClientInitMessage) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{4}
}

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*ClientMessage_Data
	//	*ClientMessage_ClientInitMessage
	Request isClientMessage_Request `protobuf_oneof:"Request"`
	// Types that are assignable to Response:
	//
	//	*ClientMessage_MessageAck
	Response isClientMessage_Response `protobuf_oneof:"Response"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{5}
}

func (m *ClientMessage) GetRequest() isClientMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ClientMessage) GetData() *Data {
	if x, ok := x.GetRequest().(*ClientMessage_Data); ok {
		return x.Data
	}
	return nil
}

func (x *ClientMessage) GetClientInitMessage() *ClientInitMessage {
	if x, ok := x.GetRequest().(*ClientMessage_ClientInitMessage); ok {
		return x.ClientInitMessage
	}
	return nil
}

func (m *ClientMessage) GetResponse() isClientMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ClientMessage) GetMessageAck() *MessageAck {
	if x, ok := x.GetResponse().(*ClientMessage_MessageAck); ok {
		return x.MessageAck
	}
	return nil
}

type isClientMessage_Request interface {
	isClientMessage_Request()
}

type ClientMessage_Data struct {
	Data *Data `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type ClientMessage_ClientInitMessage struct {
	ClientInitMessage *ClientInitMessage `protobuf:"bytes,2,opt,name=client_init_message,json=clientInitMessage,proto3,oneof"`
}

func (*ClientMessage_Data) isClientMessage_Request() {}

func (*ClientMessage_ClientInitMessage) isClientMessage_Request() {}

type isClientMessage_Response interface {
	isClientMessage_Response()
}

type ClientMessage_MessageAck struct {
	MessageAck *MessageAck `protobuf:"bytes,100,opt,name=message_ack,json=messageAck,proto3,oneof"`
}

func (*ClientMessage_MessageAck) isClientMessage_Response() {}

type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*ServerMessage_Data
	//	*ServerMessage_SyncMessage
	//	*ServerMessage_NeedClientInitMessage
	Request isServerMessage_Request `protobuf_oneof:"Request"`
	// Types that are assignable to Response:
	//
	//	*ServerMessage_MessageAck
	Response isServerMessage_Response `protobuf_oneof:"Response"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_freenet_freenext_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_freenet_freenext_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_proto_freenet_freenext_proto_rawDescGZIP(), []int{6}
}

func (m *ServerMessage) GetRequest() isServerMessage_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *ServerMessage) GetData() *Data {
	if x, ok := x.GetRequest().(*ServerMessage_Data); ok {
		return x.Data
	}
	return nil
}

func (x *ServerMessage) GetSyncMessage() *SyncMessage {
	if x, ok := x.GetRequest().(*ServerMessage_SyncMessage); ok {
		return x.SyncMessage
	}
	return nil
}

func (x *ServerMessage) GetNeedClientInitMessage() *NeedClientInitMessage {
	if x, ok := x.GetRequest().(*ServerMessage_NeedClientInitMessage); ok {
		return x.NeedClientInitMessage
	}
	return nil
}

func (m *ServerMessage) GetResponse() isServerMessage_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ServerMessage) GetMessageAck() *MessageAck {
	if x, ok := x.GetResponse().(*ServerMessage_MessageAck); ok {
		return x.MessageAck
	}
	return nil
}

type isServerMessage_Request interface {
	isServerMessage_Request()
}

type ServerMessage_Data struct {
	Data *Data `protobuf:"bytes,1,opt,name=data,proto3,oneof"`
}

type ServerMessage_SyncMessage struct {
	SyncMessage *SyncMessage `protobuf:"bytes,2,opt,name=sync_message,json=syncMessage,proto3,oneof"`
}

type ServerMessage_NeedClientInitMessage struct {
	NeedClientInitMessage *NeedClientInitMessage `protobuf:"bytes,3,opt,name=need_client_init_message,json=needClientInitMessage,proto3,oneof"`
}

func (*ServerMessage_Data) isServerMessage_Request() {}

func (*ServerMessage_SyncMessage) isServerMessage_Request() {}

func (*ServerMessage_NeedClientInitMessage) isServerMessage_Request() {}

type isServerMessage_Response interface {
	isServerMessage_Response()
}

type ServerMessage_MessageAck struct {
	MessageAck *MessageAck `protobuf:"bytes,100,opt,name=message_ack,json=messageAck,proto3,oneof"`
}

func (*ServerMessage_MessageAck) isServerMessage_Response() {}

var File_proto_freenet_freenext_proto protoreflect.FileDescriptor

var file_proto_freenet_freenext_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2f,
	0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x22, 0xd0, 0x01, 0x0a, 0x0b,
	0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x49, 0x70, 0x73, 0x12, 0x4f, 0x0a, 0x0d,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4b, 0x6e,
	0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x3f, 0x0a,
	0x11, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1c,
	0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7c, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x6c, 0x61, 0x6e, 0x49, 0x70, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x70, 0x6e, 0x5f, 0x69, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x70, 0x6e, 0x49, 0x70, 0x73, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x4e, 0x65, 0x65, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xdd, 0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x50, 0x0a, 0x13, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xa9, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x79,
	0x6e, 0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5d, 0x0a, 0x18, 0x6e, 0x65, 0x65,
	0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x72,
	0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x65, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x15, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x69,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x66, 0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x63, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65,
	0x74, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2d, 0x66,
	0x72, 0x65, 0x65, 0x6e, 0x65, 0x74, 0x2d, 0x67, 0x6f, 0x3b, 0x66, 0x72, 0x65, 0x65, 0x6e, 0x65,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_freenet_freenext_proto_rawDescOnce sync.Once
	file_proto_freenet_freenext_proto_rawDescData = file_proto_freenet_freenext_proto_rawDesc
)

func file_proto_freenet_freenext_proto_rawDescGZIP() []byte {
	file_proto_freenet_freenext_proto_rawDescOnce.Do(func() {
		file_proto_freenet_freenext_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_freenet_freenext_proto_rawDescData)
	})
	return file_proto_freenet_freenext_proto_rawDescData
}

var file_proto_freenet_freenext_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_freenet_freenext_proto_goTypes = []interface{}{
	(*SyncMessage)(nil),           // 0: freenet.api.SyncMessage
	(*MessageAck)(nil),            // 1: freenet.api.MessageAck
	(*Data)(nil),                  // 2: freenet.api.Data
	(*ClientInitMessage)(nil),     // 3: freenet.api.ClientInitMessage
	(*NeedClientInitMessage)(nil), // 4: freenet.api.NeedClientInitMessage
	(*ClientMessage)(nil),         // 5: freenet.api.ClientMessage
	(*ServerMessage)(nil),         // 6: freenet.api.ServerMessage
	nil,                           // 7: freenet.api.SyncMessage.KnownClientsEntry
}
var file_proto_freenet_freenext_proto_depIdxs = []int32{
	7, // 0: freenet.api.SyncMessage.known_clients:type_name -> freenet.api.SyncMessage.KnownClientsEntry
	2, // 1: freenet.api.ClientMessage.data:type_name -> freenet.api.Data
	3, // 2: freenet.api.ClientMessage.client_init_message:type_name -> freenet.api.ClientInitMessage
	1, // 3: freenet.api.ClientMessage.message_ack:type_name -> freenet.api.MessageAck
	2, // 4: freenet.api.ServerMessage.data:type_name -> freenet.api.Data
	0, // 5: freenet.api.ServerMessage.sync_message:type_name -> freenet.api.SyncMessage
	4, // 6: freenet.api.ServerMessage.need_client_init_message:type_name -> freenet.api.NeedClientInitMessage
	1, // 7: freenet.api.ServerMessage.message_ack:type_name -> freenet.api.MessageAck
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_freenet_freenext_proto_init() }
func file_proto_freenet_freenext_proto_init() {
	if File_proto_freenet_freenext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_freenet_freenext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncMessage); i {
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
		file_proto_freenet_freenext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageAck); i {
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
		file_proto_freenet_freenext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_freenet_freenext_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInitMessage); i {
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
		file_proto_freenet_freenext_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeedClientInitMessage); i {
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
		file_proto_freenet_freenext_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_proto_freenet_freenext_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
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
	file_proto_freenet_freenext_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ClientMessage_Data)(nil),
		(*ClientMessage_ClientInitMessage)(nil),
		(*ClientMessage_MessageAck)(nil),
	}
	file_proto_freenet_freenext_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*ServerMessage_Data)(nil),
		(*ServerMessage_SyncMessage)(nil),
		(*ServerMessage_NeedClientInitMessage)(nil),
		(*ServerMessage_MessageAck)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_freenet_freenext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_freenet_freenext_proto_goTypes,
		DependencyIndexes: file_proto_freenet_freenext_proto_depIdxs,
		MessageInfos:      file_proto_freenet_freenext_proto_msgTypes,
	}.Build()
	File_proto_freenet_freenext_proto = out.File
	file_proto_freenet_freenext_proto_rawDesc = nil
	file_proto_freenet_freenext_proto_goTypes = nil
	file_proto_freenet_freenext_proto_depIdxs = nil
}
