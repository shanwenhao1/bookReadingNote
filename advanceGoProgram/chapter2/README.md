# cgo

cgo 使用C/C++资源的三种形式:
- 直接使用源码
- 链接静态库: 静态链接,  不会产生额外的运行依赖, 也不会出现动态库特有的跨运行时资源管理的错误, 
但是静态库一般包含了全部的源代码(不同静态库可能出现符号冲突导致链接失败)
- 链接动态库: 动态库可以隔离不同动态库之间的关系, 减少链接时出现的符号冲突的风险

## 静态库和动态库调用

## Go导出成C静态库或C动态库:
- 导出成C静态库
    - 简单导出C函数
    ```go
    package main
    import "C"
    
    // CGO要求在main包中导出C函数
    func main(){
    }
    
    //export number_add_mod
    func number_add_mod(a, b, mod C.int) C.int{
        return (a + b) % mod
    }
    
    // 使用命令导出: $ go build -buildmode=c-archive -o number.a 
    ```
    - 测试导出的函数`_test_main.c`
    ```c
    #include "number.h"
    #include <stdio.h>
    
    int main(){
        int a = 10;
        int b = 5;
        int c = 12;
        
        int x = number_add_mod(a, b, c);
        printf("(%d + %d)%%%d = %d\n", a, b, c, x);
        
        return 0;
    }
    
    //$ gcc -o a.out _test_main.c number.a
    //$ ./a.out
    ```
- 导出成C动态库, 导出命令为: `go build -buildmode=c-shared -o number.so`, 可用`go help buildmode`
查看构建说明

- 导出非main包函数, 我们需要自己提供C函数对应的头文件(CGO无法为非main包的导出函数生成头文件)
```go
package number
import "C"

//export number_add_mod
func number_add_mod(a, b, mod C.int) C.int{
	return (a + b) % mod
}
```

```go
package main
import "C"

import (
	"fmt"
	_ "./number"
)
func main(){
	print("Done")
}

//export goPrintln
func goPrintln(s *C.char){
	fmt.Println("goPrintln:", C.GoString(s))
}

// $ go build -buildmode=c-archive -o main.a
// 此时main.h中并没有number包中的number_add_mod()函数的声明, 但是生成的静态库中是包含该函数的
```
```c
// 此时需要手工方式声明这两个函数
#include <stdio.h>

void goPrintln(char*);
int number_add_mod(int a, int b, int mod);

int main(){
    int a = 10;
    int b = 5;
    int c = 12;
    
    int x = number_add_mod(a, b, c);
    printf("(%d + %d)%%%d = %d\n", a, b, c, x);
    
    return 0;
}

//$ gcc -o a.out _test_main.c number.a
//$ ./a.out
```
