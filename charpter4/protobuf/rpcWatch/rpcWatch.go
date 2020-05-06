package rpcWatch

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

/*
 样例: RPC构造一个简单的内存键值数据库
*/
type KVStoreService struct {
	m      map[string]string           // map类型, 用于存储键值数据
	filter map[string]func(key string) // filter对应每个watch()调用时的过滤器函数列表
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]

	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()

	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	default:
		return nil
	}
}

func WatchServer() {
	go func() {
		// 将 KVStoreService 对象注册为一个 RPC 服务
		// 将对象中所有满足 RPC 规则的对象方法注册为 RPC 函数
		// 所有注册的方法会放在 “KVStoreService” 服务空间执行
		_ = rpc.RegisterName("KVStoreService", NewKVStoreService())
		listener, err := net.Listen("tcp", ":1234")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		// 在该 TCP 连接上为对方提供 RPC 服务
		rpc.ServeConn(conn)
	}()
}
