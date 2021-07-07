package code

/*
	问题: 用两个栈实现队列
		描述: 用两个栈来实现一个队列，分别完成在队列尾部插入整数(push)和在队列头部删除整数(pop)的功能。
			  队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

		示例:
			输入: ["PSH1","PSH2","POP","POP"]
			返回: 1,2
				解析:
					"PSH1":代表将1插入队列尾部
					"PSH2":代表将2插入队列尾部
					"POP“:代表删除一个元素，先进先出=>返回1
					"POP“:代表删除一个元素，先进先出=>返回2
*/

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	// 每次pop的时候, 如果2栈中有数据则直接pop, 如果没有则将1栈数据置入2栈中, 并重置1栈后再pop
	if len(stack2) == 0 {
		for i := 0; i <= len(stack1)-1; i++ {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = make([]int, 0)
	}
	if len(stack2) > 0 {
		ret := stack2[0]
		stack2 = stack2[1:]
		return ret
	}
	return -1
}
