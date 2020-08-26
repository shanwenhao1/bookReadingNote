// this is grpc server function
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	pb "bookReadingNote/microservicesPatterns/code/chapter3/grpcExample/mytest/proto/test"
)

const (
	port = ":6001"
)

type HelloService struct {
}

func (hs HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("你好，%s", in.Username)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	fmt.Println("-------rpc server run on port: ", port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, HelloService{})
	s.Serve(lis)
}
