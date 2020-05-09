package baseCgi

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/*
 结构树的遍历
*/
type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age    int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	if pass, _ := regexp.MatchString(
		`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, input,
	); pass {
		return true
	}
	return false
}

func Validate(v interface{}) (bool, string) {
	validateResult := true
	errMsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		k := fieldVal.Kind()

		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errMsg = "validate int failed, tag is: " + strconv.FormatInt(
					tagVal, 10,
				)
				validateResult = false
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "email":
				nestedResult := validateEmail(val)
				if nestedResult == false {
					errMsg = "validate mail failed, field val is: " + val
					validateResult = false
				}
			}
		case reflect.Struct:
			// 如果有内嵌的 struct，那么深度优先遍历
			// 就是一个递归过程
			valInter := fieldVal.Interface()
			nestedResult, msg := Validate(valInter)
			if nestedResult == false {
				validateResult = false
				errMsg = msg
			}
		}
	}
	return validateResult, errMsg
}

func ValidateRun() {
	var a = T{Age: 10, Nested: Nested{Email: "abc@abc.com"}}

	validateResult, errMsg := Validate(a)
	fmt.Println(validateResult, errMsg)
}
