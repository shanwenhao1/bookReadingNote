package channel

/*
	知识点:
			channel是用于golang goroutine之间的通信, 不可跨进程通信.
			自带读写锁, 同一时间只能有一个goroutine读或者写
			只能传递一种类型的值
*/
