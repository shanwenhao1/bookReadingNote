// --proto_path=. --proto_path=../../../../rpcDepend 自定义proto文件依赖路径,这样就可以将特定的proto依赖加入编译命令中

//go:generate protoc --proto_path=. --proto_path=../../../../rpcDepend --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. hello.proto
package helloProto
