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

package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/opensds/multi-cloud/api/pkg/utils/obs"
	//_ "github.com/opensds/multi-cloud/s3/pkg/datastore"
	handler "github.com/opensds/multi-cloud/s3/pkg/service"
	pb "github.com/opensds/multi-cloud/s3/proto"
	//"github.com/opensds/multi-cloud/s3/pkg/helper"
	//"github.com/opensds/multi-cloud/s3/pkg/meta/redis"
	//log "github.com/sirupsen/logrus"

	"github.com/opensds/multi-cloud/s3/pkg/helper"
	"github.com/opensds/multi-cloud/s3/pkg/meta/redis"
	log "github.com/sirupsen/logrus"
	"github.com/opensds/multi-cloud/s3/pkg/datastore/driver"
	"github.com/opensds/multi-cloud/s3/pkg/datastore/yig/config"
)

func main() {
	service := micro.NewService(
		micro.Name("s3"),
	)

	obs.InitLogs()
	fmt.Printf("Init s3 service\n")
	//service.Init()
	//micro.NewService()
	service.Init(micro.AfterStop(func() error {
		driver.FreeCloser()
		return nil
	}))

	helper.SetupConfig()
	//fmt.Printf("CONFIG:%+v\n", helper.CONFIG)

	log.Infof("YIG conf: %+v \n", helper.CONFIG)
	log.Infof("YIG instance ID:", helper.CONFIG.InstanceId)

	if helper.CONFIG.MetaCacheType > 0 || helper.CONFIG.EnableDataCache {
		// read common config settings
		cc, err := config.ReadCommonConfig("/etc/yig")
		if err != nil {
			return
		}
		redis.Initialize(cc)
	}

	pb.RegisterS3Handler(service.Server(), handler.NewS3Service())
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
