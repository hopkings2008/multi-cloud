// Copyright 2019 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ceph

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"

	"io"
	"io/ioutil"
	"github.com/micro/go-log"
	backendpb "github.com/opensds/multi-cloud/backend/proto"
	"github.com/opensds/multi-cloud/s3/pkg/model"
	pb "github.com/opensds/multi-cloud/s3/proto"
	"github.com/webrtcn/s3client"
	. "github.com/webrtcn/s3client"
	dscommon "github.com/opensds/multi-cloud/s3/pkg/datastore/common"
	"github.com/webrtcn/s3client/models"
	. "github.com/opensds/multi-cloud/s3/error"
	"time"
)

type CephAdapter struct {
	backend *backendpb.BackendDetail
	session *s3client.Client
}

func md5Content(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	value := base64.StdEncoding.EncodeToString(cipherStr)
	return value
}

func (ad *CephAdapter) Put(ctx context.Context, stream io.Reader, object *pb.Object) (result dscommon.PutResult, err error) {
	bucketName := ad.backend.BucketName
	objectId := object.BucketName + "/" + object.ObjectKey
	log.Logf("put object[Ceph S3], bucket:%s, objectId:%s\n", bucketName, objectId)

	bucket := ad.session.NewBucket()
	cephObject := bucket.NewObject(bucketName)
	d, err := ioutil.ReadAll(stream)
	data := []byte(d)
	contentMD5 := md5Content(data)
	length := int64(len(d))
	body := ioutil.NopCloser(bytes.NewReader(data))

	log.Logf("put object[Ceph S3] begin, objectId:%s\n", objectId)
	err = cephObject.Create(objectId, contentMD5, "", length, body, models.Private)
	log.Logf("put object[Ceph S3] end, objectId:%s\n", objectId)
	if err != nil {
		log.Logf("upload object[Ceph S3] failed, objectId:%s, err:%v", objectId, err)
		return result, ErrPutToBackendFailed
	}

	result.UpdateTime = time.Now().Unix()
	result.ObjectId = objectId
	// TODO: set ETAG
	log.Logf("upload object[Ceph S3] succeed, objectId:%s, UpdateTime is:%v\n", objectId, result.UpdateTime)

	return result,nil
}

func (ad *CephAdapter) Get(ctx context.Context, object *pb.Object, start int64, end int64) (io.ReadCloser, error) {
	log.Logf("get object[Ceph S3], bucket:%s, objectId:%s\n", object.BucketName, object.ObjectId)

	getObjectOption := GetObjectOption{}
	if start != 0 || end != 0 {
		rangeObj := Range{
			Begin: start,
			End:   end,
		}
		getObjectOption = GetObjectOption{
			Range: &rangeObj,
		}
	}

	bucket := ad.session.NewBucket()
	cephObject := bucket.NewObject(ad.backend.BucketName)
	getObject, err := cephObject.Get(object.ObjectId, &getObjectOption)
	if err != nil {
		fmt.Println(err)
		log.Logf("get object[Ceph S3], objectId:%s failed:%v", object.ObjectId, err)
		return nil, ErrGetFromBackendFailed
	}

	log.Logf("get object[Ceph S3] succeed, objectId:%s, bytes:%d\n", object.ObjectId, getObject.ContentLength)
	return getObject.Body, nil
}

func (ad *CephAdapter) Delete(ctx context.Context, object *pb.DeleteObjectInput) error {
	bucket := ad.session.NewBucket()
	objectId := object.Bucket + "/" + object.Key
	log.Logf("delete object[Ceph S3], objectId:%s, bucket:%s\n", objectId, bucket)

	cephObject := bucket.NewObject(ad.backend.BucketName)
	err := cephObject.Remove(objectId)
	if err != nil {
		log.Logf("delete object[Ceph S3] failed, objectId:%s, err:%v\n", objectId, err)
		return ErrDeleteFromBackendFailed
	}

	log.Logf("delete object[Ceph S3] succeed, objectId:%s.\n", objectId)
	return nil
}

