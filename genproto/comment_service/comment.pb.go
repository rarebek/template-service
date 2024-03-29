// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comment_service/comment.proto

package comment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Comment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	PostId               string   `protobuf:"bytes,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	OwnerId              string   `protobuf:"bytes,4,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            string   `protobuf:"bytes,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Comment) Reset()         { *m = Comment{} }
func (m *Comment) String() string { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()    {}
func (*Comment) Descriptor() ([]byte, []int) {
	return fileDescriptor_55720040aaa933ee, []int{0}
}
func (m *Comment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Comment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Comment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Comment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Comment.Merge(m, src)
}
func (m *Comment) XXX_Size() int {
	return m.Size()
}
func (m *Comment) XXX_DiscardUnknown() {
	xxx_messageInfo_Comment.DiscardUnknown(m)
}

var xxx_messageInfo_Comment proto.InternalMessageInfo

func (m *Comment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Comment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Comment) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *Comment) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Comment) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Comment) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Comment) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

type GetPostID struct {
	PostId               string   `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPostID) Reset()         { *m = GetPostID{} }
func (m *GetPostID) String() string { return proto.CompactTextString(m) }
func (*GetPostID) ProtoMessage()    {}
func (*GetPostID) Descriptor() ([]byte, []int) {
	return fileDescriptor_55720040aaa933ee, []int{1}
}
func (m *GetPostID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetPostID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetPostID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetPostID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPostID.Merge(m, src)
}
func (m *GetPostID) XXX_Size() int {
	return m.Size()
}
func (m *GetPostID) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPostID.DiscardUnknown(m)
}

var xxx_messageInfo_GetPostID proto.InternalMessageInfo

func (m *GetPostID) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

type GetOwnerID struct {
	OwnerId              string   `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOwnerID) Reset()         { *m = GetOwnerID{} }
func (m *GetOwnerID) String() string { return proto.CompactTextString(m) }
func (*GetOwnerID) ProtoMessage()    {}
func (*GetOwnerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_55720040aaa933ee, []int{2}
}
func (m *GetOwnerID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetOwnerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetOwnerID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetOwnerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOwnerID.Merge(m, src)
}
func (m *GetOwnerID) XXX_Size() int {
	return m.Size()
}
func (m *GetOwnerID) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOwnerID.DiscardUnknown(m)
}

var xxx_messageInfo_GetOwnerID proto.InternalMessageInfo

func (m *GetOwnerID) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type AllComments struct {
	Comments             []*Comment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AllComments) Reset()         { *m = AllComments{} }
func (m *AllComments) String() string { return proto.CompactTextString(m) }
func (*AllComments) ProtoMessage()    {}
func (*AllComments) Descriptor() ([]byte, []int) {
	return fileDescriptor_55720040aaa933ee, []int{3}
}
func (m *AllComments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllComments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllComments.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllComments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllComments.Merge(m, src)
}
func (m *AllComments) XXX_Size() int {
	return m.Size()
}
func (m *AllComments) XXX_DiscardUnknown() {
	xxx_messageInfo_AllComments.DiscardUnknown(m)
}

var xxx_messageInfo_AllComments proto.InternalMessageInfo

func (m *AllComments) GetComments() []*Comment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func init() {
	proto.RegisterType((*Comment)(nil), "comment.Comment")
	proto.RegisterType((*GetPostID)(nil), "comment.GetPostID")
	proto.RegisterType((*GetOwnerID)(nil), "comment.GetOwnerID")
	proto.RegisterType((*AllComments)(nil), "comment.AllComments")
}

func init() { proto.RegisterFile("comment_service/comment.proto", fileDescriptor_55720040aaa933ee) }

