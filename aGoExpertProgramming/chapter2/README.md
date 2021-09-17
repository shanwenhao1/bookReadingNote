# 常见控制结构体实现原理

## defer

```go
type	_defer	struct	{
	sp              uintptr			//函数栈指针
	pc			    uintptr			//程序计数器
	fn			    *funcval		//函数地址
	link		    *_defer			//指向自身结构的指针，用于链接多个defer
}
```
defer语句用于延迟函数的调用, 每次defer都会把一个函数压入栈中, 函数返回前再把延迟的函数取出并执行.

defer行为规则, [样例代码](code/defer_test.go):
* 延迟函数的参数在defer语句出现时就已经确定(并不会因为后续主函数对参数的修改而影响)
* 延迟函数执行按照后进先出顺序执行(类似于栈操作)
* 延迟函数可能操作主函数的具名返回值


## select

```go
type	scase	struct	{
	c					*hchan			//	chan. 当前case语句所操作的channel指针
	kind				uint16          // case语句的类型, 分为caseRecv(读channel)、caseSend(写channel)、caseDefault(default)
    elem				unsafe.Pointer	//	data	element. 缓冲区地址. 
                                        //          scase.kind	==	caseRecv	：	scase.elem表示读出channel的数据存放地址    
                                        //          scase.kind	==	caseSend	：	scase.elem表示将要写入channel的数据存放地址    
}
```
select是Golang在语言层面提供的[多路IO复用](https://zhuanlan.zhihu.com/p/115220699) 的机制, 用于检测channel
是否ready
* select语句除default外, 每个case操作一个channel, 要么读要么写
* select语句除default外, 各case执行顺序是随机的
* select语句中如果没有default语句, 则会阻塞等待任一case
* select语句中读操作要判断 是否成功读取, 关闭的channel也可以读取
    - 对于读channel的case来说， 如果`case elem, ok := <-chan1:`, 如果channel有可能被其他协程关闭的情况下, 
    一定要检测读取是否成功, 因为close的channel也有可能返回, 此时`ok == false`.


## range

* 使用index、value接收range返回值会发生一次数据拷贝
    - 遍历过程中可以视情况放弃接收index或value, 可以一定程度上提升性能
* for-range的实现实际上是C风格的for
* 遍历channel时, 如果channel中没有数据, 可能会阻塞
* 尽量避免遍历过程中修改原数据(遍历的循环次数在开始时确定)


## mutex

```go
type	Mutex	struct	{
    state	            int32       // 表示互斥锁的状态, 比如是否被锁定等
    sema		        uint32      // 信号量, 协程阻塞等待该信号量, 解锁的协程释放信号量从而唤醒等待信号量的协程
}
```
![](picture/mutex.jpg)
* Locked: 表示该Mutex是否已被锁定，0：没有锁定	1：已被锁定
* Woken: 表示是否有协程已被唤醒，0：没有协程唤醒	1：已有协程唤醒，正在加锁过程中
* Starving：表示该Mutex是否处理饥饿状态，	0：没有饥饿	1：饥饿状态，说明有协程阻塞了超过1ms. 处于饥饿状态下
不会进入自旋模式(加锁时, 当前Locked为1, 尝试加锁的协程不会马上转成阻塞态, 而是会探测一段时间, 看Locked是否变为0, 变为0
则加锁, 不变则转入阻塞, 减少协程切换的消耗)
* Waiter: 表示阻塞等待锁的协程个数，协程解锁时根据此值来判断是否需要释放信号量
协程之间抢锁实际上是抢给Locked赋值的权利，能给Locked域置1，就说明抢锁成功。抢不到的话就阻塞等待Mutex.sema信号量，
一旦持有锁的协程解锁，等待的协程会依次被唤醒

Mutext对外提供两个方法，实际上也只有这两个方法：
* Lock(): 加锁方法
* Unlock():	解锁方法

编程技巧:
* 加锁后立即使用defer对其解锁, 可以有效的避免死锁
* 加锁和解锁最好出现在同一层次的代码块中(成对出现)
* 重复解锁会引起panic, 应避免这种操作的可能性


## RWMutex

```go
type	RWMutex	struct	{
	w                   Mutex		//用于控制多个写锁，获得写锁首先要获取该锁，如果有一个写锁在进行，那么再到来的写锁将会阻塞于此
	writerSem			uint32	    //写阻塞等待的信号量，最后一个读者释放锁时会释放信号量
	readerSem			uint32	    //读阻塞的协程等待的信号量，持有写锁的协程释放锁后会释放信号量
	readerCount	        int32		//记录读者个数
	readerWait		    int32		//记录写阻塞时读者个数
}
```

RWMutex提供4个简单的接口来提供服务：
* RLock()：读锁定
* RUnlock()：解除读锁定
* Lock(): 写锁定，与Mutex完全一致
* Unlock()：解除写锁定，与Mutex完全一致

写锁:
* 写锁定操作需要做两件事:
    - 获取互斥锁
    - 阻塞等待所有读操作结束(如果有的话)
* 解除写锁定要做两件事:
    - 唤醒因读锁定而被阻塞的协程(如果有的话)
    - 解除互斥锁

读锁:
* 读锁定需要做两件事:
    - 增加读操作计数, 即readerCount++
    - 阻塞等待写操作结束(如果有的话)
* 解除读锁定需要做两件事:
    - 减少读操作计数, 即readerCount--
    - 唤醒等待写操作的协程(如果有的话), 注意: 只有最后一个释放掉的读操作协程才可以释放信号量唤醒写协程
