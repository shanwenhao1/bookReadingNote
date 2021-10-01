# 常见数据结构实现原理

## channel

```go
type	hchan	struct	{
	qcount			uint					//	当前队列中剩余元素个数.
    dataqsiz	    uint					//	环形队列长度，即可以存放的元素个数.
    buf				unsafe.Pointer	        //	环形队列指针.					
    elemsize	    uint16					//	每个元素的大小.					
    closed			uint32					//	标识关闭状态.					
    elemtype	    *_type					//	元素类型.					
    sendx			uint					//	队列下标，指示元素写入时存放到队列中的位置.					
    recvx			uint			        //	队列下标，指示元素从队列的该位置读出.					
    recvq			waitq				    //	等待读消息的goroutine队列.					
    sendq			waitq				    //	等待写消息的goroutine队列.					
    lock	        mutex					//	互斥锁，chan不允许并发读写.	
}
```
channel在数据结构上由队列(长度在创建chan时已确定)、类型信息、goroutine等待队列组成.


## slice

```go
type    slice	struct	{
    array	        unsafe.Pointer
    len			    int
    cap			    int
}
```
* 又称为动态数组, 切片依托数组实现. 属于引用类型, 本质上是数组的指针. 
* slice在底层数组容量不足时会自动扩容(但由于内存重新分配的问题, 最好创建时就酌情考虑给定足够的容量, 以防止效率低下)
    - 扩容容量遵循以下规则:
        - 如果原slice容量小于1024, 则新slice容量将扩大为原来2倍
        - 如果slice容量大于等于1024， 则新slice容量将扩大为原来1.25倍
* 使用copy来复制slice(值复制, 新的slice指向新的数组, 不影响原slice)
```go
package main

import "fmt"

func main() {
    a := []int{0, 1, 2}
    s := make([]int, 3)
    copy(s, a)
    fmt.Println(a, s)   // [0 1 2] [0 1 2]
    s[0] = 11

    fmt.Println(a, s)   // [0 1 2] [11 1 2]
}
```


## map

```go
type	hmap	struct	{
	count           int	                //	当前保存的元素个数
    ...
    B               uint8		        //	指示bucket数组的大小
    ...
    buckets			unsafe.Pointer	    //	bucket数组指针，数组的大小为2^B
    ...
}

/*
    bucket数据结构
        每个bucket可以存储8个键值对
*/
type	bmap	struct	{
	tophash         [8]uint8	        //  存储哈希值的高8位
	data			byte[1]		        //  key	value数据:key/key/key/.../value/value/value...
	overflow	    *bmap			    //  溢出bucket的地址(指向下一个bucket)
	}
```
使用[哈希表](https://baike.baidu.com/item/%E5%93%88%E5%B8%8C%E8%A1%A8/5981869?fr=aladdin)
作为底层实现, 一个哈希表有多个哈希表节点, 即bucket(每个bucket保存了map中的一个或一组键值对)  

```go
// 使用内置delete删除map
delete(exampleMap, "mapKey")
```

## struct

```go
package main

import	(
    "reflect"
    "fmt"
)

type	Server	struct	{
    ServerName      string	    `key1:	"value1"	key11:"value11"`
    ServerIP        string	    `key2:	"value2"`
}

func	main()	{
    s	:=	Server{}
    st	:=	reflect.TypeOf(s)
   
    field1	:=	st.Field(0)
    fmt.Printf("key1:%v\n",	field1.Tag.Get("key1"))         // 输出 key1:value1
    fmt.Printf("key11:%v\n",	field1.Tag.Get("key11"))    // 输出 key11:value11
    filed2	:=	st.Field(1)
    fmt.Printf("key2:%v\n",	filed2.Tag.Get("key2"))         // 输出 key2:value2
}
```

struct 声明中允许字段附带`Tag`对字段做一些标记.
* Tag主要用于反射场景
* 常见的用法主要是用作JSON数据解析、ORM映射等
    

## iota


## string
```go
type	stringStruct	struct	{
    str	            unsafe.Pointer          // 字符串的首地址
    len	            int                     // 字符串的长度
}
```
string 是8比特字节的集合，通常但并不一定是UTF-8编码的文本
* string可以为空(长度为0), 但不会是nil
* string对象不可以修改

字符串构建过程是先跟据字符串构建stringStruct，再转换成string, 其源码如下
```go
func	gostringnocopy(str	*byte)	string	{	//	跟据字符串地址构建string
	ss	:=	stringStruct{str: unsafe.Pointer(str), len:	findnull(str)}	//	先构造stringStruct
	s	:=	*(*string)(unsafe.Pointer(&ss))								//	再将stringStruct转换成string
	return	s
}
```

string和byte切片都可以互相转换, 但都需要进行内存拷贝(但[]byte转换成string有时不需要, 直接返回string, 该string指针指向
切片的内存)