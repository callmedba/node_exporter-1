# 写在前面

有很多小公司目前都使用了阿里云aliyun的rds数据库， 但是对于rds的监控基本上都是使用的阿里云的云监控系统，
没有自己的监控插件。 为了解决阿里云锁死小公司监控技术的问题，本人基于流行的Prometheus监控系统，在其
监控插件源代码的起初上进行了修改。这个开源node_exporter 的开源项目是其中之一，旨在解决rds机器层面的数据
采集问题， 调用aliyun的API， 实现一个exporter监控多个rds的功能。
本项目和本人另外一个开源项目mysqld_exporter 都是基于官方prometheus 插件直接修改，基于go语言，
稳定性和性能都可以保障，大家可以放心使用!

# Node exporter

[![CircleCI](https://circleci.com/gh/prometheus/node_exporter/tree/master.svg?style=shield)][circleci]
[![Buildkite status](https://badge.buildkite.com/94a0c1fb00b1f46883219c256efe9ce01d63b6505f3a942f9b.svg)](https://buildkite.com/prometheus/node-exporter)
[![Docker Repository on Quay](https://quay.io/repository/prometheus/node-exporter/status)][quay]
[![Docker Pulls](https://img.shields.io/docker/pulls/prom/node-exporter.svg?maxAge=604800)][hub]
[![Go Report Card](https://goreportcard.com/badge/github.com/prometheus/node_exporter)][goreportcard]

Prometheus exporter for hardware and OS metrics exposed by \*NIX kernels, written
in Go with pluggable metric collectors.

The [WMI exporter](https://github.com/martinlindhe/wmi_exporter) is recommended for Windows users.
To expose NVIDIA GPU metrics, [prometheus-dcgm
](https://github.com/NVIDIA/gpu-monitoring-tools/tree/master/exporters/prometheus-dcgm)
can be used.

## Collectors

There is varying support for collectors on each operating system. The tables
below list all existing collectors and the supported systems.

Collectors are enabled by providing a `--collector.<name>` flag.
Collectors that are enabled by default can be disabled by providing a `--no-collector.<name>` flag.

### Enabled by default

Name     | Description | OS
---------|-------------|----
arp | Exposes ARP statistics from `/proc/net/arp`. | Linux
bcache | Exposes bcache statistics from `/sys/fs/bcache/`. | Linux
bonding | Exposes the number of configured and active slaves of Linux bonding interfaces. | Linux
boottime | Exposes system boot time derived from the `kern.boottime` sysctl. | Darwin, Dragonfly, FreeBSD, NetBSD, OpenBSD, Solaris
conntrack | Shows conntrack statistics (does nothing if no `/proc/sys/net/netfilter/` present). | Linux
cpu | Exposes CPU statistics | Darwin, Dragonfly, FreeBSD, Linux, Solaris
cpufreq | Exposes CPU frequency statistics | Linux, Solaris
diskstats | Exposes disk I/O statistics. | Darwin, Linux, OpenBSD
edac | Exposes error detection and correction statistics. | Linux
entropy | Exposes available entropy. | Linux
exec | Exposes execution statistics. | Dragonfly, FreeBSD
filefd | Exposes file descriptor statistics from `/proc/sys/fs/file-nr`. | Linux
filesystem | Exposes filesystem statistics, such as disk space used. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
hwmon | Expose hardware monitoring and sensor data from `/sys/class/hwmon/`. | Linux
infiniband | Exposes network statistics specific to InfiniBand and Intel OmniPath configurations. | Linux
ipvs | Exposes IPVS status from `/proc/net/ip_vs` and stats from `/proc/net/ip_vs_stats`. | Linux
loadavg | Exposes load average. | Darwin, Dragonfly, FreeBSD, Linux, NetBSD, OpenBSD, Solaris
mdadm | Exposes statistics about devices in `/proc/mdstat` (does nothing if no `/proc/mdstat` present). | Linux
meminfo | Exposes memory statistics. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
netclass | Exposes network interface info from `/sys/class/net/` | Linux
netdev | Exposes network interface statistics such as bytes transferred. | Darwin, Dragonfly, FreeBSD, Linux, OpenBSD
netstat | Exposes network statistics from `/proc/net/netstat`. This is the same information as `netstat -s`. | Linux
nfs | Exposes NFS client statistics from `/proc/net/rpc/nfs`. This is the same information as `nfsstat -c`. | Linux
nfsd | Exposes NFS kernel server statistics from `/proc/net/rpc/nfsd`. This is the same information as `nfsstat -s`. | Linux
pressure | Exposes pressure stall statistics from `/proc/pressure/`. | Linux (kernel 4.20+ and/or [CONFIG\_PSI](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/accounting/psi.txt))
sockstat | Exposes various statistics from `/proc/net/sockstat`. | Linux
stat | Exposes various statistics from `/proc/stat`. This includes boot time, forks and interrupts. | Linux
textfile | Exposes statistics read from local disk. The `--collector.textfile.directory` flag must be set. | _any_
time | Exposes the current system time. | _any_
timex | Exposes selected adjtimex(2) system call stats. | Linux
uname | Exposes system information as provided by the uname system call. | FreeBSD, Linux
vmstat | Exposes statistics from `/proc/vmstat`. | Linux
xfs | Exposes XFS runtime statistics. | Linux (kernel 4.4+)
zfs | Exposes [ZFS](http://open-zfs.org/) performance statistics. | [Linux](http://zfsonlinux.org/), Solaris
cpuRds | Exposes Rds  cpu info runtime statistics. | [阿里云API](https://help.aliyun.com/document_detail/26226.html?spm=a2c4g.11186623.6.1485.23983189fxqVOa)
diskRdsStats | Exposes Rds  disk info runtime statistics. | [阿里云API](https://help.aliyun.com/document_detail/26226.html?spm=a2c4g.11186623.6.1485.23983189fxqVOa)
meminfoRds | Exposes Rds  memory info runtime statistics. | [阿里云API](https://help.aliyun.com/document_detail/26226.html?spm=a2c4g.11186623.6.1485.23983189fxqVOa)
iopsRds | Exposes Rds  iops info  runtime statistics. | [阿里云API](https://help.aliyun.com/document_detail/26226.html?spm=a2c4g.11186623.6.1485.23983189fxqVOa)
### Disabled by default

The perf collector may not work by default on all Linux systems due to kernel
configuration and security settings. To allow access, set the following sysctl
parameter:

```
sysctl -w kernel.perf_event_paranoid=X
```

- 2 allow only user-space measurements (default since Linux 4.6).
- 1 allow both kernel and user measurements (default before Linux 4.6).
- 0 allow access to CPU-specific data but not raw tracepoint samples.
- -1 no restrictions.

Depending on the configured value different metrics will be available, for most
cases `0` will provide the most complete set. For more information see [`man 2
perf_event_open`](http://man7.org/linux/man-pages/man2/perf_event_open.2.html).

Name     | Description | OS
---------|-------------|----
buddyinfo | Exposes statistics of memory fragments as reported by /proc/buddyinfo. | Linux
devstat | Exposes device statistics | Dragonfly, FreeBSD
drbd | Exposes Distributed Replicated Block Device statistics (to version 8.4) | Linux
interrupts | Exposes detailed interrupts statistics. | Linux, OpenBSD
ksmd | Exposes kernel and system statistics from `/sys/kernel/mm/ksm`. | Linux
logind | Exposes session counts from [logind](http://www.freedesktop.org/wiki/Software/systemd/logind/). | Linux
meminfo\_numa | Exposes memory statistics from `/proc/meminfo_numa`. | Linux
mountstats | Exposes filesystem statistics from `/proc/self/mountstats`. Exposes detailed NFS client statistics. | Linux
ntp | Exposes local NTP daemon health to check [time](./docs/TIME.md) | _any_
processes | Exposes aggregate process statistics from `/proc`. | Linux
qdisc | Exposes [queuing discipline](https://en.wikipedia.org/wiki/Network_scheduler#Linux_kernel) statistics | Linux
runit | Exposes service status from [runit](http://smarden.org/runit/). | _any_
supervisord | Exposes service status from [supervisord](http://supervisord.org/). | _any_
systemd | Exposes service and system status from [systemd](http://www.freedesktop.org/wiki/Software/systemd/). | Linux
tcpstat | Exposes TCP connection status information from `/proc/net/tcp` and `/proc/net/tcp6`. (Warning: the current version has potential performance issues in high load situations.) | Linux
wifi | Exposes WiFi device and station statistics. | Linux
perf | Exposes perf based metrics (Warning: Metrics are dependent on kernel configuration and settings). | Linux

### Textfile Collector

The textfile collector is similar to the [Pushgateway](https://github.com/prometheus/pushgateway),
in that it allows exporting of statistics from batch jobs. It can also be used
to export static metrics, such as what role a machine has. The Pushgateway
should be used for service-level metrics. The textfile module is for metrics
that are tied to a machine.

To use it, set the `--collector.textfile.directory` flag on the Node exporter. The
collector will parse all files in that directory matching the glob `*.prom`
using the [text
format](http://prometheus.io/docs/instrumenting/exposition_formats/). **Note:** Timestamps are not supported.

To atomically push completion time for a cron job:
```
echo my_batch_job_completion_time $(date +%s) > /path/to/directory/my_batch_job.prom.$$
mv /path/to/directory/my_batch_job.prom.$$ /path/to/directory/my_batch_job.prom
```

To statically set roles for a machine using labels:
```
echo 'role{role="application_server"} 1' > /path/to/directory/role.prom.$$
mv /path/to/directory/role.prom.$$ /path/to/directory/role.prom
```

### Filtering enabled collectors

The `node_exporter` will expose all metrics from enabled collectors by default.  This is the recommended way to collect metrics to avoid errors when comparing metrics of different families.

For advanced use the `node_exporter` can be passed an optional list of collectors to filter metrics. The `collect[]` parameter may be used multiple times.  In Prometheus configuration you can use this syntax under the [scrape config](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#<scrape_config>).

```
  params:
    collect[]:
      - foo
      - bar
```

This can be useful for having different Prometheus servers collect specific metrics from nodes.

## Building and running

Prerequisites:

* [Go compiler](https://golang.org/dl/)
* RHEL/CentOS: `glibc-static` package.

Building:

    go get github.com/prometheus/node_exporter
    cd ${GOPATH-$HOME/go}/src/github.com/prometheus/node_exporter
    make
    ./node_exporter <flags>

To see all available configuration flags:

    ./node_exporter -h

## Running tests

    make test


## Using Docker
The `node_exporter` is designed to monitor the host system. It's not recommended
to deploy it as a Docker container because it requires access to the host system.
Be aware that any non-root mount points you want to monitor will need to be bind-mounted
into the container.
If you start container for host monitoring, specify `path.rootfs` argument.
This argument must match path in bind-mount of host root. The node\_exporter will use
`path.rootfs` as prefix to access host filesystem.

```bash
docker run -d \
  --net="host" \
  --pid="host" \
  -v "/:/host:ro,rslave" \
  quay.io/prometheus/node-exporter \
  --path.rootfs /host
```

On some systems, the `timex` collector requires an additional Docker flag,
`--cap-add=SYS_TIME`, in order to access the required syscalls.

## Using a third-party repository for RHEL/CentOS/Fedora

There is a [community-supplied COPR repository](https://copr.fedorainfracloud.org/coprs/ibotty/prometheus-exporters/) which closely follows upstream releases.

[travis]: https://travis-ci.org/prometheus/node_exporter
[hub]: https://hub.docker.com/r/prom/node-exporter/
[circleci]: https://circleci.com/gh/prometheus/node_exporter
[quay]: https://quay.io/repository/prometheus/node-exporter
[goreportcard]: https://goreportcard.com/report/github.com/prometheus/node_exporter

# Example config for node_exporter work with prometheus(v2.5)
由于涉及到阿里云的api授权，需要提供accessKeyId和 accessSecret，所以本项目不提供二进制文件，大家将
rds_common 里的配置
```
const (
	regionId        = "cn-hangzhou"
	accessKeyId     = "accessKeyId"
	accessKeySecret = "accessKeySecret"
)
```
修改好后，自己进行编译即可。

Prometheus配置：
```
  - job_name: db-mysql-node 
    scrape_interval: 30s
    params:
        "collect[]": ["cpuRds", "meminfoRds", "diskRdsStats", "iopsRds"]

    consul_sd_configs:
      - server: 'consul-address'
        services: ["db-mysql"]
        tag: "hz-ali"
        tag_separator: "|"

    relabel_configs:

      - source_labels: [__meta_consul_service_id]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: exporter-address:9100
```

node_exporter配置：
```
./node_exporter \
--no-collector.arp \
--no-collector.bcache \
--no-collector.bonding \
--no-collector.conntrack \
--collector.cpu \
--collector.cpuRds \
--no-collector.cpufreq \
--collector.diskstats \
--no-collector.edac \
--no-collector.entropy \
--no-collector.filefd \
--no-collector.filesystem \
--no-collector.hwmon \
--no-collector.infiniband \
--no-collector.ipvs \
--no-collector.loadavg \
--no-collector.mdadm \
--collector.meminfo \
--collector.meminfoRds \
--no-collector.netclass \
--no-collector.netdev \
--no-collector.netstat \
--no-collector.nfs \
--no-collector.nfsd \
--no-collector.pressure \
--no-collector.sockstat \
--no-collector.stat \
--no-collector.textfile \
--no-collector.time \
--no-collector.timex \
--no-collector.uname \
--no-collector.vmstat \
--no-collector.xfs \
--no-collector.zfs \
--collector.iopsRds \
--web.max-requests=0 \
--log.level="info"
```

更多详细内容，参考本人的相关文档：
简书：[rds数据库本地化改造](https://www.jianshu.com/p/c38f5d039bd1) <br>
点我达博客：[rds数据库本地化改造](http://tech.dianwoda.com/2020/01/06/dian-wo-da-rdsjian-kong-xi-tong-gai-zao/) <br>
微信公众号：[rds数据库本地化改造](https://mp.weixin.qq.com/s?__biz=MzAwNDU1Njc0OA==&mid=2247483652&idx=1&sn=b1573a3d08bb53fb8df3a6b272690d82&chksm=9b2b53a7ac5cdab13dd54033f60393fdbe8e1a73dacfa63bc2bfa8f2711f03bd31b5bb0c9a20&token=1062910638&lang=zh_CN#rd)

