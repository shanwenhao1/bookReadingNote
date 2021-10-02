package machineTest

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	问题描述: 计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000.（注：字符串末尾不以空格为结尾）
		输入描述: 输入一行，代表要计算的字符串，非空，长度小于5000
		输出描述: 输出一个整数，表示输入字符串最后一个单词的长度

	示例:
		输入：hello nowcoder
		输出：8
		说明：最后一个单词为nowcoder，长度为8
*/

func StrWordLLen(str *string) int {
	// strings.Fields根据空白符作为分隔符分隔字符串
	newStr := strings.Fields(*str)
	return len(newStr[len(newStr)-1])
}

/*
func StrInputCheck() (int, error) {
	var inputStr string
	n, _ := fmt.Scan("%s", &inputStr)
	if n == 0 {
		return 0, errors.New("wrong input")
	}

	lastLen := StrWordLLen(&inputStr)
	return lastLen, nil
}
*/

/*
	Acm模式输入输出
*/
func Scan(input *string) error {
	// 捕捉控制台输入操作
	reader := bufio.NewReader(os.Stdin)
	data, _, err := reader.ReadLine()
	if err != nil {
		return err
	}
	*input = string(data)
	return nil
}

func hj1Main() {
	var inputStr string
	err := Scan(&inputStr)
	if err != nil {
		panic(err)
	}

	lastLen := StrWordLLen(&inputStr)
	fmt.Printf("%d", lastLen)
	return
}
