package cgo

/*
#include <stdio.h>

static void SayHello(const char* s){
	puts(s);
}
*/
import "C"

func BaseCgo() {
	C.SayHello(C.CString("Hello, World\n"))
}
