// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

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

type CreateRequest struct {
	User                 *Info    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
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

func (m *CreateRequest) GetUser() *Info {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateResult struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResult) Reset()         { *m = CreateResult{} }
func (m *CreateResult) String() string { return proto.CompactTextString(m) }
func (*CreateResult) ProtoMessage()    {}
func (*CreateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *CreateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResult.Unmarshal(m, b)
}
func (m *CreateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResult.Marshal(b, m, deterministic)
}
func (m *CreateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResult.Merge(m, src)
}
func (m *CreateResult) XXX_Size() int {
	return xxx_messageInfo_CreateResult.Size(m)
}
func (m *CreateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResult.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResult proto.InternalMessageInfo

func (m *CreateResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UpdateResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResult) Reset()         { *m = UpdateResult{} }
func (m *UpdateResult) String() string { return proto.CompactTextString(m) }
func (*UpdateResult) ProtoMessage()    {}
func (*UpdateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *UpdateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResult.Unmarshal(m, b)
}
func (m *UpdateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResult.Marshal(b, m, deterministic)
}
func (m *UpdateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResult.Merge(m, src)
}
func (m *UpdateResult) XXX_Size() int {
	return xxx_messageInfo_UpdateResult.Size(m)
}
func (m *UpdateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResult proto.InternalMessageInfo

type ReadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResult struct {
	User                 *Info    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResult) Reset()         { *m = ReadResult{} }
func (m *ReadResult) String() string { return proto.CompactTextString(m) }
func (*ReadResult) ProtoMessage()    {}
func (*ReadResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *ReadResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResult.Unmarshal(m, b)
}
func (m *ReadResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResult.Marshal(b, m, deterministic)
}
func (m *ReadResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResult.Merge(m, src)
}
func (m *ReadResult) XXX_Size() int {
	return xxx_messageInfo_ReadResult.Size(m)
}
func (m *ReadResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResult.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResult proto.InternalMessageInfo

func (m *ReadResult) GetUser() *Info {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
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

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResult) Reset()         { *m = DeleteResult{} }
func (m *DeleteResult) String() string { return proto.CompactTextString(m) }
func (*DeleteResult) ProtoMessage()    {}
func (*DeleteResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *DeleteResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResult.Unmarshal(m, b)
}
func (m *DeleteResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResult.Marshal(b, m, deterministic)
}
func (m *DeleteResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResult.Merge(m, src)
}
func (m *DeleteResult) XXX_Size() int {
	return xxx_messageInfo_DeleteResult.Size(m)
}
func (m *DeleteResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResult.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResult proto.InternalMessageInfo

type Info struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Secret               string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	RoleName             string   `protobuf:"bytes,5,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	Created              string   `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Updated              string   `protobuf:"bytes,7,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{8}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Info) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Info) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Info) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Info) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *Info) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Info) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.user.CreateRequest")
	proto.RegisterType((*CreateResult)(nil), "go.micro.srv.user.CreateResult")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.user.UpdateRequest")
	proto.RegisterType((*UpdateResult)(nil), "go.micro.srv.user.UpdateResult")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.user.ReadRequest")
	proto.RegisterType((*ReadResult)(nil), "go.micro.srv.user.ReadResult")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.user.DeleteRequest")
	proto.RegisterType((*DeleteResult)(nil), "go.micro.srv.user.DeleteResult")
	proto.RegisterType((*Info)(nil), "go.micro.srv.user.Info")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x4b, 0x3b, 0x31,
	0x10, 0xc6, 0xdb, 0xfd, 0x6f, 0xb7, 0xff, 0x4e, 0x5f, 0xc0, 0xe0, 0x4b, 0xa8, 0xf4, 0x85, 0x9c,
	0x04, 0x21, 0x42, 0x3d, 0x89, 0xde, 0x14, 0xc4, 0x83, 0x1e, 0x16, 0x7a, 0x96, 0xda, 0x8c, 0x5a,
	0x68, 0x9b, 0x9a, 0xec, 0xfa, 0xbd, 0xbc, 0xfb, 0xe1, 0x24, 0x93, 0xae, 0x74, 0x35, 0x15, 0x7a,
	0x59, 0x78, 0xf2, 0xcc, 0xfc, 0x98, 0xcc, 0x93, 0x85, 0x83, 0x95, 0xd1, 0x99, 0x3e, 0xcb, 0x2d,
	0x1a, 0xfa, 0x48, 0xd2, 0x6c, 0xef, 0x45, 0xcb, 0xc5, 0x6c, 0x6a, 0xb4, 0xb4, 0xe6, 0x5d, 0x3a,
	0x43, 0x5c, 0x41, 0xfb, 0xda, 0xe0, 0x24, 0xc3, 0x14, 0xdf, 0x72, 0xb4, 0x19, 0x3b, 0x85, 0xd8,
	0x19, 0xbc, 0x3a, 0xac, 0x9e, 0x34, 0x47, 0x47, 0xf2, 0x57, 0x8b, 0xbc, 0x5b, 0x3e, 0xeb, 0x94,
	0x8a, 0x44, 0x1f, 0x5a, 0x45, 0xb7, 0xcd, 0xe7, 0x19, 0xeb, 0x40, 0x34, 0x53, 0xd4, 0xda, 0x48,
	0xa3, 0x99, 0x12, 0x97, 0xd0, 0x1e, 0xaf, 0xd4, 0x06, 0xfd, 0x47, 0x01, 0xeb, 0xc2, 0x7f, 0x07,
	0x5a, 0x4e, 0x16, 0xc8, 0x23, 0x3a, 0xfd, 0xd6, 0xa2, 0x03, 0xad, 0xa2, 0xd9, 0xc1, 0x45, 0x0f,
	0x9a, 0x29, 0x4e, 0xd4, 0x16, 0x94, 0xb8, 0x00, 0xf0, 0x36, 0x4d, 0xb2, 0xd3, 0x35, 0x06, 0xd0,
	0xbe, 0xc1, 0x39, 0x6e, 0x1d, 0xd3, 0x8d, 0x52, 0x14, 0xd0, 0x28, 0x1f, 0x55, 0x88, 0x5d, 0xff,
	0x2e, 0xf7, 0x61, 0x87, 0x90, 0x58, 0x9c, 0x1a, 0xcc, 0xf8, 0x3f, 0x72, 0xd6, 0x8a, 0xed, 0x43,
	0x6d, 0xf5, 0xaa, 0x97, 0xc8, 0x63, 0x3a, 0xf6, 0x82, 0x1d, 0x43, 0xc3, 0xe8, 0x39, 0x3e, 0x12,
	0xaa, 0xe6, 0x51, 0xee, 0xe0, 0xc1, 0xa1, 0x38, 0xd4, 0xa7, 0xb4, 0x77, 0xc5, 0x13, 0xb2, 0x0a,
	0xe9, 0x9c, 0x9c, 0x96, 0xa6, 0x78, 0xdd, 0x3b, 0x6b, 0x39, 0xfa, 0x8c, 0x20, 0x1e, 0x5b, 0x34,
	0xec, 0x1e, 0x12, 0x1f, 0x1a, 0x1b, 0x06, 0xd6, 0x52, 0x7a, 0x0d, 0xdd, 0xc1, 0x1f, 0x15, 0xb4,
	0x89, 0x8a, 0xc3, 0xf9, 0x98, 0x82, 0xb8, 0x52, 0xfc, 0x41, 0x5c, 0x29, 0x63, 0xc2, 0xf9, 0x55,
	0x07, 0x71, 0xa5, 0x98, 0x82, 0xb8, 0x52, 0x4e, 0x15, 0x76, 0x0b, 0xb1, 0x7b, 0x15, 0xac, 0x1f,
	0x28, 0xdd, 0x78, 0x4d, 0xdd, 0xde, 0x56, 0xdf, 0x83, 0x9e, 0x12, 0xfa, 0x85, 0xce, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0x34, 0x6f, 0xe9, 0x5b, 0x03, 0x00, 0x00,
}
