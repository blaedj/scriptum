// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/service.proto

/*
Package scriptumproto is a generated protocol buffer package.

It is generated from these files:
	rpc/service.proto

It has these top-level messages:
	Document
	NewDocumentResponse
	SaveDocumentResponse
	DeleteDocumentRequest
	DeleteDocumentResponse
*/
package scriptumproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Document struct {
	Title    string                     `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Contents string                     `protobuf:"bytes,2,opt,name=contents" json:"contents,omitempty"`
	Id       string                     `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	SavedAt  *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=saved_at,json=savedAt" json:"saved_at,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetContents() string {
	if m != nil {
		return m.Contents
	}
	return ""
}

func (m *Document) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Document) GetSavedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.SavedAt
	}
	return nil
}

type NewDocumentResponse struct {
	Err        string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	DocumentId string `protobuf:"bytes,2,opt,name=document_id,json=documentId" json:"document_id,omitempty"`
}

func (m *NewDocumentResponse) Reset()                    { *m = NewDocumentResponse{} }
func (m *NewDocumentResponse) String() string            { return proto.CompactTextString(m) }
func (*NewDocumentResponse) ProtoMessage()               {}
func (*NewDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewDocumentResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *NewDocumentResponse) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

type SaveDocumentResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *SaveDocumentResponse) Reset()                    { *m = SaveDocumentResponse{} }
func (m *SaveDocumentResponse) String() string            { return proto.CompactTextString(m) }
func (*SaveDocumentResponse) ProtoMessage()               {}
func (*SaveDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SaveDocumentResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DeleteDocumentRequest struct {
	DocumentId string `protobuf:"bytes,1,opt,name=document_id,json=documentId" json:"document_id,omitempty"`
}

func (m *DeleteDocumentRequest) Reset()                    { *m = DeleteDocumentRequest{} }
func (m *DeleteDocumentRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDocumentRequest) ProtoMessage()               {}
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteDocumentRequest) GetDocumentId() string {
	if m != nil {
		return m.DocumentId
	}
	return ""
}

type DeleteDocumentResponse struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *DeleteDocumentResponse) Reset()                    { *m = DeleteDocumentResponse{} }
func (m *DeleteDocumentResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDocumentResponse) ProtoMessage()               {}
func (*DeleteDocumentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteDocumentResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Document)(nil), "blaedj.scriptum.proto.Document")
	proto.RegisterType((*NewDocumentResponse)(nil), "blaedj.scriptum.proto.NewDocumentResponse")
	proto.RegisterType((*SaveDocumentResponse)(nil), "blaedj.scriptum.proto.SaveDocumentResponse")
	proto.RegisterType((*DeleteDocumentRequest)(nil), "blaedj.scriptum.proto.DeleteDocumentRequest")
	proto.RegisterType((*DeleteDocumentResponse)(nil), "blaedj.scriptum.proto.DeleteDocumentResponse")
}

func init() { proto.RegisterFile("rpc/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xd3, 0xe2, 0x1f, 0x1c, 0x14, 0x74, 0x05, 0xd3, 0xf4, 0x02, 0xe9, 0x89, 0xa0, 0x96,
	0x04, 0x63, 0xe2, 0x55, 0xc3, 0x41, 0x2f, 0x1e, 0xc0, 0x13, 0x31, 0x21, 0xa5, 0x3b, 0x92, 0x35,
	0x6d, 0xb7, 0x76, 0xa7, 0xf8, 0x00, 0xbe, 0x95, 0x4f, 0x67, 0xec, 0xb6, 0x06, 0xb1, 0x0d, 0x9e,
	0xda, 0x99, 0xfd, 0xed, 0x7c, 0xdf, 0xcc, 0x2c, 0x9c, 0x24, 0xb1, 0x3f, 0x54, 0x98, 0xac, 0x84,
	0x8f, 0x6e, 0x9c, 0x48, 0x92, 0xac, 0xb3, 0x08, 0x3c, 0xe4, 0xaf, 0xae, 0xf2, 0x13, 0x11, 0x53,
	0x1a, 0xea, 0xb4, 0xdd, 0x5d, 0x4a, 0xb9, 0x0c, 0x70, 0x98, 0x45, 0x8b, 0xf4, 0x65, 0x48, 0x22,
	0x44, 0x45, 0x5e, 0x18, 0x6b, 0xc0, 0xf9, 0x30, 0xa0, 0x3e, 0x96, 0x7e, 0x1a, 0x62, 0x44, 0xac,
	0x0d, 0xbb, 0x24, 0x28, 0x40, 0xcb, 0xe8, 0x19, 0xfd, 0x83, 0x89, 0x0e, 0x98, 0x0d, 0x75, 0x5f,
	0x46, 0x84, 0x11, 0x29, 0xcb, 0xcc, 0x0e, 0x7e, 0x62, 0xd6, 0x04, 0x53, 0x70, 0xab, 0x96, 0x65,
	0x4d, 0xc1, 0xd9, 0x35, 0xd4, 0x95, 0xb7, 0x42, 0x3e, 0xf7, 0xc8, 0xda, 0xe9, 0x19, 0xfd, 0xc6,
	0xc8, 0x76, 0xb5, 0x05, 0xb7, 0xb0, 0xe0, 0x3e, 0x15, 0x16, 0x26, 0xfb, 0x19, 0x7b, 0x4b, 0xce,
	0x3d, 0x9c, 0x3e, 0xe2, 0x7b, 0xe1, 0x63, 0x82, 0x2a, 0x96, 0x91, 0x42, 0x76, 0x0c, 0x35, 0x4c,
	0x92, 0xdc, 0xcd, 0xf7, 0x2f, 0xeb, 0x42, 0x83, 0xe7, 0xd4, 0x5c, 0xf0, 0xdc, 0x0e, 0x14, 0xa9,
	0x07, 0xee, 0xf4, 0xa1, 0x3d, 0xf5, 0x56, 0xb8, 0xbd, 0x94, 0x73, 0x03, 0x9d, 0x31, 0x06, 0x48,
	0x6b, 0xec, 0x5b, 0x8a, 0x8a, 0x36, 0x35, 0x8c, 0x3f, 0x1a, 0x03, 0x38, 0xdb, 0xbc, 0x59, 0xa5,
	0x32, 0xfa, 0x34, 0xa1, 0x35, 0xcd, 0x77, 0x32, 0xd5, 0x1b, 0x63, 0x33, 0x68, 0xac, 0x75, 0xcb,
	0xba, 0x6e, 0xe9, 0xee, 0xdc, 0x02, 0xb0, 0x07, 0x15, 0x40, 0xd9, 0xc8, 0x9e, 0xe1, 0x70, 0xbd,
	0xff, 0xed, 0xc5, 0xcf, 0x2b, 0x80, 0xd2, 0x29, 0x86, 0xd0, 0xfc, 0xdd, 0x39, 0xbb, 0xa8, 0xaa,
	0x5f, 0x36, 0x5a, 0xfb, 0xf2, 0x9f, 0xb4, 0x96, 0xbb, 0x6b, 0xcd, 0x8e, 0x0a, 0x50, 0xbf, 0x9e,
	0xbd, 0xec, 0x73, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x05, 0x93, 0x62, 0x8c, 0x01, 0x03, 0x00,
	0x00,
}
