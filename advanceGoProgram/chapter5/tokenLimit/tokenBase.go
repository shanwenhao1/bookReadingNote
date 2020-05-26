package tokenLimit

/*
 该令牌桶库不支持令牌桶预热, 且无法修改初始的令牌容量, 个别极端情况下无法满足需求
*/

import (
	"fmt"
	"github.com/juju/ratelimit"
	"reflect"
	"time"
)

/*
	ratelimit令牌库用法
*/
func bucket() {
	// 创建默认令牌桶, 每隔100Millisecond向桶内放一个令牌, 桶的容量为10000且桶初始是满的, 超过桶容量部分会直接丢弃
	newBucket := ratelimit.NewBucket(100*time.Millisecond, 10000)
	// 创建令牌桶, 每隔100Millisecond向桶中放入10个令牌, 桶容量为10000
	newBucket = ratelimit.NewBucketWithQuantum(100*time.Millisecond, 10, 10000)
	// 创建令牌桶, 每秒钟向桶中放入0.1 * 1000个令牌, 桶容量为1000
	newBucket = ratelimit.NewBucketWithRate(0.1, 1000)
	// 除上述之外还有NewBucketWithClock、NewBucketWithRateAndClock等创建令牌桶的函数

	fmt.Println(reflect.TypeOf(newBucket))
}

var fillInterval = time.Millisecond * 10
var capacity = 100

// 使用buffer channel作为令牌桶
var tokenBucket = make(chan struct{}, capacity)

/*
令牌桶函数
*/
func BucketBuild() {

	fillToken := func() {
		// 使用定时触发向令牌桶发送令牌
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case <-ticker.C:
				select {
				case tokenBucket <- struct{}{}:
				default:
				}
				fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()
}

/*
获取令牌函数， `github.com/juju/ratelimit`在take令牌之前, 对令牌桶中的token数进行计算, 获取令牌数(惰性求值)， 再进行take操作
*/
func TakeAvailable(block bool) bool {
	var takenResult bool
	if block {
		select {
		case <-tokenBucket:
			takenResult = true
		}
	} else {
		select {
		case <-tokenBucket:
			takenResult = true
		default:
			takenResult = false
		}
	}

	return takenResult
}

func BucketRun() {
	BucketBuild()
	result := TakeAvailable(true)
	fmt.Println("-------------- result: ", result)
	time.Sleep(1 * time.Second)
}
