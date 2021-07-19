package utils

import (
	"strings"
	"unicode"
)

type StrUtilIF interface {
	CompareEqual(str1 string, str2 string) bool
	CompareEqualUpper(str1 string, str2 string) bool
	CkStrIsDigit(par string) bool
	CkStrIsLetter(par string) bool
	CkStrIsAlNum(par string) bool
	CkStrHasSpace(par string) bool
}

type StrUtil struct{}

// 校验字符串是否一样(区分大小写)
func (this StrUtil) CompareEqual(str1 string, str2 string) bool {
	if strings.Compare(str1, str2) != 0 {
		return false
	}
	return true
}

// 校验字符串是否一样(不区分大小写)
func (this StrUtil) CompareEqualUpper(str1 string, str2 string) bool {
	var _str1, _str2 string
	_str1 = strings.ToUpper(str1)
	_str2 = strings.ToUpper(str2)
	if strings.Compare(_str1, _str2) != 0 {
		return false
	}
	return true
}

/*
	检查字符串是否全为数字(10进制数)
*/
func (this StrUtil) CkStrIsDigit(par string) bool {
	var (
		isNumber bool
	)
	isNumber = true
	for _, p := range par {
		if !unicode.IsDigit(p) {
			isNumber = false
			break
		}
	}
	return isNumber
}

/*
	检查字符串是否全为字母
*/
func (this StrUtil) CkStrIsLetter(par string) bool {
	var (
		isLetter bool
	)
	isLetter = true
	for _, p := range par {
		if !unicode.IsLetter(p) {
			isLetter = false
			break
		}
	}
	return isLetter
}

/*
	检查字符串是否只含字母或数字(10进制数)
*/
func (this StrUtil) CkStrIsAlNum(par string) bool {
	var (
		isTrue bool
	)
	isTrue = true
	for _, p := range par {
		if !unicode.IsDigit(p) && !unicode.IsLetter(p) {
			isTrue = false
			break
		}
	}
	return isTrue
}

/*
	检查字符串是否包含空格(包括制表符)
*/
func (this StrUtil) CkStrHasSpace(par string) bool {
	var (
		hasSpace bool
	)
	hasSpace = false
	for _, p := range par {
		if unicode.IsSpace(p) {
			hasSpace = true
			break
		}
	}
	return hasSpace
}
