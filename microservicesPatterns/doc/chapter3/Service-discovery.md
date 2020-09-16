# 服务发现

## consul

Consul是一套开源的分布式服务发现和配置管理系统.
- go语言开发的服务发现、配置管理中心服务
- 基于raft协议
- 内置了服务注册与发现框架、分布式一致性协议实现、健康检查、Key/Value存储、多数据中心方案


### 安装

#### Kubernetes上的安装
[官方文档](https://www.consul.io/docs/k8s/installation/overview)

使用`Helm 3`进行默认安装
```bash
helm repo add hashicorp https://helm.releases.hashicorp.com
helm search repo hashicorp/consul
// use deafault configuration, 可使用命令helm inspect values hashicorp/consul查看consul配置
helm install consul hashicorp/consul --set global.name=consul
```