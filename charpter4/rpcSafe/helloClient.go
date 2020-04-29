package rpcSafe

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

// 约束实现接口, HelloSerClient类型必须满足HelloServiceInterface接口
var _ HelloServiceInterface = (*HelloSerClient)(nil)

type HelloSerClient struct {
	*rpc.Client
}

// 客户端方法实现
func (c *HelloSerClient) Hello(request string, reply *string) error {
	// 客户端服务调用远端程序
	return c.Client.Call(HelloServiceName+".Hello", request, reply)
}

// RPC
func DialHelloSer(network, address string) (*HelloSerClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloSerClient{Client: c}, nil
}

// Json RPC
func DialHelloJsonSer(network, address string) (*HelloSerClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloSerClient{client}, nil
}

// Go RPC使用示例
func HelloClient() {
	// 先启用HelloService
	HelloSer()
	time.Sleep(time.Second * 3)

	client, err := DialHelloSer("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

// Go Json RPC使用示例
func HelloJsonClient() {
	HelloJsonRpc()
	time.Sleep(time.Second * 3)

	client, err := DialHelloJsonSer("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

// Go http RPC使用示例
