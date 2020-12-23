package Goroutine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
	context 用于追踪上下文, 也可追踪goroutine达到控制的目的
*/

type userI interface {
	GetUserId() string
}

type User struct {
	userId string
}

func (user User) GetUserId() string {
	return user.userId
}

func workerContext(ctx context.Context, text string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("got stop signal, %s out", text))
			return
		default:
			fmt.Println(fmt.Sprintf("%s still working", text))
			time.Sleep(1 * time.Second)
		}
	}
}

func workerContextU(ctx context.Context, key string, wg *sync.WaitGroup) {
	var user *User
	defer wg.Done()
	ur := ctx.Value(key)
	if ur == nil {
		fmt.Println("goroutine get wrong context value-key")
		return
	}
	// 利用反射获取context上下文中的user
	if newUser, ok := ur.(*User); !ok {
		fmt.Println("context value is not a user struct")
		return
	} else {
		user = newUser
	}
	fmt.Println(fmt.Sprintf("goroutine get user: %s", user.GetUserId()))
}

/*
	通过context控制goroutine的主动和超时退出
*/
func ContextUse() {
	// 创建5秒后超时退出的context, 其也会触发ctx.Done()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 运行goroutine, 让其触发超时退出
	go workerContext(ctx, "--- test goroutine 1 with context ---")

	ctx2, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	// 运行goroutine, 主动退出
	go workerContext(ctx2, "--- test goroutine 2 with context ---")
	cancle()

	time.Sleep(6 * time.Second) // 为了等待ctx超时退出
	fmt.Println(ctx.Err())
	fmt.Println(ctx2.Err())
}

/*
	context 控制链, 将上下文包含的信息, 主要是user传递给goroutine
*/
func ContextLink() {
	var ur = User{
		userId: "test user id",
	}
	wg := new(sync.WaitGroup)
	// 设置context 5s超时时间
	ctx1, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 创建下一个context, 其附带user上下文信息
	wg.Add(1)
	// key可以使用一个常量来定义, 从而使得函数不需要定义该key参数, 看个人选择
	ctx := context.WithValue(ctx1, "user", &ur)

	// 启动一个goroutine, ctx包含user信息
	go workerContextU(ctx, "user", wg)
	wg.Wait()
}
