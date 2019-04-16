// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backup/data.proto

package backup

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

type Backup struct {
	Data                 map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6fa0d138206cd35, []int{0}
}

func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (m *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(m, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetData() map[string][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "backup.Backup")
	proto.RegisterMapType((map[string][]byte)(nil), "backup.Backup.DataEntry")
}

func init() { proto.RegisterFile("backup/data.proto", fileDescriptor_a6fa0d138206cd35) }

var fileDescriptor_a6fa0d138206cd35 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4a, 0x4c, 0xce,
	0x2e, 0x2d, 0xd0, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83,
	0x08, 0x29, 0xe5, 0x73, 0xb1, 0x39, 0x81, 0x59, 0x42, 0x3a, 0x5c, 0x2c, 0x20, 0x79, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x09, 0x3d, 0x88, 0x02, 0x3d, 0x88, 0xac, 0x9e, 0x4b, 0x62, 0x49,
	0xa2, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x58, 0x95, 0x94, 0x39, 0x17, 0x27, 0x5c, 0x48, 0x48,
	0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14, 0x12,
	0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x70,
	0xac, 0x98, 0x2c, 0x18, 0x93, 0xd8, 0xc0, 0xf6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0x3b, 0xd1, 0x9c, 0x94, 0x00, 0x00, 0x00,
}