var fileDescriptor_55720040aaa933ee = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0xbf, 0x81, 0x4f, 0x0a, 0x97, 0x48, 0xc8, 0x68, 0xa4, 0x92, 0xd0, 0x90, 0xc6, 0x44,
	0x16, 0x06, 0x13, 0x58, 0xba, 0xe2, 0x4f, 0x42, 0x58, 0x41, 0xf0, 0x01, 0x08, 0x76, 0xee, 0xa2,
	0x09, 0xb4, 0xa4, 0xbd, 0x6a, 0x7c, 0x13, 0x1f, 0xc8, 0x85, 0x3b, 0x7d, 0x04, 0x53, 0x5f, 0xc4,
	0x74, 0xe6, 0xb6, 0x29, 0xe2, 0xae, 0xe7, 0xfc, 0xee, 0xdc, 0x9c, 0x39, 0x1d, 0xe8, 0x78, 0xe1,
	0x6e, 0x87, 0x01, 0xad, 0x63, 0x8c, 0x9e, 0x7c, 0x0f, 0x6f, 0x59, 0xf7, 0xf7, 0x51, 0x48, 0xa1,
	0xb4, 0x58, 0xba, 0x6f, 0x02, 0xac, 0x89, 0xf9, 0x96, 0x0d, 0x28, 0xf9, 0xca, 0x16, 0x5d, 0xd1,
	0xab, 0xad, 0x4a, 0xbe, 0x92, 0x36, 0x58, 0x5e, 0x18, 0x10, 0x06, 0x64, 0x97, 0xb4, 0x99, 0x49,
	0xd9, 0x02, 0x6b, 0x1f, 0xc6, 0xb4, 0xf6, 0x95, 0x5d, 0xd6, 0xa4, 0x92, 0xca, 0xb9, 0x92, 0x97,
	0x50, 0x0d, 0x9f, 0x03, 0x8c, 0x52, 0xf2, 0xdf, 0x9c, 0xd1, 0x7a, 0xae, 0x64, 0x07, 0xc0, 0x8b,
	0x70, 0x43, 0xa8, 0xd6, 0x1b, 0xb2, 0x4f, 0x34, 0xac, 0xb1, 0x33, 0xa2, 0x14, 0x3f, 0xee, 0x55,
	0x86, 0x2b, 0x06, 0xb3, 0x63, 0xb0, 0xc2, 0x2d, 0x32, 0xb6, 0x0c, 0x66, 0x67, 0x44, 0xee, 0x15,
	0xd4, 0x66, 0x48, 0xcb, 0x34, 0xc4, 0xb4, 0x98, 0x4e, 0x14, 0xd3, 0xb9, 0xd7, 0x00, 0x33, 0xa4,
	0x85, 0x0e, 0x34, 0x3d, 0xc8, 0x2a, 0x0e, 0xb2, 0xba, 0x77, 0x50, 0x1f, 0x6d, 0xb7, 0xdc, 0x4b,
	0x2c, 0x6f, 0xa0, 0xca, 0x7d, 0xc5, 0xb6, 0xe8, 0x96, 0x7b, 0xf5, 0x41, 0xb3, 0x9f, 0xf5, 0xc9,
	0x43, 0xab, 0x7c, 0x62, 0xf0, 0x21, 0xa0, 0xc1, 0xee, 0xbd, 0x29, 0x5f, 0x0e, 0xe1, 0x74, 0xa2,
	0x6f, 0x9a, 0x55, 0x7d, 0x74, 0xbe, 0x7d, 0xe4, 0xc8, 0x31, 0x5c, 0xcc, 0x90, 0x0a, 0x39, 0xc6,
	0x2f, 0x4b, 0xd3, 0xb2, 0xcc, 0x67, 0xf3, 0x4b, 0xb7, 0xcf, 0x73, 0xaf, 0x98, 0x7c, 0x0a, 0xad,
	0xdf, 0x3b, 0x16, 0xfc, 0x3f, 0xce, 0x8a, 0x4b, 0xb8, 0x93, 0xbf, 0xb7, 0x8c, 0x9b, 0xef, 0x89,
	0x23, 0x3e, 0x13, 0x47, 0x7c, 0x25, 0x8e, 0x78, 0xfd, 0x76, 0xfe, 0x3d, 0x54, 0xf4, 0x33, 0x1a,
	0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0x56, 0xd7, 0x4c, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommentServiceClient interface {
	CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error)
	GetAllCommentsByPostId(ctx context.Context, in *GetPostID, opts ...grpc.CallOption) (*AllComments, error)
	GetAllCommentsByOwnerId(ctx context.Context, in *GetOwnerID, opts ...grpc.CallOption) (*AllComments, error)
}

type commentServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommentServiceClient(cc *grpc.ClientConn) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CreateComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/comment.CommentService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetAllCommentsByPostId(ctx context.Context, in *GetPostID, opts ...grpc.CallOption) (*AllComments, error) {
	out := new(AllComments)
	err := c.cc.Invoke(ctx, "/comment.CommentService/GetAllCommentsByPostId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) GetAllCommentsByOwnerId(ctx context.Context, in *GetOwnerID, opts ...grpc.CallOption) (*AllComments, error) {
	out := new(AllComments)
	err := c.cc.Invoke(ctx, "/comment.CommentService/GetAllCommentsByOwnerId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
type CommentServiceServer interface {
	CreateComment(context.Context, *Comment) (*Comment, error)
	GetAllCommentsByPostId(context.Context, *GetPostID) (*AllComments, error)
	GetAllCommentsByOwnerId(context.Context, *GetOwnerID) (*AllComments, error)
}

// UnimplementedCommentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (*UnimplementedCommentServiceServer) CreateComment(ctx context.Context, req *Comment) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (*UnimplementedCommentServiceServer) GetAllCommentsByPostId(ctx context.Context, req *GetPostID) (*AllComments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsByPostId not implemented")
}
func (*UnimplementedCommentServiceServer) GetAllCommentsByOwnerId(ctx context.Context, req *GetOwnerID) (*AllComments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsByOwnerId not implemented")
}

func RegisterCommentServiceServer(s *grpc.Server, srv CommentServiceServer) {
	s.RegisterService(&_CommentService_serviceDesc, srv)
}

func _CommentService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CreateComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetAllCommentsByPostId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetAllCommentsByPostId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/GetAllCommentsByPostId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetAllCommentsByPostId(ctx, req.(*GetPostID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_GetAllCommentsByOwnerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).GetAllCommentsByOwnerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.CommentService/GetAllCommentsByOwnerId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).GetAllCommentsByOwnerId(ctx, req.(*GetOwnerID))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComment",
			Handler:    _CommentService_CreateComment_Handler,
		},
		{
			MethodName: "GetAllCommentsByPostId",
			Handler:    _CommentService_GetAllCommentsByPostId_Handler,
		},
		{
			MethodName: "GetAllCommentsByOwnerId",
			Handler:    _CommentService_GetAllCommentsByOwnerId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment_service/comment.proto",
}

func (m *Comment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Comment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintComment(dAtA, i, uint64(len(m.DeletedAt)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.UpdatedAt) > 0 {
		i -= len(m.UpdatedAt)
		copy(dAtA[i:], m.UpdatedAt)
		i = encodeVarintComment(dAtA, i, uint64(len(m.UpdatedAt)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CreatedAt) > 0 {
		i -= len(m.CreatedAt)
		copy(dAtA[i:], m.CreatedAt)
		i = encodeVarintComment(dAtA, i, uint64(len(m.CreatedAt)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OwnerId) > 0 {
		i -= len(m.OwnerId)
		copy(dAtA[i:], m.OwnerId)
		i = encodeVarintComment(dAtA, i, uint64(len(m.OwnerId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintComment(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintComment(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetPostID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetPostID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetPostID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintComment(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetOwnerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetOwnerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetOwnerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.OwnerId) > 0 {
		i -= len(m.OwnerId)
		copy(dAtA[i:], m.OwnerId)
		i = encodeVarintComment(dAtA, i, uint64(len(m.OwnerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllComments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllComments) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllComments) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Comments) > 0 {
		for iNdEx := len(m.Comments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Comments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintComment(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintComment(dAtA []byte, offset int, v uint64) int {
	offset -= sovComment(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Comment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.OwnerId)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.CreatedAt)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.UpdatedAt)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	l = len(m.DeletedAt)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetPostID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetOwnerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerId)
	if l > 0 {
		n += 1 + l + sovComment(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *AllComments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovComment(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovComment(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozComment(x uint64) (n int) {
	return sovComment(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Comment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedAt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
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
func (m *GetPostID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: GetPostID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetPostID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
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
func (m *GetOwnerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: GetOwnerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetOwnerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
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
func (m *AllComments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowComment
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
			return fmt.Errorf("proto: AllComments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllComments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowComment
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
				return ErrInvalidLengthComment
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthComment
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipComment(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthComment
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
func skipComment(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
					return 0, ErrIntOverflowComment
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
				return 0, ErrInvalidLengthComment
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupComment
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthComment
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthComment        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowComment          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupComment = fmt.Errorf("proto: unexpected end of group")
)
