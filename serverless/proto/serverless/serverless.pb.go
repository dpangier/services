// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/serverless/proto/serverless/serverless.proto

package go_micro_service_serverless

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

type App struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Language             string   `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *App) Reset()         { *m = App{} }
func (m *App) String() string { return proto.CompactTextString(m) }
func (*App) ProtoMessage()    {}
func (*App) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{0}
}

func (m *App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_App.Unmarshal(m, b)
}
func (m *App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_App.Marshal(b, m, deterministic)
}
func (m *App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_App.Merge(m, src)
}
func (m *App) XXX_Size() int {
	return xxx_messageInfo_App.Size(m)
}
func (m *App) XXX_DiscardUnknown() {
	xxx_messageInfo_App.DiscardUnknown(m)
}

var xxx_messageInfo_App proto.InternalMessageInfo

func (m *App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *App) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *App) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *App) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type CreateRequest struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	App                  *App     `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{4}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Apps                 []*App   `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e0f1dc5d38c37ce, []int{6}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetApps() []*App {
	if m != nil {
		return m.Apps
	}
	return nil
}

func init() {
	proto.RegisterType((*App)(nil), "go.micro.service.serverless.App")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.service.serverless.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.service.serverless.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.service.serverless.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.service.serverless.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.service.serverless.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.service.serverless.ListResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/serverless/proto/serverless/serverless.proto", fileDescriptor_1e0f1dc5d38c37ce)
}

var fileDescriptor_1e0f1dc5d38c37ce = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0x07, 0xda, 0xf0, 0x7d, 0x0e, 0x62, 0xcc, 0x1c, 0xcc, 0x06, 0x2f, 0xa4, 0x27, 0xd4, 0x64,
	0x49, 0xd0, 0x17, 0x20, 0x10, 0x4f, 0x9e, 0xb8, 0x7b, 0x58, 0x9a, 0x49, 0x6d, 0x02, 0xdd, 0x71,
	0xa7, 0xe5, 0xb5, 0x7c, 0x45, 0xd3, 0x6d, 0x2b, 0xe0, 0xa1, 0x34, 0xf1, 0x36, 0xb3, 0xbf, 0x7f,
	0x9d, 0x5f, 0x0a, 0xaf, 0x49, 0x9a, 0x7f, 0x14, 0x5b, 0x1d, 0xdb, 0xfd, 0x7c, 0x9f, 0xc6, 0xce,
	0xce, 0x85, 0xdc, 0x21, 0x8d, 0x49, 0xfc, 0x40, 0x6e, 0x47, 0x22, 0x73, 0x76, 0x36, 0xb7, 0xa7,
	0x0f, 0xc7, 0x51, 0x7b, 0x0c, 0xef, 0x13, 0xab, 0xbd, 0x5e, 0xd7, 0x7a, 0x7d, 0xa4, 0x44, 0x09,
	0x04, 0x4b, 0x66, 0x44, 0x08, 0x33, 0xb3, 0x27, 0xd5, 0x9f, 0xf6, 0x67, 0x57, 0x1b, 0x3f, 0xa3,
	0x82, 0x7f, 0x07, 0x72, 0x92, 0xda, 0x4c, 0x0d, 0xfc, 0x73, 0xb3, 0xe2, 0x1d, 0x0c, 0xc5, 0x16,
	0x2e, 0x26, 0x15, 0x78, 0xa0, 0xde, 0x70, 0x02, 0xff, 0x77, 0x26, 0x4b, 0x0a, 0x93, 0x90, 0x0a,
	0x3d, 0xf2, 0xb3, 0x47, 0x2b, 0x18, 0xaf, 0x1c, 0x99, 0x9c, 0x36, 0xf4, 0x59, 0x90, 0xe4, 0xb8,
	0x80, 0xc0, 0x30, 0xfb, 0xc4, 0xd1, 0x62, 0xaa, 0x5b, 0x3e, 0x52, 0x2f, 0x99, 0x37, 0x25, 0x39,
	0xba, 0x85, 0x9b, 0xc6, 0x44, 0xd8, 0x66, 0xe2, 0x6d, 0xd7, 0xb4, 0xa3, 0x3f, 0xdb, 0x36, 0x26,
	0xb5, 0xed, 0x18, 0x46, 0x6f, 0xa9, 0xe4, 0xb5, 0x69, 0xb4, 0x86, 0xeb, 0x6a, 0xad, 0x60, 0x7c,
	0x81, 0xd0, 0x30, 0x8b, 0xea, 0x4f, 0x83, 0x4e, 0x29, 0x9e, 0xbd, 0xf8, 0x1a, 0x40, 0xb8, 0x64,
	0x16, 0x8c, 0x61, 0x58, 0x9d, 0x81, 0x8f, 0xad, 0xd2, 0xb3, 0xc2, 0x26, 0x4f, 0x9d, 0xb8, 0xf5,
	0x01, 0xbd, 0x32, 0xa4, 0x3a, 0xea, 0x42, 0xc8, 0x59, 0x7d, 0x17, 0x42, 0x7e, 0xb5, 0xd4, 0xc3,
	0x77, 0x08, 0xcb, 0x62, 0x70, 0xd6, 0x2a, 0x3b, 0xa9, 0x72, 0xf2, 0xd0, 0x81, 0xd9, 0xd8, 0x6f,
	0x87, 0xfe, 0x0f, 0x7e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x71, 0xc5, 0x5e, 0x80, 0x0b, 0x03,
	0x00, 0x00,
}
