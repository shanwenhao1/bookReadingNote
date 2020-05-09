# Go和Web

Go的Web框架大致可以分为两类
- Router框架: Go开源界应用最广泛的路由器是`httprouter`
    - gin
- MVC类框架;
    - beego(借鉴其他语言编程风格)
    
## httproute

- REST API设计风格 
- 支持`*`通配符, 例如: `/src/*` (通常用来做简单的HTTP静态文件服务器)
- 支持特殊情况下的回调函数支持
    ```go
    // 例如404
    r := httprouter.New()
    r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("oh no, not found"))
    })
  
    // 内部panic
    r.PanicHandler = func(w http.ResponseWriter, r *http.Request, c interface{}) {
      log.Printf("Recovering from panic, Reason: %#v", c.(error))
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(c.(error).Error()))
    }
    ```
- 数据结构使用压缩字典树(Radix Tree)

## 中间件

使用中间件剥离非业务逻辑, [Gin middleware](https://github.com/gin-gonic/contrib)

## database
- `database/sql`包提供的操作数据库接口和规范, 配合Mysql驱动
- 对象关系映射(Object Relational Mapping)
- SQL Builder

## 服务流量限制
- 磁盘IO瓶颈
- CPU计算瓶颈
- 网络带宽瓶颈

常用流量限制手段:
- 漏桶: 每隔固定的时间向外面漏一滴水, 用户获取到就可以继续服务请求, 否则等待.
- 令牌桶: 匀速向桶中添加令牌, 用户拿到令牌才可以继续服务请求(支持一定程度的并发), 
没有令牌的情况下会退化为漏桶. 
    - 令牌桶模型实际上就是对全局计数的加减法操作过程, 可以用buffered channel来完成简单的
    加令牌取令牌操作来代替加锁操作.

[代码示例](tokenLimit/tokenBase.go)

## 服务瓶颈和QoS
QoS(Quality of Service), 即服务质量, 包含:
- 可用性
- 吞吐量
- 延时
- 延时变化
- 丢失指标

## 常见大型Web项目分层
[详细](https://www.jianshu.com/p/a775836c7e25?from=groupmessage)

传统MVC框架, 划为三层:
- 控制器(Controller): 负责转发请求, 对请求进行处理.
- 视图(View): 界面设计人员进行图形界面设计.
- 模型(Model): 程序员编写程序应有的功能(实现算法等等)、数据库专家进行数据管理和数据库设计(可以实现具体的功能.

前后端分离后, 后端项目只剩下M层和C层, 前后端通过ajax来交互

纯后端API模块划分方法:
- Controller: 服务入口, 负责处理路由、参数校验、请求转发
- Logic/Service: 逻辑(服务)层, 一般是业务逻辑的入口,可以认为从这里开始,所有的请求参数一定是合法的.
业务逻辑和业务流程也都在这一层中.常见的设计中会将该层称为 Business Rules.
- DAO/Repository,这一层主要负责和数据、存储打交道.将下层存储以更简单的函数、接口形式暴露给 Logic 层来使用.
负责数据的持久化工作


DDD分层架构：
- User Interface(用户界面层): 负责向用户显示信息和解释用户命令.这里指的用户可以是另一个计算机系统,不一定是使用用户界面的人
- Application(应用层): 定义软件要完成的任务,并且指挥表达领域概念的对象来解决问题.这一层所负责的工作对业务来说意义重大,
也是与其它系统的应用层进行交互的必要渠道.应用层要尽量简单,不包含业务规则或者知识,而只为下一层中的领域对象协调任务,分配工作,
使它们互相协作.它没有反映业务情况的状态,但是却可以具有另外一种状态,为用户或程序显示某个任务的进度.
- Domain(领域层或模型层): 负责表达业务概念,业务状态信息以及业务规则.尽管保存业务状态的技术细节是由基础设施层实现的,
但是反映业务情况的状态是由本层控制并且使用的.领域层是业务软件的核心,领域模型位于这一层.
- Infrastructure(基础实施层): 向其他层提供通用的技术能力：为应用层传递消息,为领域层提供持久化机制,
为用户界面层绘制屏幕组件,等等.基础设施层还能够通过架构框架来支持四个层次间的交互模式.

DCI架构(DDD的发展和补充): 用于基于面向对象的领域建模上
- User Interface(用户接口层): 主要用于处理用户发送的Restful请求和解析用户输入的配置文件等,并将信息传递给Application层的接口.
- Application(应用层): 负责多进程管理及调度、多线程管理及调度、多协程调度和维护业务实例的状态模型.当调度层收到用户接口层的请求后,
委托Context层与本次业务相关的上下文进行处理.
- Context(环境层): 以上下文为单位,将Domain层的领域对象cast成合适的role,让role交互起来完成业务逻辑.
- Domain(领域层): 定义领域模型,不仅包括领域对象及其之间关系的建模,还包括对象的角色role的显式建模.
- Infrastructure(基础实施层): 为其他层提供通用的技术能力：业务平台,编程框架,持久化机制,消息机制,第三方库的封装,通用算法,等等.


## 灰度发布和A/B测试 

灰度发布通过两种方式实现:
- 通过分批次部署实现灰度发布: 系统旧功能进行升级迭代. 例: 15台机器遵循`1-2-4-8`两倍模式发布(避免分组过多)
- 通过业务规则进行灰度发布: 新功能上线. 常见灰度发布系统规则
    - 按城市发布
    - 按概率发布
    - 按百分比发布
    - 按白名单发布
    - 按业务线发布
    - 按UA发布(App、Web、PC)
    - 按分发渠道发布

高效的hash算法:
- MurmurHash算法的性能提升, [hashFunc.go](hashFunc/hashFunc.go)
![](hashFunc/hash.png)

## 参考
- [DDD分层架构的三种模式](https://www.jianshu.com/p/a775836c7e25?from=groupmessage)