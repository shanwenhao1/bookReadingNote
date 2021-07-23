# Go汇编语言

因为Go汇编相较于程序员来说好处有两个: 一是突破框架限制, 而是挖掘极致的性能.  因此我就不再过于深入(以后有空再了解)

## 快速入门
官方自带的Go汇编语言[入门教程](https://golang.org/doc/asm)

使用`go tool compile -S *.go`查看指定文件的汇编代码

### Go汇编常用特殊标识符
- `TEXT`: 函数标识符
    ```cgo
    // TEXT指令 函数名 可选的flags标志 函数帧大小 可选的函数参数大小
    // (SB)表示函数名相对于伪寄存器SB的偏移量(symbol(SB)作为全局标志符, 一般基于伪寄存器SB的相对位置)
    //  framesize 表示函数的局部变量所需栈空间
    TEXT symbol(SB), [flags,] $framesize[-argsize]
    ```
    ```cgo
    //go:nosplit     //该注释会禁止汇编器为汇编函数插入栈分裂的代码
    // func Swap(a, b int) (int, int)
    TEXT ·Swap(SB), NOSPLIT, $0-32
    ```