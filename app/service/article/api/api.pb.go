// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package article_service_v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ArtDetailRequest struct {
	Sn                   string   `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty" form:"name" validate:"required"`
	Status               int64    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty" form:"status" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtDetailRequest) Reset()         { *m = ArtDetailRequest{} }
func (m *ArtDetailRequest) String() string { return proto.CompactTextString(m) }
func (*ArtDetailRequest) ProtoMessage()    {}
func (*ArtDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *ArtDetailRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArtDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArtDetailRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArtDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtDetailRequest.Merge(m, src)
}
func (m *ArtDetailRequest) XXX_Size() int {
	return m.Size()
}
func (m *ArtDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ArtDetailRequest proto.InternalMessageInfo

func (m *ArtDetailRequest) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ArtDetailRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ArtDetailResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sn                   string   `protobuf:"bytes,2,opt,name=sn,proto3" json:"sn,omitempty" form:"name" validate:"required"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Img                  string   `protobuf:"bytes,4,opt,name=img,proto3" json:"img,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Status               int64    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArtDetailResp) Reset()         { *m = ArtDetailResp{} }
func (m *ArtDetailResp) String() string { return proto.CompactTextString(m) }
func (*ArtDetailResp) ProtoMessage()    {}
func (*ArtDetailResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *ArtDetailResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArtDetailResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArtDetailResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArtDetailResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtDetailResp.Merge(m, src)
}
func (m *ArtDetailResp) XXX_Size() int {
	return m.Size()
}
func (m *ArtDetailResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtDetailResp.DiscardUnknown(m)
}

var xxx_messageInfo_ArtDetailResp proto.InternalMessageInfo

func (m *ArtDetailResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArtDetailResp) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ArtDetailResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArtDetailResp) GetImg() string {
	if m != nil {
		return m.Img
	}
	return ""
}

