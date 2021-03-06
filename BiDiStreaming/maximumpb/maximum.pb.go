// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: maximumpb/maximum.proto

package maximumpb

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

type Maximum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Maximum) Reset() {
	*x = Maximum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximumpb_maximum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Maximum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Maximum) ProtoMessage() {}

func (x *Maximum) ProtoReflect() protoreflect.Message {
	mi := &file_maximumpb_maximum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Maximum.ProtoReflect.Descriptor instead.
func (*Maximum) Descriptor() ([]byte, []int) {
	return file_maximumpb_maximum_proto_rawDescGZIP(), []int{0}
}

func (x *Maximum) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum *Maximum `protobuf:"bytes,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *MaximumRequest) Reset() {
	*x = MaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximumpb_maximum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumRequest) ProtoMessage() {}

func (x *MaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maximumpb_maximum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumRequest.ProtoReflect.Descriptor instead.
func (*MaximumRequest) Descriptor() ([]byte, []int) {
	return file_maximumpb_maximum_proto_rawDescGZIP(), []int{1}
}

func (x *MaximumRequest) GetMaximum() *Maximum {
	if x != nil {
		return x.Maximum
	}
	return nil
}

type MaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max int32 `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *MaximumResponse) Reset() {
	*x = MaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maximumpb_maximum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaximumResponse) ProtoMessage() {}

func (x *MaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_maximumpb_maximum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaximumResponse.ProtoReflect.Descriptor instead.
func (*MaximumResponse) Descriptor() ([]byte, []int) {
	return file_maximumpb_maximum_proto_rawDescGZIP(), []int{2}
}

func (x *MaximumResponse) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_maximumpb_maximum_proto protoreflect.FileDescriptor

var file_maximumpb_maximum_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x22, 0x21, 0x0a, 0x07, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x22, 0x23, 0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x32, 0x56, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maximumpb_maximum_proto_rawDescOnce sync.Once
	file_maximumpb_maximum_proto_rawDescData = file_maximumpb_maximum_proto_rawDesc
)

func file_maximumpb_maximum_proto_rawDescGZIP() []byte {
	file_maximumpb_maximum_proto_rawDescOnce.Do(func() {
		file_maximumpb_maximum_proto_rawDescData = protoimpl.X.CompressGZIP(file_maximumpb_maximum_proto_rawDescData)
	})
	return file_maximumpb_maximum_proto_rawDescData
}

var file_maximumpb_maximum_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_maximumpb_maximum_proto_goTypes = []interface{}{
	(*Maximum)(nil),         // 0: maximum.Maximum
	(*MaximumRequest)(nil),  // 1: maximum.MaximumRequest
	(*MaximumResponse)(nil), // 2: maximum.MaximumResponse
}
var file_maximumpb_maximum_proto_depIdxs = []int32{
	0, // 0: maximum.MaximumRequest.maximum:type_name -> maximum.Maximum
	1, // 1: maximum.MaximumService.FindMaximum:input_type -> maximum.MaximumRequest
	2, // 2: maximum.MaximumService.FindMaximum:output_type -> maximum.MaximumResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_maximumpb_maximum_proto_init() }
func file_maximumpb_maximum_proto_init() {
	if File_maximumpb_maximum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maximumpb_maximum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Maximum); i {
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
		file_maximumpb_maximum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumRequest); i {
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
		file_maximumpb_maximum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaximumResponse); i {
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
			RawDescriptor: file_maximumpb_maximum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maximumpb_maximum_proto_goTypes,
		DependencyIndexes: file_maximumpb_maximum_proto_depIdxs,
		MessageInfos:      file_maximumpb_maximum_proto_msgTypes,
	}.Build()
	File_maximumpb_maximum_proto = out.File
	file_maximumpb_maximum_proto_rawDesc = nil
	file_maximumpb_maximum_proto_goTypes = nil
	file_maximumpb_maximum_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaximumServiceClient is the client API for MaximumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaximumServiceClient interface {
	//Bidi Streaming
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (MaximumService_FindMaximumClient, error)
}

type maximumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaximumServiceClient(cc grpc.ClientConnInterface) MaximumServiceClient {
	return &maximumServiceClient{cc}
}

func (c *maximumServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (MaximumService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaximumService_serviceDesc.Streams[0], "/maximum.MaximumService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &maximumServiceFindMaximumClient{stream}
	return x, nil
}

type MaximumService_FindMaximumClient interface {
	Send(*MaximumRequest) error
	Recv() (*MaximumResponse, error)
	grpc.ClientStream
}

type maximumServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *maximumServiceFindMaximumClient) Send(m *MaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maximumServiceFindMaximumClient) Recv() (*MaximumResponse, error) {
	m := new(MaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaximumServiceServer is the server API for MaximumService service.
type MaximumServiceServer interface {
	//Bidi Streaming
	FindMaximum(MaximumService_FindMaximumServer) error
}

// UnimplementedMaximumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMaximumServiceServer struct {
}

func (*UnimplementedMaximumServiceServer) FindMaximum(MaximumService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterMaximumServiceServer(s *grpc.Server, srv MaximumServiceServer) {
	s.RegisterService(&_MaximumService_serviceDesc, srv)
}

func _MaximumService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaximumServiceServer).FindMaximum(&maximumServiceFindMaximumServer{stream})
}

type MaximumService_FindMaximumServer interface {
	Send(*MaximumResponse) error
	Recv() (*MaximumRequest, error)
	grpc.ServerStream
}

type maximumServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *maximumServiceFindMaximumServer) Send(m *MaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maximumServiceFindMaximumServer) Recv() (*MaximumRequest, error) {
	m := new(MaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaximumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maximum.MaximumService",
	HandlerType: (*MaximumServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximum",
			Handler:       _MaximumService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "maximumpb/maximum.proto",
}
