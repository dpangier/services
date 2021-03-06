// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/kubernetes/service/proto/kubernetes.proto

package go_micro_service_kubernetes

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

type CreateNamespaceRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceRequest) Reset()         { *m = CreateNamespaceRequest{} }
func (m *CreateNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceRequest) ProtoMessage()    {}
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{0}
}

func (m *CreateNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceRequest.Unmarshal(m, b)
}
func (m *CreateNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceRequest.Merge(m, src)
}
func (m *CreateNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceRequest.Size(m)
}
func (m *CreateNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceRequest proto.InternalMessageInfo

func (m *CreateNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNamespaceResponse) Reset()         { *m = CreateNamespaceResponse{} }
func (m *CreateNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNamespaceResponse) ProtoMessage()    {}
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{1}
}

func (m *CreateNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNamespaceResponse.Unmarshal(m, b)
}
func (m *CreateNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *CreateNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNamespaceResponse.Merge(m, src)
}
func (m *CreateNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNamespaceResponse.Size(m)
}
func (m *CreateNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNamespaceResponse proto.InternalMessageInfo

type CreateImagePullSecretRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImagePullSecretRequest) Reset()         { *m = CreateImagePullSecretRequest{} }
func (m *CreateImagePullSecretRequest) String() string { return proto.CompactTextString(m) }
func (*CreateImagePullSecretRequest) ProtoMessage()    {}
func (*CreateImagePullSecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{2}
}

