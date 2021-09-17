package code

/*
	主函数拥有匿名返回值, 返回使用本地或全局变量, defer语句可引用到返回值, 但不会改变返回值
*/
func Foo() int {
	var i int = 0

	defer func() {
		i++
	}()

	return i
}

/*
	defer延迟函数操作主函数的具名返回值的典范

    该函数返回2
        执行步骤:
                result = i  (因为return不是原子操作, 只代理汇编指令ret(跳转程序执行), 此为第一步)
                result++    (defer的执行)
                return      (return 语句的第二步)
*/
func Foo1() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}
