# gRPC example

本目录是google gRPC示例目录, 演示Go中的gRPC如何使用
## 使用步骤
- 安装依赖
    ```bash
    go get google.golang.org/grpc
    # 使用mod进行包管理
    export GO111MODULE=on  # Enable module mode
    go get github.com/golang/protobuf/protoc-gen-go
    ```
- 代码运行, 代码[目录](helloworld), proto文件生成[build.bat](helloworld/helloworld/build.bat)
    ```bash
    go run greeter_server/main.go
    go run greeter_client/main.go
    ```
- mock 测试, [本项目mock示例](../../../../project/mock/README.md)
    - simple [mytest](mytest/proto/generate_mock.go) generate file, 步骤说明
        - 先定义[proto文件](mytest/proto/test.proto)
        - 根据[命令](mytest/proto/build.bat)生成rpc code 
        - 使用:
            - [服务端](mytest/server/main.go)样例代码
            - [客户端](mytest/client/main.go)样例代码
        - 测试:
            - [mock生成](mytest/proto/generate_mock.go), 生成打桩测试代码
            - 使用mock生成的code进行打桩测试, [测试文件](mytest/my_test.go)
    - gRPC example [helloworld](helloworld/helloworld/generate_mock.go) generate file


## 官方文档
- [gRPC Basics: Go](https://grpc.io/docs/languages/go/basics/)

## example
- [google example](helloworld)


## gRPC相关
- [go gRPC安装及基本使用可参考](https://blog.csdn.net/fwhezfwhez/article/details/90475510?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-1.edu_weight)
- [gRPC Document](https://www.grpc.io/docs/)