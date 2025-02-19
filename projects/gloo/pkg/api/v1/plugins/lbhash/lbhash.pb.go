// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/lbhash/lbhash.proto

package lbhash

import (
	bytes "bytes"
	fmt "fmt"
	math "math"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Specifies the route’s hashing policy if the upstream cluster uses a hashing load balancer.
// https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/route/route.proto#envoy-api-msg-route-routeaction-hashpolicy
type RouteActionHashConfig struct {
	// The list of policies Envoy will use when generating a hash key for a hashing load balancer
	HashPolicies         []*HashPolicy `protobuf:"bytes,1,rep,name=hash_policies,json=hashPolicies,proto3" json:"hash_policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteActionHashConfig) Reset()         { *m = RouteActionHashConfig{} }
func (m *RouteActionHashConfig) String() string { return proto.CompactTextString(m) }
func (*RouteActionHashConfig) ProtoMessage()    {}
func (*RouteActionHashConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_abebf15c10732758, []int{0}
}
func (m *RouteActionHashConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteActionHashConfig.Unmarshal(m, b)
}
func (m *RouteActionHashConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteActionHashConfig.Marshal(b, m, deterministic)
}
func (m *RouteActionHashConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteActionHashConfig.Merge(m, src)
}
func (m *RouteActionHashConfig) XXX_Size() int {
	return xxx_messageInfo_RouteActionHashConfig.Size(m)
}
func (m *RouteActionHashConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteActionHashConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RouteActionHashConfig proto.InternalMessageInfo

func (m *RouteActionHashConfig) GetHashPolicies() []*HashPolicy {
	if m != nil {
		return m.HashPolicies
	}
	return nil
}

// Envoy supports two types of cookie affinity:
// - Passive: Envoy reads the cookie from the headers
// - Generated: Envoy uses the cookie spec to generate a cookie
// In either case, the cookie is incorporated in the hash key.
// additional notes https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/route/route.proto#envoy-api-msg-route-routeaction-hashpolicy-cookie
type Cookie struct {
	// required, the name of the cookie to be used to obtain the hash key
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie.
	Ttl *time.Duration `protobuf:"bytes,2,opt,name=ttl,proto3,stdduration" json:"ttl,omitempty"`
	// The name of the path for the cookie. If no path is specified here, no path will be set for the cookie.
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cookie) Reset()         { *m = Cookie{} }
func (m *Cookie) String() string { return proto.CompactTextString(m) }
func (*Cookie) ProtoMessage()    {}
func (*Cookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_abebf15c10732758, []int{1}
}
func (m *Cookie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cookie.Unmarshal(m, b)
}
func (m *Cookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cookie.Marshal(b, m, deterministic)
}
func (m *Cookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cookie.Merge(m, src)
}
func (m *Cookie) XXX_Size() int {
	return xxx_messageInfo_Cookie.Size(m)
}
func (m *Cookie) XXX_DiscardUnknown() {
	xxx_messageInfo_Cookie.DiscardUnknown(m)
}

var xxx_messageInfo_Cookie proto.InternalMessageInfo

func (m *Cookie) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cookie) GetTtl() *time.Duration {
	if m != nil {
		return m.Ttl
	}
	return nil
}

func (m *Cookie) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Specifies an element of Envoy's hashing policy for hashing load balancers
type HashPolicy struct {
	// Types that are valid to be assigned to KeyType:
	//	*HashPolicy_Header
	//	*HashPolicy_Cookie
	//	*HashPolicy_SourceIp
	KeyType isHashPolicy_KeyType `protobuf_oneof:"KeyType"`
	// If set, and a hash key is available after evaluating this policy, Envoy will skip the subsequent policies and
	// use the key as it is.
	// This is useful for defining "fallback" policies and limiting the time Envoy spends generating hash keys.
	Terminal             bool     `protobuf:"varint,4,opt,name=terminal,proto3" json:"terminal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashPolicy) Reset()         { *m = HashPolicy{} }
func (m *HashPolicy) String() string { return proto.CompactTextString(m) }
func (*HashPolicy) ProtoMessage()    {}
func (*HashPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_abebf15c10732758, []int{2}
}
func (m *HashPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashPolicy.Unmarshal(m, b)
}
func (m *HashPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashPolicy.Marshal(b, m, deterministic)
}
func (m *HashPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashPolicy.Merge(m, src)
}
func (m *HashPolicy) XXX_Size() int {
	return xxx_messageInfo_HashPolicy.Size(m)
}
func (m *HashPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_HashPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_HashPolicy proto.InternalMessageInfo

type isHashPolicy_KeyType interface {
	isHashPolicy_KeyType()
	Equal(interface{}) bool
}

type HashPolicy_Header struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}
type HashPolicy_Cookie struct {
	Cookie *Cookie `protobuf:"bytes,2,opt,name=cookie,proto3,oneof"`
}
type HashPolicy_SourceIp struct {
	SourceIp bool `protobuf:"varint,3,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

func (*HashPolicy_Header) isHashPolicy_KeyType()   {}
func (*HashPolicy_Cookie) isHashPolicy_KeyType()   {}
func (*HashPolicy_SourceIp) isHashPolicy_KeyType() {}

func (m *HashPolicy) GetKeyType() isHashPolicy_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *HashPolicy) GetHeader() string {
	if x, ok := m.GetKeyType().(*HashPolicy_Header); ok {
		return x.Header
	}
	return ""
}

func (m *HashPolicy) GetCookie() *Cookie {
	if x, ok := m.GetKeyType().(*HashPolicy_Cookie); ok {
		return x.Cookie
	}
	return nil
}

func (m *HashPolicy) GetSourceIp() bool {
	if x, ok := m.GetKeyType().(*HashPolicy_SourceIp); ok {
		return x.SourceIp
	}
	return false
}

func (m *HashPolicy) GetTerminal() bool {
	if m != nil {
		return m.Terminal
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HashPolicy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HashPolicy_Header)(nil),
		(*HashPolicy_Cookie)(nil),
		(*HashPolicy_SourceIp)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteActionHashConfig)(nil), "lbhash.plugins.gloo.solo.io.RouteActionHashConfig")
	proto.RegisterType((*Cookie)(nil), "lbhash.plugins.gloo.solo.io.Cookie")
	proto.RegisterType((*HashPolicy)(nil), "lbhash.plugins.gloo.solo.io.HashPolicy")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/lbhash/lbhash.proto", fileDescriptor_abebf15c10732758)
}

var fileDescriptor_abebf15c10732758 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x1b, 0x5a, 0x85, 0xd6, 0x85, 0x8b, 0x05, 0x52, 0x28, 0xa2, 0x54, 0xe5, 0x40, 0x2f,
	0xd8, 0x6a, 0x39, 0x23, 0x44, 0x0b, 0x52, 0x10, 0x1c, 0x50, 0xc4, 0x89, 0x4b, 0xe5, 0xa4, 0xae,
	0x63, 0xea, 0xe6, 0x59, 0xb6, 0xb3, 0xa9, 0xdf, 0x64, 0x1f, 0x61, 0xd2, 0x3e, 0xd4, 0xa4, 0x7d,
	0x92, 0xc9, 0x71, 0xba, 0x5d, 0xb6, 0x6a, 0xa7, 0xfc, 0xff, 0x7e, 0xef, 0xff, 0xfc, 0x8b, 0xf5,
	0x50, 0x2a, 0xa4, 0x2b, 0xeb, 0x9c, 0x14, 0xb0, 0xa7, 0x16, 0x14, 0x7c, 0x92, 0x40, 0x85, 0x02,
	0xa0, 0xda, 0xc0, 0x7f, 0x5e, 0x38, 0x1b, 0x1c, 0xd3, 0x92, 0x9e, 0xcd, 0xa9, 0x56, 0xb5, 0x90,
	0x95, 0xa5, 0x2a, 0x2f, 0x99, 0x2d, 0xdb, 0x0f, 0xd1, 0x06, 0x1c, 0xe0, 0xb7, 0x47, 0x17, 0x7a,
	0x88, 0xcf, 0x11, 0x3f, 0x92, 0x48, 0x18, 0xbd, 0x12, 0x20, 0xa0, 0xe9, 0xa3, 0x5e, 0x85, 0xc8,
	0x68, 0x2c, 0x00, 0x84, 0xe2, 0xb4, 0x71, 0x79, 0xbd, 0xa5, 0xe7, 0x86, 0x69, 0xcd, 0x8d, 0x7d,
	0xac, 0xbe, 0xa9, 0x0d, 0x73, 0x12, 0xaa, 0x50, 0x9f, 0x72, 0xf4, 0x3a, 0x83, 0xda, 0xf1, 0x6f,
	0x85, 0x3f, 0x4c, 0x99, 0x2d, 0x57, 0x50, 0x6d, 0xa5, 0xc0, 0xbf, 0xd1, 0x4b, 0xcf, 0xb2, 0xd6,
	0xa0, 0x64, 0x21, 0xb9, 0x4d, 0xa2, 0x49, 0x77, 0x36, 0x5c, 0x7c, 0x24, 0x27, 0x18, 0x89, 0xcf,
	0xff, 0xf1, 0x81, 0x43, 0xf6, 0xa2, 0x3c, 0x6a, 0xc9, 0xed, 0xb4, 0x40, 0xf1, 0x0a, 0x60, 0x27,
	0x39, 0xc6, 0xa8, 0x57, 0xb1, 0x3d, 0x4f, 0xa2, 0x49, 0x34, 0x1b, 0x64, 0x8d, 0xc6, 0x73, 0xd4,
	0x75, 0x4e, 0x25, 0xcf, 0x26, 0xd1, 0x6c, 0xb8, 0x78, 0x43, 0x02, 0x32, 0x39, 0x22, 0x93, 0xef,
	0x2d, 0xf2, 0xb2, 0x77, 0x71, 0xfd, 0x3e, 0xca, 0x7c, 0xaf, 0x1f, 0xa3, 0x99, 0x2b, 0x93, 0x6e,
	0x18, 0xe3, 0xf5, 0xf4, 0x2a, 0x42, 0xe8, 0x9e, 0x00, 0x27, 0x28, 0x2e, 0x39, 0xdb, 0x70, 0x13,
	0xee, 0x4a, 0x3b, 0x59, 0xeb, 0xf1, 0x17, 0x14, 0x17, 0x0d, 0x4d, 0x7b, 0xe5, 0x87, 0x93, 0x3f,
	0x15, 0xc0, 0x7d, 0x3c, 0x84, 0xf0, 0x3b, 0x34, 0xb0, 0x50, 0x9b, 0x82, 0xaf, 0xa5, 0x6e, 0x00,
	0xfa, 0x69, 0x27, 0xeb, 0x87, 0xa3, 0x9f, 0x1a, 0x8f, 0x50, 0xdf, 0x71, 0xb3, 0x97, 0x15, 0x53,
	0x49, 0xcf, 0x57, 0xb3, 0x3b, 0xbf, 0x1c, 0xa0, 0xe7, 0xbf, 0xf8, 0xe1, 0xef, 0x41, 0xf3, 0xe5,
	0x8f, 0xcb, 0x9b, 0x71, 0xf4, 0xef, 0xeb, 0xd3, 0x96, 0x47, 0xef, 0xc4, 0xc3, 0x0b, 0x94, 0xc7,
	0xcd, 0x33, 0x7d, 0xbe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x4b, 0xc2, 0xf6, 0x86, 0x02, 0x00,
	0x00,
}

func (this *RouteActionHashConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteActionHashConfig)
	if !ok {
		that2, ok := that.(RouteActionHashConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.HashPolicies) != len(that1.HashPolicies) {
		return false
	}
	for i := range this.HashPolicies {
		if !this.HashPolicies[i].Equal(that1.HashPolicies[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Cookie) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Cookie)
	if !ok {
		that2, ok := that.(Cookie)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Ttl != nil && that1.Ttl != nil {
		if *this.Ttl != *that1.Ttl {
			return false
		}
	} else if this.Ttl != nil {
		return false
	} else if that1.Ttl != nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HashPolicy) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HashPolicy)
	if !ok {
		that2, ok := that.(HashPolicy)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.KeyType == nil {
		if this.KeyType != nil {
			return false
		}
	} else if this.KeyType == nil {
		return false
	} else if !this.KeyType.Equal(that1.KeyType) {
		return false
	}
	if this.Terminal != that1.Terminal {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HashPolicy_Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HashPolicy_Header)
	if !ok {
		that2, ok := that.(HashPolicy_Header)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Header != that1.Header {
		return false
	}
	return true
}
func (this *HashPolicy_Cookie) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HashPolicy_Cookie)
	if !ok {
		that2, ok := that.(HashPolicy_Cookie)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Cookie.Equal(that1.Cookie) {
		return false
	}
	return true
}
func (this *HashPolicy_SourceIp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HashPolicy_SourceIp)
	if !ok {
		that2, ok := that.(HashPolicy_SourceIp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SourceIp != that1.SourceIp {
		return false
	}
	return true
}
