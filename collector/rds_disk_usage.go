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

// +build !nodiskstats

package collector

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/prometheus/client_golang/prometheus"
)

type diskRdsstatsCollector struct {
	cname string
	descs []typedRdsFactorDesc
}

type typedRdsFactorDesc struct {
	desc      *prometheus.Desc
	valueType prometheus.ValueType
	factor    float64
}

func (d *typedRdsFactorDesc) mustNewConstMetric(value float64, labels ...string) prometheus.Metric {
	if d.factor != 0 {
		value *= d.factor
	}
	return prometheus.MustNewConstMetric(d.desc, d.valueType, value, labels...)
}

const (
	diskRdsSubsystem = "diskRds"
)

func init() {
	registerCollector("diskRdsStats", defaultEnabled, NewDiskRdsStatsCollector)
}

func DescribeResourceDiskUsage(rdsId string) *rds.DescribeResourceUsageResponse {
	client, err := rds.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)

	request := rds.CreateDescribeResourceUsageRequest()
	request.Scheme = "https"
	request.DBInstanceId = rdsId

	response, err := client.DescribeResourceUsage(request)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	if response.IsSuccess() {
		return response
	} else {
		fmt.Println("monitor value got  is null!")
		return nil
	}
	return nil
}

func NewDiskRdsStatsCollector() (Collector, error) {
	var diskLabelNames = []string{"alias"}

	return &diskRdsstatsCollector{
		cname: "execute collector " + diskRdsSubsystem,
		descs: []typedRdsFactorDesc{
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, diskRdsSubsystem, "diskUsed_total"),
					"This is the total usage of rds disk.",
					diskLabelNames,
					nil,
				), valueType: prometheus.CounterValue,
				factor: .001,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, diskRdsSubsystem, "logSize_total"),
					"This is the logSize of rds disk.",
					diskLabelNames,
					nil,
				), valueType: prometheus.CounterValue,
				factor: .001,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, diskRdsSubsystem, "dataSize_total"),
					"This is the dataSize of rds disk.",
					diskLabelNames,
					nil,
				), valueType: prometheus.CounterValue,
				factor: .001,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(namespace, diskRdsSubsystem, "sqlSize_total"),
					"This is the sqlSize of rds disk.",
					diskLabelNames,
					nil,
				), valueType: prometheus.CounterValue,
				factor: .001,
			},
		},
	}, nil
}

func (c *diskRdsstatsCollector) Update(ch chan<- prometheus.Metric) error {
	fmt.Println(c.cname)
	if RdsList == nil {
		//TODO
		fmt.Println("target is null， default is localhost")
		return nil
	}
	rdsId := RdsList[0]
	diskStats := DescribeResourceDiskUsage(rdsId)

	alias := GetRdsAlias(rdsId)

	ch <- c.descs[0].mustNewConstMetric(float64(diskStats.DiskUsed), alias)
	ch <- c.descs[1].mustNewConstMetric(float64(diskStats.LogSize), alias)
	ch <- c.descs[2].mustNewConstMetric(float64(diskStats.DataSize), alias)
	ch <- c.descs[3].mustNewConstMetric(float64(diskStats.SQLSize), alias)

	return nil
}
