package helloworld

import (
	"bookReadingNote/microservicesPatterns/code/chapter3/grpcExample/helloworld/helloworld/google.golang.org/grpc/examples/helloworld/helloworld"
	"bookReadingNote/microservicesPatterns/code/chapter3/grpcExample/helloworld/mock_helloworld"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
)

// rpcMsg implements the gomock.Matcher interface
type rpcMsg struct {
	msg proto.Message
}

func (r *rpcMsg) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}
	return proto.Equal(m, r.msg)
}

func (r *rpcMsg) String() string {
	return fmt.Sprintf("is %s", r.msg)
}

// test SayHello
func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHelloGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	req := &helloworld.HelloRequest{Name: "hello_test"}
	mockHelloGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&helloworld.HelloReply{Message: "Replay hello_test"}, nil)
	testSayHello(t, mockHelloGreeterClient)
}

func testSayHello(t *testing.T, client helloworld.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "hello_test"})
	if err != nil || r.Message != "Replay hello_test" {
		t.Errorf("test SayHello failed")
	}
	t.Log("Reply: ", r.Message)
}

// test SayHelloAgain
func TestSayHelloAgain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHelloGreeterClient := mock_helloworld.NewMockGreeterClient(ctrl)
	req := &helloworld.HelloAgainRequest{AgainName: "again_hello_test"}
	mockHelloGreeterClient.EXPECT().SayHelloAgain(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&helloworld.HelloAgainReply{AgainMessage: "Replay again_hello_test"}, nil)
	testSayHelloAgain(t, mockHelloGreeterClient)
}

func testSayHelloAgain(t *testing.T, client helloworld.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHelloAgain(ctx, &helloworld.HelloAgainRequest{AgainName: "again_hello_test"})
	if err != nil || r.AgainMessage != "Replay again_hello_test" {
		t.Errorf("test SayHelloAgain failed")
	}
	t.Log("Reply: ", r.AgainMessage)
}
