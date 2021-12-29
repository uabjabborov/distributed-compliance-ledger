// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compliance/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCertifyModel struct {
	Signer                string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty" validate:"required"`
	Vid                   int32  `protobuf:"varint,2,opt,name=vid,proto3" json:"vid,omitempty" validate:"required,gte=0,lte=65535"`
	Pid                   int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty" validate:"required,gte=0,lte=65535"`
	SoftwareVersion       uint32 `protobuf:"varint,4,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty" validate:"required"`
	SoftwareVersionString string `protobuf:"bytes,5,opt,name=software_version_string,json=softwareVersionString,proto3" json:"software_version_string,omitempty" validate:"required"`
	CertificationDate     string `protobuf:"bytes,6,opt,name=certification_date,json=certificationDate,proto3" json:"certification_date,omitempty" validate:"required"`
	CertificationType     string `protobuf:"bytes,7,opt,name=certification_type,json=certificationType,proto3" json:"certification_type,omitempty" validate:"required"`
	Reason                string `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *MsgCertifyModel) Reset()         { *m = MsgCertifyModel{} }
func (m *MsgCertifyModel) String() string { return proto.CompactTextString(m) }
func (*MsgCertifyModel) ProtoMessage()    {}
func (*MsgCertifyModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4a0d801b7ae7cc, []int{0}
}
func (m *MsgCertifyModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCertifyModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCertifyModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCertifyModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCertifyModel.Merge(m, src)
}
func (m *MsgCertifyModel) XXX_Size() int {
	return m.Size()
}
func (m *MsgCertifyModel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCertifyModel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCertifyModel proto.InternalMessageInfo

func (m *MsgCertifyModel) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgCertifyModel) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *MsgCertifyModel) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *MsgCertifyModel) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *MsgCertifyModel) GetSoftwareVersionString() string {
	if m != nil {
		return m.SoftwareVersionString
	}
	return ""
}

func (m *MsgCertifyModel) GetCertificationDate() string {
	if m != nil {
		return m.CertificationDate
	}
	return ""
}

func (m *MsgCertifyModel) GetCertificationType() string {
	if m != nil {
		return m.CertificationType
	}
	return ""
}

func (m *MsgCertifyModel) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type MsgCertifyModelResponse struct {
}

func (m *MsgCertifyModelResponse) Reset()         { *m = MsgCertifyModelResponse{} }
func (m *MsgCertifyModelResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCertifyModelResponse) ProtoMessage()    {}
func (*MsgCertifyModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4a0d801b7ae7cc, []int{1}
}
func (m *MsgCertifyModelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCertifyModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCertifyModelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCertifyModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCertifyModelResponse.Merge(m, src)
}
func (m *MsgCertifyModelResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCertifyModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCertifyModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCertifyModelResponse proto.InternalMessageInfo

type MsgRevokeModel struct {
	Signer                string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty" validate:"required"`
	Vid                   int32  `protobuf:"varint,2,opt,name=vid,proto3" json:"vid,omitempty" validate:"required,gte=0,lte=65535"`
	Pid                   int32  `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty" validate:"required,gte=0,lte=65535"`
	SoftwareVersion       uint32 `protobuf:"varint,4,opt,name=software_version,json=softwareVersion,proto3" json:"software_version,omitempty" validate:"required"`
	SoftwareVersionString string `protobuf:"bytes,5,opt,name=software_version_string,json=softwareVersionString,proto3" json:"software_version_string,omitempty" validate:"required"`
	RevocationDate        string `protobuf:"bytes,6,opt,name=revocation_date,json=revocationDate,proto3" json:"revocation_date,omitempty" validate:"required"`
	CertificationType     string `protobuf:"bytes,7,opt,name=certification_type,json=certificationType,proto3" json:"certification_type,omitempty" validate:"required"`
	Reason                string `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (m *MsgRevokeModel) Reset()         { *m = MsgRevokeModel{} }
func (m *MsgRevokeModel) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeModel) ProtoMessage()    {}
func (*MsgRevokeModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4a0d801b7ae7cc, []int{2}
}
func (m *MsgRevokeModel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeModel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeModel.Merge(m, src)
}
func (m *MsgRevokeModel) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeModel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeModel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeModel proto.InternalMessageInfo

func (m *MsgRevokeModel) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *MsgRevokeModel) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *MsgRevokeModel) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *MsgRevokeModel) GetSoftwareVersion() uint32 {
	if m != nil {
		return m.SoftwareVersion
	}
	return 0
}

func (m *MsgRevokeModel) GetSoftwareVersionString() string {
	if m != nil {
		return m.SoftwareVersionString
	}
	return ""
}

func (m *MsgRevokeModel) GetRevocationDate() string {
	if m != nil {
		return m.RevocationDate
	}
	return ""
}

func (m *MsgRevokeModel) GetCertificationType() string {
	if m != nil {
		return m.CertificationType
	}
	return ""
}

func (m *MsgRevokeModel) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type MsgRevokeModelResponse struct {
}

func (m *MsgRevokeModelResponse) Reset()         { *m = MsgRevokeModelResponse{} }
func (m *MsgRevokeModelResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeModelResponse) ProtoMessage()    {}
func (*MsgRevokeModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4a0d801b7ae7cc, []int{3}
}
func (m *MsgRevokeModelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeModelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeModelResponse.Merge(m, src)
}
func (m *MsgRevokeModelResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeModelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCertifyModel)(nil), "zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel")
	proto.RegisterType((*MsgCertifyModelResponse)(nil), "zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModelResponse")
	proto.RegisterType((*MsgRevokeModel)(nil), "zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel")
	proto.RegisterType((*MsgRevokeModelResponse)(nil), "zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModelResponse")
}

func init() { proto.RegisterFile("compliance/tx.proto", fileDescriptor_db4a0d801b7ae7cc) }

var fileDescriptor_db4a0d801b7ae7cc = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xeb, 0x26, 0x0d, 0x30, 0x40, 0x03, 0xd3, 0xd2, 0xb8, 0x59, 0xb8, 0x91, 0x59, 0x90,
	0x45, 0x63, 0x23, 0xaa, 0x20, 0x81, 0x54, 0x09, 0x0c, 0x94, 0x05, 0x32, 0x48, 0x06, 0xb1, 0x60,
	0x13, 0x39, 0x9e, 0xd7, 0x61, 0x84, 0xeb, 0x31, 0x33, 0x13, 0xd3, 0x70, 0x00, 0x76, 0x48, 0xdc,
	0x01, 0x89, 0x13, 0x70, 0x08, 0x96, 0x15, 0x62, 0xc1, 0xaa, 0x42, 0xc9, 0x0d, 0x7a, 0x02, 0x64,
	0x3b, 0x56, 0xfe, 0x54, 0x45, 0x4a, 0x55, 0x58, 0xb1, 0xf3, 0x3c, 0xbf, 0xef, 0xf7, 0xcd, 0x3c,
	0x7d, 0x9a, 0x41, 0x2b, 0x01, 0xdf, 0x8b, 0x43, 0xe6, 0x47, 0x01, 0xd8, 0x6a, 0xdf, 0x8a, 0x05,
	0x57, 0x1c, 0xb7, 0xdf, 0x33, 0xda, 0x05, 0xf0, 0xc3, 0xfc, 0x87, 0x45, 0x98, 0x54, 0x82, 0x75,
	0x7b, 0x0a, 0xc8, 0xb8, 0x3d, 0x04, 0x42, 0x41, 0x58, 0xe3, 0x42, 0x7d, 0x3d, 0xe0, 0x72, 0x8f,
	0xcb, 0x4e, 0x06, 0xb1, 0xf3, 0x45, 0x4e, 0xac, 0xaf, 0x52, 0x4e, 0x79, 0x5e, 0x4f, 0xbf, 0xf2,
	0xaa, 0xf9, 0xb1, 0x8c, 0xaa, 0xae, 0xa4, 0x0f, 0x40, 0x28, 0xb6, 0xdb, 0x77, 0x39, 0x81, 0x10,
	0x3f, 0x46, 0x15, 0xc9, 0x68, 0x04, 0x42, 0xd7, 0x1a, 0x5a, 0xf3, 0x82, 0x63, 0x1f, 0x1d, 0x6e,
	0xac, 0x24, 0x7e, 0xc8, 0x88, 0xaf, 0xe0, 0xae, 0x29, 0xe0, 0x6d, 0x8f, 0x09, 0x20, 0xe6, 0xf7,
	0xaf, 0xad, 0xd5, 0x91, 0xc5, 0x7d, 0x42, 0x04, 0x48, 0xf9, 0x5c, 0x09, 0x16, 0x51, 0x6f, 0x24,
	0xc7, 0x77, 0x50, 0x29, 0x61, 0x44, 0x5f, 0x6c, 0x68, 0xcd, 0x25, 0xe7, 0xc6, 0xd1, 0xe1, 0xc6,
	0xf5, 0xe3, 0x94, 0x4d, 0xaa, 0x60, 0xfb, 0xe6, 0x66, 0xa8, 0x60, 0xfb, 0x76, 0xbb, 0xbd, 0xd5,
	0x36, 0xbd, 0x54, 0x93, 0x4a, 0x63, 0x46, 0xf4, 0xd2, 0x9c, 0xd2, 0x98, 0x11, 0xec, 0xa0, 0x2b,
	0x92, 0xef, 0xaa, 0x77, 0xbe, 0x80, 0x4e, 0x02, 0x42, 0x32, 0x1e, 0xe9, 0xe5, 0x86, 0xd6, 0xbc,
	0xec, 0xd4, 0x4e, 0x38, 0x88, 0x57, 0x2d, 0x04, 0x2f, 0xf3, 0x7e, 0xfc, 0x0c, 0xd5, 0x66, 0x19,
	0x1d, 0x99, 0x1d, 0x4e, 0x5f, 0xca, 0x66, 0x72, 0x22, 0xea, 0xda, 0x0c, 0x2a, 0x1f, 0x09, 0xde,
	0x41, 0x38, 0xc8, 0x66, 0xcc, 0x02, 0x5f, 0xa5, 0xb4, 0x54, 0xa7, 0x57, 0xfe, 0xcc, 0xba, 0x3a,
	0x25, 0x79, 0xe8, 0x2b, 0x38, 0xce, 0x51, 0xfd, 0x18, 0xf4, 0x73, 0xf3, 0x70, 0x5e, 0xf4, 0x63,
	0xc0, 0x6b, 0xa8, 0x22, 0xc0, 0x97, 0x3c, 0xd2, 0xcf, 0xa7, 0x5a, 0x6f, 0xb4, 0x32, 0xd7, 0x51,
	0x6d, 0x26, 0x0e, 0x1e, 0xc8, 0x98, 0x47, 0x12, 0xcc, 0x0f, 0x65, 0xb4, 0xec, 0x4a, 0xea, 0x41,
	0xc2, 0xdf, 0xc0, 0xff, 0xa4, 0xfc, 0xad, 0xa4, 0xdc, 0x43, 0x55, 0x01, 0x09, 0x9f, 0x23, 0x26,
	0xcb, 0xe3, 0xfe, 0x7f, 0x92, 0x11, 0x1d, 0xad, 0x4d, 0xe7, 0xa0, 0x88, 0xc8, 0xad, 0x1f, 0x8b,
	0xa8, 0xe4, 0x4a, 0x8a, 0xbf, 0x68, 0xe8, 0xd2, 0xd4, 0x95, 0xb2, 0x63, 0x9d, 0xea, 0x3e, 0xb3,
	0x66, 0xb2, 0x58, 0x7f, 0x7a, 0x36, 0x9c, 0x62, 0xc3, 0xf8, 0xb3, 0x86, 0x2e, 0x4e, 0x06, 0xfa,
	0xd1, 0xe9, 0xf9, 0x13, 0x98, 0xba, 0x7b, 0x26, 0x98, 0x62, 0x97, 0x0e, 0x7c, 0x1b, 0x18, 0xda,
	0xc1, 0xc0, 0xd0, 0x7e, 0x0d, 0x0c, 0xed, 0xd3, 0xd0, 0x58, 0x38, 0x18, 0x1a, 0x0b, 0x3f, 0x87,
	0xc6, 0xc2, 0xab, 0x27, 0x94, 0xa9, 0xd7, 0xbd, 0x6e, 0xca, 0xb0, 0x73, 0xcb, 0x56, 0xe1, 0x69,
	0x4f, 0x78, 0xb6, 0xc6, 0x1e, 0xad, 0xdc, 0xd5, 0xde, 0xb7, 0x27, 0x9f, 0x9d, 0x7e, 0x0c, 0xb2,
	0x5b, 0xc9, 0x9e, 0x84, 0xad, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0x0c, 0x6a, 0x71, 0x91,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CertifyModel(ctx context.Context, in *MsgCertifyModel, opts ...grpc.CallOption) (*MsgCertifyModelResponse, error)
	RevokeModel(ctx context.Context, in *MsgRevokeModel, opts ...grpc.CallOption) (*MsgRevokeModelResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CertifyModel(ctx context.Context, in *MsgCertifyModel, opts ...grpc.CallOption) (*MsgCertifyModelResponse, error) {
	out := new(MsgCertifyModelResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.compliance.Msg/CertifyModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeModel(ctx context.Context, in *MsgRevokeModel, opts ...grpc.CallOption) (*MsgRevokeModelResponse, error) {
	out := new(MsgRevokeModelResponse)
	err := c.cc.Invoke(ctx, "/zigbeealliance.distributedcomplianceledger.compliance.Msg/RevokeModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CertifyModel(context.Context, *MsgCertifyModel) (*MsgCertifyModelResponse, error)
	RevokeModel(context.Context, *MsgRevokeModel) (*MsgRevokeModelResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CertifyModel(ctx context.Context, req *MsgCertifyModel) (*MsgCertifyModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CertifyModel not implemented")
}
func (*UnimplementedMsgServer) RevokeModel(ctx context.Context, req *MsgRevokeModel) (*MsgRevokeModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeModel not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CertifyModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCertifyModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CertifyModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.compliance.Msg/CertifyModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CertifyModel(ctx, req.(*MsgCertifyModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zigbeealliance.distributedcomplianceledger.compliance.Msg/RevokeModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeModel(ctx, req.(*MsgRevokeModel))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zigbeealliance.distributedcomplianceledger.compliance.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CertifyModel",
			Handler:    _Msg_CertifyModel_Handler,
		},
		{
			MethodName: "RevokeModel",
			Handler:    _Msg_RevokeModel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compliance/tx.proto",
}

func (m *MsgCertifyModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCertifyModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCertifyModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CertificationType) > 0 {
		i -= len(m.CertificationType)
		copy(dAtA[i:], m.CertificationType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertificationType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.CertificationDate) > 0 {
		i -= len(m.CertificationDate)
		copy(dAtA[i:], m.CertificationDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertificationDate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SoftwareVersionString) > 0 {
		i -= len(m.SoftwareVersionString)
		copy(dAtA[i:], m.SoftwareVersionString)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SoftwareVersionString)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x20
	}
	if m.Pid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x18
	}
	if m.Vid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCertifyModelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCertifyModelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCertifyModelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevokeModel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeModel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeModel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.CertificationType) > 0 {
		i -= len(m.CertificationType)
		copy(dAtA[i:], m.CertificationType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertificationType)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.RevocationDate) > 0 {
		i -= len(m.RevocationDate)
		copy(dAtA[i:], m.RevocationDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RevocationDate)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SoftwareVersionString) > 0 {
		i -= len(m.SoftwareVersionString)
		copy(dAtA[i:], m.SoftwareVersionString)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SoftwareVersionString)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SoftwareVersion != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SoftwareVersion))
		i--
		dAtA[i] = 0x20
	}
	if m.Pid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Pid))
		i--
		dAtA[i] = 0x18
	}
	if m.Vid != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeModelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeModelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeModelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCertifyModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Vid != 0 {
		n += 1 + sovTx(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovTx(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovTx(uint64(m.SoftwareVersion))
	}
	l = len(m.SoftwareVersionString)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertificationDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertificationType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCertifyModelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevokeModel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Vid != 0 {
		n += 1 + sovTx(uint64(m.Vid))
	}
	if m.Pid != 0 {
		n += 1 + sovTx(uint64(m.Pid))
	}
	if m.SoftwareVersion != 0 {
		n += 1 + sovTx(uint64(m.SoftwareVersion))
	}
	l = len(m.SoftwareVersionString)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RevocationDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertificationType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeModelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCertifyModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCertifyModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCertifyModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SoftwareVersionString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCertifyModelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCertifyModelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCertifyModelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRevokeModel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRevokeModel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeModel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersion", wireType)
			}
			m.SoftwareVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SoftwareVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SoftwareVersionString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SoftwareVersionString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevocationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RevocationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertificationType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertificationType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRevokeModelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRevokeModelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeModelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
