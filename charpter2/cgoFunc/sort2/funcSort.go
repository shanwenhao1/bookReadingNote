package sort2

/*
#include <stdlib.h>

typedef int(*qsort_cmp_func_t)(const void* a, const void* b);
extern int _cgo_qsort_compare(void* a, void* b);
*/
import "C"

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

type CompareFunc C.qsort_cmp_func_t

var go_qsort_compare_info struct {
	base     unsafe.Pointer
	elemNum  int
	elemSize int
	less     func(a, b int) bool
	sync.Mutex
}

/*
	新的比较函数
	改进: 使用闭包函数做为比较函数(由于闭包函数无法导出为C语言函数,  这边使用Go构造一个可以导出为C语言的代理函数)
*/
//export _cgo_qsort_compare
func _cgo_qsort_compare(a, b unsafe.Pointer) C.int {
	var (
		base     = uintptr(go_qsort_compare_info.base)
		elemSize = uintptr(go_qsort_compare_info.elemSize)
	)

	i := int((uintptr(a) - base) / elemSize)
	j := int((uintptr(b) - base) / elemSize)

	switch {
	case go_qsort_compare_info.less(i, j): // v[i] < v[j]
		return -1
	case go_qsort_compare_info.less(j, i): // v[i] > v[j]
		return +1
	default:
		return 0
	}
}

/*
	消除了对unsafe包的依赖
*/
func Slice(slice interface{}, less func(a, b int) bool) {
	// 通过反射获取qsort()函数所需要的切片信息
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("qsort called with non-slice value of type %T", slice))
	}
	if sv.Len() == 0 {
		return
	}

	// 上下文信息加锁
	go_qsort_compare_info.Lock()
	defer go_qsort_compare_info.Unlock()

	defer func() {
		go_qsort_compare_info.base = nil
		go_qsort_compare_info.elemNum = 0
		go_qsort_compare_info.elemSize = 0
		go_qsort_compare_info.less = nil
	}()

	go_qsort_compare_info.base = unsafe.Pointer(sv.Index(0).Addr().Pointer())
	go_qsort_compare_info.elemNum = sv.Len()
	go_qsort_compare_info.elemSize = int(sv.Type().Elem().Size())
	go_qsort_compare_info.less = less

	C.qsort(
		go_qsort_compare_info.base,
		C.size_t(go_qsort_compare_info.elemNum),
		C.size_t(go_qsort_compare_info.elemSize),
		C.qsort_cmp_func_t(C._cgo_qsort_compare),
	)
}
