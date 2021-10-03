package machineTest

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
	问题描述: 写出一个程序，接受一个十六进制的数，输出该数值的十进制表示
		输入描述: 输入一个十六进制的数值字符串。注意：一个用例会同时有多组输入数据
		输出描述: 输出该数值的十进制字符串。不同组的测试用例用\n隔开

			示例: 	输入: 0xA
						0xAA
					输出: 10
						170
*/

/*
	将16进制字符串转换为10进制字符串
*/
func StrToInt10(str string) (string, error) {
	// 根据0x分隔, 并取后面数值部分
	strS := strings.FieldsFunc(str, func(r rune) bool {
		return r == 'x' || r == 'X'
	})[1]
	strS = strings.ToUpper(strS)

	if strInt, err := strconv.ParseInt(strS, 16, 32); err == nil {
		return strconv.Itoa(int(strInt)), nil
	} else {
		return "", errors.New("parse error")
	}
}

func hj5Main() {
	var inputS = HJStrScan()

	for _, input := range inputS {
		strNum, err := StrToInt10(input)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", strNum)
	}
}
