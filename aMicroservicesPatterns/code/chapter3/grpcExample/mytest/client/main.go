// this is grpc client function
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"

	pb "bookReadingNote/aMicroservicesPatterns/code/chapter3/grpcExample/mytest/proto/test"
)

func main() {
	conn, err := grpc.Dial("localhost:6001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Username: "ft"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(r.Message)
}
