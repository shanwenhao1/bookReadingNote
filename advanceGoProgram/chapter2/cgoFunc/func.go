package cgoFunc

/*
#include <errno.h>

// 带返回值的函数调用(使用errno宏支持错误返回)
static int div(int a, int b){
	if(b == 0){
		errno = EINVAL;
		return 0;
	}
	return a/b;
}

// 不带返回值的void函数, 也可使用errno做特殊处理来获取C代码的状态
static void noreturn(){
}
*/
import "C"
import "fmt"

func FuncCgo() {

	// 带返回值的C Func调用
	v0, err0 := C.div(2, 1)
	fmt.Println(v0, err0)

	v1, err1 := C.div(2, 0)
	fmt.Println(v1, err1)

	// 不带返回值的C Func调用
	v, err := C.noreturn()
	fmt.Printf("%#v", v)
	fmt.Println(err)
	fmt.Println(C.noreturn())
}
