package baseCgi

import (
	"fmt"
	"github.com/go-playground/validator"
)

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username string `validate:"gt=0"`
	// 同上
	PasswordNew string `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email string `validate:"email"`
}

var validate = validator.New()

/*
 请求校验器
*/
func validateFunc(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		// doSomething()
		return err
	}
	return nil
}

func TestValidate() {
	var req = RegisterReq{
		Username:       "Xar gin",
		PasswordNew:    "oh no",
		PasswordRepeat: "ohn",
		Email:          "alex@abc.com",
	}

	var req2 = RegisterReq{
		Username:       "Xar gin",
		PasswordNew:    "oh no",
		PasswordRepeat: "oh no",
		Email:          "alex@abc.com",
	}

	err := validateFunc(req)
	fmt.Println("请求校验器返回结果1: ", err)

	err1 := validateFunc(req2)
	fmt.Println("请求校验器返回结果2: ", err1)
}
