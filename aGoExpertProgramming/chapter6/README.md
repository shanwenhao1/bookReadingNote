# 反射

[官方讲解](https://blog.golang.org/laws-of-reflection)

## 反射概念
* 反射提供一种让程序检查自身结构的能力： 反射是一种检查interface变量的底层类型和值的机制
* 反射是困惑的源泉

想深入了解反射，必须深入理解类型和接口概念
* 接口类型的变量可以存储任何实现该接口的值
    - 空interface类型变量可以存放所有值
```go
// i和j的类型是不同的, 尽管二者底层类型一样
type  Myint  int
var  i  int
var  j  Myint
```


## 反射三定律
* 反射第一定律: 反射可以将interface类型变量转换成反射对象
```go
package  main
import  (
    "fmt"
    "reflect"
)


/*
    TypeOf()和ValueOf()接受的参数都是 interface{} 类型的，也即x值是被转成了interface传入的
*/
func  main()  {
    var  x  float64  =  3.4       
    t := reflect.TypeOf(x)    //t  is  reflext.Type
    fmt.Println("type:",  t)

    v  :=  reflect.ValueOf(x)  //v  is  reflext.Value
    fmt.Println("value:",  v)
    }
```
* 反射第二定律：反射可以将反射对象还原成interface对象
```go
package main
import (
    "fmt"
    "reflect"
)
func main() {
    /*
        对象x转换成反射对象v，v又通过Interface()接口转换成interface对象，
        interface对象通过.(float64)类型断言获取float64类型的值
    */
    var x float64 = 3.4
    v := reflect.ValueOf(x) //v is reflext.Value
    var y float64 = v.Interface().(float64)
    fmt.Println("value:", y)
}
```
* 反射第三定律：反射对象可修改，value值必须是可设置的
```go
package main
import (
    "reflect"
    "fmt"
)
func main() {
    var x float64 = 3.4
    /* 
        v := reflect.ValueOf(x)
        v.SetFloat(7.1) // Error: will panic.
            panic的原因是因为: 传入reflect.ValueOf()函数的其实是x的值，而非x本身。
                             即通过v修改其值是无法影响x的，也即是无效的修改，所以golang会报错
    */
    v := reflect.ValueOf(&x)
    v.Elem().SetFloat(7.1)
    //  reflect.Value 提供了 Elem() 方法，可以获得指针向指向的 value
    fmt.Println("x :", v.Elem().Interface())    // x : 7.1
}
```

TODO: 更加深入的了解反射