func (m *CreateImagePullSecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImagePullSecretRequest.Unmarshal(m, b)
}
func (m *CreateImagePullSecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImagePullSecretRequest.Marshal(b, m, deterministic)
}
func (m *CreateImagePullSecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImagePullSecretRequest.Merge(m, src)
}
func (m *CreateImagePullSecretRequest) XXX_Size() int {
	return xxx_messageInfo_CreateImagePullSecretRequest.Size(m)
}
func (m *CreateImagePullSecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImagePullSecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImagePullSecretRequest proto.InternalMessageInfo

func (m *CreateImagePullSecretRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateImagePullSecretRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateImagePullSecretRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreateImagePullSecretResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateImagePullSecretResponse) Reset()         { *m = CreateImagePullSecretResponse{} }
func (m *CreateImagePullSecretResponse) String() string { return proto.CompactTextString(m) }
func (*CreateImagePullSecretResponse) ProtoMessage()    {}
func (*CreateImagePullSecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{3}
}

func (m *CreateImagePullSecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateImagePullSecretResponse.Unmarshal(m, b)
}
func (m *CreateImagePullSecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateImagePullSecretResponse.Marshal(b, m, deterministic)
}
func (m *CreateImagePullSecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateImagePullSecretResponse.Merge(m, src)
}
func (m *CreateImagePullSecretResponse) XXX_Size() int {
	return xxx_messageInfo_CreateImagePullSecretResponse.Size(m)
}
func (m *CreateImagePullSecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateImagePullSecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateImagePullSecretResponse proto.InternalMessageInfo

type CreateServiceAccountRequest struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ImagePullSecrets     []string `protobuf:"bytes,2,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceAccountRequest) Reset()         { *m = CreateServiceAccountRequest{} }
func (m *CreateServiceAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceAccountRequest) ProtoMessage()    {}
func (*CreateServiceAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{4}
}

func (m *CreateServiceAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceAccountRequest.Unmarshal(m, b)
}
func (m *CreateServiceAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceAccountRequest.Merge(m, src)
}
func (m *CreateServiceAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceAccountRequest.Size(m)
}
func (m *CreateServiceAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceAccountRequest proto.InternalMessageInfo

func (m *CreateServiceAccountRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateServiceAccountRequest) GetImagePullSecrets() []string {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

type CreateServiceAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceAccountResponse) Reset()         { *m = CreateServiceAccountResponse{} }
func (m *CreateServiceAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceAccountResponse) ProtoMessage()    {}
func (*CreateServiceAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09bf30676138af59, []int{5}
}

func (m *CreateServiceAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceAccountResponse.Unmarshal(m, b)
}
func (m *CreateServiceAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceAccountResponse.Merge(m, src)
}
func (m *CreateServiceAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceAccountResponse.Size(m)
}
func (m *CreateServiceAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateNamespaceRequest)(nil), "go.micro.service.kubernetes.CreateNamespaceRequest")
	proto.RegisterType((*CreateNamespaceResponse)(nil), "go.micro.service.kubernetes.CreateNamespaceResponse")
	proto.RegisterType((*CreateImagePullSecretRequest)(nil), "go.micro.service.kubernetes.CreateImagePullSecretRequest")
	proto.RegisterType((*CreateImagePullSecretResponse)(nil), "go.micro.service.kubernetes.CreateImagePullSecretResponse")
	proto.RegisterType((*CreateServiceAccountRequest)(nil), "go.micro.service.kubernetes.CreateServiceAccountRequest")
	proto.RegisterType((*CreateServiceAccountResponse)(nil), "go.micro.service.kubernetes.CreateServiceAccountResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/kubernetes/service/proto/kubernetes.proto", fileDescriptor_09bf30676138af59)
}

var fileDescriptor_09bf30676138af59 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x4e, 0xc2, 0x40,
	0x14, 0xb5, 0xa0, 0x26, 0xdc, 0x8d, 0x66, 0x82, 0x5a, 0x01, 0x95, 0x74, 0xc5, 0x82, 0x4c, 0x13,
	0x71, 0x21, 0xee, 0x0c, 0x2b, 0x63, 0x62, 0x0c, 0x7c, 0x00, 0x29, 0x93, 0x2b, 0x4e, 0x68, 0x67,
	0xea, 0x3c, 0x5c, 0xfa, 0x07, 0x26, 0xfe, 0x89, 0xbf, 0x68, 0x9c, 0x52, 0x90, 0x5a, 0x09, 0xb2,
	0x6b, 0xef, 0xe3, 0x9c, 0x73, 0xcf, 0xc9, 0xc0, 0x60, 0xca, 0xcd, 0xb3, 0x9d, 0x50, 0x26, 0x93,
	0x30, 0xe1, 0x4c, 0xc9, 0x50, 0xa3, 0x7a, 0xe5, 0x0c, 0x75, 0x38, 0xb3, 0x13, 0x54, 0x02, 0x0d,
	0xea, 0xbc, 0x16, 0xa6, 0x4a, 0x1a, 0xf9, 0xa3, 0x41, 0x5d, 0x81, 0x34, 0xa7, 0x92, 0xba, 0x65,
	0x3a, 0x1f, 0xa4, 0xcb, 0x91, 0xa0, 0x0b, 0xc7, 0x03, 0x85, 0x91, 0xc1, 0x87, 0x28, 0x41, 0x9d,
	0x46, 0x0c, 0x87, 0xf8, 0x62, 0x51, 0x1b, 0x42, 0x60, 0x57, 0x44, 0x09, 0xfa, 0x5e, 0xdb, 0xeb,
	0xd4, 0x86, 0xee, 0x3b, 0x38, 0x85, 0x93, 0x5f, 0xd3, 0x3a, 0x95, 0x42, 0x63, 0xf0, 0x04, 0xad,
	0xac, 0x75, 0x97, 0x44, 0x53, 0x7c, 0xb4, 0x71, 0x3c, 0x42, 0xa6, 0xd0, 0xac, 0x81, 0x23, 0x2d,
	0xa8, 0x89, 0x1c, 0xc8, 0xaf, 0xb8, 0xc6, 0xb2, 0x40, 0xea, 0xb0, 0x67, 0xe4, 0x0c, 0x85, 0x5f,
	0x75, 0x9d, 0xec, 0x27, 0xb8, 0x80, 0xb3, 0x3f, 0x78, 0xe6, 0x42, 0x38, 0x34, 0xb3, 0x81, 0x51,
	0x76, 0xed, 0x2d, 0x63, 0xd2, 0x8a, 0x85, 0x8e, 0x15, 0x4e, 0xaf, 0xc8, 0xd9, 0x05, 0xc2, 0xbf,
	0x71, 0xc7, 0xa9, 0x8d, 0xe3, 0xb1, 0x76, 0xc8, 0xda, 0xaf, 0xb4, 0xab, 0x9d, 0xda, 0xf0, 0x90,
	0xaf, 0x32, 0xea, 0xe0, 0x3c, 0xbf, 0xb9, 0x48, 0x95, 0x49, 0xb9, 0xfc, 0xac, 0x02, 0xdc, 0x2f,
	0xbc, 0x26, 0x6f, 0x70, 0x50, 0x70, 0x8f, 0xf4, 0xe8, 0x9a, 0x70, 0x68, 0x79, 0x32, 0x8d, 0xab,
	0xff, 0x2d, 0xcd, 0x7d, 0xd9, 0x21, 0x1f, 0x1e, 0x1c, 0x95, 0x7a, 0x47, 0xfa, 0x1b, 0x20, 0x96,
	0xe7, 0xda, 0xb8, 0xd9, 0x66, 0x75, 0x21, 0xe9, 0xdd, 0x83, 0x7a, 0x99, 0x85, 0xe4, 0x7a, 0x03,
	0xd8, 0xd2, 0x80, 0x1b, 0xfd, 0x2d, 0x36, 0x73, 0x3d, 0x93, 0x7d, 0xf7, 0x64, 0x7a, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x55, 0x75, 0x0b, 0x10, 0x79, 0x03, 0x00, 0x00,
}
