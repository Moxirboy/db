// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teacher.proto

package teacher

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

type AddTestRequest struct {
	QuestionId           string   `protobuf:"bytes,1,opt,name=QuestionId,proto3" json:"QuestionId,omitempty"`
	QuestionText         string   `protobuf:"bytes,2,opt,name=QuestionText,proto3" json:"QuestionText,omitempty"`
	TeacherID            string   `protobuf:"bytes,3,opt,name=TeacherID,proto3" json:"TeacherID,omitempty"`
	AnswerID             string   `protobuf:"bytes,4,opt,name=AnswerID,proto3" json:"AnswerID,omitempty"`
	ChoiceID             string   `protobuf:"bytes,5,opt,name=ChoiceID,proto3" json:"ChoiceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTestRequest) Reset()         { *m = AddTestRequest{} }
func (m *AddTestRequest) String() string { return proto.CompactTextString(m) }
func (*AddTestRequest) ProtoMessage()    {}
func (*AddTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c4d93a072aeccf3, []int{0}
}

func (m *AddTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTestRequest.Unmarshal(m, b)
}
func (m *AddTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTestRequest.Marshal(b, m, deterministic)
}
func (m *AddTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTestRequest.Merge(m, src)
}
func (m *AddTestRequest) XXX_Size() int {
	return xxx_messageInfo_AddTestRequest.Size(m)
}
func (m *AddTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTestRequest proto.InternalMessageInfo

func (m *AddTestRequest) GetQuestionId() string {
	if m != nil {
		return m.QuestionId
	}
	return ""
}

func (m *AddTestRequest) GetQuestionText() string {
	if m != nil {
		return m.QuestionText
	}
	return ""
}

func (m *AddTestRequest) GetTeacherID() string {
	if m != nil {
		return m.TeacherID
	}
	return ""
}

func (m *AddTestRequest) GetAnswerID() string {
	if m != nil {
		return m.AnswerID
	}
	return ""
}

func (m *AddTestRequest) GetChoiceID() string {
	if m != nil {
		return m.ChoiceID
	}
	return ""
}

type CreateClass struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TeacherID            string   `protobuf:"bytes,2,opt,name=TeacherID,proto3" json:"TeacherID,omitempty"`
	ClassID              string   `protobuf:"bytes,3,opt,name=ClassID,proto3" json:"ClassID,omitempty"`
	Period               string   `protobuf:"bytes,4,opt,name=Period,proto3" json:"Period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClass) Reset()         { *m = CreateClass{} }
func (m *CreateClass) String() string { return proto.CompactTextString(m) }
func (*CreateClass) ProtoMessage()    {}
func (*CreateClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c4d93a072aeccf3, []int{1}
}

func (m *CreateClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClass.Unmarshal(m, b)
}
func (m *CreateClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClass.Marshal(b, m, deterministic)
}
func (m *CreateClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClass.Merge(m, src)
}
func (m *CreateClass) XXX_Size() int {
	return xxx_messageInfo_CreateClass.Size(m)
}
func (m *CreateClass) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClass.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClass proto.InternalMessageInfo

func (m *CreateClass) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CreateClass) GetTeacherID() string {
	if m != nil {
		return m.TeacherID
	}
	return ""
}

func (m *CreateClass) GetClassID() string {
	if m != nil {
		return m.ClassID
	}
	return ""
}

func (m *CreateClass) GetPeriod() string {
	if m != nil {
		return m.Period
	}
	return ""
}

type Response struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c4d93a072aeccf3, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*AddTestRequest)(nil), "teacher.AddTestRequest")
	proto.RegisterType((*CreateClass)(nil), "teacher.CreateClass")
	proto.RegisterType((*Response)(nil), "teacher.Response")
}

func init() {
	proto.RegisterFile("teacher.proto", fileDescriptor_6c4d93a072aeccf3)
}

var fileDescriptor_6c4d93a072aeccf3 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0x84, 0x30,
	0x10, 0xc7, 0x03, 0xea, 0xb2, 0xcc, 0xea, 0x26, 0x36, 0x46, 0x09, 0x7e, 0xc4, 0x70, 0xf2, 0x54,
	0x12, 0x3f, 0x1e, 0x60, 0x85, 0x4b, 0x6f, 0xca, 0xe2, 0xc5, 0x1b, 0xd2, 0x49, 0x96, 0x44, 0xe9,
	0xda, 0xd6, 0x8f, 0xf8, 0x48, 0x3e, 0xa5, 0xd9, 0x6e, 0x5b, 0x17, 0xe3, 0x8d, 0xdf, 0xfc, 0x66,
	0xf8, 0x0f, 0x03, 0xec, 0x69, 0x6c, 0xda, 0x05, 0x4a, 0xba, 0x94, 0x42, 0x0b, 0x12, 0x59, 0xcc,
	0xbe, 0x03, 0x98, 0xce, 0x38, 0xaf, 0x51, 0xe9, 0x0a, 0x5f, 0xdf, 0x50, 0x69, 0x72, 0x06, 0x70,
	0xbf, 0x7a, 0xe8, 0x44, 0xcf, 0x78, 0x12, 0x9c, 0x07, 0x17, 0x71, 0xb5, 0x51, 0x21, 0x19, 0xec,
	0x3a, 0xaa, 0xf1, 0x53, 0x27, 0xa1, 0xe9, 0x18, 0xd4, 0xc8, 0x09, 0xc4, 0xf5, 0x3a, 0x81, 0x95,
	0xc9, 0x96, 0x69, 0xf8, 0x2d, 0x90, 0x14, 0xc6, 0xb3, 0x5e, 0x7d, 0x18, 0xb9, 0x6d, 0xa4, 0xe7,
	0x95, 0x2b, 0x16, 0xa2, 0x6b, 0x91, 0x95, 0xc9, 0xce, 0xda, 0x39, 0xce, 0x5e, 0x60, 0x52, 0x48,
	0x6c, 0x34, 0x16, 0xcf, 0x8d, 0x52, 0x64, 0x0a, 0x21, 0x2b, 0xed, 0x82, 0x21, 0x2b, 0x87, 0xa1,
	0xe1, 0xdf, 0xd0, 0x04, 0x22, 0x33, 0xe6, 0x17, 0x72, 0x48, 0x0e, 0x61, 0x74, 0x87, 0xb2, 0x13,
	0xdc, 0x2e, 0x63, 0x29, 0x4b, 0x61, 0x5c, 0xa1, 0x5a, 0x8a, 0x5e, 0xa1, 0xc9, 0xe2, 0x3e, 0x8b,
	0x5f, 0x7e, 0xc1, 0xe4, 0x41, 0xa1, 0x9c, 0xa3, 0x7c, 0xef, 0x5a, 0x24, 0x37, 0x10, 0xd9, 0x2b,
	0x92, 0x23, 0xea, 0x4e, 0x3d, 0xbc, 0x6b, 0xba, 0xef, 0x85, 0x7f, 0xeb, 0x35, 0xc4, 0x73, 0xdd,
	0x48, 0x6d, 0x06, 0x0f, 0xbc, 0xdf, 0xf8, 0xc8, 0x7f, 0xa6, 0x6e, 0x4f, 0x1f, 0x8f, 0x29, 0xcd,
	0x29, 0xcd, 0xad, 0xc9, 0xcd, 0x3f, 0x75, 0xf4, 0x34, 0x32, 0x78, 0xf5, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0x48, 0x7e, 0x66, 0xf3, 0x01, 0x00, 0x00,
}