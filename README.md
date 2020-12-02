# 阅读书籍笔记整理

本项目使用go mod进行管理.
- `go mod init bookReadingNote`
- 下载后的依赖存放在`$GOPATH/pkg/mod`

## [Go语言高级编程](advanceGoProgram/README.md)

## [Microservices Patterns](microservicesPatterns/README.md)

## 资源搜索
- [allitebooks](http://www.allitebooks.org/)
- [免费编程中文书籍索引大全](https://github.com/justjavac/free-programming-books-zh_CN)

## 项目build组件

### 服务发现
基于[kubernetes](Kubernetes/README.md)的[Consul服务发现](microservicesPatterns/doc/chapter3/Service-discovery.md)
+[traefik反向代理](Kubernetes/kubernetes-plugin/ingress/README.md#Traefik介绍)

### Go开发

#### [基础设施](doc/infrastructure.md)

#### [数据传输方案](project/dataTransmission/README.md)

#### 通信
- [微服务架构中的进程通信](microservicesPatterns/doc/chapter3/README.md)
- [grpc](microservicesPatterns/code/chapter3/grpcExample/README.md)

服务请求限流([令牌捅](advanceGoProgram/chapter5/tokenLimit/tokenBase.go)), [断路器和熔断](project/CircuitAndHystrix/README.md): 用于处理服务调用故障

#### Go测试
- [go mock](project/mock/README.md)

### CI/CD
[说明文档](project/ci/README.md)

### 测试

#### 压力测试
[Locust](https://locust.io/)


## 开源项目研究
- [区块链](blockChain/README.md)

#### [知乎网友推荐检索](https://zhuanlan.zhihu.com/p/23857699)

Go 1.11以后新增了[GO111MODULE](https://learnku.com/go/t/39086)用模块管理, 淘汰GoPath
- [Go Modules 详解使用](https://learnku.com/articles/27401)
- [Goland使用Module](https://www.cnblogs.com/bbllw/p/12377155.html)