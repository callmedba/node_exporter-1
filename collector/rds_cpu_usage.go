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
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/prometheus/client_golang/prometheus"
)

type cpuRDSCollector struct {
	cname       string
	fs          string
	cpuRdsUsage *prometheus.Desc
}

const (
	cpuCollectorRdsSubsystem = "cpuRds"
)

func init() {
	registerCollector("cpuRds", defaultEnabled, NewCPURDSCollector)
}

func DescribeResourceCpuUsage(rdsId string) float64 {
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
		value, _ := strconv.ParseFloat(valueStr[0:strings.IndexByte(valueStr, '&')], 64)
		return value
	} else {
		fmt.Println("monitor value got  is null!")
		return 0
	}
	return 0
}

func NewCPURDSCollector() (Collector, error) {
	fs := ""
	return &cpuRDSCollector{
		cname: "execute collector " + cpuCollectorRdsSubsystem,
		fs:    fs,
		cpuRdsUsage: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, cpuCollectorRdsSubsystem, "usage_total"),
			"Seconds the cpus spent in guests (VMs) for each mode.",
			[]string{"cpuRds", "alias"}, nil,
		),
	}, nil
}

func (c *cpuRDSCollector) Update(ch chan<- prometheus.Metric) error {
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

func (c *cpuRDSCollector) updateStat(ch chan<- prometheus.Metric) error {
	value := DescribeResourceCpuUsage(c.fs)
	fmt.Println(c.fs)
	alias := GetRdsAlias(c.fs)
	ch <- prometheus.MustNewConstMetric(c.cpuRdsUsage, prometheus.CounterValue, value, "total", alias)
	return nil
}
