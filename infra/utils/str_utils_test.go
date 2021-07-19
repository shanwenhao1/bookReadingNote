package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrUtil_CompareEqual(t *testing.T) {
	var strU StrUtil
	a := assert.New(t)

	str1 := "test str 1"
	str2 := "test str 2"
	str3 := "TEST str 1"
	str4 := "test str 1"

	a.Equal(strU.CompareEqual(str1, str2), false, "CompareEqual function error")
	a.Equal(strU.CompareEqual(str1, str3), false, "CompareEqual function error")
	a.Equal(strU.CompareEqual(str1, str4), true, "CompareEqual function error")
}

func TestStrUtil_CompareEqualUpper(t *testing.T) {
	var strU StrUtil
	a := assert.New(t)

	str1 := "test str 1"
	str2 := "test str 2"
	str3 := "TEST str 1"
	str4 := "test str 1"

	a.Equal(strU.CompareEqualUpper(str1, str2), false, "CompareEqualUpper function error")
	a.Equal(strU.CompareEqualUpper(str1, str3), true, "CompareEqualUpper function error")
	a.Equal(strU.CompareEqualUpper(str1, str4), true, "CompareEqualUpper function error")
}

func TestStrUtil_CkStrIsDigit(t *testing.T) {
	var (
		strU StrUtil
	)
	a := assert.New(t)

	str1 := "10001"
	str2 := "0x10"

	a.Equal(strU.CkStrIsDigit(str1), true, "CkStrIsDigit function error")
	a.Equal(strU.CkStrIsDigit(str2), false, "CkStrIsDigit function error")
}

func TestStrUtil_CkStrIsLetter(t *testing.T) {
	var (
		strU StrUtil
	)
	a := assert.New(t)

	str1 := "ceshistr"
	str2 := "ce shi str"

	a.Equal(strU.CkStrIsLetter(str1), true, "CkStrIsLetter function error")
	a.Equal(strU.CkStrIsLetter(str2), false, "CkStrIsLetter function error")
}

func TestStrUtil_CkStrIsAlNum(t *testing.T) {
	var (
		strU StrUtil
	)
	a := assert.New(t)

	str1 := "ceshi01str"
	str2 := "ce shi 01str"

	a.Equal(strU.CkStrIsAlNum(str1), true, "CkStrIsAlNum function error")
	a.Equal(strU.CkStrIsAlNum(str2), false, "CkStrIsAlNum function error")
}

func TestStrUtil_CkStrHasSpace(t *testing.T) {
	var (
		strU StrUtil
	)
	a := assert.New(t)

	str1 := "ceshistr"
	str2 := "ce shi str"
	str3 := "ce\tshistr"

	a.Equal(strU.CkStrHasSpace(str1), false, "CkStrHasSpace function error")
	a.Equal(strU.CkStrHasSpace(str2), true, "CkStrHasSpace function error")
	a.Equal(strU.CkStrHasSpace(str3), true, "CkStrHasSpace function error")
}
