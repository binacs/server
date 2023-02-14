# prometheus

## 1. 概述

metric 统计组件；监控告警解决方案。

## 2. 关键概念

### 2.1 存储 `时序数据` ，其时序数据格式为：

```xml
<metric name>{<label name>=<label value>, ...}
```

### 2.2 Metric

几种 Metric 类型：

- `Counter`: 一种累加的`metric`，如请求的个数，结束的任务数，出现的错误数等
- `Gauge`: 常规的metric,如温度，可任意加减。其为瞬时的，与时间没有关系的，可以任意变化的数据。
- `Histogram`: 柱状图，用于观察结果采样，分组及统计，如：请求持续时间，响应大小。其主要用于表示一段时间内对数据的采样，并能够对其指定区间及总数进行统计。**根据统计区间计算**
- `Summary`: 类似`Histogram`，用于表示一段时间内数据采样结果，其直接存储quantile数据，而不是根据统计区间计算出来的。**不需要计算，直接存储结果**。在项目中往往用于统计请求相关度量。

### 2.3 PromQL

PromQL (Prometheus Query Language) 是 Prometheus 自己开发的数据查询 DSL 语言。

查询结果类型：

- 瞬时数据 (Instant vector): 包含一组时序，每个时序只有一个点，例如：`http_requests_total`
- 区间数据 (Range vector): 包含一组时序，每个时序有多个点，例如：`http_requests_total[5m]`
- 纯量数据 (Scalar): 纯量只有一个数字，没有时序，例如：`count(http_requests_total)`

查询语句请参阅官方文档。

### 2.4 prometheus采集策略

只能抓取指定 `target` 暴露的数据，即对 `target` 开放端口有要求；

对于不对外暴露固定端口的 `target` ，需采取上报至 `push gateway` ，再由 prometheus 主动抓取 `push gateway` 暴露端口的办法实现数据采集。



## 3. 使用示例

常规配置即可，每个项的含义参考官方文档。

**注意：**

**采用 `push gateway` 解决方案时，需指定 `push gateway` 的 `honor_labels: true` 。**

*如果 `honor_labels` 设置为 `"true"` ，则通过保留已抓取数据的标签值并忽略冲突的服务器端标签来解决标签冲突。*



## 4. 配套组件

`NodeExporter` ：提供 http 路由以供 prometheus 抓取其在单个主机上采集得到的各类数据。

`PushGateway` ：采用上报方案时各个节点上报的目标，也是 prometheus 抓取的目标。

`AlertManager` ：接收 prometheus 依据规则产生的告警信息，可以转发至 `webhook` 以及直接触发邮件等告警机制。**由于 `alert manager` 的告警渠道支持并不完善，在需要同时支持多个告警渠道的场景下建议写一个新的组件作为 `webhook` ，由新组件触发各场景下的告警消息推送。如需要同时支持邮件、微信、钉钉、消息队列等渠道时。**

在本项目中，prometheus 采集的数据均来自push gateway，而后者的数据均由物理节点运行的 [ProPush](https://github.com/binacs/ProPush) 主动上报。

## 5. Advanced

关于[如何估算 Prometheus 所消耗的资源](https://www.robustperception.io/how-much-ram-does-prometheus-2-x-need-for-cardinality-and-ingestion)
