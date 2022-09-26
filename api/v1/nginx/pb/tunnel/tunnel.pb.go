// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tunnel.proto

package tunnel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SetTunnelRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTunnelRequest) Reset()         { *m = SetTunnelRequest{} }
func (m *SetTunnelRequest) String() string { return proto.CompactTextString(m) }
func (*SetTunnelRequest) ProtoMessage()    {}
func (*SetTunnelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}

func (m *SetTunnelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTunnelRequest.Unmarshal(m, b)
}
func (m *SetTunnelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTunnelRequest.Marshal(b, m, deterministic)
}
func (m *SetTunnelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTunnelRequest.Merge(m, src)
}
func (m *SetTunnelRequest) XXX_Size() int {
	return xxx_messageInfo_SetTunnelRequest.Size(m)
}
func (m *SetTunnelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTunnelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetTunnelRequest proto.InternalMessageInfo

func (m *SetTunnelRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tunnel struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 string   `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Domain               string   `protobuf:"bytes,4,opt,name=domain,proto3" json:"domain,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tunnel) Reset()         { *m = Tunnel{} }
func (m *Tunnel) String() string { return proto.CompactTextString(m) }
func (*Tunnel) ProtoMessage()    {}
func (*Tunnel) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}

func (m *Tunnel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tunnel.Unmarshal(m, b)
}
func (m *Tunnel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tunnel.Marshal(b, m, deterministic)
}
func (m *Tunnel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tunnel.Merge(m, src)
}
func (m *Tunnel) XXX_Size() int {
	return xxx_messageInfo_Tunnel.Size(m)
}
func (m *Tunnel) XXX_DiscardUnknown() {
	xxx_messageInfo_Tunnel.DiscardUnknown(m)
}

var xxx_messageInfo_Tunnel proto.InternalMessageInfo

func (m *Tunnel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tunnel) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Tunnel) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Tunnel) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Tunnel) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetTunnelResponse struct {
	Message              []*Tunnel `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
	Status               int32     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTunnelResponse) Reset()         { *m = GetTunnelResponse{} }
func (m *GetTunnelResponse) String() string { return proto.CompactTextString(m) }
func (*GetTunnelResponse) ProtoMessage()    {}
func (*GetTunnelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{2}
}

func (m *GetTunnelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTunnelResponse.Unmarshal(m, b)
}
func (m *GetTunnelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTunnelResponse.Marshal(b, m, deterministic)
}
func (m *GetTunnelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTunnelResponse.Merge(m, src)
}
func (m *GetTunnelResponse) XXX_Size() int {
	return xxx_messageInfo_GetTunnelResponse.Size(m)
}
func (m *GetTunnelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTunnelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTunnelResponse proto.InternalMessageInfo

func (m *GetTunnelResponse) GetMessage() []*Tunnel {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *GetTunnelResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type SetTunnelResponse struct {
	Message              *Tunnel  `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetTunnelResponse) Reset()         { *m = SetTunnelResponse{} }
func (m *SetTunnelResponse) String() string { return proto.CompactTextString(m) }
func (*SetTunnelResponse) ProtoMessage()    {}
func (*SetTunnelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{3}
}

