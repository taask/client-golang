// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package model

import (
	fmt "fmt"
	simplcrypto "github.com/cohix/simplcrypto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Meta                 *TaskMeta            `protobuf:"bytes,1,opt,name=Meta,proto3" json:"Meta,omitempty"`
	UUID                 string               `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Kind                 string               `protobuf:"bytes,3,opt,name=Kind,proto3" json:"Kind,omitempty"`
	Status               string               `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	EncBody              *simplcrypto.Message `protobuf:"bytes,5,opt,name=EncBody,proto3" json:"EncBody,omitempty"`
	EncResult            *simplcrypto.Message `protobuf:"bytes,6,opt,name=EncResult,proto3" json:"EncResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetMeta() *TaskMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Task) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Task) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Task) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Task) GetEncBody() *simplcrypto.Message {
	if m != nil {
		return m.EncBody
	}
	return nil
}

func (m *Task) GetEncResult() *simplcrypto.Message {
	if m != nil {
		return m.EncResult
	}
	return nil
}

type TaskMeta struct {
	Annotations          []string                        `protobuf:"bytes,1,rep,name=Annotations,proto3" json:"Annotations,omitempty"`
	RunnerUUID           string                          `protobuf:"bytes,2,opt,name=RunnerUUID,proto3" json:"RunnerUUID,omitempty"`
	PartnerUUID          string                          `protobuf:"bytes,3,opt,name=PartnerUUID,proto3" json:"PartnerUUID,omitempty"`
	ClientKeyKID         string                          `protobuf:"bytes,4,opt,name=ClientKeyKID,proto3" json:"ClientKeyKID,omitempty"`
	EncTaskKeys          map[string]*simplcrypto.Message `protobuf:"bytes,5,rep,name=EncTaskKeys,proto3" json:"EncTaskKeys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RetrySeconds         int32                           `protobuf:"varint,6,opt,name=RetrySeconds,proto3" json:"RetrySeconds,omitempty"`
	TimeoutSeconds       int32                           `protobuf:"varint,7,opt,name=TimeoutSeconds,proto3" json:"TimeoutSeconds,omitempty"`
	Version              int32                           `protobuf:"varint,8,opt,name=Version,proto3" json:"Version,omitempty"`
	GroupUUID            string                          `protobuf:"bytes,9,opt,name=GroupUUID,proto3" json:"GroupUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TaskMeta) Reset()         { *m = TaskMeta{} }
func (m *TaskMeta) String() string { return proto.CompactTextString(m) }
func (*TaskMeta) ProtoMessage()    {}
func (*TaskMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{1}
}

func (m *TaskMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskMeta.Unmarshal(m, b)
}
func (m *TaskMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskMeta.Marshal(b, m, deterministic)
}
func (m *TaskMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMeta.Merge(m, src)
}
func (m *TaskMeta) XXX_Size() int {
	return xxx_messageInfo_TaskMeta.Size(m)
}
func (m *TaskMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMeta proto.InternalMessageInfo

func (m *TaskMeta) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskMeta) GetRunnerUUID() string {
	if m != nil {
		return m.RunnerUUID
	}
	return ""
}

func (m *TaskMeta) GetPartnerUUID() string {
	if m != nil {
		return m.PartnerUUID
	}
	return ""
}

func (m *TaskMeta) GetClientKeyKID() string {
	if m != nil {
		return m.ClientKeyKID
	}
	return ""
}

func (m *TaskMeta) GetEncTaskKeys() map[string]*simplcrypto.Message {
	if m != nil {
		return m.EncTaskKeys
	}
	return nil
}

func (m *TaskMeta) GetRetrySeconds() int32 {
	if m != nil {
		return m.RetrySeconds
	}
	return 0
}

func (m *TaskMeta) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *TaskMeta) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TaskMeta) GetGroupUUID() string {
	if m != nil {
		return m.GroupUUID
	}
	return ""
}

type TaskUpdate struct {
	UUID                 string       `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Version              int32        `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	Changes              *TaskChanges `protobuf:"bytes,3,opt,name=Changes,proto3" json:"Changes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TaskUpdate) Reset()         { *m = TaskUpdate{} }
func (m *TaskUpdate) String() string { return proto.CompactTextString(m) }
func (*TaskUpdate) ProtoMessage()    {}
func (*TaskUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{2}
}

func (m *TaskUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskUpdate.Unmarshal(m, b)
}
func (m *TaskUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskUpdate.Marshal(b, m, deterministic)
}
func (m *TaskUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskUpdate.Merge(m, src)
}
func (m *TaskUpdate) XXX_Size() int {
	return xxx_messageInfo_TaskUpdate.Size(m)
}
func (m *TaskUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_TaskUpdate proto.InternalMessageInfo

func (m *TaskUpdate) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *TaskUpdate) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TaskUpdate) GetChanges() *TaskChanges {
	if m != nil {
		return m.Changes
	}
	return nil
}

