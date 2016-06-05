// Code generated by protoc-gen-go.
// source: Task.proto
// DO NOT EDIT!

package Transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task struct {
	Uuid         string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Finisheduuid string `protobuf:"bytes,2,opt,name=finisheduuid" json:"finisheduuid,omitempty"`
	State        int32  `protobuf:"varint,3,opt,name=state" json:"state,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*Task)(nil), "Transport.Task")
}

var fileDescriptor1 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x0a, 0x49, 0x2c, 0xce,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x29, 0x4a, 0xcc, 0x2b, 0x2e, 0xc8, 0x2f,
	0x2a, 0x51, 0x4a, 0xe1, 0x62, 0x01, 0x49, 0x08, 0x09, 0x71, 0xb1, 0x94, 0x96, 0x66, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x4a, 0x5c, 0x3c, 0x69, 0x99, 0x79, 0x99,
	0xc5, 0x19, 0xa9, 0x29, 0x60, 0x39, 0x26, 0xb0, 0x1c, 0x8a, 0x98, 0x90, 0x08, 0x17, 0x6b, 0x71,
	0x49, 0x62, 0x49, 0xaa, 0x04, 0x33, 0x50, 0x92, 0x35, 0x08, 0xc2, 0x11, 0xe2, 0xe3, 0x62, 0x02,
	0xaa, 0x67, 0x01, 0x0b, 0x01, 0x59, 0x49, 0x6c, 0x60, 0x7b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0xfe, 0x85, 0xbe, 0x85, 0x00, 0x00, 0x00,
}