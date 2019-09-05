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

package azure

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/micro/go-log"
	backendpb "github.com/opensds/multi-cloud/backend/proto"
	. "github.com/opensds/multi-cloud/s3/error"
	dscommon "github.com/opensds/multi-cloud/s3/pkg/datastore/common"
	"github.com/opensds/multi-cloud/s3/pkg/model"
	pb "github.com/opensds/multi-cloud/s3/proto"
)

// TryTimeout indicates the maximum time allowed for any single try of an HTTP request.
var MaxTimeForSingleHttpRequest = 50 * time.Minute

type AzureAdapter struct {
	backend      *backendpb.BackendDetail
	containerURL azblob.ContainerURL
}

/*func Init(backend *backendpb.BackendDetail) *AzureAdapter {
	endpoint := backend.Endpoint
	AccessKeyID := backend.Access
	AccessKeySecret := backend.Security
	ad := AzureAdapter{}
	containerURL, err := ad.createContainerURL(endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		log.Logf("AzureAdapter Init container URL faild:%v\n", err)
		return nil
	}
	adap := &AzureAdapter{backend: backend, containerURL: containerURL}
	log.Log("AzureAdapter Init succeed, container URL:", containerURL.String())
	return adap
}*/

func (ad *AzureAdapter) createContainerURL(endpoint string, acountName string, accountKey string) (azblob.ContainerURL,
	error) {
	credential, err := azblob.NewSharedKeyCredential(acountName, accountKey)

	if err != nil {
		log.Logf("create credential[Azure Blob] failed, err:%v\n", err)
		return azblob.ContainerURL{}, err
	}

	//create containerURL
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{
		Retry: azblob.RetryOptions{
			TryTimeout: MaxTimeForSingleHttpRequest,
		},
	})
	URL, _ := url.Parse(endpoint)

	return azblob.NewContainerURL(*URL, p), nil
}

func (ad *AzureAdapter) Put(ctx context.Context, stream io.Reader, object *pb.Object) (dscommon.PutResult, error) {
	objectId := object.BucketName + "/" + object.ObjectKey
	blobURL := ad.containerURL.NewBlockBlobURL(objectId)
	log.Logf("put object[Azure Blob], objectId:%s, blobURL is %v\n", objectId, blobURL)

	bytess, _ := ioutil.ReadAll(stream)
	result := dscommon.PutResult{}
	log.Logf("put object[Azure Blob] begin, objectId:%s\n", objectId)
	uploadResp, err := blobURL.Upload(ctx, bytes.NewReader(bytess), azblob.BlobHTTPHeaders{}, nil,
		azblob.BlobAccessConditions{})
	log.Logf("put object[Azure Blob] end, objectId:%s\n", objectId)
	if err != nil {
		log.Logf("put object[Azure Blob] failed, objectId:%s, err = %v\n", objectId, err)
		return result, ErrPutToBackendFailed
	}

	if uploadResp.StatusCode() != http.StatusCreated {
		log.Logf("put object[Azure Blob], objectId:%s, StatusCode:%d\n", objectId, uploadResp.StatusCode())
		return result, ErrPutToBackendFailed
	}

	// Currently, only support Hot
	_, err = blobURL.SetTier(ctx, azblob.AccessTierHot, azblob.LeaseAccessConditions{})
	if err != nil {
		log.Logf("set azure blob tier failed:%v\n", err)
		return result, ErrPutToBackendFailed
	}

	result.UpdateTime = time.Now().Unix()
	result.ObjectId = objectId
	// TODO: set ETAG
	log.Logf("upload object[Azure Blob] succeed, objectId:%s, UpdateTime is:%v\n", objectId, result.UpdateTime)

	return result, nil
}

