# 生成proto序列化, 反序列化代码的文件
protoc --go_out=. ./helloworld.proto
# 生成服务器和客户端通讯、实现的公共代码(如果写客户端和服务端的通信, 用第二个编译方式, 如果只是作为序列化和反序列化的工具, 第一个命令就可以)
protoc --go_out=plugins=grpc:. ./helloworld.proto

# 或者
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld