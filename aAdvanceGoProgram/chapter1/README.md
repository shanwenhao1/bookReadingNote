# goroutine and context
- [goroutine](#goroutine)
- [context](#context)

## goroutine
  Go中通过`goroutine`启用轻量级线程来支持并发, 其以很小的栈启动(可能是2KB或4KB)并按需自动扩容.以此来实现高质量的并发.
  
  注意并发并不一定是越高越好, 有时我们也需要控制并发数量,预留资源.
  可通过`管道channel`来控制并发的数量.
    ```go
    // 示例, 控制一个最大并发数量3的并发程序
    var limit = make(chan int, 3)
    
    func main() {
        for _, w := range work {
            go func() {
                limit <- 1
                w()
                <-limit
            }()
        }
        select{}
    }
    ```
  例如`godoc`中的`vfs`(对应虚拟文件系统), 其下有一个`gatefs`子包就是控制访问该文件系统的
  最大并发数.
  ```go
    type gate chan bool

    func (g gate) enter() { g <- true }
    func (g gate) leave() { <-g }
    
    type gatefs struct {
        fs vfs.FileSystem
        gate
    }

    func (fs gatefs) Lstat(p string) (os.FileInfo, error) {
        fs.enter()
        defer fs.leave()
        return fs.fs.Lstat(p)
    }
  ```

## context
`context`是go 中的一个标准包,主要用于在`goroutine`传递上下文信息. 它可以携带截止时间、取消信号以及其他请求范围内的值.
context包提供了两种方式创建根context
- `context.Background()`: 通常用于主函数、初始化和测试代码中，作为顶层的context
- `context.TODO()`: 则用于不确定使用哪个context时
