// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: s3.proto

/*
Package s3 is a generated protocol buffer package.

It is generated from these files:
	s3.proto

It has these top-level messages:
	CopyObjectRequest
	PutObjectRequest
	PutBucketVersioningRequest
	PutBucketACLRequest
	BucketACL
	BucketVersioning
	BaseBucketRequest
	BaseObjRequest
	PutObjACLRequest
	ObjACL
	CopyObjPartRequest
	CopyObjPartResponse
	ServerSideEncryption
	VersioningConfiguration
	RedirectAllRequestsTo
	Redirect
	Condition
	RoutingRules
	WebsiteConfiguration
	CORSConfiguration
	Destination
	ReplicationRole
	ReplicationConfiguration
	Tag
	LifecycleFilter
	Action
	AbortMultipartUpload
	LifecycleRule
	ReplicationInfo
	Bucket
	Partion
	Version
	Object
	ListBucketsResponse
	BaseResponse
	BaseRequest
	ListObjectsRequest
	ListObjectResponse
	CountObjectsResponse
	DeleteObjectInput
	GetObjectInput
	MultipartUpload
	ListParts
	TList
	Tier2ClassName
	GetTierMapResponse
	UpdateObjMetaRequest
	StorageClass
	GetStorageClassesResponse
	GetBackendTypeByTierRequest
	GetBackendTypeByTierResponse
	DeleteLifecycleInput
	MultipartUploadRecord
	ListBucketPartsRequest
	ListBucketPartsResponse
*/
package s3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for S3 service

