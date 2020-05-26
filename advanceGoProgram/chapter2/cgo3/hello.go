package cgo3

/*
#include <hello.h>
*/
import "C"

func Hello() {
	C.SayHello(C.CString("Hello, World\n"))
}
