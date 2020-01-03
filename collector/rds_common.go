// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collector

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

//自定义，存储rds id
var RdsList []string

var RdsDuration = make(map[string]time.Duration)

var RdsAttribute = make(map[string]*rds.DBInstanceAttribute)

const (
	regionId        = "cn-hangzhou"
	accessKeyId     = "accessKeyId"
	accessKeySecret = "accessKeySecret"
)

func GetRdsDuration(rdsId string) time.Duration {
	dura, ok := RdsDuration[rdsId]
	if ok {
		if dura != 0 {
			return dura
		} else {
			RdsDuration[rdsId] = DescribeDBInstanceMonitor(rdsId)
			return RdsDuration[rdsId]
		}
	} else {
		RdsDuration[rdsId] = DescribeDBInstanceMonitor(rdsId)
		return RdsDuration[rdsId]
	}
}

func DescribeDBInstanceMonitor(rdsId string) time.Duration {

	client, err := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	request := rds.CreateDescribeDBInstanceMonitorRequest()
	request.Scheme = "https"
	request.DBInstanceId = rdsId

	response, err := client.DescribeDBInstanceMonitor(request)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}
	period, _ := strconv.Atoi(response.Period)

	interval := -1 * time.Minute

	//interval = interval/60
	switch period {
	case 0:
		fmt.Println("Error, monitor interval time is 0")
		break
	case 60:
		interval = -1 * time.Minute
		fmt.Println("monitor interval time is 1 minute")
		break
	case 300:
		interval = -5 * time.Minute
		fmt.Println("monitor interval time is 5 minute")
		break
	default:
		//默认间隔时间为6分钟
		interval = -6 * time.Minute
		fmt.Println("monitor interval time is default 5 minute")
	}

	return interval
}

func GetRdsAlias(rdsId string) string {
	attr, ok := RdsAttribute[rdsId]
	if ok {
		if attr.DBInstanceDescription != "" {
			return attr.DBInstanceDescription
		} else {
			RdsAttribute[rdsId] = DescribeDBInstanceAttribute(rdsId)
			return GetRdsAlias(rdsId)
		}
	} else {
		RdsAttribute[rdsId] = DescribeDBInstanceAttribute(rdsId)
		return GetRdsAlias(rdsId)
	}
}

func DescribeDBInstanceAttribute(rdsId string) *rds.DBInstanceAttribute {

	client, _ := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	request := rds.CreateDescribeDBInstanceAttributeRequest()

	request.DBInstanceId = rdsId

	resp, _ := client.DescribeDBInstanceAttribute(request)

	if resp.GetHttpStatus() == 200 {
		return &resp.Items.DBInstanceAttribute[0]

	} else {
		return nil
	}

}
