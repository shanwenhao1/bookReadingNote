package rpcSafe

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 明确服务的名字和接口, 防止客户端弄错RPC方法或者参数类型不匹配等低级错误
const HelloServiceName = "path/to/HelloService" // rpc服务名字中增加了包路径前缀, 并非完全等价于Go语言路径

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func HelloSer() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)
	}()
}

/*
	Json编码实现的RPC服务
*/
func HelloJsonRpc() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}()
}

/*
	Http协议上提供jsonrpc服务
	使用命令, curl localhost:1234/jsonrpc -X POST --data `{"method":"HelloService.Hello","params":["hello"],"id":0}`
*/
func HelloHttpJsonRpc() {
	RegisterHelloService(new(HelloService))

	go func() {
		http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
			var conn io.ReadWriteCloser = struct {
				io.Writer
				io.ReadCloser
			}{
				ReadCloser: r.Body,
				Writer:     w,
			}
			rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		})
		http.ListenAndServe(":1234", nil)
	}()
}
