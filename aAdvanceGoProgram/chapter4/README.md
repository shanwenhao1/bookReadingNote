# RPC和Protobuf

## proto安装
protoc编译器不支持Go语言, 需要安装官方的
[protoc工具](https://github.com/protocolbuffers/protobuf/releases)
- 将解压后的bin目录下的`protoc.exe`放入`$GOPATH/bin`目录下(也可使用环境变量的方式)
- 使用命令安装`protoc-gen-go`
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

**注意: 本项目中使用的一些依赖rpc放置于[rpcDepend](../../rpcDepend/README.md)**

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

示例: [hello.proto](./protobuf/helloProto/hello.proto)定义`HelloService`
```proto
// 生成Go代码
// protoc --go_out=.hello.proto
syntax = "proto3";

package helloProto;
message String{
    string value = 1;
}

// protoc --go_out=plugins=grpc:. hello.proto
service HelloService{
    rpc Hello (String) returns (String);
}
```
**这里我们使用了方便的go generate[自定义generate.go文件](./protobuf/helloProto/generate.go)生成方式.**

也可用命令行执行

```
// --proto_path=. --proto_path=../../../../rpcDepend 自定义proto文件依赖路径,这样就可以将特定的proto依赖加入编译命令中
protoc --proto_path=. --proto_path=../../../../rpcDepend --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. hello.proto
```

使用依赖生成proto文件的示例: [testProto](./protobuf/testProto/test.proto)