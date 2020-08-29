# 断路器和熔断

## 断路器

[断路器](https://microservices.io/patterns/reliability/circuit-breaker.html)
是分布式系统中处理服务故障的一种模式: 在一个远程调用的代理, 在该过程中, 当失败次数超过指定阈值后的
一段时间内, 这个代理会拒绝其他调用.

故障发生时的断路处理方案:
- 服务只是向其客户端返回错误
- 返回备用值(使用API组合提供服务)

### Go Circuit

[sony/gobreaker](https://github.com/sony/gobreaker)
```bash
go get github.com/sony/gobreaker
```

有三个状态:
- `Closed`: 指容断器放行所有请求
-  `Open`: 达到一定数量的错误计数, 进入Open状态, 指容断发生, 下游出现错误, 不能再放行请求.
- `Half-Open`: 经过一段Interval时间后, 自动进入Half-Open状态, 然后开始尝试对接成功请求计数. 进入Half-Open后,
根据成功/失败计数(一段时间)情况, 会自动进入`Closed`或`Open`.


## 熔断

[afex/hystrix-go](https://github.com/afex/hystrix-go)