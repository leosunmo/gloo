// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/jwt/solo_jwt_authn.proto

package jwt

import (
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SoloJwtAuthnPerRoute struct {
	Requirement     string                                          `protobuf:"bytes,1,opt,name=requirement,proto3" json:"requirement,omitempty"`
	ClaimsToHeaders map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders `protobuf:"bytes,2,rep,name=claims_to_headers,json=claimsToHeaders,proto3" json:"claims_to_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// clear the route cache if claims were added to the header
	ClearRouteCache      bool     `protobuf:"varint,3,opt,name=clear_route_cache,json=clearRouteCache,proto3" json:"clear_route_cache,omitempty"`
	PayloadInMetadata    string   `protobuf:"bytes,4,opt,name=payload_in_metadata,json=payloadInMetadata,proto3" json:"payload_in_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoloJwtAuthnPerRoute) Reset()         { *m = SoloJwtAuthnPerRoute{} }
func (m *SoloJwtAuthnPerRoute) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8dcaaf88f211b6a, []int{0}
}
func (m *SoloJwtAuthnPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute.Size(m)
}
func (m *SoloJwtAuthnPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute) GetRequirement() string {
	if m != nil {
		return m.Requirement
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute) GetClaimsToHeaders() map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders {
	if m != nil {
		return m.ClaimsToHeaders
	}
	return nil
}

func (m *SoloJwtAuthnPerRoute) GetClearRouteCache() bool {
	if m != nil {
		return m.ClearRouteCache
	}
	return false
}

func (m *SoloJwtAuthnPerRoute) GetPayloadInMetadata() string {
	if m != nil {
		return m.PayloadInMetadata
	}
	return ""
}

// If this is specified, one of the claims will be copied to a header
// and the route cache will be cleared.
type SoloJwtAuthnPerRoute_ClaimToHeader struct {
	Claim                string   `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	Header               string   `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Append               bool     `protobuf:"varint,3,opt,name=append,proto3" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) Reset()         { *m = SoloJwtAuthnPerRoute_ClaimToHeader{} }
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute_ClaimToHeader) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute_ClaimToHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8dcaaf88f211b6a, []int{0, 0}
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.Size(m)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeader proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetClaim() string {
	if m != nil {
		return m.Claim
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeader) GetAppend() bool {
	if m != nil {
		return m.Append
	}
	return false
}

type SoloJwtAuthnPerRoute_ClaimToHeaders struct {
	Claims               []*SoloJwtAuthnPerRoute_ClaimToHeader `protobuf:"bytes,1,rep,name=claims,proto3" json:"claims,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) Reset()         { *m = SoloJwtAuthnPerRoute_ClaimToHeaders{} }
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) String() string { return proto.CompactTextString(m) }
func (*SoloJwtAuthnPerRoute_ClaimToHeaders) ProtoMessage()    {}
func (*SoloJwtAuthnPerRoute_ClaimToHeaders) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8dcaaf88f211b6a, []int{0, 1}
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Unmarshal(m, b)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Marshal(b, m, deterministic)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Merge(m, src)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_Size() int {
	return xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.Size(m)
}
func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) XXX_DiscardUnknown() {
	xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders.DiscardUnknown(m)
}

var xxx_messageInfo_SoloJwtAuthnPerRoute_ClaimToHeaders proto.InternalMessageInfo

func (m *SoloJwtAuthnPerRoute_ClaimToHeaders) GetClaims() []*SoloJwtAuthnPerRoute_ClaimToHeader {
	if m != nil {
		return m.Claims
	}
	return nil
}

func init() {
	proto.RegisterType((*SoloJwtAuthnPerRoute)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute")
	proto.RegisterMapType((map[string]*SoloJwtAuthnPerRoute_ClaimToHeaders)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimsToHeadersEntry")
	proto.RegisterType((*SoloJwtAuthnPerRoute_ClaimToHeader)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeader")
	proto.RegisterType((*SoloJwtAuthnPerRoute_ClaimToHeaders)(nil), "envoy.config.filter.http.solo_jwt_authn.v2.SoloJwtAuthnPerRoute.ClaimToHeaders")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/external/envoy/jwt/solo_jwt_authn.proto", fileDescriptor_f8dcaaf88f211b6a)
}

var fileDescriptor_f8dcaaf88f211b6a = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x27, 0xdb, 0xdd, 0xa2, 0x53, 0xfc, 0xd3, 0xb1, 0x68, 0xe8, 0xa9, 0x78, 0x2a, 0x0b, 0x4e,
	0xa0, 0x5e, 0x16, 0x4f, 0xba, 0x8b, 0xa0, 0x82, 0xba, 0x46, 0xf7, 0xe2, 0x25, 0xcc, 0x26, 0xaf,
	0xcd, 0x74, 0xa7, 0xf3, 0xc6, 0xc9, 0x4b, 0xdb, 0x1c, 0xfd, 0x08, 0xe2, 0x07, 0x56, 0x32, 0x93,
	0x82, 0x95, 0x15, 0x56, 0xd8, 0x5b, 0xde, 0xef, 0xbd, 0xf9, 0xfd, 0x79, 0xe4, 0xb1, 0x4f, 0x0b,
	0x45, 0x65, 0x7d, 0x29, 0x72, 0x5c, 0x25, 0x15, 0x6a, 0x7c, 0xa6, 0x30, 0x59, 0x68, 0xc4, 0xc4,
	0x3a, 0x5c, 0x42, 0x4e, 0x55, 0xa8, 0xa4, 0x55, 0x09, 0x6c, 0x09, 0x9c, 0x91, 0x3a, 0x01, 0xb3,
	0xc6, 0x26, 0x59, 0x6e, 0xc8, 0xbf, 0xc8, 0x96, 0x1b, 0xca, 0x64, 0x4d, 0xa5, 0x11, 0xd6, 0x21,
	0x21, 0x3f, 0xf6, 0x7d, 0x91, 0xa3, 0x99, 0xab, 0x85, 0x98, 0x2b, 0x4d, 0xe0, 0x44, 0x49, 0x64,
	0xc5, 0x5f, 0xe3, 0xeb, 0xd9, 0xf8, 0xc9, 0x5a, 0x6a, 0x55, 0x48, 0x82, 0x64, 0xf7, 0x11, 0x48,
	0x9e, 0xfe, 0x3a, 0x64, 0xa3, 0xcf, 0xa8, 0xf1, 0xdd, 0x86, 0x5e, 0xb5, 0xc3, 0xe7, 0xe0, 0x52,
	0xac, 0x09, 0xf8, 0x84, 0x0d, 0x1c, 0x7c, 0xab, 0x95, 0x83, 0x15, 0x18, 0x8a, 0xa3, 0x49, 0x34,
	0xbd, 0x9b, 0xfe, 0x09, 0xf1, 0xef, 0x11, 0x1b, 0xe6, 0x5a, 0xaa, 0x55, 0x95, 0x11, 0x66, 0x25,
	0xc8, 0x02, 0x5c, 0x15, 0x1f, 0x4c, 0x7a, 0xd3, 0xc1, 0xec, 0x42, 0xdc, 0xdc, 0x9c, 0xb8, 0x4e,
	0x5f, 0x9c, 0x79, 0xe6, 0x2f, 0xf8, 0x26, 0xf0, 0xbe, 0x36, 0xe4, 0x9a, 0xf4, 0x41, 0xbe, 0x8f,
	0xf2, 0xe3, 0xd6, 0x02, 0x48, 0x97, 0xb9, 0xf6, 0x51, 0x96, 0xcb, 0xbc, 0x84, 0xb8, 0x37, 0x89,
	0xa6, 0x77, 0xda, 0x59, 0x90, 0x81, 0xec, 0xac, 0x85, 0xb9, 0x60, 0x8f, 0xac, 0x6c, 0x34, 0xca,
	0x22, 0x53, 0x26, 0x5b, 0x01, 0xc9, 0x42, 0x92, 0x8c, 0x0f, 0x7d, 0xb2, 0x61, 0xd7, 0x7a, 0x6b,
	0xde, 0x77, 0x8d, 0xf1, 0x05, 0xbb, 0xe7, 0x4d, 0xec, 0xd4, 0xf8, 0x88, 0x1d, 0x79, 0xfd, 0x6e,
	0x19, 0xa1, 0xe0, 0x8f, 0x59, 0x3f, 0x64, 0x8f, 0x0f, 0x3c, 0xdc, 0x55, 0x2d, 0x2e, 0xad, 0x05,
	0x53, 0x74, 0x7e, 0xba, 0x6a, 0xbc, 0x65, 0xf7, 0xf7, 0x68, 0x2b, 0x3e, 0x67, 0xfd, 0x90, 0x2b,
	0x8e, 0xfc, 0xf2, 0x3e, 0xdc, 0xce, 0xf2, 0x76, 0x02, 0x69, 0xc7, 0x3e, 0xfe, 0x19, 0xb1, 0xd1,
	0x75, 0x6b, 0xe5, 0x0f, 0x59, 0xef, 0x0a, 0x9a, 0x2e, 0x56, 0xfb, 0xc9, 0x81, 0x1d, 0xad, 0xa5,
	0xae, 0xc1, 0x67, 0x1a, 0xcc, 0x3e, 0xde, 0xae, 0xa3, 0x2a, 0x0d, 0xec, 0x2f, 0x0e, 0x4e, 0xa2,
	0xd3, 0x1f, 0x11, 0x3b, 0x51, 0x18, 0x04, 0xac, 0xc3, 0x6d, 0xf3, 0x1f, 0x5a, 0xa7, 0xc3, 0x3d,
	0xb1, 0xf6, 0x8f, 0x3e, 0x8f, 0xbe, 0xbe, 0xbc, 0xd9, 0xad, 0xd9, 0xab, 0xc5, 0x3f, 0xee, 0xed,
	0xb2, 0xef, 0x8f, 0xe3, 0xf9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xbb, 0xaf, 0xdc, 0xb6,
	0x03, 0x00, 0x00,
}
