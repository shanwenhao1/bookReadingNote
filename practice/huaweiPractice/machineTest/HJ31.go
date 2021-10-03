package machineTest

import (
	"fmt"
	"regexp"
)

/*
	问题描述: 对字符串中的所有单词进行倒排。
				说明: 构成单词的字符只有26个大写或小写英文字母
					 非构成单词的字符均视为单词间隔符
					 要求倒排后的单词间隔符以一个空格表示；如果原字符串中相邻单词间有多个间隔符时，倒排转换后也只允许出现一个空格间隔符
		输入描述: 输入一行以空格来分隔的句子
		输出描述: 输出句子的逆序

			示例1: 	输入: I am a student
					输出：student a am I
			示例2: 	输入: $bo*y gi!r#l
					输出：l r gi y bo
*/

/*
	反转字符串切片
		reverseSlice([]string{"a", "b"})
		return []string{"b", "a"}
*/
func reverseSlice(strSlice []string) []string {
	sLen := len(strSlice)
	for i := 0; i < sLen/2; i++ {
		strSlice[i], strSlice[sLen-i-1] = strSlice[sLen-i-1], strSlice[i]
	}
	return strSlice
}

func SplitStrReverse(str string) string {
	var res string
	// 使用正则匹配
	r, _ := regexp.Compile(`[a-z|A-Z]+`)
	// FindAllString 返回[]string. 所有匹配到的都会返回, n为匹配次数, 负值表示全部匹配
	strSplit := r.FindAllString(str, -1)
	// 反转匹配的str切片
	strSplit = reverseSlice(strSplit)
	// 拼接字符串
	for pos, _str := range strSplit {
		if pos == 0 {
			res = _str
			continue
		}
		res = fmt.Sprintf("%s %s", res, _str)
	}
	return res
}

func hj31Main() {
	var inputS = HJStrScan()

	for _, input := range inputS {
		res := SplitStrReverse(input)
		fmt.Printf("%s\n", res)
	}
}
