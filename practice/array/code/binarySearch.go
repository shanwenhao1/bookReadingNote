package code

/*
	问题: 二分查找
		描述: 请实现有重复数字的升序数组的二分查找
			 给定一个 元素有序的（升序）整型数组 nums 和一个目标值 target  ，
			 写一个函数搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1

		示例:
			输入：[1,2,4,4,5],4
			返回值：2
			说明：从左到右，查找到第1个为4的，下标为2，返回2
*/

func Search(nums []int, target int) int {
	// write code here
	var mid, ret = 0, -1
	var left, right = 0, len(nums) - 1
	// 二分法找到target目标的值, 但不确定左边是否还存在target
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			ret = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// 寻找最左边的target(如果存在的话)
	for ret >= 1 {
		if nums[ret-1] == target {
			ret -= 1
		} else {
			break
		}
	}
	return ret
}
