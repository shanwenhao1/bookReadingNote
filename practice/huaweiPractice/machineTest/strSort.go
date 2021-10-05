package machineTest

import (
	"fmt"
	"sort"
	"strings"
)

/*
	将输入的字符串排序后返回(空格作为间隔符)
*/
func StrSort(str string) string {
	var resS string
	var strS = strings.Fields(str)

	fmt.Println(sort.StringSlice(strS))
	// sort.StringSlice 将数据转换为StringSlice类型(本质上还是[]string), 因为
	// sort.Sort可以排序任何实现了 sort.Inferface 接口的对象
	sort.Sort(sort.StringSlice(strS))
	for pos, data := range strS {
		if pos == 0 {
			resS = data
			continue
		}
		resS = fmt.Sprintf("%s %s", resS, data)
	}
	return resS
}
