// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.7.1
// source: grpc/aura.proto

package api

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

type GetNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNumberRequest) Reset() {
	*x = GetNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberRequest) ProtoMessage() {}

func (x *GetNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberRequest.ProtoReflect.Descriptor instead.
func (*GetNumberRequest) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{0}
}

type GetNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetNumberResponse) Reset() {
	*x = GetNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNumberResponse) ProtoMessage() {}

func (x *GetNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNumberResponse.ProtoReflect.Descriptor instead.
func (*GetNumberResponse) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{1}
}

func (x *GetNumberResponse) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type IncrementNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncrementNumberRequest) Reset() {
	*x = IncrementNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementNumberRequest) ProtoMessage() {}

func (x *IncrementNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementNumberRequest.ProtoReflect.Descriptor instead.
func (*IncrementNumberRequest) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{2}
}

type IncrementNumberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncrementNumberResponse) Reset() {
	*x = IncrementNumberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncrementNumberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementNumberResponse) ProtoMessage() {}

func (x *IncrementNumberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementNumberResponse.ProtoReflect.Descriptor instead.
func (*IncrementNumberResponse) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{3}
}

type SetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxValue      uint64 `protobuf:"varint,1,opt,name=maxValue,proto3" json:"maxValue,omitempty"`
	IncrementStep uint64 `protobuf:"varint,2,opt,name=incrementStep,proto3" json:"incrementStep,omitempty"`
}

func (x *SetSettingsRequest) Reset() {
	*x = SetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSettingsRequest) ProtoMessage() {}

func (x *SetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSettingsRequest.ProtoReflect.Descriptor instead.
func (*SetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{4}
}

func (x *SetSettingsRequest) GetMaxValue() uint64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *SetSettingsRequest) GetIncrementStep() uint64 {
	if x != nil {
		return x.IncrementStep
	}
	return 0
}

type SetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetSettingsResponse) Reset() {
	*x = SetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_aura_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSettingsResponse) ProtoMessage() {}

func (x *SetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_aura_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSettingsResponse.ProtoReflect.Descriptor instead.
func (*SetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_aura_proto_rawDescGZIP(), []int{5}
}

var File_grpc_aura_proto protoreflect.FileDescriptor

var file_grpc_aura_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x72, 0x61, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x19, 0x0a, 0x17, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x56, 0x0a, 0x12, 0x53,
	0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x65, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x65, 0x70, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe5, 0x01, 0x0a, 0x0b, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x72, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x61, 0x75, 0x72, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x61, 0x75, 0x72, 0x61, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x75,
	0x72, 0x61, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b,
	0x53, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x2e, 0x61, 0x75,
	0x72, 0x61, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x72, 0x61, 0x2e, 0x53, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_aura_proto_rawDescOnce sync.Once
	file_grpc_aura_proto_rawDescData = file_grpc_aura_proto_rawDesc
)

func file_grpc_aura_proto_rawDescGZIP() []byte {
	file_grpc_aura_proto_rawDescOnce.Do(func() {
		file_grpc_aura_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_aura_proto_rawDescData)
	})
	return file_grpc_aura_proto_rawDescData
}

