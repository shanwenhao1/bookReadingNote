# 并发控制

Go并发中三种控制并发的方案:

## Channel
- 优点: 实现简单，清晰易懂
- 缺点: 当需要大量创建协程时就需要有相同数量的channel，而且对于子协程继续派生出来的协程不方便控制

channel一般用于协程之间的通信，channel也可以用于并发控制。比如主协程启动N个子协程，
主协程等待所有子协程退出后再继续后续流程. [示例](code/channel.go)

## WaitGroup
子协程个数动态可调整

WaitGroup，可理解为Wait-Goroutine-Group，即等待一组goroutine结束. 使用信号量实现:
* 信号量可以理解为一个数值
    - 当信号量>0时，表示资源可用，获取信号量时系统自动将信号量减1
    - 当信号量==0时，表示资源暂不可用，获取信号量时，当前线程会进入睡眠，当信号量为正时被唤醒

```go
type WaitGroup struct {
    state1 [3]uint32           // 包含了state(state实际上是两个计数器)和一个信号量
                                /*
                                    counter： 当前还未执行结束的goroutine计数器
                                    waiter count: 等待goroutine-group结束的goroutine数量，即有多少个等候者s
                                    semaphore: 信号量
                                */

}
```

WaitGroup对外提供三个接口:
* Add(delta int):
    - 将delta值加到counter中
    - 当counter值变为0时，跟据waiter数值释放等量的信号量，把等待的goroutine全部唤醒
* Wait()： waiter递增1，并阻塞等待信号量semaphore
    - 累加waiter
    - 阻塞等待信号量
* Done()： counter递减1，按照waiter数值释放相应次数信号量

## Context
对子协程派生出来的孙子协程的控制. 它可以控制一组呈树状结构的goroutine, 每个goroutine拥有相同的上下文

context实际上只定义了接口
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)   // 该方法返回一个deadline和标识是否已设置deadline的bool值，
                                                    // 如果没有设置deadline，则ok == false，此时deadline为一个初始值的time.Time值
    Done() <-chan struct{}              // 该方法返回一个channel，需要在select-case语句中使用，如”case <-context.Done():”
    Err() error                       // 该方法描述context关闭的原因。关闭原因由context实现控制，不需要用户设置
    Value(key interface{}) interface{}  // 用于在树状分布的goroutine间传递信息， 根据key值查询map中的value
}
```

### 空context
使用`context.Background()`可获取到context自定义的`background`(类型为: `emptyCtx`). context提供了四个方法创建不同类型的context
(如果没有父context, 则需要传入background)
* WithCancel()
* WithDeadline()
* WithTimeout()
* WithValue()

[示例代码](code/context.go)


### cancelCtx
```go
type cancelCtx struct {
     Context
        mu          sync.Mutex              // protects following fields
        done        chan struct{}           // created lazily, closed by first cancel call
        children    map[canceler]struct{}   // set to nil by the first cancel call
        err         error                   // set to non-nil by the first cancel call
}
```

[样例代码](code/contextCancel.go)


### timerCtx
在cancelCtx基础上增加了deadline用于标示自动cancel的最终时间，而timer就是一个触发自动cancel的定时器.
主要有两个方法:
* WithDeadline()
* WithTimeout(): WithTimeout()实际调用了WithDeadline，二者实现原理一致

[样例代码](code/contextTimeD.go)


### valueCtx

```go
type valueCtx struct {
    Context
    key, val interface{}
}
```
valueCtx不要实现cancel, 专注于在协程中传递value(键值对的方式)

