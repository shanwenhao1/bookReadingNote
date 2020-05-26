package cgo2

/*
// CGO中增加了_GoString_预定义的C语言类型
void SayHello(_GoString_ s);
*/
import "C"
import "fmt"

func Hello() {
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}
