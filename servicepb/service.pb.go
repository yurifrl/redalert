// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package servicepb is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	CheckListRequest
	CheckListResponse
	CheckEnableRequest
	CheckEnableResponse
	CheckDisableRequest
	CheckDisableResponse
	Check
*/
package servicepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Check_Status int32

const (
	Check_DISABLED   Check_Status = 0
	Check_UNKNOWN    Check_Status = 1
	Check_FAILING    Check_Status = 2
	Check_SUCCESSFUL Check_Status = 3
)

var Check_Status_name = map[int32]string{
	0: "DISABLED",
	1: "UNKNOWN",
	2: "FAILING",
	3: "SUCCESSFUL",
}
var Check_Status_value = map[string]int32{
	"DISABLED":   0,
	"UNKNOWN":    1,
	"FAILING":    2,
	"SUCCESSFUL": 3,
}

func (x Check_Status) String() string {
	return proto.EnumName(Check_Status_name, int32(x))
}
func (Check_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type CheckListRequest struct {
}

func (m *CheckListRequest) Reset()                    { *m = CheckListRequest{} }
func (m *CheckListRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckListRequest) ProtoMessage()               {}
func (*CheckListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CheckListResponse struct {
	Members []*Check `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
}

func (m *CheckListResponse) Reset()                    { *m = CheckListResponse{} }
func (m *CheckListResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckListResponse) ProtoMessage()               {}
func (*CheckListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckListResponse) GetMembers() []*Check {
	if m != nil {
		return m.Members
	}
	return nil
}

type CheckEnableRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CheckEnableRequest) Reset()                    { *m = CheckEnableRequest{} }
func (m *CheckEnableRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckEnableRequest) ProtoMessage()               {}
func (*CheckEnableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckEnableRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CheckEnableResponse struct {
}

func (m *CheckEnableResponse) Reset()                    { *m = CheckEnableResponse{} }
func (m *CheckEnableResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckEnableResponse) ProtoMessage()               {}
func (*CheckEnableResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CheckDisableRequest struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CheckDisableRequest) Reset()                    { *m = CheckDisableRequest{} }
func (m *CheckDisableRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckDisableRequest) ProtoMessage()               {}
func (*CheckDisableRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckDisableRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type CheckDisableResponse struct {
}

func (m *CheckDisableResponse) Reset()                    { *m = CheckDisableResponse{} }
func (m *CheckDisableResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckDisableResponse) ProtoMessage()               {}
func (*CheckDisableResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Check struct {
	ID      string       `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name    string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type    string       `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Status  Check_Status `protobuf:"varint,4,opt,name=status,enum=servicepb.Check_Status" json:"status,omitempty"`
	Enabled bool         `protobuf:"varint,5,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *Check) Reset()                    { *m = Check{} }
func (m *Check) String() string            { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()               {}
func (*Check) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Check) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Check) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Check) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Check) GetStatus() Check_Status {
	if m != nil {
		return m.Status
	}
	return Check_DISABLED
}

