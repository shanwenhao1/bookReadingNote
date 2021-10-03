package machineTest

import (
	"fmt"
)

/*
	问题描述: 连续输入字符串，请按长度为8拆分每个输入字符串并进行输出
			长度不是8整数倍的字符串请在后面补数字0，空字符串不处理
		输入描述: 连续输入字符串(输入多次,每个字符串长度小于等于100)
		输出描述: 依次输出所有分割后的长度为8的新字符串

		示例:	输入：abc
					123456789
				输出: abc00000
					12345678
					90000000
*/

func StrSplit(str string) []string {
	var newStrS []string
	var zeroByte = []byte("0")[0]

	for {
		// 尾部无字符直接退出分隔循环
		if len(str) == 0 {
			break
		}
		// 尾部不足八个字符需要补足
		if len(str) < 8 {
			// 要补足的字符
			var addByte = make([]byte, 8-len(str))
			for pos, _ := range addByte {
				addByte[pos] = zeroByte
			}
			newStrS = append(newStrS, fmt.Sprintf("%s%s", str, string(addByte)))
			break
		}
		// 切分8位的字符串
		newStrS = append(newStrS, str[:8])
		str = str[8:]
	}
	return newStrS
}

func hj4Main() {
	var inputS = HJStrScan()

	for _, input := range inputS {
		splitS := StrSplit(input)
		for _, _str := range splitS {
			fmt.Printf("%s\n", _str)
		}
	}
}
