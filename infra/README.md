# DDD: infra基础设施层

- DDD: infrastructure 基础设施层, 最底层(可以与所有层进行交互)
向其他层提供通用的基础的功能, 比如工具类, 第三方库类支持, 基本常用配置, 数据库和缓存底层实现等
包含以下内容:
    - 为应用层: 传递消息(通知等)
    - 为领域层: 提供持久化机制(最底层实现)
    - 为用户界面层: 提供组件配置
    - 基础设施层还能通过架构框架来支持四个层次之间的交互模式
    
## Go基础设施

### 插件相关
- Hash: 高性能[MurmurHash](../aAdvanceGoProgram/chapter5/hashFunc/hashFunc.go)

### 数据库
- [go 自带sql Base](../aAdvanceGoProgram/chapter5/database/sqlBase.go)
- [gorm](db/README.md)

### Redis
[redis](redis/README.md)部署及使用
