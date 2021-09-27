package code

/*
	MakeSliceWithPreAlloc:
		构造一个容量为100000的切片
		不预先分配内存
*/
func MakeSliceWithoutAlloc() []int {
	var newSlice []int

	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}

	return newSlice
}

/*
	MakeSliceWithPreAlloc:
		构造一个容量为100000的切片
		预先分配内存
*/
func MakeSliceWithPreAlloc() []int {
	var newSlice []int

	newSlice = make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		newSlice = append(newSlice, i)
	}

	return newSlice
}
