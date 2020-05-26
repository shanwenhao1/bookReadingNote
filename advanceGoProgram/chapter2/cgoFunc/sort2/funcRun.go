package sort2

import "fmt"

func SortRun() {
	values := []int32{42, 9, 101, 95, 27, 25}

	Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	fmt.Println(values)
}
