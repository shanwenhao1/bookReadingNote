package loadBalance

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
 普通伪随机洗牌算法:
	洗牌不均匀，会导致整个数组第一个节点有大概率被选中，并且多个节点的负载分布不均衡
*/
func shuffle1(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

/*
	fisher-yates算法
*/
func shuffle2(indexes []int) {
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}

func ShuffleCompare() {
	var cnt1 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle1(sl)
		cnt1[sl[0]]++
	}

	var cnt2 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var sl = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle2(sl)
		cnt2[sl[0]]++
	}

	fmt.Println(cnt1, "\n", cnt2)
}
