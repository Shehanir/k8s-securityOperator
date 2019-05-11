// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/application.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import date "google.golang.org/genproto/googleapis/type/date"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum that represents the application status.
type Application_ApplicationState int32

const (
	// Default value.
	Application_APPLICATION_STATE_UNSPECIFIED Application_ApplicationState = 0
	// The current stage is in progress or pending, for example, interviews in
	// progress.
	Application_IN_PROGRESS Application_ApplicationState = 1
	// The current stage was terminated by a candidate decision.
	Application_CANDIDATE_WITHDREW Application_ApplicationState = 2
	// The current stage was terminated by an employer or agency decision.
	Application_EMPLOYER_WITHDREW Application_ApplicationState = 3
	// The current stage is successfully completed, but the next stage (if
	// applicable) has not begun.
	Application_COMPLETED Application_ApplicationState = 4
	// The current stage was closed without an exception, or terminated for
	// reasons unrealated to the candidate.
	Application_CLOSED Application_ApplicationState = 5
)

var Application_ApplicationState_name = map[int32]string{
	0: "APPLICATION_STATE_UNSPECIFIED",
	1: "IN_PROGRESS",
	2: "CANDIDATE_WITHDREW",
	3: "EMPLOYER_WITHDREW",
	4: "COMPLETED",
	5: "CLOSED",
}
var Application_ApplicationState_value = map[string]int32{
	"APPLICATION_STATE_UNSPECIFIED": 0,
	"IN_PROGRESS":                   1,
	"CANDIDATE_WITHDREW":            2,
	"EMPLOYER_WITHDREW":             3,
	"COMPLETED":                     4,
	"CLOSED":                        5,
}

func (x Application_ApplicationState) String() string {
	return proto.EnumName(Application_ApplicationState_name, int32(x))
}
func (Application_ApplicationState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_application_dbba57710742dc0d, []int{0, 0}
}

// The stage of the application.
type Application_ApplicationStage int32

const (
	// Default value.
	Application_APPLICATION_STAGE_UNSPECIFIED Application_ApplicationStage = 0
	// Candidate has applied or a recruiter put candidate into consideration but
	// candidate is not yet screened / no decision has been made to move or not
	// move the candidate to the next stage.
	Application_NEW Application_ApplicationStage = 1
	// A recruiter decided to screen the candidate for this role.
	Application_SCREEN Application_ApplicationStage = 2
	// Candidate is being / was sent to the customer / hiring manager for
	// detailed review.
	Application_HIRING_MANAGER_REVIEW Application_ApplicationStage = 3
	// Candidate was approved by the client / hiring manager and is being / was
	// interviewed for the role.
	Application_INTERVIEW Application_ApplicationStage = 4
	// Candidate will be / has been given an offer of employment.
	Application_OFFER_EXTENDED Application_ApplicationStage = 5
	// Candidate has accepted their offer of employment.
	Application_OFFER_ACCEPTED Application_ApplicationStage = 6
	// Candidate has begun (or completed) their employment or assignment with
	// the employer.
	Application_STARTED Application_ApplicationStage = 7
)

var Application_ApplicationStage_name = map[int32]string{
	0: "APPLICATION_STAGE_UNSPECIFIED",
	1: "NEW",
	2: "SCREEN",
	3: "HIRING_MANAGER_REVIEW",
	4: "INTERVIEW",
	5: "OFFER_EXTENDED",
	6: "OFFER_ACCEPTED",
	7: "STARTED",
}
var Application_ApplicationStage_value = map[string]int32{
	"APPLICATION_STAGE_UNSPECIFIED": 0,
	"NEW":                           1,
	"SCREEN":                        2,
	"HIRING_MANAGER_REVIEW":         3,
	"INTERVIEW":                     4,
	"OFFER_EXTENDED":                5,
	"OFFER_ACCEPTED":                6,
	"STARTED":                       7,
}

func (x Application_ApplicationStage) String() string {
	return proto.EnumName(Application_ApplicationStage_name, int32(x))
}
func (Application_ApplicationStage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_application_dbba57710742dc0d, []int{0, 1}
}