func (m *ArtDetailResp) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ArtDetailResp) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ArtDetailResp) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ArtDetailResp) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ArtDetailResp) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*ArtDetailRequest)(nil), "article.service.v1.ArtDetailRequest")
	proto.RegisterType((*ArtDetailResp)(nil), "article.service.v1.ArtDetailResp")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x8b, 0x13, 0x3f,
	0x18, 0xc6, 0xbf, 0x93, 0xee, 0xb6, 0xdf, 0x06, 0x95, 0x12, 0x64, 0x19, 0xaa, 0xb6, 0xdd, 0x51,
	0xa1, 0x17, 0x33, 0xea, 0x82, 0x87, 0x05, 0x0f, 0x53, 0x14, 0xaf, 0xd2, 0x3d, 0x7a, 0x90, 0x74,
	0xe6, 0xdd, 0x31, 0x30, 0x93, 0xcc, 0x26, 0xef, 0x14, 0x96, 0x65, 0x11, 0x3c, 0x78, 0xf1, 0xe8,
	0xc5, 0x9b, 0xff, 0x8e, 0x47, 0xc1, 0xfb, 0x22, 0xd5, 0xbf, 0x60, 0xff, 0x02, 0x69, 0x92, 0xaa,
	0xf8, 0x03, 0x11, 0x3c, 0x35, 0xcf, 0xf3, 0x79, 0x9f, 0xf2, 0x64, 0xde, 0xd0, 0xbe, 0x68, 0x24,
	0x6f, 0x8c, 0x46, 0xcd, 0x98, 0x30, 0x28, 0xf3, 0x0a, 0xb8, 0x05, 0xb3, 0x94, 0x39, 0xf0, 0xe5,
	0x9d, 0xe1, 0xad, 0x52, 0xe2, 0xb3, 0x76, 0xc1, 0x73, 0x5d, 0xa7, 0xa5, 0x2e, 0x75, 0xea, 0x46,
	0x17, 0xed, 0xa1, 0x53, 0x4e, 0xb8, 0x93, 0xff, 0x8b, 0xe1, 0x95, 0x52, 0xeb, 0xb2, 0x82, 0x6f,
	0x53, 0x50, 0x37, 0x78, 0x1c, 0xe0, 0xd5, 0x00, 0x45, 0x23, 0x53, 0xa1, 0x94, 0x46, 0x81, 0x52,
	0x2b, 0xeb, 0x69, 0xf2, 0x32, 0xa2, 0x83, 0xcc, 0xe0, 0x03, 0x40, 0x21, 0xab, 0x39, 0x1c, 0xb5,
	0x60, 0x91, 0xed, 0x51, 0x62, 0x55, 0x1c, 0x4d, 0xa2, 0x69, 0x7f, 0x76, 0xfd, 0xfc, 0x6c, 0x3c,
	0x3e, 0xd4, 0xa6, 0xde, 0x4f, 0x94, 0xa8, 0x21, 0x99, 0x2c, 0x45, 0x25, 0x0b, 0x81, 0xb0, 0x9f,
	0x18, 0x38, 0x6a, 0xa5, 0x81, 0x22, 0x99, 0x13, 0xab, 0xd8, 0x7d, 0xda, 0xb5, 0x28, 0xb0, 0xb5,
	0x31, 0x99, 0x44, 0xd3, 0xce, 0xec, 0xe6, 0xf9, 0xd9, 0x78, 0xd7, 0x07, 0xbd, 0xff, 0xeb, 0x68,
	0x08, 0x25, 0xaf, 0x08, 0xbd, 0xf8, 0x5d, 0x11, 0xdb, 0xb0, 0x4b, 0x94, 0xc8, 0xc2, 0xb5, 0xe8,
	0xcc, 0x89, 0x2c, 0x42, 0x2b, 0xf2, 0x77, 0xad, 0x2e, 0xd3, 0x6d, 0x94, 0x58, 0x41, 0xdc, 0x59,
	0xe7, 0xe6, 0x5e, 0xb0, 0x01, 0xed, 0xc8, 0xba, 0x8c, 0xb7, 0x9c, 0xb7, 0x3e, 0xb2, 0x98, 0xf6,
	0x72, 0xad, 0x10, 0x14, 0xc6, 0xdb, 0xce, 0xdd, 0x48, 0xb6, 0xf3, 0xf5, 0x5e, 0x5d, 0x57, 0x25,
	0x28, 0x76, 0x8d, 0xd2, 0xdc, 0x80, 0x40, 0x28, 0x9e, 0x0a, 0x8c, 0x7b, 0x2e, 0xd4, 0x0f, 0x4e,
	0x86, 0x6b, 0xdc, 0x36, 0xc5, 0x06, 0xff, 0xef, 0x71, 0x70, 0x3c, 0x2e, 0xa0, 0x82, 0x80, 0xfb,
	0x1e, 0x07, 0x27, 0xc3, 0xbb, 0x6f, 0x09, 0xed, 0x65, 0xfe, 0x5d, 0xb0, 0x7b, 0x74, 0xeb, 0xb1,
	0x54, 0x25, 0xdb, 0xe1, 0x7e, 0x93, 0x7c, 0xb3, 0x66, 0xfe, 0x70, 0xbd, 0xe6, 0xe1, 0x6f, 0x7c,
	0xf6, 0x9c, 0xd2, 0x47, 0x80, 0x99, 0xc1, 0xd9, 0xf1, 0x81, 0x62, 0x37, 0xf8, 0xcf, 0xef, 0x8c,
	0xff, 0xb8, 0xf9, 0xe1, 0xee, 0x1f, 0xa6, 0x6c, 0x93, 0x4c, 0x5f, 0x7c, 0xf8, 0xfc, 0x9a, 0x24,
	0x6c, 0x92, 0x86, 0xd1, 0xd4, 0xaa, 0xf4, 0xc4, 0xaa, 0xd3, 0xd4, 0x7f, 0x98, 0xf4, 0xc4, 0xff,
	0x9e, 0xb2, 0x27, 0xf4, 0x82, 0x2f, 0x70, 0x80, 0x06, 0x44, 0xfd, 0xcf, 0x2a, 0x4c, 0xa3, 0xdb,
	0xd1, 0x6c, 0xf0, 0x6e, 0x35, 0x8a, 0xde, 0xaf, 0x46, 0xd1, 0xc7, 0xd5, 0x28, 0x7a, 0xf3, 0x69,
	0xf4, 0xdf, 0xa2, 0xeb, 0xee, 0xbf, 0xf7, 0x25, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xff, 0x09, 0xe8,
	0x5c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArticleClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	GetArtBySn(ctx context.Context, in *ArtDetailRequest, opts ...grpc.CallOption) (*ArtDetailResp, error)
	GetArtStream(ctx context.Context, opts ...grpc.CallOption) (Article_GetArtStreamClient, error)
}

