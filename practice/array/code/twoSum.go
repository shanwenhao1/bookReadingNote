package code

import "fmt"

/**
 *  问题: 给出一个整数数组，请在数组中找出两个加起来等于目标值的数，
	你给出的函数twoSum 需要返回这两个数字的下标（index1，index2），需要满足 index1 小于index2.。注意：下标是从1开始的
	假设给出的数组中只存在唯一解
	例如：
	给出的数组为 {20, 70, 110, 150},目标值为90
	输出 index1=1, index2=2
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
*/
func twoSum(numbers []int, target int) []int {
	// write code here
	var vMap = make(map[int]int, len(numbers))
	for k, v := range numbers {
		vMap[v] = k + 1
	}
	for k, v := range numbers {
		if index, ok := vMap[target-v]; ok && k+1 != index {
			return []int{k + 1, index}
		}
	}
	return nil
}

func TwoSumRun() {
	var num = []int{3, 2, 4}
	var target = 6
	fmt.Println("result is: ", twoSum(num, target))
}
