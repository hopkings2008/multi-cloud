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

package s3

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/emicklei/go-restful"
        log "github.com/sirupsen/logrus"
	"github.com/opensds/multi-cloud/api/pkg/common"
	"github.com/opensds/multi-cloud/s3/error"
	"github.com/opensds/multi-cloud/s3/proto"
)

var ChunkSize int = 2048

//ObjectPut -
func (s *APIService) ObjectPut(request *restful.Request, response *restful.Response) {
	bucketName := request.PathParameter(common.REQUEST_PATH_BUCKET_NAME)
	objectKey := request.PathParameter(common.REQUEST_PATH_OBJECT_KEY)
	backendName := request.HeaderParameter(common.REQUEST_HEADER_STORAGE_CLASS)
	url := request.Request.URL
	if strings.HasSuffix(url.String(), "/") {
		objectKey = objectKey + "/"
	}
	log.Infof("received request: PUT object, objectkey=%s, bucketName=%s\n:",
		objectKey, bucketName)

	// get size
	size, err := getSize(request, response)
	if err != nil {
		return
	}

	// check if specific bucket exist
	ctx := common.InitCtxWithAuthInfo(request)
	bucketMeta := s.getBucketMeta(ctx, bucketName)
	if bucketMeta == nil {
		response.WriteError(http.StatusInternalServerError, s3error.ErrGetBucketFailed)
		return
	}

	location := bucketMeta.DefaultLocation
	if backendName != "" {
		// check if backend exist
		if s.isBackendExist(ctx, backendName) == false {
			response.WriteError(http.StatusBadRequest, s3error.ErrGetBackendFailed)
		}
		location = backendName
	}

	md := map[string]string{
		common.CTX_KEY_OBJECT_KEY:  objectKey,
		common.CTX_KEY_BUCKET_NAME: bucketName,
		common.CTX_KEY_SIZE:        strconv.FormatInt(size, 10),
		common.CTX_KEY_LOCATION:    location,
	}
	ctx = common.InitCtxWithVal(request, md)

	var limitedDataReader io.Reader
	if size > 0 { // request.ContentLength is -1 if length is unknown
		limitedDataReader = io.LimitReader(request.Request.Body, size)
	} else {
		limitedDataReader = request.Request.Body
	}

	buf := make([]byte, ChunkSize)
	eof := false
	stream, err := s.s3Client.PutObject(ctx)
	defer stream.Close()
	for !eof {
		n, err := limitedDataReader.Read(buf)
		if err == io.EOF {
			eof = true
			break
		}
		if err != nil {
			log.Infof("read error:%v\n", err)
			break
		}
		err = stream.Send(&s3.PutObjectRequest{Data: buf[:n]})
		if err != nil {
			log.Infof("stream send error: %v\n", err)
			break
		}
	}

	// if read or send data failed, then close stream and return error
	if !eof {
		response.WriteError(http.StatusInternalServerError, s3error.ErrInternalError)
		return
	}

	// TODO: is this the right way to get response?
	rsp := s3.PutObjectResponse{}
	err = stream.RecvMsg(rsp)
	if err != nil {
		log.Infof("stream receive message failed:%v\n", err)
		response.WriteError(http.StatusInternalServerError, s3error.ErrInternalError)
	}

	log.Info("PUT object successfully.")
	response.WriteEntity(rsp)
}

func getSize(request *restful.Request, response *restful.Response) (int64, error) {
	// get content-length
	contentLenght := request.HeaderParameter(common.REQUEST_HEADER_CONTENT_LENGTH)
	size, err := strconv.ParseInt(contentLenght, 10, 64)
	if err != nil {
		log.Infof("parse contentLenght[%s] failed, err:%v\n", contentLenght, err)
		response.WriteError(http.StatusLengthRequired, s3error.ErrMissingContentLength)
		return 0, err
	}

	log.Infof("object size is %v\n", size)

	if size == 0 || size > common.MaxObjectSize {
		log.Infof("invalid contentLenght:%s\n", contentLenght)
		errMsg := fmt.Sprintf("invalid contentLenght[%s], it should be less than %d and more than 0",
			contentLenght, common.MaxObjectSize)
		err := errors.New(errMsg)
		response.WriteError(http.StatusBadRequest, err)
		return size, err
	}

	return size, nil
}

