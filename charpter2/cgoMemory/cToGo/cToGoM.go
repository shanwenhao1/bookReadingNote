package cToGo

/*
#include <stdlib.h>
void* makeSlice(size_t memSize){
	return malloc(memSize);
}
*/
import "C"
import (
	"unsafe"
)

func makeByteSlice(n int) []byte {
	p := C.makeSlice(C.size_t(n))
	return ((*[1 << 31]byte)(p))[0:n:n]
}

func freeByteSlice(p []byte) {
	C.free(unsafe.Pointer(&p[0]))
}

func MemoryRun() {
	// 利用C创建大于4GB内存的切片供Go使用
	s := makeByteSlice(1 << 31)
	s[len(s)-1] = 123
	print(s[len(s)-1])
	// 释放内存
	freeByteSlice(s)
}
