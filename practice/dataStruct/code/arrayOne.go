package code

/*
	问题一: 只出现一次的数字
		描述: 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
			(你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗?)
		示例: 输入: [2,2,1]
   			 输出: 1


	问题二: 只出现一次的数字
		描述: 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素
			(你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗?)
		示例: 输入：nums = [2,2,3,2]
			 输出：3
*/

/*
	解题思路: 使用Hash映射可以快速的解决此问题(但是会使用额外空间).
			因此这里使用异或运算 a⊕a=0 0⊕a=a 可以完美的解决此问题
*/
func SingleNumber(nums []int) int {
	var ans = 0
	for _, _num := range nums {
		ans ^= _num
	}
	return ans
}

/*
	解题思路:
*/
func SingleNumberTwo(nums []int) int {
	return 0
}
