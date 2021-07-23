package sort

/*
#include <stdlib.h>

typedef int(*qsort_cmp_func_t)(const void* a, const void* b);
*/
import "C"
import "unsafe"

type CompareFunc C.qsort_cmp_func_t

/*
 此排序函数需要通过CGO实现C语言版本的比较函数go_qsort_compare
*/
func Sort(base unsafe.Pointer, num, size int, cmp CompareFunc) {
	C.qsort(base, C.size_t(num), C.size_t(size), C.qsort_cmp_func_t(cmp))
}
