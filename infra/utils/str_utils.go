package utils

import "strings"

type StrCompare struct{}

// 校验字符串是否一样(区分大小写)
func (strC StrCompare) CompareEqual(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) != 0 {
		return false
	}
	return true
}

// 校验字符串是否一样(不区分大小写)
func (strC StrCompare) CompareEqualUpper(str1 string, str2 string) bool {
	var _str1, _str2 string
	_str1 = strings.ToUpper(str1)
	_str2 = strings.ToUpper(str2)
	if strings.Compare(_str1, _str2) != 0 {
		return false
	}
	return true
}
