# httptest

httptest可以方便的模拟各种Web服务器和客户端，以达到测试目的

## Go提供的两种定时器

### 一次性定时器
```go
/*
     Timer代表一次定时，时间到来后仅发生一个事件
        Timer实质上是把一个定时任务交给专门的协程(每个go应用程序都有该协程专门负责管理所有的Timer)进行监控
        
*/
type Timer struct { 
    C <-chan Time               // 管道，上层应用跟据此管道接收事件
    r runtimeTimer              // runtime定时器，该定时器即系统管理的定时器，对上层应用不可见
}


type runtimeTimer struct {
 tb      uintptr                     // 系统底层存储runtimeTimer的数组地址
 i       int                         // 当前runtimeTimer在tb数组中的下标

 when    int64                       // 当前定时器触发时间
    period  int64                       // 当前定时器周期触发间隔
    f       func(interface{}, uintptr)  // 定时器触发时执行的回调函数，回调函数接收两个参数
    arg     interface{}                 // 定时器触发时执行函数传递的参数一
    seq     uintptr                     // 定时器触发时执行函数传递的参数二(该参数只在网络收发场景下使用)10. }
```
[Timer使用示例](code/timerUse.go)

timer的几个方法：
* `func NewTimer(d Duration) *Timer`: 创建定时器
* `func (t *Timer) Stop() bool`: 停止定时器
    - true: 定时器超时前停止，后续不会再有事件发送
    - false: 定时器超时后停止
* `func (t *Timer) Reset(d Duration) bool`: 重置定时器. 实际动作是先停掉定时器，再启动(返回值为停掉定时器的返回)

#### 一次性匿名定时器
`time.After`: 创建一个定时器并返回定时器的管道, [使用示例](../chapter5/code/channel.go)



### 周期性定时器

`Ticker`是周期性定时器, 即周期性的触发一个事件, 通过Ticker本身提供的管道将事件传递出去
```go
type Ticker struct {
    C <-chan Time
    r   runtimeTimer
}
```
[使用示例](code/tickerUse.go)

Ticker：
* `func NewTicker(d Duration) *Ticker`: 创建周期性定时器
* `func (t *Ticker) Stop()`: 停止定时器