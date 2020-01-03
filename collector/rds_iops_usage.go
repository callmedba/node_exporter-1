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

// +build !nocpuRds

package collector

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/prometheus/client_golang/prometheus"
)

type iopsRDSCollector struct {
	cname   string
	fs      string
	iopsRds *prometheus.Desc
}

const (
	iopsCollectorRdsSubsystem = "iopsRds"
)

func init() {
	registerCollector("iopsRds", defaultEnabled, NewIOPSRDSCollector)
}

func DescribeResourceIOPS(rdsId string) float64 {
	client, err := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	period := GetRdsDuration(rdsId)

	request := rds.CreateDescribeDBInstancePerformanceRequest()
	request.Scheme = "https"
	request.DBInstanceId = rdsId
	request.Key = "MySQL_IOPS"

	request.StartTime = time.Now().UTC().Add(period).Format("2006-01-02T15:04Z")
	request.EndTime = time.Now().UTC().Format("2006-01-02T15:04Z")

	response, err := client.DescribeDBInstancePerformance(request)
	if err != nil {
		fmt.Print(err.Error())
		return 0
	}
	if len(response.PerformanceKeys.PerformanceKey[0].Values.PerformanceValue) != 0 {
		valueStr := response.PerformanceKeys.PerformanceKey[0].Values.PerformanceValue[0].Value
		value, _ := strconv.ParseFloat(valueStr, 64)
		return value
	} else {
		fmt.Println("monitor value got  is null!")
		return 0
	}
	return 0
}

func NewIOPSRDSCollector() (Collector, error) {
	fs := ""
	return &iopsRDSCollector{
		cname: "execute collector " + iopsCollectorRdsSubsystem,
		fs:    fs,
		iopsRds: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, iopsCollectorRdsSubsystem, "usage_total"),
			"Seconds the iops spent in guests (VMs) for each mode.",
			[]string{"iopsRds", "alias"}, nil,
		),
	}, nil
}

func (c *iopsRDSCollector) Update(ch chan<- prometheus.Metric) error {
	fmt.Println(c.cname)
	if RdsList == nil {
		//TODO
		fmt.Println("target is null， default is localhost")
		return nil
	}

	c.fs = RdsList[0]

	if err := c.updateStat(ch); err != nil {
		return err
	}
	return nil
}

func (c *iopsRDSCollector) updateStat(ch chan<- prometheus.Metric) error {
	value := DescribeResourceIOPS(c.fs)
	alias := GetRdsAlias(c.fs)
	ch <- prometheus.MustNewConstMetric(c.iopsRds, prometheus.CounterValue, value, "total", alias)
	return nil
}
