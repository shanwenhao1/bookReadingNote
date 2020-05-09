package rpcWatch

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()

	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)
}

func WatchClient() {
	// 先运行watch service
	WatchServer()
	time.Sleep(time.Second * 3)

	// quit为结束信号, 用于控制程序结束
	quit := make(chan struct{})
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	// 启动独立的 Goroutine 监控 key 的变化，同步阻塞，直到有 key 发生变化或者超时
	go func() {
		var keyChanged string
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: ", keyChanged)
		time.Sleep(time.Second * 2)
		quit <- struct{}{}
	}()

	var key string
	// 获取某个 key 的值
	err = client.Call("KVStoreService.Get", "abc", &key)
	if err != nil {
		log.Println("get key-1 error: ", err)
	} else {
		fmt.Println("get key-1: ", key)
	}
	// 设置某个 key 的值，因为原来对应的为空，调用该方法，会触发 Watch 返回
	err = client.Call("KVStoreService.Set", [2]string{"abc", "abc-valueTest"}, new(struct{}))
	if err != nil {
		log.Fatal(err)
	}
	// 再次获取key 的值
	err = client.Call("KVStoreService.Get", "abc", &key)
	if err != nil {
		log.Println("get key-1 error: ", err)
	} else {
		fmt.Println("get key-2: ", key)
	}
	<-quit
}
