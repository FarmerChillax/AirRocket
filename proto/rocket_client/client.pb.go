// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/client.proto

package rocket_client

import (
	context "context"
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

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URI string `protobuf:"bytes,1,opt,name=URI,proto3" json:"URI,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_client_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetURI() string {
	if x != nil {
		return x.URI
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_proto_client_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_client_proto protoreflect.FileDescriptor

var file_proto_client_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x23, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x49, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x49, 0x22, 0x26, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x60, 0x0a, 0x0f, 0x41, 0x69, 0x72, 0x52, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12,
	0x1e, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_client_proto_rawDescOnce sync.Once
	file_proto_client_proto_rawDescData = file_proto_client_proto_rawDesc
)

func file_proto_client_proto_rawDescGZIP() []byte {
	file_proto_client_proto_rawDescOnce.Do(func() {
		file_proto_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_client_proto_rawDescData)
	})
	return file_proto_client_proto_rawDescData
}

var file_proto_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_client_proto_goTypes = []interface{}{
	(*TransferRequest)(nil),  // 0: rocket_client.TransferRequest
	(*TransferResponse)(nil), // 1: rocket_client.TransferResponse
}
var file_proto_client_proto_depIdxs = []int32{
	0, // 0: rocket_client.AirRocketClient.Transfer:input_type -> rocket_client.TransferRequest
	1, // 1: rocket_client.AirRocketClient.Transfer:output_type -> rocket_client.TransferResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_client_proto_init() }
func file_proto_client_proto_init() {
	if File_proto_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_proto_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
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
			RawDescriptor: file_proto_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_client_proto_goTypes,
		DependencyIndexes: file_proto_client_proto_depIdxs,
		MessageInfos:      file_proto_client_proto_msgTypes,
	}.Build()
	File_proto_client_proto = out.File
	file_proto_client_proto_rawDesc = nil
	file_proto_client_proto_goTypes = nil
	file_proto_client_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AirRocketClientClient is the client API for AirRocketClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AirRocketClientClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (AirRocketClient_TransferClient, error)
}

type airRocketClientClient struct {
	cc grpc.ClientConnInterface
}

func NewAirRocketClientClient(cc grpc.ClientConnInterface) AirRocketClientClient {
	return &airRocketClientClient{cc}
}

func (c *airRocketClientClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (AirRocketClient_TransferClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AirRocketClient_serviceDesc.Streams[0], "/rocket_client.AirRocketClient/Transfer", opts...)
	if err != nil {
		return nil, err
	}
	x := &airRocketClientTransferClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AirRocketClient_TransferClient interface {
	Recv() (*TransferResponse, error)
	grpc.ClientStream
}

type airRocketClientTransferClient struct {
	grpc.ClientStream
}

func (x *airRocketClientTransferClient) Recv() (*TransferResponse, error) {
	m := new(TransferResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AirRocketClientServer is the server API for AirRocketClient service.
type AirRocketClientServer interface {
	Transfer(*TransferRequest, AirRocketClient_TransferServer) error
}

// UnimplementedAirRocketClientServer can be embedded to have forward compatible implementations.
type UnimplementedAirRocketClientServer struct {
}

func (*UnimplementedAirRocketClientServer) Transfer(*TransferRequest, AirRocketClient_TransferServer) error {
	return status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterAirRocketClientServer(s *grpc.Server, srv AirRocketClientServer) {
	s.RegisterService(&_AirRocketClient_serviceDesc, srv)
}

func _AirRocketClient_Transfer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransferRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AirRocketClientServer).Transfer(m, &airRocketClientTransferServer{stream})
}

type AirRocketClient_TransferServer interface {
	Send(*TransferResponse) error
	grpc.ServerStream
}

type airRocketClientTransferServer struct {
	grpc.ServerStream
}

func (x *airRocketClientTransferServer) Send(m *TransferResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AirRocketClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rocket_client.AirRocketClient",
	HandlerType: (*AirRocketClientServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transfer",
			Handler:       _AirRocketClient_Transfer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/client.proto",
}
