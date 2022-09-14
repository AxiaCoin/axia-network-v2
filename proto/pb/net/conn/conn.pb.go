// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: net/conn/conn.proto

package conn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// length of the request in bytes
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_conn_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// read is the payload in bytes
	Read []byte `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	// error is an error message
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// errored is true if an error has been set
	Errored bool `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_conn_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetRead() []byte {
	if x != nil {
		return x.Read
	}
	return nil
}

func (x *ReadResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ReadResponse) GetErrored() bool {
	if x != nil {
		return x.Errored
	}
	return false
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// payload is the write request in bytes
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_conn_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// length of the response in bytes
	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// error is an error message
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// errored is true if an error has been set
	Errored bool `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_conn_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{3}
}

func (x *WriteResponse) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *WriteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *WriteResponse) GetErrored() bool {
	if x != nil {
		return x.Errored
	}
	return false
}

type SetDeadlineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// time represents an instant in time in bytes
	Time []byte `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *SetDeadlineRequest) Reset() {
	*x = SetDeadlineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_net_conn_conn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDeadlineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeadlineRequest) ProtoMessage() {}

func (x *SetDeadlineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_net_conn_conn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDeadlineRequest.ProtoReflect.Descriptor instead.
func (*SetDeadlineRequest) Descriptor() ([]byte, []int) {
	return file_net_conn_conn_proto_rawDescGZIP(), []int{4}
}

func (x *SetDeadlineRequest) GetTime() []byte {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_net_conn_conn_proto protoreflect.FileDescriptor

var file_net_conn_conn_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x52, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x57, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x32, 0x88, 0x03, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x35, 0x0a,
	0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e,
	0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76,
	0x61, 0x2d, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x65, 0x74, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_net_conn_conn_proto_rawDescOnce sync.Once
	file_net_conn_conn_proto_rawDescData = file_net_conn_conn_proto_rawDesc
)

func file_net_conn_conn_proto_rawDescGZIP() []byte {
	file_net_conn_conn_proto_rawDescOnce.Do(func() {
		file_net_conn_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_net_conn_conn_proto_rawDescData)
	})
	return file_net_conn_conn_proto_rawDescData
}

var file_net_conn_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_net_conn_conn_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),        // 0: net.conn.ReadRequest
	(*ReadResponse)(nil),       // 1: net.conn.ReadResponse
	(*WriteRequest)(nil),       // 2: net.conn.WriteRequest
	(*WriteResponse)(nil),      // 3: net.conn.WriteResponse
	(*SetDeadlineRequest)(nil), // 4: net.conn.SetDeadlineRequest
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_net_conn_conn_proto_depIdxs = []int32{
	0, // 0: net.conn.Conn.Read:input_type -> net.conn.ReadRequest
	2, // 1: net.conn.Conn.Write:input_type -> net.conn.WriteRequest
	5, // 2: net.conn.Conn.Close:input_type -> google.protobuf.Empty
	4, // 3: net.conn.Conn.SetDeadline:input_type -> net.conn.SetDeadlineRequest
	4, // 4: net.conn.Conn.SetReadDeadline:input_type -> net.conn.SetDeadlineRequest
	4, // 5: net.conn.Conn.SetWriteDeadline:input_type -> net.conn.SetDeadlineRequest
	1, // 6: net.conn.Conn.Read:output_type -> net.conn.ReadResponse
	3, // 7: net.conn.Conn.Write:output_type -> net.conn.WriteResponse
	5, // 8: net.conn.Conn.Close:output_type -> google.protobuf.Empty
	5, // 9: net.conn.Conn.SetDeadline:output_type -> google.protobuf.Empty
	5, // 10: net.conn.Conn.SetReadDeadline:output_type -> google.protobuf.Empty
	5, // 11: net.conn.Conn.SetWriteDeadline:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_net_conn_conn_proto_init() }
func file_net_conn_conn_proto_init() {
	if File_net_conn_conn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_net_conn_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_net_conn_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_net_conn_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_net_conn_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_net_conn_conn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDeadlineRequest); i {
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
			RawDescriptor: file_net_conn_conn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_net_conn_conn_proto_goTypes,
		DependencyIndexes: file_net_conn_conn_proto_depIdxs,
		MessageInfos:      file_net_conn_conn_proto_msgTypes,
	}.Build()
	File_net_conn_conn_proto = out.File
	file_net_conn_conn_proto_rawDesc = nil
	file_net_conn_conn_proto_goTypes = nil
	file_net_conn_conn_proto_depIdxs = nil
}
