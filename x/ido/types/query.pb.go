// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nebula/ido/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// ====== IDO
type IDORequest struct {
	// Project unique id of each project
	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" yaml:"project_id"`
}

func (m *IDORequest) Reset()         { *m = IDORequest{} }
func (m *IDORequest) String() string { return proto.CompactTextString(m) }
func (*IDORequest) ProtoMessage()    {}
func (*IDORequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{2}
}
func (m *IDORequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDORequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDORequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDORequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDORequest.Merge(m, src)
}
func (m *IDORequest) XXX_Size() int {
	return m.Size()
}
func (m *IDORequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDORequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDORequest proto.InternalMessageInfo

func (m *IDORequest) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

type IDOResponse struct {
	Ido *IDO `protobuf:"bytes,1,opt,name=ido,proto3" json:"ido,omitempty" yaml:"ido"`
}

func (m *IDOResponse) Reset()         { *m = IDOResponse{} }
func (m *IDOResponse) String() string { return proto.CompactTextString(m) }
func (*IDOResponse) ProtoMessage()    {}
func (*IDOResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{3}
}
func (m *IDOResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDOResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDOResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDOResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDOResponse.Merge(m, src)
}
func (m *IDOResponse) XXX_Size() int {
	return m.Size()
}
func (m *IDOResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDOResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDOResponse proto.InternalMessageInfo

func (m *IDOResponse) GetIdo() *IDO {
	if m != nil {
		return m.Ido
	}
	return nil
}

// ====== Total IDO
type TotalIDORequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *TotalIDORequest) Reset()         { *m = TotalIDORequest{} }
func (m *TotalIDORequest) String() string { return proto.CompactTextString(m) }
func (*TotalIDORequest) ProtoMessage()    {}
func (*TotalIDORequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{4}
}
func (m *TotalIDORequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalIDORequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalIDORequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalIDORequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalIDORequest.Merge(m, src)
}
func (m *TotalIDORequest) XXX_Size() int {
	return m.Size()
}
func (m *TotalIDORequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalIDORequest.DiscardUnknown(m)
}

var xxx_messageInfo_TotalIDORequest proto.InternalMessageInfo

func (m *TotalIDORequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type TotalIDOResponse struct {
	Idos       []IDO               `protobuf:"bytes,1,rep,name=idos,proto3" json:"idos"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *TotalIDOResponse) Reset()         { *m = TotalIDOResponse{} }
func (m *TotalIDOResponse) String() string { return proto.CompactTextString(m) }
func (*TotalIDOResponse) ProtoMessage()    {}
func (*TotalIDOResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_612569411179be11, []int{5}
}
func (m *TotalIDOResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalIDOResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalIDOResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalIDOResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalIDOResponse.Merge(m, src)
}
func (m *TotalIDOResponse) XXX_Size() int {
	return m.Size()
}
func (m *TotalIDOResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalIDOResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TotalIDOResponse proto.InternalMessageInfo

func (m *TotalIDOResponse) GetIdos() []IDO {
	if m != nil {
		return m.Idos
	}
	return nil
}

func (m *TotalIDOResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nebula.ido.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nebula.ido.QueryParamsResponse")
	proto.RegisterType((*IDORequest)(nil), "nebula.ido.IDORequest")
	proto.RegisterType((*IDOResponse)(nil), "nebula.ido.IDOResponse")
	proto.RegisterType((*TotalIDORequest)(nil), "nebula.ido.TotalIDORequest")
	proto.RegisterType((*TotalIDOResponse)(nil), "nebula.ido.TotalIDOResponse")
}

func init() { proto.RegisterFile("nebula/ido/query.proto", fileDescriptor_612569411179be11) }

var fileDescriptor_612569411179be11 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0x4d, 0x0c, 0xfa, 0x0a, 0x56, 0xc7, 0xb4, 0x29, 0x31, 0x6e, 0x64, 0x10, 0xb5,
	0x42, 0x77, 0x4c, 0xf5, 0xe4, 0x49, 0x62, 0xb0, 0xe4, 0x94, 0x1a, 0xbc, 0x28, 0x82, 0xcc, 0x66,
	0x86, 0x75, 0x64, 0xb3, 0x6f, 0x9b, 0xdd, 0x88, 0x41, 0xbc, 0x78, 0xf0, 0x2c, 0xf8, 0x4d, 0xfc,
	0x14, 0x3d, 0x16, 0xbc, 0x78, 0x0a, 0x92, 0xf8, 0x09, 0xfa, 0x09, 0x64, 0x67, 0x26, 0xec, 0xb6,
	0x89, 0x78, 0x4b, 0xde, 0xfb, 0xcf, 0xef, 0xfd, 0xe7, 0xff, 0x66, 0x61, 0x27, 0x92, 0xfe, 0x24,
	0xe4, 0x4c, 0x09, 0x64, 0xc7, 0x13, 0x39, 0x9e, 0x7a, 0xf1, 0x18, 0x53, 0x24, 0x60, 0xea, 0x9e,
	0x12, 0xd8, 0xa8, 0x05, 0x18, 0xa0, 0x2e, 0xb3, 0xec, 0x97, 0x51, 0x34, 0x9a, 0x01, 0x62, 0x10,
	0x4a, 0xc6, 0x63, 0xc5, 0x78, 0x14, 0x61, 0xca, 0x53, 0x85, 0x51, 0x62, 0xbb, 0x0f, 0x86, 0x98,
	0x8c, 0x30, 0x61, 0x3e, 0x4f, 0xa4, 0x01, 0xb3, 0x0f, 0x6d, 0x5f, 0xa6, 0xbc, 0xcd, 0x62, 0x1e,
	0xa8, 0x48, 0x8b, 0xad, 0xb6, 0x5e, 0xf0, 0x10, 0xf3, 0x31, 0x1f, 0x2d, 0x21, 0xb5, 0x42, 0x43,
	0x09, 0x3b, 0x98, 0xd6, 0x80, 0xbc, 0xc8, 0x80, 0x47, 0x5a, 0x3a, 0x90, 0xc7, 0x13, 0x99, 0xa4,
	0xf4, 0x10, 0x6e, 0x9c, 0xab, 0x26, 0x31, 0x46, 0x89, 0x24, 0x0f, 0xa1, 0x6a, 0x90, 0xbb, 0xce,
	0x6d, 0xe7, 0xfe, 0xe6, 0x01, 0xf1, 0xf2, 0x8b, 0x79, 0x46, 0xdb, 0xa9, 0x9c, 0xcc, 0x5a, 0xa5,
	0x81, 0xd5, 0xd1, 0x0e, 0x40, 0xaf, 0xdb, 0xb7, 0x58, 0xf2, 0x18, 0x20, 0x1e, 0xe3, 0x7b, 0x39,
	0x4c, 0xdf, 0x2a, 0xa1, 0x19, 0x95, 0xce, 0xf6, 0xd9, 0xac, 0x75, 0x7d, 0xca, 0x47, 0xe1, 0x13,
	0x9a, 0xf7, 0xe8, 0xe0, 0x8a, 0xfd, 0xd3, 0x13, 0xf4, 0x29, 0x6c, 0x6a, 0x86, 0x35, 0xd1, 0x86,
	0xb2, 0x12, 0x68, 0x1d, 0x6c, 0x15, 0x1d, 0xf4, 0xba, 0xfd, 0xce, 0xd5, 0xb3, 0x59, 0x0b, 0x0c,
	0x4e, 0x09, 0xa4, 0x83, 0x4c, 0x4b, 0x5f, 0xc1, 0xd6, 0x4b, 0x4c, 0x79, 0x58, 0xb0, 0xf2, 0x1c,
	0x20, 0x8f, 0x6e, 0x77, 0x43, 0xc3, 0xee, 0x7a, 0x26, 0x67, 0x2f, 0xcb, 0xd9, 0x33, 0x0b, 0xb4,
	0x39, 0x7b, 0x47, 0x3c, 0x90, 0xf6, 0xec, 0xa0, 0x70, 0x92, 0x7e, 0x75, 0xe0, 0x5a, 0xce, 0xb6,
	0x16, 0xf7, 0xa0, 0xa2, 0x04, 0x66, 0x29, 0x95, 0xd7, 0x79, 0x34, 0x11, 0x69, 0x09, 0x39, 0x5c,
	0xe3, 0xe3, 0xde, 0x7f, 0x7d, 0x98, 0x39, 0x45, 0x23, 0x07, 0x3f, 0x36, 0xe0, 0x92, 0xde, 0x19,
	0x91, 0x50, 0x35, 0xbb, 0x20, 0x6e, 0x71, 0xf2, 0xea, 0x9a, 0x1b, 0xad, 0x7f, 0xf6, 0xcd, 0x00,
	0xda, 0xf8, 0xf2, 0xf3, 0xcf, 0xf7, 0x8d, 0x1a, 0x21, 0x6c, 0xe5, 0x55, 0x91, 0x37, 0x50, 0xee,
	0x75, 0xfb, 0x64, 0xe7, 0xc2, 0xed, 0x96, 0xec, 0xfa, 0x4a, 0xdd, 0x32, 0xef, 0x68, 0xa6, 0x4b,
	0x9a, 0xec, 0xfc, 0x83, 0x64, 0x9f, 0xf2, 0xfd, 0x7f, 0x26, 0x02, 0x2e, 0x2f, 0x63, 0x25, 0x37,
	0x8b, 0xa8, 0x0b, 0x8b, 0x6c, 0x34, 0xd7, 0x37, 0xed, 0xb0, 0x5b, 0x7a, 0x58, 0x9d, 0x6c, 0x17,
	0x87, 0xa5, 0x99, 0x6a, 0x5f, 0x09, 0xec, 0x3c, 0x3b, 0x99, 0xbb, 0xce, 0xe9, 0xdc, 0x75, 0x7e,
	0xcf, 0x5d, 0xe7, 0xdb, 0xc2, 0x2d, 0x9d, 0x2e, 0xdc, 0xd2, 0xaf, 0x85, 0x5b, 0x7a, 0xbd, 0x17,
	0xa8, 0xf4, 0xdd, 0xc4, 0xf7, 0x86, 0x38, 0xb2, 0x47, 0xf7, 0x43, 0xee, 0x27, 0x4b, 0xcc, 0x47,
	0x03, 0x9a, 0xc6, 0x32, 0xf1, 0xab, 0xfa, 0x4b, 0x7a, 0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x01,
	0xe1, 0x2b, 0x70, 0xfe, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	IDO(ctx context.Context, in *IDORequest, opts ...grpc.CallOption) (*IDOResponse, error)
	TotalIDO(ctx context.Context, in *TotalIDORequest, opts ...grpc.CallOption) (*TotalIDOResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nebula.ido.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IDO(ctx context.Context, in *IDORequest, opts ...grpc.CallOption) (*IDOResponse, error) {
	out := new(IDOResponse)
	err := c.cc.Invoke(ctx, "/nebula.ido.Query/IDO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalIDO(ctx context.Context, in *TotalIDORequest, opts ...grpc.CallOption) (*TotalIDOResponse, error) {
	out := new(TotalIDOResponse)
	err := c.cc.Invoke(ctx, "/nebula.ido.Query/TotalIDO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	IDO(context.Context, *IDORequest) (*IDOResponse, error)
	TotalIDO(context.Context, *TotalIDORequest) (*TotalIDOResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) IDO(ctx context.Context, req *IDORequest) (*IDOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IDO not implemented")
}
func (*UnimplementedQueryServer) TotalIDO(ctx context.Context, req *TotalIDORequest) (*TotalIDOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalIDO not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebula.ido.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IDO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IDO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebula.ido.Query/IDO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IDO(ctx, req.(*IDORequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalIDO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TotalIDORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalIDO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nebula.ido.Query/TotalIDO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalIDO(ctx, req.(*TotalIDORequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nebula.ido.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "IDO",
			Handler:    _Query_IDO_Handler,
		},
		{
			MethodName: "TotalIDO",
			Handler:    _Query_TotalIDO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nebula/ido/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *IDORequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDORequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDORequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProjectId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProjectId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IDOResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDOResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDOResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ido != nil {
		{
			size, err := m.Ido.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TotalIDORequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalIDORequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalIDORequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *TotalIDOResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalIDOResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalIDOResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Idos) > 0 {
		for iNdEx := len(m.Idos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Idos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *IDORequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectId != 0 {
		n += 1 + sovQuery(uint64(m.ProjectId))
	}
	return n
}

func (m *IDOResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ido != nil {
		l = m.Ido.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *TotalIDORequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *TotalIDOResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Idos) > 0 {
		for _, e := range m.Idos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IDORequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IDORequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDORequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
			}
			m.ProjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IDOResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IDOResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDOResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ido", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ido == nil {
				m.Ido = &IDO{}
			}
			if err := m.Ido.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TotalIDORequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TotalIDORequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalIDORequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TotalIDOResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TotalIDOResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalIDOResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Idos = append(m.Idos, IDO{})
			if err := m.Idos[len(m.Idos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
