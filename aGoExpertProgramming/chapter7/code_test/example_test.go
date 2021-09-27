package code

import (
	"bookReadingNote/aGoExpertProgramming/chapter7/code"
	"fmt"
)

// 检测单行输出
func ExampleSayHello() {
	code.SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	code.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

// 检测乱序输出
func ExamplePrintNames() {
	code.PrintNames()
	// Unordered output:
	// Jim
	// Bob
	// Tom
	// Sue
}

/*
	检测int数据返回
*/
func ExampleOutInt() {
	num := code.OutInt()
	fmt.Println(num)
	// output:
	// 1
}
