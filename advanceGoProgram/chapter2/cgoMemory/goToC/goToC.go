package goToC

/*
#include <stdio.h>
void printString(const char* s, int n){
	int i;
	for(i = 0; i < n; i++){
		putchar(s[i]);
	}
	putchar('\n');
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

/*
	CGO规定, 在调用的C语言函数返回前, CGO保证传入的Go语言内存在此期间不会发生移动
*/

func printString(s string) {
	p := (*reflect.StringHeader)(unsafe.Pointer(&s))
	C.printString((*C.char)(unsafe.Pointer(p.Data)), C.int(len(s)))
}

func MemoryRun() {
	s := "hello"
	printString(s)
}
