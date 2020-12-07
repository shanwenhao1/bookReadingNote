package code

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
 *
 * @param number int整型
 * @return int整型
 */
func JumpFloor(number int) int {
	// write code here
	if number <= 0 {
		return 0
	} else if number <= 2 {
		return number
	} else {
		return JumpFloor(number-1) + JumpFloor(number-2)
	}
}
