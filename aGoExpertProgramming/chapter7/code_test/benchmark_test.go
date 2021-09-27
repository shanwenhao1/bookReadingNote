package code_test

import (
	"bookReadingNote/aGoExpertProgramming/chapter7/code"
	"testing"
	"time"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	/* 	b.N	表示循环执行的次数
	N值是动态调整的，直到可靠的算出程序执行时间后才会停止
	*/
	for i := 0; i < b.N; i++ {
		code.MakeSliceWithoutAlloc()
	}
}

/*
	预先分配内存可以进一步提升性能
*/
func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		code.MakeSliceWithPreAlloc()
	}
}

func BenchmarkSetBytes(b *testing.B) {
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second) // 模拟待测函数
	}
}
