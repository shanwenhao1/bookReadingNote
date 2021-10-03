package machineTest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
		问题描述: 正整数A和正整数B 的最小公倍数是指 能被A和B整除的最小的正整数值，设计一个算法，求输入A和B的最小公倍数
			输入描述: 输入两个正整数A和B
			输出描述: 输出A和B的最小公倍数
				示例： 	输入: 5 7
						输出: 35

s
		解题思路: 利用辗转相除法: https://blog.csdn.net/huyr_123/article/details/81670972
			算法解析: 1. 最小公倍数 = 两数的乘积 / 最大公约数
					2. 要求x、y (x>y) 两数的最大公约数f(x,y), 假设x = ky + b, 其中`k=x/y`、`b=x%y`.
							2.1 如果一个数能同时整除x、y, 那么公式中的`ky + b`能够被该数整除, 那么b也能被该数整除.
								即： y和b也能同时被该数整除. 那么可以得出结论x和y的公约数、最大公约数与b和y一致
								f(x, y) = f(y, x%y)  y>0
							2.2 根据2.1的理论, 我们可以把两个较大数的最大公约数问题转换为两个较小数的最大公约数问题,
								直到其中一个数为0, 剩下的另外一个数即为最大公约数了
*/

/*
	使用辗转相除法获取两数的最大公约数
*/
func maxG(a, b int) int {
	if a < b {
		a, b = b, a
	}
	// 当一数为0时结束递归(此时已经获取到了最大公约数)
	if b == 0 {
		return a
	}
	num := maxG(b, a%b)
	return num
}

func LeastComMul(a, b int) int {
	g := maxG(a, b)
	lCM := a * b / g
	return lCM
}

func hj108Scan() [][]int {
	var scanner = bufio.NewScanner(os.Stdin)
	var inputI [][]int

	// 获取输入
	for {
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		}
		inputS := strings.Fields(input)
		var inputOne []int
		for _, data := range inputS {
			num, _ := strconv.Atoi(data)
			inputOne = append(inputOne, num)
		}
		inputI = append(inputI, inputOne)
	}
	return inputI
}

func hj108Main() {
	var inputI = hj108Scan()

	for _, input := range inputI {
		fmt.Printf("%d\n", LeastComMul(input[0], input[1]))
	}
}
