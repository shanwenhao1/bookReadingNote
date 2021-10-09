package machineTest

import (
	"errors"
)

/*
	问题描述: A、B两人分苹果. 第一行输入苹果数量(例如: 3), 第二行输入苹果的重量(例如: 3, 5, 6)
				A的分法是要求等分, 但是重量计算使用二进制进行相加但不进位比如011 + 001 = 010
				B的分法是在满足A的情况下尽可能使自己获取的苹果总重量最大, 采用10进制相加可进位
		求输入样例是否满足B的分法, 如果有返回B获取苹果的最大重量, 如果没有返回-1

		示例1: 输入: 3
					3 5 6
				输出: 11
		示例2: 输入: 5

	解题思路: 1. 所有样本可能性使用排列组合得到
			2. 使用异或运算得出二进制不进位运算相加结果
			3， 根据对比找出符合A分配的B的最大重量

*/

/*
	将left中的1移动到最左端
*/
func moveToLeft(left []int) []int {
	var oneNum int
	// 统计1的个数
	for _, data := range left {
		if data == 1 {
			oneNum += 1
		}
	}
	var moveD = make([]int, len(left))
	for i := 0; i < oneNum; i++ {
		moveD[i] = 1
	}
	return moveD
}

/*
	记录A、B分组后的下标
*/
func posNote(dataPosS []int, aNum int, dataL int) ([]int, []int) {
	var aPos = make([]int, 0, aNum)
	var bPos = make([]int, 0, dataL-aNum)
	for pos, _data := range dataPosS {
		if _data == 1 {
			aPos = append(aPos, pos)
		} else {
			bPos = append(bPos, pos)
		}
	}
	return aPos, bPos
}

/*
	排列组合问题:
	将切片分为两部分, aNum为第一部分的数量, 返回分配后的所有可能性
		只记录下标, 节约内存
*/
func DvdExHa(data []int, aNum int) ([][][]int, error) {
	var dataL = len(data)
	var dvDAf [][][]int
	var allAPos [][]int
	var allBPos [][]int

	if aNum < 1 || aNum > dataL-1 {
		return dvDAf, errors.New("wrong parameter")
	}

	// 创建有n个元素切片，切片元素的值为1表示选中，为0则没选中
	var dataPosS = make([]int, dataL)
	// 初始化第一种组合, 从左到右aNum个数置位1(即第一种组合), 示例: 1 1 1 0 0 (五选三)
	for pos, _ := range dataPosS {
		if pos < aNum {
			dataPosS[pos] = 1
		} else {
			break
		}
	}
	aPos, bPos := posNote(dataPosS, aNum, dataL)
	allAPos = append(allAPos, aPos)
	allBPos = append(allBPos, bPos)

	/*
		从左到右扫描数组元素值的“10”组合，找到第一个“10”组合后将其变为“01”组合，同时将其左边的所有“1”全部移动到数组的最左端
			这里以五选三为例, 穷举步骤如下:
			1 1 1 0 0
			1 1 0 1 0
			1 0 1 1 0
			0 1 1 1 0
			1 1 0 0 1
			1 0 1 0 1
			0 1 1 0 1
			1 0 0 1 1
			0 1 0 1 1
			0 0 1 1 1
	*/
	for {
		find := false
		for i := 0; i < dataL-1; i++ {
			// 找到第一个10, 将其改为01, 并将其左边的1移动至最左端
			if dataPosS[i] == 1 && dataPosS[i+1] == 0 {
				find = true
				rightPart := dataPosS[i:]
				rightPart[0] = 0
				rightPart[1] = 1

				leftPart := moveToLeft(dataPosS[:i])

				// 记录该组合
				newDataPos := make([]int, 0, dataL)
				newDataPos = append(newDataPos, leftPart...)
				newDataPos = append(newDataPos, rightPart...)
				aPos, bPos := posNote(newDataPos, aNum, dataL)
				allAPos = append(allAPos, aPos)
				allBPos = append(allBPos, bPos)

				dataPosS = newDataPos
				break
			}
		}
		// 未找到10说明穷举完毕, 退出循环
		if find == false {
			break
		}
	}
	for pos, _ := range allAPos {
		dvDAf = append(dvDAf, [][]int{allAPos[pos], allBPos[pos]})
	}
	return dvDAf, nil
}

/*
	按照A的计算方法进行相加, 判断是否符合A的分配
		rsData []int: 苹果重量数据
		ex [][]int: 苹果分配后A、B拥有苹果对应rsData的下标
*/
func BinaryCheck(rsData []int, ex [][]int) bool {
	var aPos, bPos = ex[0], ex[1]
	var aSum, bSum int
	for pos, aD := range aPos {
		if pos == 0 {
			aSum = rsData[aD]
			continue
		}
		// 异或运算即可符合A的二进制不进位加法
		aSum = aSum ^ rsData[aD]
	}
	for pos, bD := range bPos {
		if pos == 0 {
			bSum = rsData[bD]
			continue
		}
		bSum = bSum ^ rsData[bD]
	}
	if aSum == bSum {
		return true
	}
	return false
}

/*
	分苹果问题
*/
func AppleDivide(rsData []int) (int, error) {
	var allDvd [][][]int
	// 步骤一: 先穷举a、b所分苹果的可能性.
	//		这里以A获取的苹果数量为锚点, 从1-len(rsData)/2即可, 再大的话其实重复了
	for i := 0; i < (len(rsData)+1)/2; i++ {
		curCom, err := DvdExHa(rsData, i+1)
		if err != nil {
			return -1, err
		}
		allDvd = append(allDvd, curCom...)
	}

	// 筛选出所有符合A分配的组合, 如果需要考虑性能的话, 这里可以使用多并发, 然后使用channel 返回结果
	var accordA [][][]int
	for _, dvEx := range allDvd {
		if BinaryCheck(rsData, dvEx) {
			accordA = append(accordA, dvEx)
		}
	}

	// 从符合A分配的组合中挑选b重量最大的
	var maxH = 0
	for _, accord := range accordA {
		sum := 0
		for _, bA := range accord[1] {
			sum += rsData[bA]
		}
		if sum > maxH {
			maxH = sum
			continue
		}
	}
	if maxH > 0 {
		return maxH, nil
	}
	return -1, nil
}
