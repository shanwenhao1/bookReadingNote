package mytest

import (
	"bookReadingNote/aMicroservicesPatterns/code/chapter3/grpcExample/mytest/mock_mytest"
	"bookReadingNote/aMicroservicesPatterns/code/chapter3/grpcExample/mytest/proto/test"
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

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockHelloServiceClient := mock_test.NewMockHelloServiceClient(ctrl)
	req := &test.HelloRequest{Username: "my_test"}
	mockHelloServiceClient.EXPECT().SayHello(
		gomock.Any(),
		&rpcMsg{msg: req},
	).Return(&test.HelloResponse{Message: "Mocked Interface"}, nil)
	testSayHello(t, mockHelloServiceClient)
}

func testSayHello(t *testing.T, client test.HelloServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &test.HelloRequest{Username: "my_test"})
	if err != nil || r.Message != "Mocked Interface" {
		t.Errorf("test SayHello failed")
	}
	t.Log("Reply: ", r.Message)
}