func (ad *AzureAdapter) Get(ctx context.Context, object *pb.Object, start int64, end int64) (io.ReadCloser, error) {
	bucket := ad.backend.BucketName
	log.Logf("get object[Azure Blob], bucket:%s, objectId:%s\n", bucket, object.ObjectId)

	blobURL := ad.containerURL.NewBlobURL(object.ObjectId)
	log.Logf("blobURL:%v, size:%d\n", blobURL, object.Size)

	len := object.Size
	var buf []byte
	if start != 0 || end != 0 {
		count := end - start + 1
		buf = make([]byte, count)
		err := azblob.DownloadBlobToBuffer(ctx, blobURL, start, count, buf, azblob.DownloadFromBlobOptions{})
		if err != nil {
			log.Logf("get object[Azure Blob] failed, objectId:%s, err:%v\n", object.ObjectId, err)
			return nil, ErrGetFromBackendFailed
		}
		body := bytes.NewReader(buf)
		ioReaderClose := ioutil.NopCloser(body)
		return ioReaderClose, nil
	} else {
		buf = make([]byte, len)
		//err := azblob.DownloadBlobToBuffer(context, blobURL, 0, 0, buf, azblob.DownloadFromBlobOptions{})
		downloadResp, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{},
			false)
		_, readErr := downloadResp.Response().Body.Read(buf)
		if readErr != nil {
			log.Logf("readErr[objkey:%s]=%v\n", object.ObjectId, readErr)
			return nil, ErrGetFromBackendFailed
		}
		if err != nil {
			log.Logf("get object[Azure Blob] failed, objectId:%s, err:%v\n", object.ObjectId, err)
			return nil, ErrGetFromBackendFailed
		}
		body := bytes.NewReader(buf)
		ioReaderClose := ioutil.NopCloser(body)
		return ioReaderClose, nil
	}

	log.Log("get object[Azure Blob]: should not be here")
	return nil, ErrInternalError
}

func (ad *AzureAdapter) Delete(ctx context.Context, input *pb.DeleteObjectInput) error {
	bucket := ad.backend.BucketName
	objectId := input.Bucket + "/" + input.Key
	log.Logf("delete object[Azure Blob], objectId:%s, bucket:%s\n", objectId, bucket)

	blobURL := ad.containerURL.NewBlockBlobURL(objectId)
	log.Logf("blobURL is %v\n", blobURL)
	delRsp, err := blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionInclude, azblob.BlobAccessConditions{})
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok { // This error is a Service-specific
			log.Logf("delete service code:%s\n", serr.ServiceCode())
			if string(serr.ServiceCode()) == string(azblob.StorageErrorCodeBlobNotFound) {
				return nil
			}
		}

		log.Logf("delete object[Azure Blob] failed, objectId:%s, err:%v\n", objectId, err)
		return ErrDeleteFromBackendFailed
	}

	if delRsp.StatusCode() != http.StatusOK && delRsp.StatusCode() != http.StatusAccepted {
		log.Logf("delete object[Azure Blob] failed, objectId:%s, status code:%d\n", objectId, delRsp.StatusCode())
		return ErrDeleteFromBackendFailed
	}

	log.Logf("delete object[Azure Blob] succeed, objectId:%s\n", objectId)
	return nil
}

func (ad *AzureAdapter) GetObjectInfo(bucketName string, key string, context context.Context) (*pb.Object, error) {
	object := pb.Object{}
	object.BucketName = bucketName
	object.ObjectKey = key
	return &object, nil
}

func (ad *AzureAdapter) InitMultipartUpload(ctx context.Context, object *pb.Object) (*pb.MultipartUpload, error) {
	bucket := ad.backend.BucketName
	log.Logf("bucket is %v\n", bucket)
	multipartUpload := &pb.MultipartUpload{}
	multipartUpload.Key = object.ObjectKey

	multipartUpload.Bucket = object.BucketName
	multipartUpload.UploadId = object.ObjectKey
	multipartUpload.ObjectId = object.BucketName + "/" + object.ObjectKey
	return multipartUpload, nil
}

