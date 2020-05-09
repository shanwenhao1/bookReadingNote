package begin

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func HelloClient() {
	// 先启用HelloService
	HelloSer()
	time.Sleep(time.Second * 3)

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