type S3Service interface {
	// TODO: Redefine parameters
	ListBuckets(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*ListBucketsResponse, error)
	CreateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error)
	DeleteBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error)
	GetBucket(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*Bucket, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*ListObjectResponse, error)
	CountObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*CountObjectsResponse, error)
	// rpc CreateObject(Object) returns (BaseResponse) {}
	PutObject(ctx context.Context, in *PutObjectRequest, opts ...client.CallOption) (*BaseResponse, error)
	UpdateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error)
	GetObject(ctx context.Context, in *GetObjectInput, opts ...client.CallOption) (*Object, error)
	DeleteObject(ctx context.Context, in *DeleteObjectInput, opts ...client.CallOption) (*BaseResponse, error)
	GetTierMap(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*GetTierMapResponse, error)
	UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetStorageClasses(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*GetStorageClassesResponse, error)
	GetBackendTypeByTier(ctx context.Context, in *GetBackendTypeByTierRequest, opts ...client.CallOption) (*GetBackendTypeByTierResponse, error)
	DeleteBucketLifecycle(ctx context.Context, in *DeleteLifecycleInput, opts ...client.CallOption) (*BaseResponse, error)
	UpdateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error)
	ListBucketUploadRecords(ctx context.Context, in *ListBucketPartsRequest, opts ...client.CallOption) (*ListBucketPartsResponse, error)
	InitMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AbortMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	CompleteMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	UploadPart(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	ListObjectParts(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	AppendObject(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	PostObject(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	// For lifecycle, may need some change.
	AddUploadRecord(ctx context.Context, in *MultipartUploadRecord, opts ...client.CallOption) (*BaseResponse, error)
	DeleteUploadRecord(ctx context.Context, in *MultipartUploadRecord, opts ...client.CallOption) (*BaseResponse, error)
	HeadObject(ctx context.Context, in *BaseObjRequest, opts ...client.CallOption) (*Object, error)
	CopyObject(ctx context.Context, in *CopyObjectRequest, opts ...client.CallOption) (*BaseResponse, error)
	CopyObjPart(ctx context.Context, in *CopyObjPartRequest, opts ...client.CallOption) (*CopyObjPartResponse, error)
	PutObjACL(ctx context.Context, in *PutObjACLRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetObjACL(ctx context.Context, in *BaseObjRequest, opts ...client.CallOption) (*ObjACL, error)
	GetBucketLocation(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetBucketVersioning(ctx context.Context, in *BaseBucketRequest, opts ...client.CallOption) (*BucketVersioning, error)
	PutBucketVersioning(ctx context.Context, in *PutBucketVersioningRequest, opts ...client.CallOption) (*BaseResponse, error)
	PutBucketACL(ctx context.Context, in *PutBucketACLRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetBucketACL(ctx context.Context, in *BaseBucketRequest, opts ...client.CallOption) (*BucketACL, error)
	PutBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeleteBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	PutBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	GetBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	DeleteBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error)
	HeadBucket(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*Bucket, error)
}

type s3Service struct {
	c    client.Client
	name string
}

func NewS3Service(name string, c client.Client) S3Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "s3"
	}
	return &s3Service{
		c:    c,
		name: name,
	}
}

func (c *s3Service) ListBuckets(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*ListBucketsResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListBuckets", in)
	out := new(ListBucketsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CreateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CreateBucket", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteBucket", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucket(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*Bucket, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucket", in)
	out := new(Bucket)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*ListObjectResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListObjects", in)
	out := new(ListObjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CountObjects(ctx context.Context, in *ListObjectsRequest, opts ...client.CallOption) (*CountObjectsResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CountObjects", in)
	out := new(CountObjectsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutObject(ctx context.Context, in *PutObjectRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UpdateObject(ctx context.Context, in *Object, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UpdateObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetObject(ctx context.Context, in *GetObjectInput, opts ...client.CallOption) (*Object, error) {
	req := c.c.NewRequest(c.name, "S3.GetObject", in)
	out := new(Object)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteObject(ctx context.Context, in *DeleteObjectInput, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetTierMap(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*GetTierMapResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetTierMap", in)
	out := new(GetTierMapResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UpdateObjMeta", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetStorageClasses(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*GetStorageClassesResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetStorageClasses", in)
	out := new(GetStorageClassesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBackendTypeByTier(ctx context.Context, in *GetBackendTypeByTierRequest, opts ...client.CallOption) (*GetBackendTypeByTierResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetBackendTypeByTier", in)
	out := new(GetBackendTypeByTierResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteBucketLifecycle(ctx context.Context, in *DeleteLifecycleInput, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteBucketLifecycle", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UpdateBucket(ctx context.Context, in *Bucket, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UpdateBucket", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) ListBucketUploadRecords(ctx context.Context, in *ListBucketPartsRequest, opts ...client.CallOption) (*ListBucketPartsResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListBucketUploadRecords", in)
	out := new(ListBucketPartsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) InitMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.InitMultipartUpload", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) AbortMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.AbortMultipartUpload", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CompleteMultipartUpload(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CompleteMultipartUpload", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) UploadPart(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.UploadPart", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) ListObjectParts(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.ListObjectParts", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) AppendObject(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.AppendObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PostObject(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PostObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) AddUploadRecord(ctx context.Context, in *MultipartUploadRecord, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.AddUploadRecord", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteUploadRecord(ctx context.Context, in *MultipartUploadRecord, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteUploadRecord", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) HeadObject(ctx context.Context, in *BaseObjRequest, opts ...client.CallOption) (*Object, error) {
	req := c.c.NewRequest(c.name, "S3.HeadObject", in)
	out := new(Object)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CopyObject(ctx context.Context, in *CopyObjectRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CopyObject", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) CopyObjPart(ctx context.Context, in *CopyObjPartRequest, opts ...client.CallOption) (*CopyObjPartResponse, error) {
	req := c.c.NewRequest(c.name, "S3.CopyObjPart", in)
	out := new(CopyObjPartResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutObjACL(ctx context.Context, in *PutObjACLRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutObjACL", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetObjACL(ctx context.Context, in *BaseObjRequest, opts ...client.CallOption) (*ObjACL, error) {
	req := c.c.NewRequest(c.name, "S3.GetObjACL", in)
	out := new(ObjACL)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucketLocation(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucketLocation", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucketVersioning(ctx context.Context, in *BaseBucketRequest, opts ...client.CallOption) (*BucketVersioning, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucketVersioning", in)
	out := new(BucketVersioning)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutBucketVersioning(ctx context.Context, in *PutBucketVersioningRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutBucketVersioning", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutBucketACL(ctx context.Context, in *PutBucketACLRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutBucketACL", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucketACL(ctx context.Context, in *BaseBucketRequest, opts ...client.CallOption) (*BucketACL, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucketACL", in)
	out := new(BucketACL)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutBucketCORS", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucketCORS", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteBucketCORS(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteBucketCORS", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) PutBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.PutBucketPolicy", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) GetBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.GetBucketPolicy", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) DeleteBucketPolicy(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "S3.DeleteBucketPolicy", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Service) HeadBucket(ctx context.Context, in *BaseRequest, opts ...client.CallOption) (*Bucket, error) {
	req := c.c.NewRequest(c.name, "S3.HeadBucket", in)
	out := new(Bucket)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S3 service

type S3Handler interface {
	// TODO: Redefine parameters
	ListBuckets(context.Context, *BaseRequest, *ListBucketsResponse) error
	CreateBucket(context.Context, *Bucket, *BaseResponse) error
	DeleteBucket(context.Context, *Bucket, *BaseResponse) error
	GetBucket(context.Context, *BaseRequest, *Bucket) error
	ListObjects(context.Context, *ListObjectsRequest, *ListObjectResponse) error
	CountObjects(context.Context, *ListObjectsRequest, *CountObjectsResponse) error
	// rpc CreateObject(Object) returns (BaseResponse) {}
	PutObject(context.Context, *PutObjectRequest, *BaseResponse) error
	UpdateObject(context.Context, *Object, *BaseResponse) error
	GetObject(context.Context, *GetObjectInput, *Object) error
	DeleteObject(context.Context, *DeleteObjectInput, *BaseResponse) error
	GetTierMap(context.Context, *BaseRequest, *GetTierMapResponse) error
	UpdateObjMeta(context.Context, *UpdateObjMetaRequest, *BaseResponse) error
	GetStorageClasses(context.Context, *BaseRequest, *GetStorageClassesResponse) error
	GetBackendTypeByTier(context.Context, *GetBackendTypeByTierRequest, *GetBackendTypeByTierResponse) error
	DeleteBucketLifecycle(context.Context, *DeleteLifecycleInput, *BaseResponse) error
	UpdateBucket(context.Context, *Bucket, *BaseResponse) error
	ListBucketUploadRecords(context.Context, *ListBucketPartsRequest, *ListBucketPartsResponse) error
	InitMultipartUpload(context.Context, *BaseRequest, *BaseResponse) error
	AbortMultipartUpload(context.Context, *BaseRequest, *BaseResponse) error
	CompleteMultipartUpload(context.Context, *BaseRequest, *BaseResponse) error
	UploadPart(context.Context, *BaseRequest, *BaseResponse) error
	ListObjectParts(context.Context, *BaseRequest, *BaseResponse) error
	AppendObject(context.Context, *BaseRequest, *BaseResponse) error
	PostObject(context.Context, *BaseRequest, *BaseResponse) error
	// For lifecycle, may need some change.
	AddUploadRecord(context.Context, *MultipartUploadRecord, *BaseResponse) error
	DeleteUploadRecord(context.Context, *MultipartUploadRecord, *BaseResponse) error
	HeadObject(context.Context, *BaseObjRequest, *Object) error
	CopyObject(context.Context, *CopyObjectRequest, *BaseResponse) error
	CopyObjPart(context.Context, *CopyObjPartRequest, *CopyObjPartResponse) error
	PutObjACL(context.Context, *PutObjACLRequest, *BaseResponse) error
	GetObjACL(context.Context, *BaseObjRequest, *ObjACL) error
	GetBucketLocation(context.Context, *BaseRequest, *BaseResponse) error
	GetBucketVersioning(context.Context, *BaseBucketRequest, *BucketVersioning) error
	PutBucketVersioning(context.Context, *PutBucketVersioningRequest, *BaseResponse) error
	PutBucketACL(context.Context, *PutBucketACLRequest, *BaseResponse) error
	GetBucketACL(context.Context, *BaseBucketRequest, *BucketACL) error
	PutBucketCORS(context.Context, *BaseRequest, *BaseResponse) error
	GetBucketCORS(context.Context, *BaseRequest, *BaseResponse) error
	DeleteBucketCORS(context.Context, *BaseRequest, *BaseResponse) error
	PutBucketPolicy(context.Context, *BaseRequest, *BaseResponse) error
	GetBucketPolicy(context.Context, *BaseRequest, *BaseResponse) error
	DeleteBucketPolicy(context.Context, *BaseRequest, *BaseResponse) error
	HeadBucket(context.Context, *BaseRequest, *Bucket) error
}

func RegisterS3Handler(s server.Server, hdlr S3Handler, opts ...server.HandlerOption) error {
	type s3 interface {
		ListBuckets(ctx context.Context, in *BaseRequest, out *ListBucketsResponse) error
		CreateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error
		DeleteBucket(ctx context.Context, in *Bucket, out *BaseResponse) error
		GetBucket(ctx context.Context, in *BaseRequest, out *Bucket) error
		ListObjects(ctx context.Context, in *ListObjectsRequest, out *ListObjectResponse) error
		CountObjects(ctx context.Context, in *ListObjectsRequest, out *CountObjectsResponse) error
		PutObject(ctx context.Context, in *PutObjectRequest, out *BaseResponse) error
		UpdateObject(ctx context.Context, in *Object, out *BaseResponse) error
		GetObject(ctx context.Context, in *GetObjectInput, out *Object) error
		DeleteObject(ctx context.Context, in *DeleteObjectInput, out *BaseResponse) error
		GetTierMap(ctx context.Context, in *BaseRequest, out *GetTierMapResponse) error
		UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, out *BaseResponse) error
		GetStorageClasses(ctx context.Context, in *BaseRequest, out *GetStorageClassesResponse) error
		GetBackendTypeByTier(ctx context.Context, in *GetBackendTypeByTierRequest, out *GetBackendTypeByTierResponse) error
		DeleteBucketLifecycle(ctx context.Context, in *DeleteLifecycleInput, out *BaseResponse) error
		UpdateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error
		ListBucketUploadRecords(ctx context.Context, in *ListBucketPartsRequest, out *ListBucketPartsResponse) error
		InitMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AbortMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		CompleteMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		UploadPart(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		ListObjectParts(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AppendObject(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		PostObject(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		AddUploadRecord(ctx context.Context, in *MultipartUploadRecord, out *BaseResponse) error
		DeleteUploadRecord(ctx context.Context, in *MultipartUploadRecord, out *BaseResponse) error
		HeadObject(ctx context.Context, in *BaseObjRequest, out *Object) error
		CopyObject(ctx context.Context, in *CopyObjectRequest, out *BaseResponse) error
		CopyObjPart(ctx context.Context, in *CopyObjPartRequest, out *CopyObjPartResponse) error
		PutObjACL(ctx context.Context, in *PutObjACLRequest, out *BaseResponse) error
		GetObjACL(ctx context.Context, in *BaseObjRequest, out *ObjACL) error
		GetBucketLocation(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetBucketVersioning(ctx context.Context, in *BaseBucketRequest, out *BucketVersioning) error
		PutBucketVersioning(ctx context.Context, in *PutBucketVersioningRequest, out *BaseResponse) error
		PutBucketACL(ctx context.Context, in *PutBucketACLRequest, out *BaseResponse) error
		GetBucketACL(ctx context.Context, in *BaseBucketRequest, out *BucketACL) error
		PutBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeleteBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		PutBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		GetBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		DeleteBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error
		HeadBucket(ctx context.Context, in *BaseRequest, out *Bucket) error
	}
	type S3 struct {
		s3
	}
	h := &s3Handler{hdlr}
	return s.Handle(s.NewHandler(&S3{h}, opts...))
}

type s3Handler struct {
	S3Handler
}

func (h *s3Handler) ListBuckets(ctx context.Context, in *BaseRequest, out *ListBucketsResponse) error {
	return h.S3Handler.ListBuckets(ctx, in, out)
}

func (h *s3Handler) CreateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error {
	return h.S3Handler.CreateBucket(ctx, in, out)
}

func (h *s3Handler) DeleteBucket(ctx context.Context, in *Bucket, out *BaseResponse) error {
	return h.S3Handler.DeleteBucket(ctx, in, out)
}

func (h *s3Handler) GetBucket(ctx context.Context, in *BaseRequest, out *Bucket) error {
	return h.S3Handler.GetBucket(ctx, in, out)
}

func (h *s3Handler) ListObjects(ctx context.Context, in *ListObjectsRequest, out *ListObjectResponse) error {
	return h.S3Handler.ListObjects(ctx, in, out)
}

func (h *s3Handler) CountObjects(ctx context.Context, in *ListObjectsRequest, out *CountObjectsResponse) error {
	return h.S3Handler.CountObjects(ctx, in, out)
}

func (h *s3Handler) PutObject(ctx context.Context, in *PutObjectRequest, out *BaseResponse) error {
	return h.S3Handler.PutObject(ctx, in, out)
}

func (h *s3Handler) UpdateObject(ctx context.Context, in *Object, out *BaseResponse) error {
	return h.S3Handler.UpdateObject(ctx, in, out)
}

func (h *s3Handler) GetObject(ctx context.Context, in *GetObjectInput, out *Object) error {
	return h.S3Handler.GetObject(ctx, in, out)
}

func (h *s3Handler) DeleteObject(ctx context.Context, in *DeleteObjectInput, out *BaseResponse) error {
	return h.S3Handler.DeleteObject(ctx, in, out)
}

func (h *s3Handler) GetTierMap(ctx context.Context, in *BaseRequest, out *GetTierMapResponse) error {
	return h.S3Handler.GetTierMap(ctx, in, out)
}

func (h *s3Handler) UpdateObjMeta(ctx context.Context, in *UpdateObjMetaRequest, out *BaseResponse) error {
	return h.S3Handler.UpdateObjMeta(ctx, in, out)
}

func (h *s3Handler) GetStorageClasses(ctx context.Context, in *BaseRequest, out *GetStorageClassesResponse) error {
	return h.S3Handler.GetStorageClasses(ctx, in, out)
}

func (h *s3Handler) GetBackendTypeByTier(ctx context.Context, in *GetBackendTypeByTierRequest, out *GetBackendTypeByTierResponse) error {
	return h.S3Handler.GetBackendTypeByTier(ctx, in, out)
}

func (h *s3Handler) DeleteBucketLifecycle(ctx context.Context, in *DeleteLifecycleInput, out *BaseResponse) error {
	return h.S3Handler.DeleteBucketLifecycle(ctx, in, out)
}

func (h *s3Handler) UpdateBucket(ctx context.Context, in *Bucket, out *BaseResponse) error {
	return h.S3Handler.UpdateBucket(ctx, in, out)
}

func (h *s3Handler) ListBucketUploadRecords(ctx context.Context, in *ListBucketPartsRequest, out *ListBucketPartsResponse) error {
	return h.S3Handler.ListBucketUploadRecords(ctx, in, out)
}

func (h *s3Handler) InitMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.InitMultipartUpload(ctx, in, out)
}

func (h *s3Handler) AbortMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.AbortMultipartUpload(ctx, in, out)
}

func (h *s3Handler) CompleteMultipartUpload(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.CompleteMultipartUpload(ctx, in, out)
}

func (h *s3Handler) UploadPart(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.UploadPart(ctx, in, out)
}

func (h *s3Handler) ListObjectParts(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.ListObjectParts(ctx, in, out)
}

func (h *s3Handler) AppendObject(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.AppendObject(ctx, in, out)
}

func (h *s3Handler) PostObject(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.PostObject(ctx, in, out)
}

func (h *s3Handler) AddUploadRecord(ctx context.Context, in *MultipartUploadRecord, out *BaseResponse) error {
	return h.S3Handler.AddUploadRecord(ctx, in, out)
}

func (h *s3Handler) DeleteUploadRecord(ctx context.Context, in *MultipartUploadRecord, out *BaseResponse) error {
	return h.S3Handler.DeleteUploadRecord(ctx, in, out)
}

func (h *s3Handler) HeadObject(ctx context.Context, in *BaseObjRequest, out *Object) error {
	return h.S3Handler.HeadObject(ctx, in, out)
}

func (h *s3Handler) CopyObject(ctx context.Context, in *CopyObjectRequest, out *BaseResponse) error {
	return h.S3Handler.CopyObject(ctx, in, out)
}

func (h *s3Handler) CopyObjPart(ctx context.Context, in *CopyObjPartRequest, out *CopyObjPartResponse) error {
	return h.S3Handler.CopyObjPart(ctx, in, out)
}

func (h *s3Handler) PutObjACL(ctx context.Context, in *PutObjACLRequest, out *BaseResponse) error {
	return h.S3Handler.PutObjACL(ctx, in, out)
}

func (h *s3Handler) GetObjACL(ctx context.Context, in *BaseObjRequest, out *ObjACL) error {
	return h.S3Handler.GetObjACL(ctx, in, out)
}

func (h *s3Handler) GetBucketLocation(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.GetBucketLocation(ctx, in, out)
}

func (h *s3Handler) GetBucketVersioning(ctx context.Context, in *BaseBucketRequest, out *BucketVersioning) error {
	return h.S3Handler.GetBucketVersioning(ctx, in, out)
}

func (h *s3Handler) PutBucketVersioning(ctx context.Context, in *PutBucketVersioningRequest, out *BaseResponse) error {
	return h.S3Handler.PutBucketVersioning(ctx, in, out)
}

func (h *s3Handler) PutBucketACL(ctx context.Context, in *PutBucketACLRequest, out *BaseResponse) error {
	return h.S3Handler.PutBucketACL(ctx, in, out)
}

func (h *s3Handler) GetBucketACL(ctx context.Context, in *BaseBucketRequest, out *BucketACL) error {
	return h.S3Handler.GetBucketACL(ctx, in, out)
}

func (h *s3Handler) PutBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.PutBucketCORS(ctx, in, out)
}

func (h *s3Handler) GetBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.GetBucketCORS(ctx, in, out)
}

func (h *s3Handler) DeleteBucketCORS(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.DeleteBucketCORS(ctx, in, out)
}

func (h *s3Handler) PutBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.PutBucketPolicy(ctx, in, out)
}

func (h *s3Handler) GetBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.GetBucketPolicy(ctx, in, out)
}

func (h *s3Handler) DeleteBucketPolicy(ctx context.Context, in *BaseRequest, out *BaseResponse) error {
	return h.S3Handler.DeleteBucketPolicy(ctx, in, out)
}

func (h *s3Handler) HeadBucket(ctx context.Context, in *BaseRequest, out *Bucket) error {
	return h.S3Handler.HeadBucket(ctx, in, out)
}