func (ad *AzureAdapter) Int64ToBase64(blockID int64) string {
	buf := (&[8]byte{})[:]
	binary.LittleEndian.PutUint64(buf, uint64(blockID))
	return ad.BinaryToBase64(buf)
}

func (ad *AzureAdapter) BinaryToBase64(binaryID []byte) string {
	return base64.StdEncoding.EncodeToString(binaryID)
}

func (ad *AzureAdapter) Base64ToInt64(base64ID string) int64 {
	bin, _ := base64.StdEncoding.DecodeString(base64ID)
	return int64(binary.LittleEndian.Uint64(bin))
}

func (ad *AzureAdapter) UploadPart(ctx context.Context, stream io.Reader, multipartUpload *pb.MultipartUpload,
	partNumber int64, upBytes int64) (*model.UploadPartResult, error) {
	bucket := ad.backend.BucketName
	log.Logf("upload part[Azure Blob], bucket:%s, objectId:%s\n", bucket, multipartUpload.ObjectId)

	blobURL := ad.containerURL.NewBlockBlobURL(multipartUpload.UploadId)
	base64ID := ad.Int64ToBase64(partNumber)
	bytess, _ := ioutil.ReadAll(stream)
	_, err := blobURL.StageBlock(ctx, base64ID, bytes.NewReader(bytess), azblob.LeaseAccessConditions{}, nil)
	if err != nil {
		log.Logf("stage block[#%d,base64ID:%s] failed:%v\n", partNumber, base64ID, err)
		return nil, ErrPutToBackendFailed
	}

	log.Logf("stage block[#%d,base64ID:%s] succeed.\n", partNumber, base64ID)
	result := &model.UploadPartResult{PartNumber: partNumber, ETag: multipartUpload.UploadId}

	return result, nil
}

func (ad *AzureAdapter) CompleteMultipartUpload(ctx context.Context, multipartUpload *pb.MultipartUpload,
	completeUpload *model.CompleteMultipartUpload) (*model.CompleteMultipartUploadResult, error) {
	bucket := ad.backend.BucketName
	result := model.CompleteMultipartUploadResult{}
	result.Bucket = multipartUpload.Bucket
	result.Key = multipartUpload.Key
	result.Location = ad.backend.Name
	log.Logf("complete multipart upload[Azure Blob], bucket:%s, objectId:%s\n", bucket, multipartUpload.ObjectId)

	blobURL := ad.containerURL.NewBlockBlobURL(multipartUpload.ObjectId)
	var completeParts []string
	for _, p := range completeUpload.Part {
		base64ID := ad.Int64ToBase64(p.PartNumber)
		completeParts = append(completeParts, base64ID)
	}

	_, err := blobURL.CommitBlockList(ctx, completeParts, azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
	if err != nil {
		log.Logf("commit blocks[bucket:%s, objectId:%s] failed:%v\n", bucket, multipartUpload.ObjectId, err)
		return nil, ErrBackendCompleteMultipartFailed
	} else {
		log.Logf("commit blocks succeed.\n")
		// Currently, only support Hot
		_, err = blobURL.SetTier(ctx, azblob.AccessTierHot, azblob.LeaseAccessConditions{})
		if err != nil {
			log.Logf("set blob[objectId:%s] tier failed:%v\n", multipartUpload.ObjectId, err)
			return nil, ErrBackendCompleteMultipartFailed
		}
	}

	log.Logf("complete multipart upload[Azure Blob], bucket:%s, objectId:%s succeed\n",
		bucket, multipartUpload.ObjectId)
	return &result, nil
}

func (ad *AzureAdapter) AbortMultipartUpload(ctx context.Context, multipartUpload *pb.MultipartUpload) error {
	bucket := ad.backend.BucketName
	log.Logf("no need to abort multipart upload[objkey:%s].\n", bucket)
	return nil
}

func (ad *AzureAdapter) Close(ctx context.Context) error {
	// TODO:
	return nil
}
