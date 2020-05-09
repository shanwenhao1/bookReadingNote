package sort

//extern int go_qsort_compare(void* a, void* b);
import "C"
import "unsafe"
import "fmt"

/*
go 实现的比较函数,  供C使用
*/
//export go_qsort_compare
func go_qsort_compare(a, b unsafe.Pointer) C.int {
	pa, pb := (*C.int)(a), (*C.int)(b)
	return C.int(*pa - *pb)
}

func SortRun() {
	values := []int32{42, 9, 101, 95, 27, 25}
	Sort(unsafe.Pointer(&values[0]), len(values), int(unsafe.Sizeof(values[0])), CompareFunc(C.go_qsort_compare))
	fmt.Println(values)
}
