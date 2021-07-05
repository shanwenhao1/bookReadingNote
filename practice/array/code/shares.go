package code

import (
	"fmt"
)

/**
 * 问题: 假设你有一个数组，其中第 i 个元素是股票在第i 天的价格。
		你有一次买入和卖出的机会。（只有买入了股票以后才能卖出）。请你设计一个算法来计算可以获得的最大收益。
 *
 * @param prices int整型一维数组
 * @return int整型
*/
func maxProfit(prices []int) int {
	// write code here
	if prices == nil {
		return 0
	}
	var min, max = prices[0], 0
	for _, v := range prices {
		if v > min {
			if v-min > max {
				max = v - min
			}
		} else {
			min = v
		}
	}
	return max
}

func ProfitRun() {
	fmt.Println("-----Max Profit: ", maxProfit([]int{1, 4, 2}))
	fmt.Println("-----Max Profit: ", maxProfit([]int{2, 4, 1}))
}