func (m *SetTunnelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetTunnelResponse.Unmarshal(m, b)
}
func (m *SetTunnelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetTunnelResponse.Marshal(b, m, deterministic)
}
func (m *SetTunnelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetTunnelResponse.Merge(m, src)
}
func (m *SetTunnelResponse) XXX_Size() int {
	return xxx_messageInfo_SetTunnelResponse.Size(m)
}
func (m *SetTunnelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetTunnelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetTunnelResponse proto.InternalMessageInfo

func (m *SetTunnelResponse) GetMessage() *Tunnel {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SetTunnelResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type DeleteTunnelResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTunnelResponse) Reset()         { *m = DeleteTunnelResponse{} }
func (m *DeleteTunnelResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTunnelResponse) ProtoMessage()    {}
func (*DeleteTunnelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{4}
}

func (m *DeleteTunnelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTunnelResponse.Unmarshal(m, b)
}
func (m *DeleteTunnelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTunnelResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTunnelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTunnelResponse.Merge(m, src)
}
func (m *DeleteTunnelResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTunnelResponse.Size(m)
}
func (m *DeleteTunnelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTunnelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTunnelResponse proto.InternalMessageInfo

func (m *DeleteTunnelResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DeleteTunnelResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
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
	return fileDescriptor_6f51ddaa7891a711, []int{5}
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

func init() {
	proto.RegisterType((*SetTunnelRequest)(nil), "tunnelservice.SetTunnelRequest")
	proto.RegisterType((*Tunnel)(nil), "tunnelservice.Tunnel")
	proto.RegisterType((*GetTunnelResponse)(nil), "tunnelservice.GetTunnelResponse")
	proto.RegisterType((*SetTunnelResponse)(nil), "tunnelservice.SetTunnelResponse")
	proto.RegisterType((*DeleteTunnelResponse)(nil), "tunnelservice.DeleteTunnelResponse")
	proto.RegisterType((*Empty)(nil), "tunnelservice.Empty")
}

func init() {
	proto.RegisterFile("tunnel.proto", fileDescriptor_6f51ddaa7891a711)
}

var fileDescriptor_6f51ddaa7891a711 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5d, 0x4b, 0x3a, 0x41,
	0x14, 0xc6, 0xff, 0xfe, 0x75, 0x95, 0x3d, 0x9a, 0xe5, 0x60, 0xb1, 0x44, 0x90, 0x6c, 0x10, 0x5e,
	0x29, 0xd8, 0x27, 0xc8, 0x5e, 0x0c, 0x02, 0x09, 0x37, 0xbb, 0xa8, 0x6e, 0x56, 0xf7, 0x20, 0x82,
	0xfb, 0xd2, 0xce, 0x31, 0xf2, 0xa6, 0xcf, 0xd4, 0x47, 0x8c, 0x99, 0xf1, 0x75, 0x94, 0x2c, 0xdc,
	0xbb, 0x99, 0xe7, 0x3c, 0xf3, 0x9b, 0x67, 0x1f, 0x86, 0x85, 0x02, 0x8d, 0x83, 0x00, 0x47, 0xb5,
	0x28, 0x0e, 0x29, 0x64, 0x7b, 0x6a, 0xc7, 0x31, 0x7e, 0x1f, 0xf6, 0xd1, 0x3e, 0x87, 0x03, 0x07,
	0xe9, 0x51, 0x6a, 0x1d, 0x7c, 0x1b, 0x23, 0x27, 0xc6, 0x20, 0x13, 0xb8, 0x3e, 0x5a, 0xa9, 0x4a,
	0xaa, 0x6a, 0x76, 0xe4, 0xda, 0xfe, 0x84, 0xac, 0x32, 0x6d, 0x9a, 0x0a, 0x2d, 0x0a, 0x63, 0xb2,
	0xfe, 0x2b, 0x4d, 0xac, 0xd9, 0x09, 0x98, 0xfd, 0x18, 0x5d, 0x42, 0xef, 0x92, 0xac, 0xb4, 0x1c,
	0x2c, 0x04, 0x76, 0x04, 0x59, 0x2f, 0xf4, 0xdd, 0x61, 0x60, 0x65, 0xe4, 0x68, 0xba, 0x13, 0x3a,
	0x27, 0x97, 0xc6, 0xdc, 0x32, 0x94, 0xae, 0x76, 0xf6, 0x2b, 0x94, 0x5a, 0x8b, 0x9c, 0x3c, 0x0a,
	0x03, 0x8e, 0xac, 0x0e, 0x39, 0x1f, 0x39, 0x77, 0x07, 0x22, 0x4d, 0xba, 0x9a, 0x6f, 0x1c, 0xd6,
	0x56, 0xbe, 0xae, 0x36, 0xf5, 0xcf, 0x5c, 0x4b, 0x74, 0x91, 0xd4, 0x58, 0xa6, 0x3b, 0x3f, 0xd3,
	0x53, 0x3b, 0xd0, 0xef, 0xa0, 0x7c, 0x8d, 0x23, 0x24, 0xd4, 0x2e, 0xb0, 0x56, 0x2f, 0x30, 0xb7,
	0x93, 0x72, 0x60, 0xdc, 0xf8, 0x11, 0x4d, 0x1a, 0x5f, 0x06, 0xb0, 0xf6, 0x60, 0x18, 0x7c, 0x28,
	0xa4, 0xa3, 0x12, 0xb1, 0x5b, 0x80, 0x79, 0x4b, 0x9c, 0x95, 0xb5, 0xbc, 0xf2, 0xe8, 0x71, 0x45,
	0x53, 0xd7, 0x6a, 0xb5, 0xff, 0xb1, 0x07, 0x30, 0xe7, 0x7d, 0xb0, 0x53, 0xed, 0x80, 0xfe, 0x5e,
	0xd6, 0x88, 0xce, 0x06, 0xe2, 0x13, 0x14, 0x96, 0x3b, 0xd8, 0x0e, 0x3d, 0xd3, 0x0c, 0x9b, 0x1a,
	0x54, 0x49, 0x5b, 0x48, 0xcd, 0x49, 0x5b, 0x3c, 0xc3, 0x44, 0x92, 0xde, 0xc3, 0x7e, 0x0b, 0xe9,
	0xca, 0xf5, 0xbc, 0xc9, 0xee, 0x45, 0x76, 0xa1, 0xe8, 0xac, 0xc0, 0x92, 0xc9, 0xf8, 0x02, 0x25,
	0xd5, 0xc7, 0x9f, 0xc8, 0xbf, 0xac, 0xb4, 0x0b, 0xc5, 0x59, 0x01, 0x09, 0xf6, 0xda, 0xcc, 0x3f,
	0x9b, 0x51, 0xaf, 0xae, 0x7c, 0xbd, 0xac, 0xfc, 0x19, 0x5d, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xed, 0x5d, 0x5e, 0x4d, 0x9c, 0x04, 0x00, 0x00,
}