type TaskChanges struct {
	Status               string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	EncResult            *simplcrypto.Message   `protobuf:"bytes,2,opt,name=EncResult,proto3" json:"EncResult,omitempty"`
	AddedEncTaskKeys     []*simplcrypto.Message `protobuf:"bytes,3,rep,name=AddedEncTaskKeys,proto3" json:"AddedEncTaskKeys,omitempty"`
	RunnerUUID           string                 `protobuf:"bytes,4,opt,name=RunnerUUID,proto3" json:"RunnerUUID,omitempty"`
	RetrySeconds         int32                  `protobuf:"varint,5,opt,name=RetrySeconds,proto3" json:"RetrySeconds,omitempty"`
	PartnerUUID          string                 `protobuf:"bytes,6,opt,name=PartnerUUID,proto3" json:"PartnerUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TaskChanges) Reset()         { *m = TaskChanges{} }
func (m *TaskChanges) String() string { return proto.CompactTextString(m) }
func (*TaskChanges) ProtoMessage()    {}
func (*TaskChanges) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce5d8dd45b4a91ff, []int{3}
}

func (m *TaskChanges) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskChanges.Unmarshal(m, b)
}
func (m *TaskChanges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskChanges.Marshal(b, m, deterministic)
}
func (m *TaskChanges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskChanges.Merge(m, src)
}
func (m *TaskChanges) XXX_Size() int {
	return xxx_messageInfo_TaskChanges.Size(m)
}
func (m *TaskChanges) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskChanges.DiscardUnknown(m)
}

var xxx_messageInfo_TaskChanges proto.InternalMessageInfo

func (m *TaskChanges) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TaskChanges) GetEncResult() *simplcrypto.Message {
	if m != nil {
		return m.EncResult
	}
	return nil
}

func (m *TaskChanges) GetAddedEncTaskKeys() []*simplcrypto.Message {
	if m != nil {
		return m.AddedEncTaskKeys
	}
	return nil
}

func (m *TaskChanges) GetRunnerUUID() string {
	if m != nil {
		return m.RunnerUUID
	}
	return ""
}

func (m *TaskChanges) GetRetrySeconds() int32 {
	if m != nil {
		return m.RetrySeconds
	}
	return 0
}

func (m *TaskChanges) GetPartnerUUID() string {
	if m != nil {
		return m.PartnerUUID
	}
	return ""
}

func init() {
	proto.RegisterType((*Task)(nil), "taask.server.model.Task")
	proto.RegisterType((*TaskMeta)(nil), "taask.server.model.TaskMeta")
	proto.RegisterMapType((map[string]*simplcrypto.Message)(nil), "taask.server.model.TaskMeta.EncTaskKeysEntry")
	proto.RegisterType((*TaskUpdate)(nil), "taask.server.model.TaskUpdate")
	proto.RegisterType((*TaskChanges)(nil), "taask.server.model.TaskChanges")
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_ce5d8dd45b4a91ff) }

var fileDescriptor_ce5d8dd45b4a91ff = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x71, 0x9c, 0x3f, 0xf5, 0xf1, 0x18, 0x41, 0xb0, 0x61, 0x4a, 0xd9, 0x42, 0xc6, 0x46,
	0x60, 0xd4, 0x2e, 0xe9, 0xcd, 0xb6, 0xbb, 0xb4, 0x0d, 0xa3, 0x84, 0xb2, 0xa1, 0x36, 0xbb, 0xd8,
	0x9d, 0x6a, 0x1f, 0x12, 0x93, 0x44, 0x0a, 0x96, 0x5c, 0xe6, 0xeb, 0xbd, 0xd7, 0x1e, 0x65, 0xcf,
	0x32, 0x24, 0xc5, 0xab, 0x92, 0x2c, 0x65, 0x37, 0x41, 0xfa, 0xf2, 0xe9, 0xe8, 0x3b, 0xc7, 0x3f,
	0x04, 0xa0, 0x98, 0x5c, 0xc4, 0xeb, 0x42, 0x28, 0x41, 0x88, 0x62, 0x7a, 0x23, 0xb1, 0x78, 0xc0,
	0x22, 0x5e, 0x89, 0x0c, 0x97, 0xc7, 0x67, 0xb3, 0x5c, 0xcd, 0xcb, 0xfb, 0x38, 0x15, 0xab, 0x24,
	0x15, 0xf3, 0xfc, 0x47, 0x22, 0xf3, 0xd5, 0x7a, 0x99, 0x16, 0xd5, 0x5a, 0x89, 0xc4, 0x9c, 0x4b,
	0x56, 0x28, 0x25, 0x9b, 0xa1, 0xad, 0xd2, 0xff, 0xed, 0x41, 0xf3, 0x8e, 0xc9, 0x05, 0x39, 0x83,
	0xe6, 0x0d, 0x2a, 0x16, 0x79, 0x3d, 0x6f, 0x10, 0x0e, 0x4f, 0xe2, 0xfd, 0xea, 0xb1, 0xf6, 0x69,
	0x0f, 0x35, 0x4e, 0x42, 0xa0, 0x39, 0x9d, 0x5e, 0x5f, 0x45, 0x8d, 0x9e, 0x37, 0x08, 0xa8, 0x59,
	0x6b, 0x6d, 0x92, 0xf3, 0x2c, 0xf2, 0xad, 0xa6, 0xd7, 0xe4, 0x25, 0xb4, 0x6f, 0x15, 0x53, 0xa5,
	0x8c, 0x9a, 0x46, 0xdd, 0xec, 0x48, 0x02, 0x9d, 0x31, 0x4f, 0x2f, 0x44, 0x56, 0x45, 0x2d, 0x73,
	0xe9, 0x8b, 0xd8, 0xa4, 0x8d, 0x6d, 0xdc, 0xf8, 0xc6, 0x06, 0xa5, 0xb5, 0x8b, 0x9c, 0x43, 0x30,
	0xe6, 0x29, 0x45, 0x59, 0x2e, 0x55, 0xd4, 0x7e, 0xea, 0xc8, 0xa3, 0xaf, 0xff, 0xcb, 0x87, 0xa3,
	0x3a, 0x38, 0xe9, 0x41, 0x38, 0xe2, 0x5c, 0x28, 0xa6, 0x72, 0xc1, 0x65, 0xe4, 0xf5, 0xfc, 0x41,
	0x40, 0x5d, 0x89, 0xbc, 0x02, 0xa0, 0x25, 0xe7, 0x58, 0x38, 0xad, 0x39, 0x8a, 0xae, 0xf0, 0x95,
	0x15, 0xaa, 0x36, 0xd8, 0x3e, 0x5d, 0x89, 0xf4, 0xe1, 0xd9, 0xe5, 0x32, 0x47, 0xae, 0x26, 0x58,
	0x4d, 0xae, 0xaf, 0x36, 0x4d, 0x6f, 0x69, 0xe4, 0x0b, 0x84, 0x63, 0x9e, 0xea, 0x58, 0x13, 0xac,
	0x64, 0xd4, 0xea, 0xf9, 0x83, 0x70, 0x78, 0xfa, 0xd4, 0xcc, 0x63, 0xc7, 0x3f, 0xe6, 0xaa, 0xa8,
	0xa8, 0x5b, 0x41, 0x5f, 0x4a, 0x51, 0x15, 0xd5, 0x2d, 0xa6, 0x82, 0x67, 0xd2, 0x4c, 0xa7, 0x45,
	0xb7, 0x34, 0xf2, 0x0e, 0x9e, 0xdf, 0xe5, 0x2b, 0x14, 0xa5, 0xaa, 0x5d, 0x1d, 0xe3, 0xda, 0x51,
	0x49, 0x04, 0x9d, 0x6f, 0x58, 0xc8, 0x5c, 0xf0, 0xe8, 0xc8, 0x18, 0xea, 0x2d, 0x39, 0x81, 0xe0,
	0x73, 0x21, 0xca, 0xb5, 0x69, 0x3d, 0x30, 0x7d, 0x3d, 0x0a, 0xc7, 0x53, 0xe8, 0xee, 0x86, 0x24,
	0x5d, 0xf0, 0x17, 0x58, 0x19, 0xa8, 0x02, 0xaa, 0x97, 0xe4, 0x3d, 0xb4, 0x1e, 0xd8, 0xb2, 0x44,
	0x33, 0xdb, 0x83, 0x1f, 0xd0, 0x7a, 0x3e, 0x35, 0x3e, 0x78, 0xfd, 0x12, 0x40, 0xd7, 0x9c, 0xae,
	0x33, 0xa6, 0xf0, 0x2f, 0x74, 0x9e, 0x03, 0x9d, 0x13, 0xb8, 0xb1, 0x1d, 0xf8, 0x23, 0x74, 0x2e,
	0xe7, 0x8c, 0xcf, 0x50, 0x9a, 0x2f, 0x15, 0x0e, 0x5f, 0x1f, 0x9a, 0xf1, 0xc6, 0x46, 0x6b, 0x7f,
	0xff, 0x67, 0x03, 0x42, 0xe7, 0x0f, 0x4d, 0xb1, 0xb4, 0x14, 0xdb, 0xab, 0x37, 0xbb, 0x6d, 0x28,
	0x1b, 0xff, 0x07, 0x25, 0x19, 0x41, 0x77, 0x94, 0x65, 0x98, 0xb9, 0x10, 0xf8, 0x06, 0x82, 0x03,
	0x67, 0xf7, 0xec, 0x3b, 0xa0, 0x36, 0xf7, 0x40, 0xdd, 0x25, 0xa2, 0xf5, 0x0f, 0x22, 0x76, 0x60,
	0x6e, 0xef, 0xc1, 0x7c, 0xf1, 0xf6, 0xfb, 0x1b, 0xe7, 0x49, 0x31, 0xb3, 0xb3, 0xbf, 0xa7, 0x76,
	0x82, 0x89, 0x99, 0xe0, 0x7d, 0xdb, 0x3c, 0x26, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x37, 0x8c, 0xfa, 0xa0, 0x04, 0x00, 0x00,
}
