// Code generated by protoc-gen-go.
// source: entry.proto
// DO NOT EDIT!

package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EntryKind int32

const (
	EntryKind_Directory EntryKind = 0
	EntryKind_File      EntryKind = 1
)

var EntryKind_name = map[int32]string{
	0: "Directory",
	1: "File",
}
var EntryKind_value = map[string]int32{
	"Directory": 0,
	"File":      1,
}

func (x EntryKind) String() string {
	return proto.EnumName(EntryKind_name, int32(x))
}
func (EntryKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type Entry struct {
	Kind       EntryKind         `protobuf:"varint,1,opt,name=kind,enum=sync.EntryKind" json:"kind,omitempty"`
	Executable bool              `protobuf:"varint,2,opt,name=executable" json:"executable,omitempty"`
	Digest     []byte            `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	Contents   map[string]*Entry `protobuf:"bytes,4,rep,name=contents" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Entry) GetKind() EntryKind {
	if m != nil {
		return m.Kind
	}
	return EntryKind_Directory
}

func (m *Entry) GetExecutable() bool {
	if m != nil {
		return m.Executable
	}
	return false
}

func (m *Entry) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *Entry) GetContents() map[string]*Entry {
	if m != nil {
		return m.Contents
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "sync.Entry")
	proto.RegisterEnum("sync.EntryKind", EntryKind_name, EntryKind_value)
}

func init() { proto.RegisterFile("entry.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x2b, 0x29,
	0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xae, 0xcc, 0x4b, 0x56, 0x7a, 0xca,
	0xc8, 0xc5, 0xea, 0x0a, 0x12, 0x15, 0x52, 0xe6, 0x62, 0xc9, 0xce, 0xcc, 0x4b, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x33, 0xe2, 0xd7, 0x03, 0x49, 0xeb, 0x81, 0xa5, 0xbc, 0x33, 0xf3, 0x52, 0x82,
	0xc0, 0x92, 0x42, 0x72, 0x5c, 0x5c, 0xa9, 0x15, 0xa9, 0xc9, 0xa5, 0x25, 0x89, 0x49, 0x39, 0xa9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x48, 0x22, 0x42, 0x62, 0x5c, 0x6c, 0x29, 0x99, 0xe9,
	0xa9, 0xc5, 0x25, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x50, 0x9e, 0x90, 0x29, 0x17, 0x47,
	0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x49, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x24,
	0x92, 0x05, 0x7a, 0xce, 0x50, 0x39, 0x30, 0x2f, 0x08, 0xae, 0x54, 0xca, 0x83, 0x8b, 0x17, 0x45,
	0x4a, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x12, 0xec, 0x46, 0xce, 0x20, 0x10, 0x53, 0x48, 0x91,
	0x8b, 0xb5, 0x2c, 0x31, 0xa7, 0x14, 0xe2, 0x18, 0x6e, 0x23, 0x6e, 0x24, 0x63, 0x83, 0x20, 0x32,
	0x56, 0x4c, 0x16, 0x8c, 0x5a, 0x2a, 0x5c, 0x9c, 0x70, 0xbf, 0x08, 0xf1, 0x72, 0x71, 0xba, 0x64,
	0x16, 0xa5, 0x26, 0x97, 0xe4, 0x17, 0x55, 0x0a, 0x30, 0x08, 0x71, 0x70, 0xb1, 0xb8, 0x65, 0xe6,
	0xa4, 0x0a, 0x30, 0x26, 0xb1, 0x81, 0x83, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x5d,
	0xb1, 0x20, 0x29, 0x01, 0x00, 0x00,
}