/*func (ad *CephAdapter) GetObjectInfo(context context.Context, bucketName string, key string) (*pb.Object, error) {
	bucket := ad.backend.BucketName
	newKey := bucketName + "/" + key

	bucketO := ad.session.NewBucket()
	bucketResp, err := bucketO.Get(bucket, newKey, "", "", 1000)
	if err != nil {
		log.Logf("error occured during get Object Info, err:%v\n", err)
		return nil, err
	}

	for _, content := range bucketResp.Contents {
		realKey := bucketName + "/" + key
		if realKey != content.Key {
			break
		}
		obj := &pb.Object{
			BucketName: bucketName,
			ObjectKey:  key,
			Size:       content.Size,
		}

		return obj, nil
	}

	log.Logf("can not find specified object(%s).\n", key)
	return nil, NoSuchObject.Error()
}*/

func (ad *CephAdapter) InitMultipartUpload(ctx context.Context, object *pb.Object) (*pb.MultipartUpload, error) {
	bucket := ad.session.NewBucket()
	objectId := object.BucketName + "/" + object.ObjectKey
	log.Logf("init multipart upload[Ceph S3], bucket = %v,objectId = %v\n", bucket, objectId)
	cephObject := bucket.NewObject(ad.backend.BucketName)
	uploader := cephObject.NewUploads(objectId)
	multipartUpload := &pb.MultipartUpload{}

	res, err := uploader.Initiate(nil)

	if err != nil {
		log.Fatalf("init multipart upload[Ceph S3] failed, objectId:%s, err:%v\n", objectId, err)
		return nil, err
	} else {
		log.Logf("init multipart upload[Ceph S3] succeed, objectId:%s, UploadId:%s\n", objectId, res.UploadID)
		multipartUpload.Bucket = object.BucketName
		multipartUpload.Key = object.ObjectKey
		multipartUpload.UploadId = res.UploadID
		multipartUpload.ObjectId = objectId
		return multipartUpload, nil
	}
}

func (ad *CephAdapter) UploadPart(ctx context.Context, stream io.Reader, multipartUpload *pb.MultipartUpload,
	partNumber int64, upBytes int64) (*model.UploadPartResult, error) {
	bucket := ad.session.NewBucket()
	log.Logf("upload part[Ceph S3], objectId:%s, bucket:%s\n", multipartUpload.ObjectId, bucket)

	cephObject := bucket.NewObject(ad.backend.BucketName)
	uploader := cephObject.NewUploads(multipartUpload.ObjectId)
	tries := 1
	for tries <= 3 {
		d, err := ioutil.ReadAll(stream)
		data := []byte(d)
		body := ioutil.NopCloser(bytes.NewReader(data))
		contentMD5 := md5Content(data)
		//length := int64(len(data))
		part, err := uploader.UploadPart(int(partNumber), multipartUpload.UploadId, contentMD5, "", upBytes, body)

		if err != nil {
			if tries == 3 {
				log.Logf("upload part[Ceph S3] failed, err:%v\n", err)
				return nil, ErrPutToBackendFailed
			}
			log.Logf("retrying to upload[Ceph S3] part#%d ,err:%s\n", partNumber, err)
			tries++
		} else {
			log.Logf("uploaded part[Ceph S3] #%d successfully, ETag:%s\n", partNumber, part.Etag)
			result := &model.UploadPartResult{
				Xmlns:      model.Xmlns,
				ETag:       part.Etag,
				PartNumber: partNumber}
			return result, nil
		}
	}

	log.Log("upload part[Ceph S3]: should not be here.")
	return nil, ErrInternalError
}

