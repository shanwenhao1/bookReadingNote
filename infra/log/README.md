# log 使用说明

- 使用前请参考[log4go.xml](../../config/log4go.xml)配置文件, 使用[log.go](log.go)
中的`InitializedLog4go`函数进行初始化日志操作
```go
// 一般在main.go中使用 
import (
 "scanPen/infra/log"
)
    

// 日志初始化
log.InitializedLog4go("config/log4go.xml")
```
- 使用推荐使用`Tag`函数, 为日志添加标志
```go
import (
 "scanPen/infra/log"
)

log.Tag(InfoLog, InitSer, "日志框架初始化完成")
```