// Resource that represents a job application record of a candidate.
type Application struct {
	// Required during application update.
	//
	// Resource name assigned to an application by the API.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}/applications/{application_id}",
	// for example,
	// "projects/api-test-project/tenants/foo/profiles/bar/applications/baz".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// Client side application identifier, used to uniquely identify the
	// application.
	//
	// The maximum number of allowed characters is 255.
	ExternalId string `protobuf:"bytes,31,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// Output only. Resource name of the candidate of this application.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/profiles/{profile_id}",
	// for example, "projects/api-test-project/tenants/foo/profiles/bar".
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	// One of either a job or a company is required.
	//
	// Resource name of the job which the candidate applied for.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}",
	// for example, "projects/api-test-project/tenants/foo/jobs/bar".
	Job string `protobuf:"bytes,4,opt,name=job,proto3" json:"job,omitempty"`
	// One of either a job or a company is required.
	//
	// Resource name of the company which the candidate applied for.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}",
	// for example, "projects/api-test-project/tenants/foo/companies/bar".
	Company string `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	// Optional.
	//
	// The application date.
	ApplicationDate *date.Date `protobuf:"bytes,7,opt,name=application_date,json=applicationDate,proto3" json:"application_date,omitempty"`
	// Required.
	//
	// What is the most recent stage of the application (that is, new, screen,
	// send cv, hired, finished work)?  This field is intentionally not
	// comprehensive of every possible status, but instead, represents statuses
	// that would be used to indicate to the ML models good / bad matches.
	Stage Application_ApplicationStage `protobuf:"varint,11,opt,name=stage,proto3,enum=google.cloud.talent.v4beta1.Application_ApplicationStage" json:"stage,omitempty"`
	// Optional.
	//
	// The application state.
	State Application_ApplicationState `protobuf:"varint,13,opt,name=state,proto3,enum=google.cloud.talent.v4beta1.Application_ApplicationState" json:"state,omitempty"`
	// Optional.
	//
	// All interviews (screen, onsite, and so on) conducted as part of this
	// application (includes details such as user conducting the interview,
	// timestamp, feedback, and so on).
	Interviews []*Interview `protobuf:"bytes,16,rep,name=interviews,proto3" json:"interviews,omitempty"`
	// Optional.
	//
	// If the candidate is referred by a employee.
	Referral *wrappers.BoolValue `protobuf:"bytes,18,opt,name=referral,proto3" json:"referral,omitempty"`
	// Required.
	//
	// Reflects the time that the application was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,19,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Optional.
	//
	// The last update timestamp.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,20,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Optional.
	//
	// Free text reason behind the recruitement outcome (for example, reason for
	// withdraw / reject, reason for an unsuccessful finish, and so on).
	//
	// Number of characters allowed is 100.
	OutcomeNotes string `protobuf:"bytes,21,opt,name=outcome_notes,json=outcomeNotes,proto3" json:"outcome_notes,omitempty"`
	// Optional.
	//
	// Outcome positiveness shows how positive the outcome is.
	Outcome Outcome `protobuf:"varint,22,opt,name=outcome,proto3,enum=google.cloud.talent.v4beta1.Outcome" json:"outcome,omitempty"`
	// Output only. Indicates whether this job application is a match to
	// application related filters. This value is only applicable in profile
	// search response.
	IsMatch *wrappers.BoolValue `protobuf:"bytes,28,opt,name=is_match,json=isMatch,proto3" json:"is_match,omitempty"`
	// Output only. Job title snippet shows how the job title is related to a
	// search query. It's empty if the job title isn't related to the search
	// query.
	JobTitleSnippet      string   `protobuf:"bytes,29,opt,name=job_title_snippet,json=jobTitleSnippet,proto3" json:"job_title_snippet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_dbba57710742dc0d, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Application) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Application) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

func (m *Application) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *Application) GetApplicationDate() *date.Date {
	if m != nil {
		return m.ApplicationDate
	}
	return nil
}

func (m *Application) GetStage() Application_ApplicationStage {
	if m != nil {
		return m.Stage
	}
	return Application_APPLICATION_STAGE_UNSPECIFIED
}

func (m *Application) GetState() Application_ApplicationState {
	if m != nil {
		return m.State
	}
	return Application_APPLICATION_STATE_UNSPECIFIED
}

func (m *Application) GetInterviews() []*Interview {
	if m != nil {
		return m.Interviews
	}
	return nil
}

func (m *Application) GetReferral() *wrappers.BoolValue {
	if m != nil {
		return m.Referral
	}
	return nil
}

func (m *Application) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Application) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Application) GetOutcomeNotes() string {
	if m != nil {
		return m.OutcomeNotes
	}
	return ""
}

func (m *Application) GetOutcome() Outcome {
	if m != nil {
		return m.Outcome
	}
	return Outcome_OUTCOME_UNSPECIFIED
}

func (m *Application) GetIsMatch() *wrappers.BoolValue {
	if m != nil {
		return m.IsMatch
	}
	return nil
}

func (m *Application) GetJobTitleSnippet() string {
	if m != nil {
		return m.JobTitleSnippet
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "google.cloud.talent.v4beta1.Application")
	proto.RegisterEnum("google.cloud.talent.v4beta1.Application_ApplicationState", Application_ApplicationState_name, Application_ApplicationState_value)
	proto.RegisterEnum("google.cloud.talent.v4beta1.Application_ApplicationStage", Application_ApplicationStage_name, Application_ApplicationStage_value)
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/application.proto", fileDescriptor_application_dbba57710742dc0d)
}

var fileDescriptor_application_dbba57710742dc0d = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x71, 0xd3, 0x36, 0xdd, 0x31, 0xdd, 0xba, 0x03, 0xad, 0x86, 0xb2, 0x4b, 0x42, 0xf9,
	0x50, 0x84, 0x84, 0x2d, 0xca, 0x87, 0x84, 0x16, 0x21, 0xb9, 0xf6, 0x24, 0x6b, 0x29, 0xb1, 0xad,
	0xb1, 0xd9, 0x02, 0x37, 0xd6, 0xc4, 0x99, 0x1a, 0x47, 0xb6, 0xc7, 0xb2, 0x27, 0xbb, 0xec, 0x0b,
	0x70, 0xcd, 0x33, 0x70, 0xc7, 0xa3, 0xf1, 0x16, 0x68, 0xfc, 0xb1, 0x8d, 0xca, 0x92, 0xbd, 0xd8,
	0x3b, 0x9f, 0xff, 0xfc, 0xfe, 0x7f, 0x9d, 0x33, 0x67, 0x64, 0xf0, 0x65, 0xc2, 0x79, 0x92, 0x31,
	0x23, 0xce, 0xf8, 0x66, 0x65, 0x08, 0x9a, 0xb1, 0x42, 0x18, 0xcf, 0xbf, 0x59, 0x32, 0x41, 0xbf,
	0x32, 0x68, 0x59, 0x66, 0x69, 0x4c, 0x45, 0xca, 0x0b, 0xbd, 0xac, 0xb8, 0xe0, 0xf0, 0xc3, 0x16,
	0xd7, 0x1b, 0x5c, 0x6f, 0x71, 0xbd, 0xc3, 0x2f, 0x26, 0xbb, 0xb2, 0x62, 0x9e, 0xe7, 0x7d, 0xcc,
	0xc5, 0x67, 0xbb, 0xc8, 0x35, 0x5f, 0x76, 0xd8, 0xa8, 0xc3, 0x9a, 0x6a, 0xb9, 0xb9, 0x35, 0x44,
	0x9a, 0xb3, 0x5a, 0xd0, 0xbc, 0xec, 0x80, 0x8f, 0xee, 0x03, 0x2f, 0x2a, 0x5a, 0x96, 0xac, 0xaa,
	0xbb, 0xf3, 0xf3, 0xee, 0x5c, 0xbc, 0x2c, 0x99, 0xb1, 0xa2, 0x82, 0x75, 0xfa, 0xa3, 0x4e, 0xa7,
	0x65, 0x6a, 0xd0, 0xa2, 0xe0, 0xa2, 0x99, 0xb1, 0x73, 0x5d, 0xfe, 0x73, 0x04, 0x54, 0xf3, 0x6e,
	0x74, 0x08, 0xc1, 0x7e, 0x41, 0x73, 0x86, 0x94, 0xb1, 0x32, 0x79, 0x40, 0x9a, 0x6f, 0x38, 0x02,
	0x2a, 0xfb, 0x5d, 0xb0, 0xaa, 0xa0, 0x59, 0x94, 0xae, 0xd0, 0xa8, 0x39, 0x02, 0xbd, 0xe4, 0xac,
	0x20, 0x02, 0xc3, 0xb2, 0xe2, 0xb7, 0x69, 0xc6, 0xd0, 0x5e, 0x73, 0xd8, 0x97, 0x50, 0x03, 0x83,
	0x35, 0x5f, 0xa2, 0xfd, 0x46, 0x95, 0x9f, 0x92, 0x8d, 0x79, 0x5e, 0xd2, 0xe2, 0x25, 0x3a, 0x68,
	0xd9, 0xae, 0x84, 0x3f, 0x00, 0x6d, 0x6b, 0x09, 0x91, 0x1c, 0x01, 0x0d, 0xc7, 0xca, 0x44, 0xbd,
	0x3a, 0xd5, 0xbb, 0x55, 0xc8, 0xd9, 0x74, 0x9b, 0x0a, 0x46, 0x4e, 0xb6, 0x50, 0x29, 0x40, 0x0f,
	0x1c, 0xd4, 0x82, 0x26, 0x0c, 0xa9, 0x63, 0x65, 0xf2, 0xf0, 0xea, 0x7b, 0x7d, 0xc7, 0xf6, 0xf4,
	0xad, 0x89, 0xb7, 0xbf, 0x03, 0x19, 0x40, 0xda, 0x9c, 0x2e, 0x50, 0x30, 0x74, 0xfc, 0x56, 0x81,
	0xa2, 0x0d, 0x14, 0x0c, 0x4e, 0x01, 0x48, 0x0b, 0xc1, 0xaa, 0xe7, 0x29, 0x7b, 0x51, 0x23, 0x6d,
	0x3c, 0x98, 0xa8, 0x57, 0x9f, 0xef, 0x4c, 0x75, 0x7a, 0x9c, 0x6c, 0x39, 0xe1, 0x77, 0xe0, 0xa8,
	0x62, 0xb7, 0xac, 0xaa, 0x68, 0x86, 0x60, 0x73, 0x3f, 0x17, 0x7d, 0x4a, 0xff, 0x36, 0xf4, 0x6b,
	0xce, 0xb3, 0x67, 0x34, 0xdb, 0x30, 0xf2, 0x8a, 0x85, 0x4f, 0x80, 0x1a, 0x57, 0x8c, 0x0a, 0x16,
	0xc9, 0xa7, 0x85, 0xde, 0xfb, 0x1f, 0x6b, 0xd8, 0xbf, 0x3b, 0x02, 0x5a, 0x5c, 0x0a, 0xd2, 0xbc,
	0x29, 0x57, 0xaf, 0xcc, 0xef, 0xbf, 0xd9, 0xdc, 0xe2, 0x8d, 0xf9, 0x13, 0x70, 0xcc, 0x37, 0x22,
	0xe6, 0x39, 0x8b, 0x0a, 0x2e, 0x58, 0x8d, 0xce, 0x9a, 0xcd, 0xbf, 0xdb, 0x89, 0xae, 0xd4, 0xe0,
	0x8f, 0x60, 0xd8, 0xd5, 0xe8, 0xbc, 0xb9, 0xf1, 0x4f, 0x77, 0xde, 0x8d, 0xd7, 0xb2, 0xa4, 0x37,
	0xc1, 0x6f, 0xc1, 0x51, 0x5a, 0x47, 0x39, 0x15, 0xf1, 0x6f, 0xe8, 0xd1, 0x1b, 0xaf, 0x65, 0x98,
	0xd6, 0x0b, 0x89, 0xc2, 0x2f, 0xc0, 0xe9, 0x9a, 0x2f, 0x23, 0x91, 0x8a, 0x8c, 0x45, 0x75, 0x91,
	0x96, 0x25, 0x13, 0xe8, 0x71, 0xd3, 0xdf, 0xc9, 0x9a, 0x2f, 0x43, 0xa9, 0x07, 0xad, 0x7c, 0xf9,
	0xa7, 0x02, 0xb4, 0xfb, 0xdb, 0x85, 0x1f, 0x83, 0xc7, 0xa6, 0xef, 0xcf, 0x1d, 0xcb, 0x0c, 0x1d,
	0xcf, 0x8d, 0x82, 0xd0, 0x0c, 0x71, 0xf4, 0x93, 0x1b, 0xf8, 0xd8, 0x72, 0xa6, 0x0e, 0xb6, 0xb5,
	0x77, 0xe0, 0x09, 0x50, 0x1d, 0x37, 0xf2, 0x89, 0x37, 0x23, 0x38, 0x08, 0x34, 0x05, 0x9e, 0x03,
	0x68, 0x99, 0xae, 0xed, 0xd8, 0x92, 0xbd, 0x71, 0xc2, 0xa7, 0x36, 0xc1, 0x37, 0xda, 0x1e, 0x3c,
	0x03, 0xa7, 0x78, 0xe1, 0xcf, 0xbd, 0x5f, 0x30, 0xb9, 0x93, 0x07, 0xf0, 0x18, 0x3c, 0xb0, 0xbc,
	0x85, 0x3f, 0xc7, 0x21, 0xb6, 0xb5, 0x7d, 0x08, 0xc0, 0xa1, 0x35, 0xf7, 0x02, 0x6c, 0x6b, 0x07,
	0x97, 0x7f, 0xff, 0xa7, 0xa5, 0xe4, 0x75, 0x2d, 0xcd, 0xee, 0xb7, 0x34, 0x04, 0x03, 0x17, 0xdf,
	0x68, 0x8a, 0x0c, 0x0b, 0x2c, 0x82, 0xb1, 0xab, 0xed, 0xc1, 0x0f, 0xc0, 0xd9, 0x53, 0x87, 0x38,
	0xee, 0x2c, 0x5a, 0x98, 0xae, 0x39, 0xc3, 0x24, 0x22, 0xf8, 0x99, 0xd3, 0xb7, 0xe0, 0xb8, 0x21,
	0x26, 0x4d, 0xb9, 0x0f, 0x21, 0x78, 0xe8, 0x4d, 0xa7, 0x98, 0x44, 0xf8, 0xe7, 0x10, 0xbb, 0xb6,
	0x6c, 0xe5, 0x4e, 0x33, 0x2d, 0x0b, 0xfb, 0xb2, 0xd5, 0x43, 0xa8, 0x82, 0x61, 0x10, 0x9a, 0x44,
	0x16, 0xc3, 0xeb, 0x3f, 0x14, 0x30, 0x8a, 0x79, 0xbe, 0x6b, 0xad, 0xd7, 0x68, 0x6b, 0x18, 0xc2,
	0x6a, 0xbe, 0xa9, 0x62, 0xe6, 0xcb, 0xf5, 0xf9, 0xca, 0xaf, 0x66, 0x67, 0x4c, 0x78, 0x46, 0x8b,
	0x44, 0xe7, 0x55, 0x62, 0x24, 0xac, 0x68, 0x96, 0x6b, 0xb4, 0x47, 0xb4, 0x4c, 0xeb, 0xd7, 0xfe,
	0x68, 0x9f, 0xb4, 0xe5, 0x5f, 0x7b, 0x03, 0x2b, 0x0c, 0x96, 0x87, 0x8d, 0xe7, 0xeb, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x72, 0x18, 0xca, 0xff, 0x11, 0x06, 0x00, 0x00,
}