func (m *Check) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*CheckListRequest)(nil), "servicepb.CheckListRequest")
	proto.RegisterType((*CheckListResponse)(nil), "servicepb.CheckListResponse")
	proto.RegisterType((*CheckEnableRequest)(nil), "servicepb.CheckEnableRequest")
	proto.RegisterType((*CheckEnableResponse)(nil), "servicepb.CheckEnableResponse")
	proto.RegisterType((*CheckDisableRequest)(nil), "servicepb.CheckDisableRequest")
	proto.RegisterType((*CheckDisableResponse)(nil), "servicepb.CheckDisableResponse")
	proto.RegisterType((*Check)(nil), "servicepb.Check")
	proto.RegisterEnum("servicepb.Check_Status", Check_Status_name, Check_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
<<<<<<< HEAD
// const _ = grpc.SupportPackageIsVersion3
=======
const _ = grpc.SupportPackageIsVersion4
>>>>>>> 033f1573675e466f748642385e547e7139534836

// Client API for RedalertService service

type RedalertServiceClient interface {
	// CheckList lists all the checks in configuration
	CheckList(ctx context.Context, in *CheckListRequest, opts ...grpc.CallOption) (*CheckListResponse, error)
	// CheckEnable enables a check
	CheckEnable(ctx context.Context, in *CheckEnableRequest, opts ...grpc.CallOption) (*CheckEnableResponse, error)
	// CheckDisable disables a check
	CheckDisable(ctx context.Context, in *CheckDisableRequest, opts ...grpc.CallOption) (*CheckDisableResponse, error)
}

type redalertServiceClient struct {
	cc *grpc.ClientConn
}

func NewRedalertServiceClient(cc *grpc.ClientConn) RedalertServiceClient {
	return &redalertServiceClient{cc}
}

func (c *redalertServiceClient) CheckList(ctx context.Context, in *CheckListRequest, opts ...grpc.CallOption) (*CheckListResponse, error) {
	out := new(CheckListResponse)
	err := grpc.Invoke(ctx, "/servicepb.RedalertService/CheckList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redalertServiceClient) CheckEnable(ctx context.Context, in *CheckEnableRequest, opts ...grpc.CallOption) (*CheckEnableResponse, error) {
	out := new(CheckEnableResponse)
	err := grpc.Invoke(ctx, "/servicepb.RedalertService/CheckEnable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redalertServiceClient) CheckDisable(ctx context.Context, in *CheckDisableRequest, opts ...grpc.CallOption) (*CheckDisableResponse, error) {
	out := new(CheckDisableResponse)
	err := grpc.Invoke(ctx, "/servicepb.RedalertService/CheckDisable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RedalertService service

type RedalertServiceServer interface {
	// CheckList lists all the checks in configuration
	CheckList(context.Context, *CheckListRequest) (*CheckListResponse, error)
	// CheckEnable enables a check
	CheckEnable(context.Context, *CheckEnableRequest) (*CheckEnableResponse, error)
	// CheckDisable disables a check
	CheckDisable(context.Context, *CheckDisableRequest) (*CheckDisableResponse, error)
}

func RegisterRedalertServiceServer(s *grpc.Server, srv RedalertServiceServer) {
	s.RegisterService(&_RedalertService_serviceDesc, srv)
}

func _RedalertService_CheckList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedalertServiceServer).CheckList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicepb.RedalertService/CheckList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedalertServiceServer).CheckList(ctx, req.(*CheckListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedalertService_CheckEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedalertServiceServer).CheckEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicepb.RedalertService/CheckEnable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedalertServiceServer).CheckEnable(ctx, req.(*CheckEnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedalertService_CheckDisable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedalertServiceServer).CheckDisable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicepb.RedalertService/CheckDisable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedalertServiceServer).CheckDisable(ctx, req.(*CheckDisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RedalertService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "servicepb.RedalertService",
	HandlerType: (*RedalertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckList",
			Handler:    _RedalertService_CheckList_Handler,
		},
		{
			MethodName: "CheckEnable",
			Handler:    _RedalertService_CheckEnable_Handler,
		},
		{
			MethodName: "CheckDisable",
			Handler:    _RedalertService_CheckDisable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xe9, 0x4e, 0xab, 0x40,
	0x14, 0xee, 0xd0, 0xfd, 0xb4, 0xb7, 0x97, 0x7b, 0xae, 0x0b, 0xa9, 0x1b, 0x21, 0x9a, 0x10, 0x7f,
	0x60, 0x52, 0x1f, 0x40, 0x6b, 0x69, 0x95, 0x48, 0x30, 0x42, 0x1a, 0x7f, 0x43, 0x3b, 0x89, 0x8d,
	0x5d, 0x90, 0x99, 0x9a, 0xf8, 0x8e, 0x3e, 0x93, 0x31, 0x1d, 0x68, 0x03, 0x68, 0xfd, 0x37, 0xe7,
	0xdb, 0x98, 0xc3, 0x37, 0xf0, 0x87, 0xd1, 0xe8, 0x6d, 0x32, 0xa2, 0x46, 0x18, 0x2d, 0xf8, 0x02,
	0xeb, 0xc9, 0x18, 0x06, 0x1a, 0x82, 0xdc, 0x7b, 0xa6, 0xa3, 0x17, 0x7b, 0xc2, 0xb8, 0x4b, 0x5f,
	0x97, 0x94, 0x71, 0xed, 0x0a, 0xfe, 0xa5, 0x30, 0x16, 0x2e, 0xe6, 0x8c, 0xe2, 0x39, 0x54, 0x67,
	0x74, 0x16, 0xd0, 0x88, 0x29, 0x44, 0x2d, 0xea, 0x8d, 0x8e, 0x6c, 0x6c, 0x52, 0x0c, 0x21, 0x77,
	0xd7, 0x02, 0xed, 0x14, 0x50, 0x20, 0xfd, 0xb9, 0x1f, 0x4c, 0x69, 0x12, 0x8b, 0x2d, 0x90, 0x2c,
	0x53, 0x21, 0x2a, 0xd1, 0xeb, 0xae, 0x64, 0x99, 0xda, 0x2e, 0xfc, 0xcf, 0xa8, 0xe2, 0x0f, 0x69,
	0x67, 0x09, 0x6c, 0x4e, 0xd8, 0x6f, 0xee, 0x3d, 0xd8, 0xc9, 0xca, 0x12, 0xfb, 0x07, 0x81, 0xb2,
	0x20, 0xf2, 0x0e, 0x44, 0x28, 0xcd, 0xfd, 0x19, 0x55, 0x24, 0x81, 0x88, 0xf3, 0x0a, 0xe3, 0xef,
	0x21, 0x55, 0x8a, 0x31, 0xb6, 0x3a, 0xe3, 0x05, 0x54, 0x18, 0xf7, 0xf9, 0x92, 0x29, 0x25, 0x95,
	0xe8, 0xad, 0xce, 0x7e, 0x7e, 0x51, 0xc3, 0x13, 0xb4, 0x9b, 0xc8, 0x50, 0x81, 0x2a, 0x15, 0x3b,
	0x8c, 0x95, 0xb2, 0x4a, 0xf4, 0x9a, 0xbb, 0x1e, 0xb5, 0x6b, 0xa8, 0xc4, 0x5a, 0x6c, 0x42, 0xcd,
	0xb4, 0xbc, 0xee, 0x8d, 0xdd, 0x37, 0xe5, 0x02, 0x36, 0xa0, 0x3a, 0x74, 0xee, 0x9d, 0x87, 0x27,
	0x47, 0x26, 0xab, 0x61, 0xd0, 0xb5, 0x6c, 0xcb, 0xb9, 0x95, 0x25, 0x6c, 0x01, 0x78, 0xc3, 0x5e,
	0xaf, 0xef, 0x79, 0x83, 0xa1, 0x2d, 0x17, 0x3b, 0x9f, 0x04, 0xfe, 0xba, 0x74, 0xec, 0x4f, 0x69,
	0xc4, 0xbd, 0xf8, 0x1a, 0x78, 0x07, 0xf5, 0x4d, 0x3f, 0x78, 0x90, 0xbf, 0x5d, 0xaa, 0xc9, 0xf6,
	0xe1, 0xcf, 0x64, 0xf2, 0xab, 0x0a, 0xe8, 0x40, 0x23, 0x55, 0x01, 0x1e, 0xe5, 0xe5, 0x99, 0x02,
	0xdb, 0xc7, 0xdb, 0xe8, 0x4d, 0xde, 0x23, 0x34, 0xd3, 0xa5, 0xe0, 0x37, 0x47, 0xb6, 0xd4, 0xf6,
	0xc9, 0x56, 0x7e, 0x1d, 0x19, 0x54, 0xc4, 0x93, 0xbd, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x7c,
	0xbf, 0xf3, 0x7e, 0xc3, 0x02, 0x00, 0x00,
}