func (ad *CephAdapter) CompleteMultipartUpload(ctx context.Context, multipartUpload *pb.MultipartUpload,
	completeUpload *model.CompleteMultipartUpload) (*model.CompleteMultipartUploadResult, error) {
	bucket := ad.session.NewBucket()
	log.Logf("complete multipart upload[Ceph S3], objectId:%s, bucket:%s\n", multipartUpload.ObjectId, bucket)

	cephObject := bucket.NewObject(ad.backend.BucketName)
	uploader := cephObject.NewUploads(multipartUpload.ObjectId)
	var completeParts []CompletePart
	for _, p := range completeUpload.Part {
		completePart := CompletePart{
			Etag:       p.ETag,
			PartNumber: int(p.PartNumber),
		}
		completeParts = append(completeParts, completePart)
	}
	resp, err := uploader.Complete(multipartUpload.UploadId, completeParts)
	if err != nil {
		log.Logf("complete multipart upload[Ceph S3] failed, objectId:%s, err:%v\n", multipartUpload.ObjectId, err)
		return nil, ErrBackendCompleteMultipartFailed
	}
	result := &model.CompleteMultipartUploadResult{
		Xmlns:    model.Xmlns,
		Location: ad.backend.Endpoint,
		Bucket:   multipartUpload.Bucket,
		Key:      multipartUpload.Key,
		ETag:     resp.Etag,
	}

	log.Logf("complete multipart upload[Ceph S3] succeed, objectId:%s, resp:%v\n", multipartUpload.ObjectId, resp)
	return result, nil
}

func (ad *CephAdapter) AbortMultipartUpload(ctx context.Context, multipartUpload *pb.MultipartUpload) error {
	bucket := ad.session.NewBucket()
	cephObject := bucket.NewObject(ad.backend.BucketName)
	uploader := cephObject.NewUploads(multipartUpload.ObjectId)
	log.Logf("abort multipart upload[Ceph S3], objectId:%s, bucket:%s\n", multipartUpload.ObjectId, bucket)

	err := uploader.RemoveUploads(multipartUpload.UploadId)
	if err != nil {
		log.Logf("abort multipart upload[Ceph S3] failed, objectId:%s, err:%v\n", multipartUpload.ObjectId, err)
		return ErrBackendAbortMultipartFailed
	} else {
		log.Logf("abort multipart upload[Ceph S3] succeed, objectId:%s, err:%v\n", multipartUpload.ObjectId, err)
	}

	return nil
}

/*func (ad *CephAdapter) ListParts(context context.Context, listParts *pb.ListParts) (*model.ListPartsOutput, error) {
	newObjectKey := listParts.Bucket + "/" + listParts.Key
	bucket := ad.session.NewBucket()
	cephObject := bucket.NewObject(ad.backend.BucketName)
	uploader := cephObject.NewUploads(newObjectKey)

	listPartsResult, err := uploader.ListPart(listParts.UploadId)
	if err != nil {
		log.Logf("list parts failed, err:%v\n", err)
		return nil, S3Error{Code: 500, Description: err.Error()}.Error()
	} else {
		log.Logf("List parts successful\n")
		var parts []model.Part
		for _, p := range listPartsResult.Parts {
			part := model.Part{
				ETag:       p.Etag,
				PartNumber: int64(p.PartNumber),
			}
			parts = append(parts, part)
		}
		listPartsOutput := &model.ListPartsOutput{
			Xmlns:       model.Xmlns,
			Key:         listPartsResult.Key,
			Bucket:      listParts.Bucket,
			IsTruncated: listPartsResult.IsTruncated,
			MaxParts:    listPartsResult.MaxParts,
			Owner: model.Owner{
				ID:          listPartsResult.Owner.OwnerID,
				DisplayName: listPartsResult.Owner.DisplayName,
			},
			UploadId: listPartsResult.UploadID,
			Parts:    parts,
		}

		return listPartsOutput, nil
	}
}*/

func (ad *CephAdapter) Close(ctx context.Context) error {
	// TODO
	return nil
}

