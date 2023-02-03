package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	/*
		// 使用样例: https://zhuanlan.zhihu.com/p/150985877
		Minutes：分钟，取值范围[0-59]，支持特殊字符* / , -；
		Hours：小时，取值范围[0-23]，支持特殊字符* / , -；
		Day of month：每月的第几天，取值范围[1-31]，支持特殊字符* / , - ?；
		Month：月，取值范围[1-12]或者使用月份名字缩写[JAN-DEC]，支持特殊字符* / , -；
		Day of week：周历，取值范围[0-6]或名字缩写[JUN-SAT]，支持特殊字符* / , - ?。
			*：使用*的域可以匹配任何值，例如将月份域（第 4 个）设置为*，表示每个月；
			/：用来指定范围的步长，例如将小时域（第 2 个）设置为3-59/15表示第 3 分钟触发，以后每隔 15 分钟触发一次，因此第 2 次触发为第 18 分钟，第 3 次为 33 分钟。。。直到分钟大于 59；
			,：用来列举一些离散的值和多个范围，例如将周历的域（第 5 个）设置为MON,WED,FRI表示周一、三和五；
			-：用来表示范围，例如将小时的域（第 1 个）设置为9-17表示上午 9 点到下午 17 点（包括 9 和 17）；
			?：只能用在月历和周历的域中，用来代替*，表示每月/周的任意一天。

		example:
			30 * * * *：分钟域为 30，其他域都是*表示任意。每小时的 30 分触发；
			30 3-6,20-23 * * *：分钟域为 30，小时域的3-6,20-23表示 3 点到 6 点和 20 点到 23 点。3,4,5,6,20,21,22,23 时的 30 分触发；
			0 0 1 1 *：1（第 4 个） 月 1（第 3 个） 号的 0（第 2 个） 时 0（第 1 个） 分触发。
	*/
	cron2 := cron.New() //创建一个cron实例

	//执行定时任务（每两分钟执行一次）
	id, err := cron2.AddFunc("*/2 0 * * *", print5)
	//_, err := cron2.AddFunc("@every 5s", print5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("cron task: %d is up...", id))
	id2, err2 := cron2.AddFunc("*/1 * * * *", print6)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(fmt.Sprintf("cron task: %d is up...", id2))

	//启动/关闭
	cron2.Start()
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

// 执行函数
func print5() {
	fmt.Println("每两分钟执行一次cron task", time.Now())
}

// 执行函数
func print6() {
	fmt.Println("每分钟执行一次cron task", time.Now())
}
