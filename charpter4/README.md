# RPC和Protobuf

## RPC
RPC(Remote Procedure Call)是远程过程调用的简称, 是分布式系统中不同节点间流行的通信方式.

Go的RPC包路径为`net/rpc`(默认采用Go语言特有的Gob编码协议, 因此其他语言难以调用Go语言实现的RPC服务, 
如需跨语言调用, 请查看[Go RPC的跨语言支持](#Go RPC的跨语言支持)), Go语言RPC规则
- 方法只能有两个可序列化的参数, 其中第二个参数是指针类型, 并且返回一个error类型, 同时必须是公开的方法(首字母大写)


Go RPC Service构建步骤:
- 明确服务的名字和接口

### Go RPC的跨语言支持
Go的RPC框架特色设计:
- RPC数据打包时可以通过插件实现自定义的编码和解码
- RPC建立在抽象的io.ReadWriteCloser接口上, 我们可以将RPC架设在不同的通信协议上

Go如何跨语言调用RPC:
- 通过官方自带的`net/rpc/jsonrpc`扩展实现

## Protobuf 
protobuf是谷歌公司开发的一种数据描述语言. 基本数据单元是`message`

protoc编译器不支持Go语言, 需要安装官方的[protoc工具](https://github.com/protocolbuffers/protobuf/releases), 
- 将解压后的bin目录下的`protoc.exe`放入`$GOPATH/bin`目录下(也可使用环境变量的方式)
- 再安装代码生成插件`go get github.com/golang/protobuf/protoc-gen-go`
    ```proto
    // 生成Go代码
    protoc --go_out=. hello.proto
    ```