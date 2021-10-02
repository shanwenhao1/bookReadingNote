package machineTest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	问题描述: 写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。
			不区分大小写，字符串长度小于500

		输入描述: 第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母
		输出描述: 输出输入字符串中含有该字符的个数
	示例: 输入：ABCabc
				A
		输出：2
*/

func LetterNum(str *string, letter *string) int {
	// 输入小写
	num := strings.Count(strings.ToLower(*str), strings.ToLower(*letter))
	return num
}

func letterScan(str *string, le *string) error {
	reader := bufio.NewReader(os.Stdin)
	strData, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	leData, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	*str = string(strData)
	*le = string(leData)
	return nil
}

func hj2Main() {
	var inputStr, letter string

	err := letterScan(&inputStr, &letter)
	if err != nil {
		panic(err)
	}

	leNum := LetterNum(&inputStr, &letter)
	fmt.Printf("%d", leNum)
}
