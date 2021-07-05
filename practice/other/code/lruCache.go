package code

import "fmt"

/*
	问题: 设计LRU缓存结构
		问题描述:
			设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
			set(key, value)：将记录(key, value)插入该结构
			get(key)：返回key对应的value值
			[要求]
			set和get方法的时间复杂度为O(1)
			某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
			当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
			若opt=1，接下来两个整数x, y，表示set(x, y)
			若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
			对于每个操作2，输出一个答案
		示例:
			输入：
			[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
			复制
			返回值：
			[1,-1]
			复制
			说明：
			第一次操作后：最常使用的记录为("1", 1)
			第二次操作后：最常使用的记录为("2", 2)，("1", 1)变为最不常用的
			第三次操作后：最常使用的记录为("3", 2)，("1", 1)还是最不常用的
			第四次操作后：最常用的记录为("1", 1)，("2", 2)变为最不常用的
			第五次操作后：大小超过了3，所以移除此时最不常使用的记录("2", 2)，加入记录("4", 4)，并且为最常使用的记录，然后("3", 2)变为最不常使用的记录
*/

/*
	operators: 输入
	k: 缓存大小限制
*/
func lruCache(operators [][]int, k int) []int {
	var index int
	// 构造存放操作结果的map
	res := make([]int, 0, len(operators))
	// 使用切片来模拟堆栈操作
	key := make([]int, 0, k)
	value := make([]int, 0, k)
	for _, v := range operators {
		// 入栈操作
		if v[0] == 1 {
			if len(key) == k {
				// 剔除最不常用数据
				key = key[1:]
				value = value[1:]
			}
			key = append(key, v[1])
			value = append(value, v[2])
		} else if v[0] == 2 {
			// 获取操作, 需要更新常用和不常用数据, 并记录返回res的值
			index = -1
			// 寻找get需要的数据
			for i := 0; i < len(key); i++ {
				if v[1] == key[i] {
					index = i
					break
				}
			}
			if index == -1 {
				res = append(res, -1)
			} else {
				res = append(res, value[index])
				// 需要更新常用和不常用数据
				if index < k-1 {
					key = append(key[:index], append(key[index+1:], key[index])...)
					value = append(value[:index], append(value[index+1:], value[index])...)
				}
			}
		}
	}
	return res
}

func LRUCacheRun() {
	var operatorCache [][]int = [][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {2, 1}, {1, 4, 4}, {2, 2}}
	result := lruCache(operatorCache, 3)
	fmt.Println("-----------LRU Cache Run result: ", result)
}
