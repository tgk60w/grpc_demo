// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: transfer.proto

package protos

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type WebJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WebJson) Reset() {
	*x = WebJson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebJson) ProtoMessage() {}

func (x *WebJson) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebJson.ProtoReflect.Descriptor instead.
func (*WebJson) Descriptor() ([]byte, []int) {
	return file_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *WebJson) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *WebJson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_transfer_proto protoreflect.FileDescriptor

var file_transfer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x2d, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x4a,
	0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x70, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x4a, 0x73,
	0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x4a,
	0x73, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x57,
	0x65, 0x62, 0x4a, 0x73, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x57, 0x65, 0x62, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_transfer_proto_rawDescOnce sync.Once
	file_transfer_proto_rawDescData = file_transfer_proto_rawDesc
)

func file_transfer_proto_rawDescGZIP() []byte {
	file_transfer_proto_rawDescOnce.Do(func() {
		file_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_proto_rawDescData)
	})
	return file_transfer_proto_rawDescData
}

var file_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transfer_proto_goTypes = []interface{}{
	(*WebJson)(nil), // 0: protos.WebJson
}
var file_transfer_proto_depIdxs = []int32{
	0, // 0: protos.Transfer.WebRequest:input_type -> protos.WebJson
	0, // 1: protos.Transfer.ReturnData:input_type -> protos.WebJson
	0, // 2: protos.Transfer.WebRequest:output_type -> protos.WebJson
	0, // 3: protos.Transfer.ReturnData:output_type -> protos.WebJson
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transfer_proto_init() }
func file_transfer_proto_init() {
	if File_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebJson); i {
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
			RawDescriptor: file_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfer_proto_goTypes,
		DependencyIndexes: file_transfer_proto_depIdxs,
		MessageInfos:      file_transfer_proto_msgTypes,
	}.Build()
	File_transfer_proto = out.File
	file_transfer_proto_rawDesc = nil
	file_transfer_proto_goTypes = nil
	file_transfer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferClient interface {
	//stream  设定grpc模式
	WebRequest(ctx context.Context, in *WebJson, opts ...grpc.CallOption) (Transfer_WebRequestClient, error)
	ReturnData(ctx context.Context, in *WebJson, opts ...grpc.CallOption) (*WebJson, error)
}

type transferClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferClient(cc grpc.ClientConnInterface) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) WebRequest(ctx context.Context, in *WebJson, opts ...grpc.CallOption) (Transfer_WebRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transfer_serviceDesc.Streams[0], "/protos.Transfer/WebRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferWebRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transfer_WebRequestClient interface {
	Recv() (*WebJson, error)
	grpc.ClientStream
}

type transferWebRequestClient struct {
	grpc.ClientStream
}

func (x *transferWebRequestClient) Recv() (*WebJson, error) {
	m := new(WebJson)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transferClient) ReturnData(ctx context.Context, in *WebJson, opts ...grpc.CallOption) (*WebJson, error) {
	out := new(WebJson)
	err := c.cc.Invoke(ctx, "/protos.Transfer/ReturnData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServer is the server API for Transfer service.
type TransferServer interface {
	//stream  设定grpc模式
	WebRequest(*WebJson, Transfer_WebRequestServer) error
	ReturnData(context.Context, *WebJson) (*WebJson, error)
}

// UnimplementedTransferServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (*UnimplementedTransferServer) WebRequest(*WebJson, Transfer_WebRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method WebRequest not implemented")
}
func (*UnimplementedTransferServer) ReturnData(context.Context, *WebJson) (*WebJson, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnData not implemented")
}

func RegisterTransferServer(s *grpc.Server, srv TransferServer) {
	s.RegisterService(&_Transfer_serviceDesc, srv)
}

func _Transfer_WebRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WebJson)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransferServer).WebRequest(m, &transferWebRequestServer{stream})
}

type Transfer_WebRequestServer interface {
	Send(*WebJson) error
	grpc.ServerStream
}

type transferWebRequestServer struct {
	grpc.ServerStream
}

func (x *transferWebRequestServer) Send(m *WebJson) error {
	return x.ServerStream.SendMsg(m)
}

func _Transfer_ReturnData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebJson)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServer).ReturnData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Transfer/ReturnData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServer).ReturnData(ctx, req.(*WebJson))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReturnData",
			Handler:    _Transfer_ReturnData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WebRequest",
			Handler:       _Transfer_WebRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transfer.proto",
}
