# 服务发现
主要有两种实现方式:
- 服务及客户端直接与服务注册表交互
- 通过部署基础设施来处理服务发现

第一种方式以`应用层服务发现模式`为例: Eureka
- 优点: 可以处理多平台部署问题(服务发现机制与具体的部署平台无关)
- 缺点: 
    - 需要为你使用的每种编程语言(可能还有框架)提供服务发现库
    - 开发者负责设置和管理服务注册表, 分散精力(最好使用部署基础设施提供的服务发现机制)

第二种方式以`平台层服务发现模式`为例: Docker、Kubernetes: 部署平台为每个服务提供
DNS名称、虚拟IP(VIP)地址和解析为VIP地址的DNS名称, 客户端向DNS名称和VIP发出请求, 部署
平台自动将请求路由到其中一个可用服务实例(服务注册、服务发现和请求路由完全由部署平台处理)


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