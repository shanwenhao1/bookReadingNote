# Go语言高级编程阅读笔记

[官方在线书籍](https://chai2010.gitbooks.io/advanced-go-programming-book/)

## [cgo](chapter2/README.md)
## [Go RPC](chapter4/README.md)
## [Go and Web](chapter5/README.md)
## [分布式系统](chapter6/README.md)

## 附录

### Go语言常见坑
- 可变参数是空接口类型
- 数组是值传递
- map遍历顺序不固定
- recover必须在defer函数中运行
- 独占CPU导致其它Goroutine饿死
    ```go
    func main() {
        runtime.GOMAXPROCS(1)
    
        go func() {
            for i := 0; i < 10; i++ {
                fmt.Println(i)
            }
        }()
    
        for {} // 占用CPU
    }
  
  
  
    // 解决方案
  func main() {
      runtime.GOMAXPROCS(1)
  
      go func() {
          for i := 0; i < 10; i++ {
              fmt.Println(i)
          }
      }()
  
      for {
          runtime.Gosched()       // 调度函数 
      }
  }
    ```