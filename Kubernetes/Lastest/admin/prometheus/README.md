# Prometheus

[Prometheus](https://prometheus.io/docs/introduction/overview/) 相比于其他传统监控工具主要有以下几个特点:
- 具有由 metric 名称和键/值对标识的时间序列数据的多维数据模型
- 有一个灵活的查询语言
- 不依赖分布式存储，只和本地磁盘有关
- 通过 HTTP 的服务拉取时间序列数据
- 也支持推送的方式来添加时间序列数据
- 还支持通过服务发现或静态配置发现目标
- 多种图形和仪表板支持

## 组件
Prometheus 由多个组件组成，但是其中许多组件是可选的：

- Prometheus Server：用于抓取指标、存储时间序列数据
- exporter：暴露指标让任务来抓
- pushgateway：push 的方式将指标数据推送到该网关
- alertmanager：处理报警的报警组件
- adhoc：用于数据查询

## 架构
![](picture/architecture.png)

## 安装

[安装步骤](install.md)


## 参考
- [k8s 集群手动安装prometheus](https://www.qikqiak.com/k8s-book/docs/52.Prometheus%E5%9F%BA%E6%9C%AC%E4%BD%BF%E7%94%A8.html)