var file_grpc_aura_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_aura_proto_goTypes = []interface{}{
	(*GetNumberRequest)(nil),        // 0: aura.GetNumberRequest
	(*GetNumberResponse)(nil),       // 1: aura.GetNumberResponse
	(*IncrementNumberRequest)(nil),  // 2: aura.IncrementNumberRequest
	(*IncrementNumberResponse)(nil), // 3: aura.IncrementNumberResponse
	(*SetSettingsRequest)(nil),      // 4: aura.SetSettingsRequest
	(*SetSettingsResponse)(nil),     // 5: aura.SetSettingsResponse
}
var file_grpc_aura_proto_depIdxs = []int32{
	0, // 0: aura.Incrementer.GetNumber:input_type -> aura.GetNumberRequest
	2, // 1: aura.Incrementer.IncrementNumber:input_type -> aura.IncrementNumberRequest
	4, // 2: aura.Incrementer.SetSettings:input_type -> aura.SetSettingsRequest
	1, // 3: aura.Incrementer.GetNumber:output_type -> aura.GetNumberResponse
	3, // 4: aura.Incrementer.IncrementNumber:output_type -> aura.IncrementNumberResponse
	5, // 5: aura.Incrementer.SetSettings:output_type -> aura.SetSettingsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_aura_proto_init() }
func file_grpc_aura_proto_init() {
	if File_grpc_aura_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_aura_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberRequest); i {
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
		file_grpc_aura_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNumberResponse); i {
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
		file_grpc_aura_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementNumberRequest); i {
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
		file_grpc_aura_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncrementNumberResponse); i {
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
		file_grpc_aura_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSettingsRequest); i {
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
		file_grpc_aura_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSettingsResponse); i {
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
			RawDescriptor: file_grpc_aura_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_aura_proto_goTypes,
		DependencyIndexes: file_grpc_aura_proto_depIdxs,
		MessageInfos:      file_grpc_aura_proto_msgTypes,
	}.Build()
	File_grpc_aura_proto = out.File
	file_grpc_aura_proto_rawDesc = nil
	file_grpc_aura_proto_goTypes = nil
	file_grpc_aura_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IncrementerClient is the client API for Incrementer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IncrementerClient interface {
	GetNumber(ctx context.Context, in *GetNumberRequest, opts ...grpc.CallOption) (*GetNumberResponse, error)
	IncrementNumber(ctx context.Context, in *IncrementNumberRequest, opts ...grpc.CallOption) (*IncrementNumberResponse, error)
	SetSettings(ctx context.Context, in *SetSettingsRequest, opts ...grpc.CallOption) (*SetSettingsResponse, error)
}

type incrementerClient struct {
	cc grpc.ClientConnInterface
}

func NewIncrementerClient(cc grpc.ClientConnInterface) IncrementerClient {
	return &incrementerClient{cc}
}

func (c *incrementerClient) GetNumber(ctx context.Context, in *GetNumberRequest, opts ...grpc.CallOption) (*GetNumberResponse, error) {
	out := new(GetNumberResponse)
	err := c.cc.Invoke(ctx, "/aura.Incrementer/GetNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementerClient) IncrementNumber(ctx context.Context, in *IncrementNumberRequest, opts ...grpc.CallOption) (*IncrementNumberResponse, error) {
	out := new(IncrementNumberResponse)
	err := c.cc.Invoke(ctx, "/aura.Incrementer/IncrementNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incrementerClient) SetSettings(ctx context.Context, in *SetSettingsRequest, opts ...grpc.CallOption) (*SetSettingsResponse, error) {
	out := new(SetSettingsResponse)
	err := c.cc.Invoke(ctx, "/aura.Incrementer/SetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncrementerServer is the server API for Incrementer service.
type IncrementerServer interface {
	GetNumber(context.Context, *GetNumberRequest) (*GetNumberResponse, error)
	IncrementNumber(context.Context, *IncrementNumberRequest) (*IncrementNumberResponse, error)
	SetSettings(context.Context, *SetSettingsRequest) (*SetSettingsResponse, error)
}

// UnimplementedIncrementerServer can be embedded to have forward compatible implementations.
type UnimplementedIncrementerServer struct {
}

func (*UnimplementedIncrementerServer) GetNumber(context.Context, *GetNumberRequest) (*GetNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumber not implemented")
}
func (*UnimplementedIncrementerServer) IncrementNumber(context.Context, *IncrementNumberRequest) (*IncrementNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementNumber not implemented")
}
func (*UnimplementedIncrementerServer) SetSettings(context.Context, *SetSettingsRequest) (*SetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSettings not implemented")
}

func RegisterIncrementerServer(s *grpc.Server, srv IncrementerServer) {
	s.RegisterService(&_Incrementer_serviceDesc, srv)
}

func _Incrementer_GetNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementerServer).GetNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aura.Incrementer/GetNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementerServer).GetNumber(ctx, req.(*GetNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incrementer_IncrementNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementerServer).IncrementNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aura.Incrementer/IncrementNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementerServer).IncrementNumber(ctx, req.(*IncrementNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Incrementer_SetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncrementerServer).SetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aura.Incrementer/SetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncrementerServer).SetSettings(ctx, req.(*SetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Incrementer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "aura.Incrementer",
	HandlerType: (*IncrementerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNumber",
			Handler:    _Incrementer_GetNumber_Handler,
		},
		{
			MethodName: "IncrementNumber",
			Handler:    _Incrementer_IncrementNumber_Handler,
		},
		{
			MethodName: "SetSettings",
			Handler:    _Incrementer_SetSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/aura.proto",
}