type articleClient struct {
	cc *grpc.ClientConn
}

func NewArticleClient(cc *grpc.ClientConn) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/article.service.v1.Article/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArtBySn(ctx context.Context, in *ArtDetailRequest, opts ...grpc.CallOption) (*ArtDetailResp, error) {
	out := new(ArtDetailResp)
	err := c.cc.Invoke(ctx, "/article.service.v1.Article/GetArtBySn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArtStream(ctx context.Context, opts ...grpc.CallOption) (Article_GetArtStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Article_serviceDesc.Streams[0], "/article.service.v1.Article/GetArtStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &articleGetArtStreamClient{stream}
	return x, nil
}

type Article_GetArtStreamClient interface {
	Send(*ArtDetailRequest) error
	Recv() (*ArtDetailResp, error)
	grpc.ClientStream
}

type articleGetArtStreamClient struct {
	grpc.ClientStream
}

func (x *articleGetArtStreamClient) Send(m *ArtDetailRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *articleGetArtStreamClient) Recv() (*ArtDetailResp, error) {
	m := new(ArtDetailResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArticleServer is the server API for Article service.
type ArticleServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	GetArtBySn(context.Context, *ArtDetailRequest) (*ArtDetailResp, error)
	GetArtStream(Article_GetArtStreamServer) error
}

// UnimplementedArticleServer can be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (*UnimplementedArticleServer) Ping(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedArticleServer) GetArtBySn(ctx context.Context, req *ArtDetailRequest) (*ArtDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtBySn not implemented")
}
func (*UnimplementedArticleServer) GetArtStream(srv Article_GetArtStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArtStream not implemented")
}

func RegisterArticleServer(s *grpc.Server, srv ArticleServer) {
	s.RegisterService(&_Article_serviceDesc, srv)
}

func _Article_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.service.v1.Article/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArtBySn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArtDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArtBySn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/article.service.v1.Article/GetArtBySn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArtBySn(ctx, req.(*ArtDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArtStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ArticleServer).GetArtStream(&articleGetArtStreamServer{stream})
}

type Article_GetArtStreamServer interface {
	Send(*ArtDetailResp) error
	Recv() (*ArtDetailRequest, error)
	grpc.ServerStream
}

type articleGetArtStreamServer struct {
	grpc.ServerStream
}

func (x *articleGetArtStreamServer) Send(m *ArtDetailResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *articleGetArtStreamServer) Recv() (*ArtDetailRequest, error) {
	m := new(ArtDetailRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Article_serviceDesc = grpc.ServiceDesc{
	ServiceName: "article.service.v1.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Article_Ping_Handler,
		},
		{
			MethodName: "GetArtBySn",
			Handler:    _Article_GetArtBySn_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArtStream",
			Handler:       _Article_GetArtStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func (m *ArtDetailRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArtDetailRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArtDetailRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sn) > 0 {
		i -= len(m.Sn)
		copy(dAtA[i:], m.Sn)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Sn)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ArtDetailResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArtDetailResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArtDetailResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.DeletedAt) > 0 {
		i -= len(m.DeletedAt)
		copy(dAtA[i:], m.DeletedAt)
		i = encodeVarintApi(dAtA, i, uint64(len(m.DeletedAt)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintApi(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintApi(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Img) > 0 {
		i -= len(m.Img)
		copy(dAtA[i:], m.Img)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Img)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Sn) > 0 {
		i -= len(m.Sn)
		copy(dAtA[i:], m.Sn)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Sn)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ArtDetailRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sn)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovApi(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ArtDetailResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApi(uint64(m.Id))
	}
	l = len(m.Sn)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Img)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovApi(uint64(m.Status))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.DeletedAt)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArtDetailRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ArtDetailRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArtDetailRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ArtDetailResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ArtDetailResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArtDetailResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Img", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Img = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)