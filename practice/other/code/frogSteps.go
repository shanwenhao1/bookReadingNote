package code

import (
	"bookReadingNote/infra/utils"
	"fmt"
)

/*
问题: 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）

解题思路: 可以用动态规划来求解该题
		跳到第n个台阶，只有两种可能
			从第n-1个台阶跳1个台阶
			从第n-2个台阶跳2个台阶
			只需求出跳到第n-1个台阶和第n-2个台阶的可能跳法即可
				F（n）:n个台阶的跳法
					递推公式：F（n）=F（n-1）+F（n-2）
					不难发现这是一个斐波那契数列
					起始条件为F（0）=1，F（1）=1
*/

/**
 * 递归方式, 效率慢
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	if number <= 0 {
		return 0
	} else if number <= 2 {
		return number
	} else {
		return jumpFloor(number-1) + jumpFloor(number-2)
	}
}

/**
 * 迭代算法, 效率较高
 * @param number int整型
 * @return int整型
 */
func jumpFloor2(number int) int {
	// write code here
	var former1, former2, target = 1, 2, 3
	if number <= 0 {
		return 0
	} else if number <= 2 {
		return number
	} else {
		for i := 3; i <= number; i++ {
			target = former1 + former2
			former1 = former2
			former2 = target
		}
		return target
	}
}

func FrogRun() {
	// 青蛙跳台阶问题, 递归方式
	t1 := utils.GetCurTimeUtc()
	fmt.Println(jumpFloor(39))
	t2 := utils.GetCurTimeUtc()
	// 青蛙跳台阶问题, 迭代方式
	fmt.Println(jumpFloor2(39))
	t3 := utils.GetCurTimeUtc()
	fmt.Println("recursive algorithm cost time: ", utils.GetTimeSub(t1, t2), "iterative algorithm cost time: ",
		utils.GetTimeSub(t2, t3))
}
