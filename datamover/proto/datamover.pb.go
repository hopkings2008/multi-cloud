// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamover.proto

package datamover

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{0}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (dst *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(dst, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Filter struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Tag                  []*KV    `protobuf:"bytes,2,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{1}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Filter) GetTag() []*KV {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Connector struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	ConnConfig           []*KV    `protobuf:"bytes,3,rep,name=ConnConfig,proto3" json:"ConnConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connector) Reset()         { *m = Connector{} }
func (m *Connector) String() string { return proto.CompactTextString(m) }
func (*Connector) ProtoMessage()    {}
func (*Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{2}
}
func (m *Connector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connector.Unmarshal(m, b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
}
func (dst *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(dst, src)
}
func (m *Connector) XXX_Size() int {
	return xxx_messageInfo_Connector.Size(m)
}
func (m *Connector) XXX_DiscardUnknown() {
	xxx_messageInfo_Connector.DiscardUnknown(m)
}

var xxx_messageInfo_Connector proto.InternalMessageInfo

func (m *Connector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Connector) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *Connector) GetConnConfig() []*KV {
	if m != nil {
		return m.ConnConfig
	}
	return nil
}

type RunJobRequest struct {
	Context              string     `protobuf:"bytes,1,opt,name=Context,proto3" json:"Context,omitempty"`
	Id                   string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	SourceConn           *Connector `protobuf:"bytes,3,opt,name=sourceConn,proto3" json:"sourceConn,omitempty"`
	DestConn             *Connector `protobuf:"bytes,4,opt,name=destConn,proto3" json:"destConn,omitempty"`
	Filt                 *Filter    `protobuf:"bytes,5,opt,name=filt,proto3" json:"filt,omitempty"`
	RemainSource         bool       `protobuf:"varint,6,opt,name=remainSource,proto3" json:"remainSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunJobRequest) Reset()         { *m = RunJobRequest{} }
func (m *RunJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunJobRequest) ProtoMessage()    {}
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{3}
}
func (m *RunJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobRequest.Unmarshal(m, b)
}
func (m *RunJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobRequest.Marshal(b, m, deterministic)
}
func (dst *RunJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobRequest.Merge(dst, src)
}
func (m *RunJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunJobRequest.Size(m)
}
func (m *RunJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobRequest proto.InternalMessageInfo

func (m *RunJobRequest) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *RunJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RunJobRequest) GetSourceConn() *Connector {
	if m != nil {
		return m.SourceConn
	}
	return nil
}

func (m *RunJobRequest) GetDestConn() *Connector {
	if m != nil {
		return m.DestConn
	}
	return nil
}

func (m *RunJobRequest) GetFilt() *Filter {
	if m != nil {
		return m.Filt
	}
	return nil
}

func (m *RunJobRequest) GetRemainSource() bool {
	if m != nil {
		return m.RemainSource
	}
	return false
}

type RunJobResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunJobResponse) Reset()         { *m = RunJobResponse{} }
func (m *RunJobResponse) String() string { return proto.CompactTextString(m) }
func (*RunJobResponse) ProtoMessage()    {}
func (*RunJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{4}
}
func (m *RunJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobResponse.Unmarshal(m, b)
}
func (m *RunJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobResponse.Marshal(b, m, deterministic)
}
func (dst *RunJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobResponse.Merge(dst, src)
}
func (m *RunJobResponse) XXX_Size() int {
	return xxx_messageInfo_RunJobResponse.Size(m)
}
func (m *RunJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobResponse proto.InternalMessageInfo

func (m *RunJobResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type LifecycleActionRequest struct {
	Actx                 string   `protobuf:"bytes,1,opt,name=actx,proto3" json:"actx,omitempty"`
	ObjKey               string   `protobuf:"bytes,2,opt,name=objKey,proto3" json:"objKey,omitempty"`
	BucketName           string   `protobuf:"bytes,3,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Action               int32    `protobuf:"varint,4,opt,name=action,proto3" json:"action,omitempty"`
	SourceTier           int32    `protobuf:"varint,5,opt,name=sourceTier,proto3" json:"sourceTier,omitempty"`
	TargetTier           int32    `protobuf:"varint,6,opt,name=targetTier,proto3" json:"targetTier,omitempty"`
	SourceBackend        string   `protobuf:"bytes,7,opt,name=sourceBackend,proto3" json:"sourceBackend,omitempty"`
	TargetBackend        string   `protobuf:"bytes,8,opt,name=targetBackend,proto3" json:"targetBackend,omitempty"`
	ObjSize              int64    `protobuf:"varint,9,opt,name=objSize,proto3" json:"objSize,omitempty"`
	LastModified         int64    `protobuf:"varint,10,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	UploadId             string   `protobuf:"bytes,11,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionRequest) Reset()         { *m = LifecycleActionRequest{} }
func (m *LifecycleActionRequest) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionRequest) ProtoMessage()    {}
func (*LifecycleActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{5}
}
func (m *LifecycleActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionRequest.Unmarshal(m, b)
}
func (m *LifecycleActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionRequest.Marshal(b, m, deterministic)
}
func (dst *LifecycleActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionRequest.Merge(dst, src)
}
func (m *LifecycleActionRequest) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionRequest.Size(m)
}
func (m *LifecycleActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionRequest proto.InternalMessageInfo

func (m *LifecycleActionRequest) GetActx() string {
	if m != nil {
		return m.Actx
	}
	return ""
}

func (m *LifecycleActionRequest) GetObjKey() string {
	if m != nil {
		return m.ObjKey
	}
	return ""
}

func (m *LifecycleActionRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *LifecycleActionRequest) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceTier() int32 {
	if m != nil {
		return m.SourceTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetTargetTier() int32 {
	if m != nil {
		return m.TargetTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceBackend() string {
	if m != nil {
		return m.SourceBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetTargetBackend() string {
	if m != nil {
		return m.TargetBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetObjSize() int64 {
	if m != nil {
		return m.ObjSize
	}
	return 0
}

func (m *LifecycleActionRequest) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

func (m *LifecycleActionRequest) GetUploadId() string {
	if m != nil {
		return m.UploadId
	}
	return ""
}

type LifecycleActionResonse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionResonse) Reset()         { *m = LifecycleActionResonse{} }
func (m *LifecycleActionResonse) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionResonse) ProtoMessage()    {}
func (*LifecycleActionResonse) Descriptor() ([]byte, []int) {
	return fileDescriptor_datamover_74634334dc056cb6, []int{6}
}
func (m *LifecycleActionResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionResonse.Unmarshal(m, b)
}
func (m *LifecycleActionResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionResonse.Marshal(b, m, deterministic)
}
func (dst *LifecycleActionResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionResonse.Merge(dst, src)
}
func (m *LifecycleActionResonse) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionResonse.Size(m)
}
func (m *LifecycleActionResonse) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionResonse.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionResonse proto.InternalMessageInfo

func (m *LifecycleActionResonse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*KV)(nil), "KV")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*Connector)(nil), "Connector")
	proto.RegisterType((*RunJobRequest)(nil), "RunJobRequest")
	proto.RegisterType((*RunJobResponse)(nil), "RunJobResponse")
	proto.RegisterType((*LifecycleActionRequest)(nil), "LifecycleActionRequest")
	proto.RegisterType((*LifecycleActionResonse)(nil), "LifecycleActionResonse")
}

func init() { proto.RegisterFile("datamover.proto", fileDescriptor_datamover_74634334dc056cb6) }

var fileDescriptor_datamover_74634334dc056cb6 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x51, 0x8f, 0xd2, 0x4c,
	0x14, 0xfd, 0xda, 0x42, 0x81, 0xcb, 0xb7, 0xac, 0x4e, 0x74, 0x9d, 0x60, 0x62, 0x48, 0x35, 0x86,
	0xac, 0xa6, 0x0f, 0xf8, 0xe0, 0xb3, 0x8b, 0xd1, 0x28, 0xea, 0xc3, 0xec, 0x66, 0xdf, 0xa7, 0xed,
	0x85, 0x0c, 0x94, 0x4e, 0x9d, 0x4e, 0x37, 0x8b, 0x6f, 0xfe, 0x41, 0x7f, 0x84, 0xbf, 0xc4, 0xcc,
	0xb4, 0x05, 0xaa, 0xbc, 0xcd, 0x39, 0xf7, 0xdc, 0x73, 0xb9, 0xf7, 0x50, 0x38, 0x4f, 0xb8, 0xe6,
	0x5b, 0x79, 0x87, 0x2a, 0xcc, 0x95, 0xd4, 0x32, 0x78, 0x0d, 0xee, 0xe2, 0x96, 0x3c, 0x00, 0x6f,
	0x83, 0x3b, 0xea, 0x4c, 0x9c, 0xe9, 0x80, 0x99, 0x27, 0x79, 0x04, 0xdd, 0x3b, 0x9e, 0x96, 0x48,
	0x5d, 0xcb, 0x55, 0x20, 0x78, 0x0b, 0xfe, 0x07, 0x91, 0x6a, 0x54, 0xe4, 0x02, 0xfc, 0x5c, 0xe1,
	0x52, 0xdc, 0xd7, 0x4d, 0x35, 0x22, 0x8f, 0xc1, 0xd3, 0x7c, 0x45, 0xdd, 0x89, 0x37, 0x1d, 0xce,
	0xbc, 0x70, 0x71, 0xcb, 0x0c, 0x0e, 0x12, 0x18, 0xcc, 0x65, 0x96, 0x61, 0xac, 0xa5, 0x22, 0x04,
	0x3a, 0x37, 0xbb, 0x1c, 0xeb, 0x4e, 0xfb, 0x26, 0xcf, 0x00, 0xae, 0xca, 0x78, 0x83, 0xfa, 0x1b,
	0xdf, 0x36, 0x43, 0x8f, 0x18, 0xf2, 0x1c, 0xc0, 0x18, 0xcc, 0x65, 0xb6, 0x14, 0x2b, 0xea, 0x1d,
	0xec, 0x8f, 0xe8, 0xe0, 0x97, 0x03, 0x67, 0xac, 0xcc, 0x3e, 0xcb, 0x88, 0xe1, 0xf7, 0x12, 0x0b,
	0x4d, 0x28, 0xf4, 0xe6, 0x32, 0xd3, 0x78, 0xaf, 0xeb, 0x69, 0x0d, 0x24, 0x23, 0x70, 0x45, 0x52,
	0x0f, 0x72, 0x45, 0x42, 0x2e, 0x01, 0x0a, 0x59, 0xaa, 0x18, 0x8d, 0x1f, 0xf5, 0x26, 0xce, 0x74,
	0x38, 0x83, 0x70, 0xff, 0xa3, 0xd9, 0x51, 0x95, 0xbc, 0x84, 0x7e, 0x82, 0x85, 0xb6, 0xca, 0xce,
	0x3f, 0xca, 0x7d, 0x8d, 0x3c, 0x85, 0xce, 0x52, 0xa4, 0x9a, 0x76, 0xad, 0xa6, 0x17, 0x56, 0xb7,
	0x63, 0x96, 0x24, 0x01, 0xfc, 0xaf, 0x70, 0xcb, 0x45, 0x76, 0x6d, 0x8d, 0xa9, 0x3f, 0x71, 0xa6,
	0x7d, 0xd6, 0xe2, 0x82, 0x00, 0x46, 0xcd, 0x3e, 0x45, 0x2e, 0xb3, 0x02, 0x4d, 0x52, 0xa8, 0x54,
	0x93, 0x14, 0x2a, 0x15, 0xfc, 0x76, 0xe1, 0xe2, 0x8b, 0x58, 0x62, 0xbc, 0x8b, 0x53, 0x7c, 0x17,
	0x6b, 0x21, 0xb3, 0x66, 0x7b, 0x02, 0x1d, 0x1e, 0xeb, 0x26, 0x22, 0xfb, 0x36, 0xc1, 0xc9, 0x68,
	0xbd, 0xc0, 0x5d, 0xbd, 0x7b, 0x8d, 0x4c, 0x00, 0xd1, 0x21, 0x00, 0xaf, 0x0a, 0xe0, 0xc0, 0x98,
	0x3e, 0x6e, 0xcd, 0xed, 0xc6, 0x5d, 0x56, 0x23, 0xd3, 0x57, 0x5d, 0xe6, 0x46, 0xa0, 0xb2, 0x9b,
	0x76, 0xd9, 0x11, 0x63, 0xea, 0x9a, 0xab, 0x15, 0x6a, 0x5b, 0xf7, 0xab, 0xfa, 0x81, 0x21, 0x2f,
	0xe0, 0xac, 0x52, 0x5f, 0xf1, 0x78, 0x83, 0x59, 0x42, 0x7b, 0x76, 0x74, 0x9b, 0x34, 0xaa, 0xaa,
	0xa7, 0x51, 0xf5, 0x2b, 0x55, 0x8b, 0x34, 0x69, 0xcb, 0x68, 0x7d, 0x2d, 0x7e, 0x20, 0x1d, 0x4c,
	0x9c, 0xa9, 0xc7, 0x1a, 0x68, 0x8e, 0x9d, 0xf2, 0x42, 0x7f, 0x95, 0x89, 0x58, 0x0a, 0x4c, 0x28,
	0xd8, 0x72, 0x8b, 0x23, 0x63, 0xe8, 0x97, 0x79, 0x2a, 0x79, 0xf2, 0x29, 0xa1, 0x43, 0x6b, 0xbf,
	0xc7, 0xc1, 0xe5, 0x89, 0x1b, 0x17, 0xa7, 0x03, 0x99, 0xfd, 0x74, 0x60, 0xb0, 0xff, 0xcc, 0xc8,
	0x2b, 0xf0, 0x59, 0x99, 0xad, 0x65, 0x44, 0x46, 0x61, 0xeb, 0xbf, 0x39, 0x3e, 0x0f, 0xdb, 0xd9,
	0x06, 0xff, 0x91, 0x8f, 0xf0, 0xf0, 0xbd, 0xfc, 0x6b, 0x10, 0x79, 0x12, 0x9e, 0x8e, 0x77, 0x7c,
	0xa2, 0x50, 0x54, 0x46, 0x91, 0x6f, 0xbf, 0xee, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0x3f, 0xdd, 0x76, 0xf0, 0x03, 0x00, 0x00,
}
