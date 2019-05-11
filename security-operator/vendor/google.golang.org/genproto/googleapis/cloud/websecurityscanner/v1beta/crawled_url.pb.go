// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/crawled_url.proto

package websecurityscanner // import "google.golang.org/genproto/googleapis/cloud/websecurityscanner/v1beta"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A CrawledUrl resource represents a URL that was crawled during a ScanRun. Web
// Security Scanner Service crawls the web applications, following all links
// within the scope of sites, to find the URLs to test against.
type CrawledUrl struct {
	// Output only.
	// The http method of the request that was used to visit the URL, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,1,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only.
	// The URL that was crawled.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Output only.
	// The body of the request that was used to visit the URL.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrawledUrl) Reset()         { *m = CrawledUrl{} }
func (m *CrawledUrl) String() string { return proto.CompactTextString(m) }
func (*CrawledUrl) ProtoMessage()    {}
func (*CrawledUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_crawled_url_47969d3dbeaecdf6, []int{0}
}
func (m *CrawledUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrawledUrl.Unmarshal(m, b)
}
func (m *CrawledUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrawledUrl.Marshal(b, m, deterministic)
}
func (dst *CrawledUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrawledUrl.Merge(dst, src)
}
func (m *CrawledUrl) XXX_Size() int {
	return xxx_messageInfo_CrawledUrl.Size(m)
}
func (m *CrawledUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_CrawledUrl.DiscardUnknown(m)
}

var xxx_messageInfo_CrawledUrl proto.InternalMessageInfo

func (m *CrawledUrl) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *CrawledUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CrawledUrl) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CrawledUrl)(nil), "google.cloud.websecurityscanner.v1beta.CrawledUrl")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/crawled_url.proto", fileDescriptor_crawled_url_47969d3dbeaecdf6)
}

var fileDescriptor_crawled_url_47969d3dbeaecdf6 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x4f, 0x04, 0xc7, 0x42, 0x49, 0xb5, 0x88, 0xa0, 0x58, 0x88, 0x58, 0x24, 0x88,
	0x8d, 0x60, 0x77, 0xd6, 0xc2, 0xa1, 0x08, 0x62, 0x73, 0xcc, 0x66, 0x87, 0xdc, 0x42, 0x2e, 0xb3,
	0xcc, 0xce, 0x7a, 0xdc, 0x4f, 0xf1, 0xdf, 0xca, 0x26, 0x07, 0x16, 0x57, 0x68, 0x37, 0xbc, 0xc7,
	0xf7, 0xf2, 0x11, 0x78, 0x0c, 0xcc, 0x21, 0x92, 0xf3, 0x91, 0xc7, 0xd6, 0x6d, 0xa8, 0x19, 0xc8,
	0x8f, 0xd2, 0xe9, 0x76, 0xf0, 0x98, 0x12, 0x89, 0xfb, 0xba, 0x6f, 0x48, 0xd1, 0x79, 0xc1, 0x4d,
	0xa4, 0x76, 0x39, 0x4a, 0xb4, 0xbd, 0xb0, 0xb2, 0xb9, 0x29, 0xa4, 0xcd, 0xa4, 0xdd, 0x27, 0x6d,
	0x21, 0xcf, 0x2f, 0x76, 0x2f, 0x60, 0xdf, 0x39, 0x4c, 0x89, 0x15, 0xb5, 0xe3, 0x34, 0x94, 0x95,
	0xeb, 0x37, 0x80, 0xe7, 0x32, 0xfd, 0x2e, 0xd1, 0x5c, 0xc2, 0xc9, 0x4a, 0xb5, 0x5f, 0xae, 0x49,
	0x57, 0xdc, 0xd6, 0xd5, 0x55, 0x75, 0x7b, 0xfc, 0x0a, 0x53, 0xf4, 0x92, 0x13, 0x73, 0x06, 0xb3,
	0x51, 0x62, 0x7d, 0x90, 0x8b, 0xe9, 0x34, 0x06, 0x0e, 0x1b, 0x6e, 0xb7, 0xf5, 0x2c, 0x47, 0xf9,
	0x9e, 0x7f, 0x57, 0x70, 0xe7, 0x79, 0x6d, 0xff, 0x67, 0x38, 0x3f, 0xfd, 0x35, 0x58, 0x4c, 0x52,
	0x8b, 0xea, 0xf3, 0x63, 0x87, 0x06, 0x8e, 0x98, 0x82, 0x65, 0x09, 0x2e, 0x50, 0xca, 0xca, 0xae,
	0x54, 0xd8, 0x77, 0xc3, 0x5f, 0xbf, 0xf6, 0xb4, 0xdf, 0x34, 0x47, 0x79, 0xe4, 0xe1, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xac, 0xe1, 0xa8, 0x52, 0x79, 0x01, 0x00, 0x00,
}
