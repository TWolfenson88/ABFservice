// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/api.proto

package api

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

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip    string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Pass  string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserData) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UserData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserData) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

type AllowUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AllowUser) Reset() {
	*x = AllowUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowUser) ProtoMessage() {}

func (x *AllowUser) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllowUser.ProtoReflect.Descriptor instead.
func (*AllowUser) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *AllowUser) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type StatusMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StatusMsg) Reset() {
	*x = StatusMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusMsg) ProtoMessage() {}

func (x *StatusMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusMsg.ProtoReflect.Descriptor instead.
func (*StatusMsg) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *StatusMsg) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type IpData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr string `protobuf:"bytes,1,opt,name=IpAddr,proto3" json:"IpAddr,omitempty"`
	List   string `protobuf:"bytes,2,opt,name=List,proto3" json:"List,omitempty"`
}

func (x *IpData) Reset() {
	*x = IpData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpData) ProtoMessage() {}

func (x *IpData) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpData.ProtoReflect.Descriptor instead.
func (*IpData) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *IpData) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *IpData) GetList() string {
	if x != nil {
		return x.List
	}
	return ""
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x44, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x21, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x34, 0x0a, 0x06, 0x49,
	0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xd1, 0x01, 0x0a, 0x0a, 0x41, 0x42, 0x46, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x30, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x0e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x70, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x70, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_api_proto_goTypes = []interface{}{
	(*UserData)(nil),  // 0: chat.UserData
	(*AllowUser)(nil), // 1: chat.AllowUser
	(*StatusMsg)(nil), // 2: chat.StatusMsg
	(*IpData)(nil),    // 3: chat.IpData
}
var file_api_api_proto_depIdxs = []int32{
	0, // 0: chat.ABFService.AuthRequest:input_type -> chat.UserData
	0, // 1: chat.ABFService.ResetBucket:input_type -> chat.UserData
	3, // 2: chat.ABFService.AddToList:input_type -> chat.IpData
	3, // 3: chat.ABFService.RemoveFromList:input_type -> chat.IpData
	1, // 4: chat.ABFService.AuthRequest:output_type -> chat.AllowUser
	2, // 5: chat.ABFService.ResetBucket:output_type -> chat.StatusMsg
	2, // 6: chat.ABFService.AddToList:output_type -> chat.StatusMsg
	2, // 7: chat.ABFService.RemoveFromList:output_type -> chat.StatusMsg
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowUser); i {
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
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusMsg); i {
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
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpData); i {
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
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ABFServiceClient is the client API for ABFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ABFServiceClient interface {
	AuthRequest(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*AllowUser, error)
	ResetBucket(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*StatusMsg, error)
	AddToList(ctx context.Context, in *IpData, opts ...grpc.CallOption) (*StatusMsg, error)
	RemoveFromList(ctx context.Context, in *IpData, opts ...grpc.CallOption) (*StatusMsg, error)
}

type aBFServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewABFServiceClient(cc grpc.ClientConnInterface) ABFServiceClient {
	return &aBFServiceClient{cc}
}

func (c *aBFServiceClient) AuthRequest(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*AllowUser, error) {
	out := new(AllowUser)
	err := c.cc.Invoke(ctx, "/chat.ABFService/AuthRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBFServiceClient) ResetBucket(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/chat.ABFService/ResetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBFServiceClient) AddToList(ctx context.Context, in *IpData, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/chat.ABFService/AddToList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBFServiceClient) RemoveFromList(ctx context.Context, in *IpData, opts ...grpc.CallOption) (*StatusMsg, error) {
	out := new(StatusMsg)
	err := c.cc.Invoke(ctx, "/chat.ABFService/RemoveFromList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABFServiceServer is the server API for ABFService service.
type ABFServiceServer interface {
	AuthRequest(context.Context, *UserData) (*AllowUser, error)
	ResetBucket(context.Context, *UserData) (*StatusMsg, error)
	AddToList(context.Context, *IpData) (*StatusMsg, error)
	RemoveFromList(context.Context, *IpData) (*StatusMsg, error)
}

// UnimplementedABFServiceServer can be embedded to have forward compatible implementations.
type UnimplementedABFServiceServer struct {
}

func (*UnimplementedABFServiceServer) AuthRequest(context.Context, *UserData) (*AllowUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthRequest not implemented")
}
func (*UnimplementedABFServiceServer) ResetBucket(context.Context, *UserData) (*StatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetBucket not implemented")
}
func (*UnimplementedABFServiceServer) AddToList(context.Context, *IpData) (*StatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToList not implemented")
}
func (*UnimplementedABFServiceServer) RemoveFromList(context.Context, *IpData) (*StatusMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromList not implemented")
}

func RegisterABFServiceServer(s *grpc.Server, srv ABFServiceServer) {
	s.RegisterService(&_ABFService_serviceDesc, srv)
}

func _ABFService_AuthRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABFServiceServer).AuthRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ABFService/AuthRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABFServiceServer).AuthRequest(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABFService_ResetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABFServiceServer).ResetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ABFService/ResetBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABFServiceServer).ResetBucket(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABFService_AddToList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABFServiceServer).AddToList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ABFService/AddToList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABFServiceServer).AddToList(ctx, req.(*IpData))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABFService_RemoveFromList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABFServiceServer).RemoveFromList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ABFService/RemoveFromList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABFServiceServer).RemoveFromList(ctx, req.(*IpData))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABFService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ABFService",
	HandlerType: (*ABFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthRequest",
			Handler:    _ABFService_AuthRequest_Handler,
		},
		{
			MethodName: "ResetBucket",
			Handler:    _ABFService_ResetBucket_Handler,
		},
		{
			MethodName: "AddToList",
			Handler:    _ABFService_AddToList_Handler,
		},
		{
			MethodName: "RemoveFromList",
			Handler:    _ABFService_RemoveFromList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}