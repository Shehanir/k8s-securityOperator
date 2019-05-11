// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/user_list_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

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

// Enum containing possible user list types.
type UserListTypeEnum_UserListType int32

const (
	// Not specified.
	UserListTypeEnum_UNSPECIFIED UserListTypeEnum_UserListType = 0
	// Used for return value only. Represents value unknown in this version.
	UserListTypeEnum_UNKNOWN UserListTypeEnum_UserListType = 1
	// UserList represented as a collection of conversion types.
	UserListTypeEnum_REMARKETING UserListTypeEnum_UserListType = 2
	// UserList represented as a combination of other user lists/interests.
	UserListTypeEnum_LOGICAL UserListTypeEnum_UserListType = 3
	// UserList created in the Google Ad Manager platform.
	UserListTypeEnum_EXTERNAL_REMARKETING UserListTypeEnum_UserListType = 4
	// UserList associated with a rule.
	UserListTypeEnum_RULE_BASED UserListTypeEnum_UserListType = 5
	// UserList with users similar to users of another UserList.
	UserListTypeEnum_SIMILAR UserListTypeEnum_UserListType = 6
	// UserList of first-party CRM data provided by advertiser in the form of
	// emails or other formats.
	UserListTypeEnum_CRM_BASED UserListTypeEnum_UserListType = 7
)

var UserListTypeEnum_UserListType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REMARKETING",
	3: "LOGICAL",
	4: "EXTERNAL_REMARKETING",
	5: "RULE_BASED",
	6: "SIMILAR",
	7: "CRM_BASED",
}
var UserListTypeEnum_UserListType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"REMARKETING":          2,
	"LOGICAL":              3,
	"EXTERNAL_REMARKETING": 4,
	"RULE_BASED":           5,
	"SIMILAR":              6,
	"CRM_BASED":            7,
}

func (x UserListTypeEnum_UserListType) String() string {
	return proto.EnumName(UserListTypeEnum_UserListType_name, int32(x))
}
func (UserListTypeEnum_UserListType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_user_list_type_a2df5dfff3017dbe, []int{0, 0}
}

// The user list types.
type UserListTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListTypeEnum) Reset()         { *m = UserListTypeEnum{} }
func (m *UserListTypeEnum) String() string { return proto.CompactTextString(m) }
func (*UserListTypeEnum) ProtoMessage()    {}
func (*UserListTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_list_type_a2df5dfff3017dbe, []int{0}
}
func (m *UserListTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListTypeEnum.Unmarshal(m, b)
}
func (m *UserListTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListTypeEnum.Marshal(b, m, deterministic)
}
func (dst *UserListTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListTypeEnum.Merge(dst, src)
}
func (m *UserListTypeEnum) XXX_Size() int {
	return xxx_messageInfo_UserListTypeEnum.Size(m)
}
func (m *UserListTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserListTypeEnum)(nil), "google.ads.googleads.v1.enums.UserListTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.UserListTypeEnum_UserListType", UserListTypeEnum_UserListType_name, UserListTypeEnum_UserListType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/user_list_type.proto", fileDescriptor_user_list_type_a2df5dfff3017dbe)
}

