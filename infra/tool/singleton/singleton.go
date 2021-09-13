package singleton

import "sync"

/*
	go 单例实现示例

	单例实现思路:
			一个类能返回对象一个引用(永远是同一个)和一个获得该实例的方法(必须是静态方法). 调用时如果类持有的引用不为空就返回该引用,
		否则创建该类的实例并将实例的引用赋予该类保持的引用.
			该类的构造函数定义为私有方法, 这样其他处的代码就只能通过该类提供的静态方法来得到该类的唯一实例

	单例使用场景:
				网站计数器
				应用程序的日志应用(一个实例去操作)
				配置对象的读取
				数据库连接池
				多线程的线程池
				等等场景

	单例模式应用场景一般发生在以下场景中:
			资源共享的情况下, 避免由于资源操作时导致的性能或损耗等
			控制资源的情况下, 方便资源之间的互相通信.
*/
type singleton struct {
}

var (
	instance *singleton
	once     sync.Once
)

func Instance() *singleton {
	// once.Do 调用的函数只执行一次
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
