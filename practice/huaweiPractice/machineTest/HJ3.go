package machineTest

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

/*
	问题描述: 明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到1000之间的随机整数（N≤1000），
			对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。然后再把这些数从小到大排序，
			按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序”的工作(同一个测试用例里可能会有多组数据(用于不同的调查，
			希望大家能正确处理)
				注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。当没有新的输入时，说明输入结束
		输入描述: 注意：输入可能有多组数据(用于不同的调查)。每组数据都包括多行，第一行先输入随机整数的个数N，
				接下来的N行再输入相应个数的整数。具体格式请看下面的"示例"
		输出描述: 返回多行，处理后的结果

	示例:
		输入:
			3					第一个样例N=3
			2
			2
			1					第一个样例随机数截止
			11					第二个样例N=11
			10
			20
			40
			32
			67
			40
			20
			89
			300
			400
			15					第二个样例随机数截止
		输出:
			1
			2
			10
			15
			20
			32
			40
			67
			89
			300
			400

*/

/*
	使用map的思想(利用空间换时间)进行去重
*/
func removeDuplicate(data *[]int, originD *[]int) {
	rmMap := make(map[int]struct{})
	// 去重
	for _, _data := range *data {
		if _, ok := rmMap[_data]; ok {
			continue
		}
		*originD = append(*originD, _data)
		rmMap[_data] = struct{}{}
	}
}

func DeDuAndSort(data *[]int) *[]int {
	var newData []int

	// 去重
	removeDuplicate(data, &newData)
	// 排序
	sort.Ints(newData)
	return &newData
}

/*
	获取控制台输入, 单个样本输入
*/
func oneScan(scanner *bufio.Scanner, hc chan int, dataC chan *[]int) {
	/* 输入完毕, 正常应该检测EOF
	if err := scanner.Err(); err != nil{
		if err != io.EOF{
			hc <- -1
		}else{
			hc <- 1
		}
	}
	*/

	// 读取一行
	scanner.Scan()
	input := scanner.Text()
	// 输入完毕(正常应该检测EOF信号, 这里为了通过只能检测是否为空)
	if input == "" {
		hc <- 1
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		hc <- -2
	}
	curInput := make([]int, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		data, err := strconv.Atoi(scanner.Text())
		if err != nil {
			hc <- -2
		}
		curInput = append(curInput, data)
	}

	hc <- 0
	dataC <- &curInput
}

/*
	获取控制台输入, 包含多个样本输入
*/
func hj3Scan(inputData *[][]int) error {
	var scanner = bufio.NewScanner(os.Stdin)
	var timer = time.NewTimer(5 * time.Second)

	hChan := make(chan int, 1)
	dataChan := make(chan *[]int, 1)

	for {
		// 一次样本输入协程启动
		go oneScan(scanner, hChan, dataChan)

		select {
		// 输入超时控制
		case <-timer.C:
			return errors.New("timeout, break")
		case status := <-hChan:
			if status == -1 {
				return errors.New("input error")
			} else if status == -2 {
				return errors.New("input wrong")
			} else if status == 1 {
				// 无输入退出
				return nil
			} else {
				addrD := <-dataChan
				*inputData = append(*inputData, *addrD)
				// 刷新超时
				timer.Reset(5 * time.Second)
				continue
			}
		}
	}
}

func hj3Main() {
	var inputData [][]int

	// 获取输入
	err := hj3Scan(&inputData)
	if err != nil {
		panic(err)
	}

	for _, numSlice := range inputData {
		sortD := DeDuAndSort(&numSlice)
		for _, _data := range *sortD {
			fmt.Printf("%d\n", _data)
		}
	}
}
