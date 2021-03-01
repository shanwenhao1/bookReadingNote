# 阅读书籍笔记整理

本项目使用go mod进行管理.
- `go mod init bookReadingNote`
- 下载后的依赖存放在`$GOPATH/pkg/mod`

## [Go语言高级编程](advanceGoProgram/README.md)

## [Microservices Patterns](microservicesPatterns/README.md)

## 项目build组件

### 服务发现
基于[kubernetes](Kubernetes/README.md)的[Consul服务发现](microservicesPatterns/doc/chapter3/Service-discovery.md)
+[traefik反向代理](Kubernetes/kubernetes-plugin/ingress/README.md#Traefik介绍)


### 日志收集
采用`EFK`进行[日志收集](doc/log/README.md)

### Go开发
[Go 相关依赖包索引](https://godoc.org/)

#### [其他](doc/other.md) 包含中间件等

#### DDD
##### [基础设施](infra/README.md)

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

### 负载均衡
[目前简单的编写](advanceGoProgram/chapter6/README.md#负载均衡)


## [Go刷题练习](practice/README.md)

## [运维](doc/DevOps/README.md)

## 资源搜索
- [allitebooks](http://www.allitebooks.org/)
- [免费编程中文书籍索引大全](https://github.com/justjavac/free-programming-books-zh_CN)

## 开源项目研究
- [区块链](blockChain/README.md)

#### [知乎网友推荐检索](https://zhuanlan.zhihu.com/p/23857699)

Go 1.11以后新增了[GO111MODULE](https://learnku.com/go/t/39086) 用模块管理, 淘汰GoPath
- [Goland Module 初始化](https://www.cnblogs.com/bbllw/p/12377155.html)
- [Go Modules 详解使用](https://learnku.com/articles/27401)
```bash
# go module 连接不到外网问题, 按照如下命令使用代理
go env -w GOSUMDB=off
go env -w GOPROXY=https://goproxy.cn,direct
# https://blog.csdn.net/suoyudong/article/details/107922682
```