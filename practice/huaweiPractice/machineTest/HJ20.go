package machineTest

import (
	"fmt"
	"unicode"
)

/*
	问题描述: 密码要求:
				长度超过8位
				包括大小写字母.数字.其它符号,以上四种至少三种
				不能有相同长度大于2的子串重复

		输入描述: 一组或多组长度超过2的字符串。每组占一行
		输出描述: 如果符合要求输出：OK，否则输出NG

			示例: 输入: 021Abc9000
					 021Abc9Abc1
					 021ABC9000
					 021$bc9000
				 输出: OK
						NG
						NG
						OK
*/

/*
	检查字符串是否有超过2个字符的子串重复
*/
func checkDu(str string) bool {
	var ckM = make(map[string]struct{})

	// str转为byte方便截取子字符串
	strB := []byte(str)
	strL := len(str)
	for pos, _ := range strB {
		// 到字符串末尾退出
		if pos > strL-1-2 {
			break
		}
		subS := string(strB[pos : pos+3])
		if _, ok := ckM[subS]; !ok {
			ckM[subS] = struct{}{}
		} else {
			return false
		}
	}
	return true
}

// 统计字符类型是否超过3
func checkChar(str string) bool {
	var typeN int

	var typeM = make(map[string]int)
	for _, strB := range str {
		// 数字则记录类型
		if unicode.IsDigit(strB) {
			// 数字则记录类型
			if _, ok := typeM["digit"]; !ok {
				typeM["digit"] = 1
				typeN += 1
			} else {
				continue
			}
		} else if unicode.IsUpper(strB) {
			// 大写字母记录类型
			if _, ok := typeM["upper"]; !ok {
				typeM["upper"] = 1
				typeN += 1
			} else {
				continue
			}
		} else if unicode.IsLower(strB) {
			// 小写字母记录类型
			if _, ok := typeM["lower"]; !ok {
				typeM["lower"] = 1
				typeN += 1
			} else {
				continue
			}
		} else {
			if _, ok := typeM["other"]; !ok {
				typeM["other"] = 1
				typeN += 1
			} else {
				continue
			}
		}
	}
	if typeN < 3 {
		return false
	}
	return true
}

func PassCheck(str string) bool {
	// 检查长度是否超过8位
	if len(str) <= 8 {
		return false
	}
	// 检查是否包括大小写字母.数字.其它符号,以上四种至少三种
	if !checkChar(str) {
		return false
	}
	// 检查是否有长度超过2的子串重复
	if !checkDu(str) {
		return false
	}
	return true
}

func hj20Main() {
	var inputS = HJStrScan()

	for _, input := range inputS {
		if PassCheck(input) {
			fmt.Println("OK")
		} else {
			fmt.Println("NG")
		}
	}
}
