package code

import "fmt"

/**
 * 输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为 O(n).
 * @param array int整型一维数组
 * @return int整型

 * 注: 注意考虑最大值为负值的情况
 */
func FindGreatestSumOfSubArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	// write code here
	var sum, max = 0, array[0]
	for _, v := range array {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func SubArraySumRun() {
	var data = []int{1, -2, 3, 10, -4, 7, 2, -5}
	var result int
	result = FindGreatestSumOfSubArray(data)
	fmt.Println("max sub array sum: ", result)
}