var fileDescriptor_user_list_type_a2df5dfff3017dbe = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x6f, 0xd2, 0x7b, 0x5b, 0xee, 0xd4, 0x3f, 0x31, 0xb8, 0x10, 0xb1, 0x8b, 0xf6, 0x01,
	0x26, 0x44, 0x77, 0xe3, 0x6a, 0xd2, 0x8e, 0x21, 0x34, 0x4d, 0x4b, 0xda, 0x54, 0x91, 0x40, 0x88,
	0x66, 0x08, 0x81, 0x76, 0x26, 0x64, 0xd2, 0x42, 0xdf, 0xc2, 0x67, 0x10, 0xdc, 0xf8, 0x28, 0x3e,
	0x88, 0x0b, 0x9f, 0x42, 0x26, 0x69, 0x4b, 0x37, 0xba, 0x19, 0x0e, 0xf3, 0x3b, 0xe7, 0xe3, 0xfb,
	0x0e, 0xb8, 0x4e, 0x39, 0x4f, 0x17, 0xd4, 0x88, 0x13, 0x61, 0xd4, 0x52, 0xaa, 0xb5, 0x69, 0x50,
	0xb6, 0x5a, 0x0a, 0x63, 0x25, 0x68, 0x11, 0x2d, 0x32, 0x51, 0x46, 0xe5, 0x26, 0xa7, 0x30, 0x2f,
	0x78, 0xc9, 0xf5, 0x4e, 0x6d, 0x84, 0x71, 0x22, 0xe0, 0x3e, 0x03, 0xd7, 0x26, 0xac, 0x32, 0x97,
	0x57, 0xbb, 0x91, 0x79, 0x66, 0xc4, 0x8c, 0xf1, 0x32, 0x2e, 0x33, 0xce, 0x44, 0x1d, 0xee, 0xbd,
	0x29, 0x40, 0x0b, 0x04, 0x2d, 0xdc, 0x4c, 0x94, 0xb3, 0x4d, 0x4e, 0x09, 0x5b, 0x2d, 0x7b, 0x2f,
	0x0a, 0x38, 0x3a, 0xfc, 0xd4, 0x4f, 0x41, 0x3b, 0xf0, 0xa6, 0x13, 0xd2, 0x77, 0xee, 0x1c, 0x32,
	0xd0, 0xfe, 0xe8, 0x6d, 0xd0, 0x0a, 0xbc, 0xa1, 0x37, 0xbe, 0xf7, 0x34, 0x45, 0x52, 0x9f, 0x8c,
	0xb0, 0x3f, 0x24, 0x33, 0xc7, 0xb3, 0x35, 0x55, 0x52, 0x77, 0x6c, 0x3b, 0x7d, 0xec, 0x6a, 0x0d,
	0xfd, 0x02, 0x9c, 0x93, 0x87, 0x19, 0xf1, 0x3d, 0xec, 0x46, 0x87, 0xb6, 0xbf, 0xfa, 0x09, 0x00,
	0x7e, 0xe0, 0x92, 0xc8, 0xc2, 0x53, 0x32, 0xd0, 0xfe, 0xc9, 0xd8, 0xd4, 0x19, 0x39, 0x2e, 0xf6,
	0xb5, 0xa6, 0x7e, 0x0c, 0xfe, 0xf7, 0xfd, 0xd1, 0x96, 0xb5, 0xac, 0x4f, 0x05, 0x74, 0x9f, 0xf9,
	0x12, 0xfe, 0x7a, 0xab, 0x75, 0x76, 0xb8, 0xf5, 0x44, 0x1e, 0x38, 0x51, 0x1e, 0xad, 0x6d, 0x26,
	0xe5, 0x8b, 0x98, 0xa5, 0x90, 0x17, 0xa9, 0x91, 0x52, 0x56, 0x9d, 0xbf, 0xeb, 0x38, 0xcf, 0xc4,
	0x0f, 0x95, 0xdf, 0x56, 0xef, 0xab, 0xda, 0xb0, 0x31, 0x7e, 0x57, 0x3b, 0x76, 0x3d, 0x0a, 0x27,
	0x02, 0xd6, 0x52, 0xaa, 0xb9, 0x09, 0x65, 0x6d, 0xe2, 0x63, 0xc7, 0x43, 0x9c, 0x88, 0x70, 0xcf,
	0xc3, 0xb9, 0x19, 0x56, 0xfc, 0x4b, 0xed, 0xd6, 0x9f, 0x08, 0xe1, 0x44, 0x20, 0xb4, 0x77, 0x20,
	0x34, 0x37, 0x11, 0xaa, 0x3c, 0x4f, 0xcd, 0x6a, 0xb1, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x09, 0x1d, 0x7e, 0x46, 0x0a, 0x02, 0x00, 0x00,
}
