// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_FAILURE Status = 0
	Status_SUCCESS Status = 1
)

var Status_name = map[int32]string{
	0: "FAILURE",
	1: "SUCCESS",
}

var Status_value = map[string]int32{
	"FAILURE": 0,
	"SUCCESS": 1,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

type BalanceRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceRequest.Unmarshal(m, b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceRequest.Size(m)
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type BalanceResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Balance              string    `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *BalanceResponse) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

type Response struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=account.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_FAILURE
}

func init() {
	proto.RegisterEnum("account.Status", Status_name, Status_value)
	proto.RegisterType((*BalanceRequest)(nil), "account.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "account.BalanceResponse")
	proto.RegisterType((*Response)(nil), "account.Response")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xb4, 0xb8,
	0xf8, 0x9c, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x60, 0x5c, 0xa5, 0x28, 0x2e, 0x7e, 0xb8, 0xda, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21,
	0x5d, 0x2e, 0x8e, 0x22, 0x28, 0x1b, 0xac, 0x9a, 0xdb, 0x48, 0x50, 0x0f, 0x66, 0x13, 0x4c, 0x51,
	0x10, 0x5c, 0x09, 0xc8, 0xec, 0x24, 0x88, 0x09, 0x12, 0x4c, 0x10, 0xb3, 0xa1, 0x5c, 0x25, 0x63,
	0x2e, 0x0e, 0xb8, 0xa1, 0xea, 0x5c, 0x6c, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0x10, 0x07, 0xf0, 0x19,
	0xf1, 0xc3, 0x8d, 0x0c, 0x06, 0x0b, 0x07, 0x41, 0xa5, 0xb5, 0x94, 0xb8, 0xd8, 0x20, 0x22, 0x42,
	0xdc, 0x5c, 0xec, 0x6e, 0x8e, 0x9e, 0x3e, 0xa1, 0x41, 0xae, 0x02, 0x0c, 0x20, 0x4e, 0x70, 0xa8,
	0xb3, 0xb3, 0x6b, 0x70, 0xb0, 0x00, 0xa3, 0x91, 0x17, 0x17, 0xbb, 0x23, 0x44, 0xb7, 0x90, 0x3d,
	0x17, 0x57, 0x7a, 0x6a, 0x09, 0xd4, 0x0b, 0x42, 0xe2, 0x70, 0x53, 0x51, 0x03, 0x40, 0x4a, 0x02,
	0x53, 0x02, 0xe2, 0xb0, 0x24, 0x36, 0x70, 0xe0, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb1,
	0xfc, 0xf3, 0x2d, 0x4d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) GetBalance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/account.Account/getBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	GetBalance(context.Context, *BalanceRequest) (*BalanceResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) GetBalance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.Account/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetBalance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getBalance",
			Handler:    _Account_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
