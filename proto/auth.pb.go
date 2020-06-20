// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AcessToken           string               `protobuf:"bytes,1,opt,name=acessToken,proto3" json:"acessToken,omitempty"`
	RefreshToken         string               `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	AccessTokenExpiry    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=accessTokenExpiry,proto3" json:"accessTokenExpiry,omitempty"`
	RefreshTokenExpiry   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=refreshTokenExpiry,proto3" json:"refreshTokenExpiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAcessToken() string {
	if m != nil {
		return m.AcessToken
	}
	return ""
}

func (m *LoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *LoginResponse) GetAccessTokenExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.AccessTokenExpiry
	}
	return nil
}

func (m *LoginResponse) GetRefreshTokenExpiry() *timestamp.Timestamp {
	if m != nil {
		return m.RefreshTokenExpiry
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Timestamp struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type Status struct {
	IsUserLoggedIn       bool     `protobuf:"varint,1,opt,name=isUserLoggedIn,proto3" json:"isUserLoggedIn,omitempty"`
	AccessToken          string   `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken         string   `protobuf:"bytes,3,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetIsUserLoggedIn() bool {
	if m != nil {
		return m.IsUserLoggedIn
	}
	return false
}

func (m *Status) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Status) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*Empty)(nil), "auth.Empty")
	proto.RegisterType((*Timestamp)(nil), "auth.Timestamp")
	proto.RegisterType((*Status)(nil), "auth.Status")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x49, 0x93, 0xb6, 0x93, 0x80, 0xc4, 0xc0, 0x21, 0xf2, 0x01, 0x22, 0x1f, 0x50, 0x39,
	0xc4, 0x86, 0x72, 0xe4, 0x54, 0x50, 0x10, 0x41, 0x81, 0x83, 0x63, 0x2e, 0xdc, 0xb6, 0xf6, 0x74,
	0x6d, 0x51, 0x7b, 0xcd, 0xce, 0x6e, 0xa0, 0x7c, 0x04, 0xff, 0xc9, 0x5f, 0x20, 0x7b, 0xed, 0xc8,
	0x21, 0x48, 0x9c, 0xec, 0xf7, 0x66, 0xde, 0x93, 0xde, 0xdb, 0x01, 0x10, 0xd6, 0xe4, 0x61, 0xad,
	0x95, 0x51, 0x78, 0xd2, 0xfc, 0xfb, 0x4f, 0xa5, 0x52, 0xf2, 0x96, 0xa2, 0x96, 0xbb, 0xb6, 0x37,
	0x91, 0x29, 0x4a, 0x62, 0x23, 0xca, 0xda, 0xad, 0x05, 0xef, 0x60, 0xb6, 0x51, 0xb2, 0xa8, 0x62,
	0xfa, 0x66, 0x89, 0x0d, 0xfa, 0x70, 0x66, 0x99, 0x74, 0x25, 0x4a, 0x9a, 0x7b, 0x0b, 0xef, 0xe2,
	0x3c, 0xde, 0xe3, 0x66, 0x56, 0x0b, 0xe6, 0xef, 0x4a, 0x67, 0xf3, 0x7b, 0x6e, 0xd6, 0xe3, 0xe0,
	0xb7, 0x07, 0xf7, 0x3b, 0x23, 0xae, 0x55, 0xc5, 0x84, 0x4f, 0x00, 0x44, 0x4a, 0xcc, 0x89, 0xfa,
	0x4a, 0x55, 0xe7, 0x35, 0x60, 0x30, 0x80, 0x99, 0xa6, 0x1b, 0x4d, 0x9c, 0xbb, 0x0d, 0xe7, 0x78,
	0xc0, 0xe1, 0x7b, 0x78, 0x28, 0xd2, 0xbd, 0x64, 0xf5, 0xa3, 0x2e, 0xf4, 0xdd, 0x7c, 0xb4, 0xf0,
	0x2e, 0xa6, 0x97, 0x7e, 0xe8, 0xa2, 0x85, 0x7d, 0xb4, 0x30, 0xe9, 0xa3, 0xc5, 0xc7, 0x22, 0xfc,
	0x00, 0x38, 0x74, 0xee, 0xac, 0x4e, 0xfe, 0x6b, 0xf5, 0x0f, 0x55, 0x70, 0x0a, 0xe3, 0x55, 0x59,
	0x9b, 0xbb, 0xe0, 0x35, 0x9c, 0xef, 0x37, 0x71, 0x0e, 0xa7, 0x4c, 0xa9, 0xaa, 0x32, 0x6e, 0xc3,
	0x8e, 0xe2, 0x1e, 0xe2, 0x63, 0x18, 0x57, 0xa2, 0x52, 0xdc, 0x46, 0x1c, 0xc7, 0x0e, 0x04, 0x3b,
	0x98, 0x6c, 0x8d, 0x30, 0x96, 0xf1, 0x19, 0x3c, 0x28, 0xf8, 0x33, 0x93, 0xde, 0x28, 0x29, 0x29,
	0x5b, 0xbb, 0xb6, 0xce, 0xe2, 0xbf, 0x58, 0x5c, 0xc0, 0x74, 0x10, 0xac, 0x2b, 0x6c, 0x48, 0x1d,
	0x75, 0x3a, 0x3a, 0xee, 0xf4, 0xf2, 0x97, 0x07, 0xd3, 0x2b, 0x6b, 0xf2, 0x2d, 0xe9, 0x5d, 0x91,
	0x12, 0xbe, 0x80, 0x71, 0xfb, 0x70, 0x88, 0x61, 0x7b, 0x3e, 0xc3, 0x73, 0xf0, 0x1f, 0x1d, 0x70,
	0xdd, 0xcb, 0x06, 0x30, 0xd9, 0x28, 0xa9, 0xac, 0xc1, 0xa9, 0x1b, 0xb7, 0x6d, 0xf8, 0x43, 0x80,
	0xcf, 0x61, 0xb6, 0x13, 0xb7, 0x45, 0x26, 0x0c, 0x35, 0x19, 0x0e, 0x37, 0x67, 0x0e, 0xb8, 0xf8,
	0x6f, 0x5e, 0x7e, 0x89, 0x64, 0x61, 0x72, 0x7b, 0x1d, 0xa6, 0xaa, 0x8c, 0xd6, 0xdb, 0x64, 0xb5,
	0xdc, 0xbe, 0x5d, 0x7e, 0xbc, 0xfa, 0xb4, 0x4e, 0xa2, 0x92, 0xa4, 0x30, 0x9a, 0x54, 0x6d, 0x7f,
	0x2e, 0x1b, 0x4d, 0x77, 0xc8, 0x93, 0xf6, 0xf3, 0xea, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x1f, 0xac, 0x46, 0xf1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	ValidateUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateUser(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/auth.AuthService/validateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *Empty) (*Empty, error)
	ValidateUser(context.Context, *Empty) (*Status, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Logout(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (*UnimplementedAuthServiceServer) ValidateUser(ctx context.Context, req *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUser not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ValidateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateUser(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthService_Logout_Handler,
		},
		{
			MethodName: "validateUser",
			Handler:    _AuthService_ValidateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
