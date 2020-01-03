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

// +build !nomeminfoRds

package collector

import (
	"fmt"
	"github.com/prometheus/common/log"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	memInfoRdsSubsystem = "memoryRds"
)

type meminfoRDSCollector struct {
	cname string
}

func init() {
	registerCollector("meminfoRds", defaultEnabled, NewMemInfoRDSCollector)
}

func DescribeResourceMemUsage(rdsId string) float64 {
	client, err := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	period := GetRdsDuration(rdsId)

	request := rds.CreateDescribeDBInstancePerformanceRequest()
	request.Scheme = "https"
	request.DBInstanceId = rdsId
	request.Key = "MySQL_MemCpuUsage"

	request.StartTime = time.Now().UTC().Add(period).Format("2006-01-02T15:04Z")
	request.EndTime = time.Now().UTC().Format("2006-01-02T15:04Z")

	response, err := client.DescribeDBInstancePerformance(request)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}

	if len(response.PerformanceKeys.PerformanceKey[0].Values.PerformanceValue) != 0 {
		valueStr := response.PerformanceKeys.PerformanceKey[0].Values.PerformanceValue[0].Value
		value, _ := strconv.ParseFloat(valueStr[strings.IndexByte(valueStr, '&')+1:], 64)
		return value
	} else {
		fmt.Println("monitor value got  is null!")
		return 0
	}
	return 0
}

func NewMemInfoRDSCollector() (Collector, error) {
	return &meminfoRDSCollector{
		cname: "execute collector " + memInfoRdsSubsystem,
	}, nil
}

func (c *meminfoRDSCollector) Update(ch chan<- prometheus.Metric) error {
	fmt.Println(c.cname)
	if RdsList == nil {
		//TODO
		fmt.Println("target is null， default is localhost")
		return nil
	}
	rdsId := RdsList[0]
	var metricType prometheus.ValueType
	memUsageInfo := DescribeResourceMemUsage(rdsId)

	log.Debugf("Set node_mem: %#v", memUsageInfo)

	metricType = prometheus.CounterValue

	alias := GetRdsAlias(rdsId)

	ch <- prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(namespace, memInfoRdsSubsystem, "MemUsageTotal"),
			fmt.Sprintf("Memory information field %s.", "MemUsageTotal"),
			[]string{"alias"}, nil,
		),
		metricType, memUsageInfo, alias)

	return nil
}
