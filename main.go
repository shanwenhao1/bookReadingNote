package main

import (
	"bookReadingNote/infra/db"
	"bookReadingNote/infra/log"
	"bookReadingNote/infra/redis"
	"bookReadingNote/infra/tool/file/xmlFile"
	"fmt"
)

/*
	服务启动前的准备, 包括redis连接
*/
func InitServer() {
	/*
		---------------------------本地配置文件方式加载配置----------------------------------
	*/
	var (
		err error
	)

	// 初始化mysql
	var sqlCfg = new(xmlFile.MysqlConfig)
	err = xmlFile.XmlParse("config/dbConfig.xml", new(xmlFile.MysqlConfig), sqlCfg)
	if err != nil {
		log.Tag(log.ERROR, log.InitSer, "读取数据库配置文件异常:[%v]", err)
		panic(err)
	}
	db.InitMysql(sqlCfg)
	// 同步数据库表结构  (domain.AutoMigrate负责调用各个domain模块下的db migrate函数)
	// domain.AutoMigrate()

	// 加载redis 配置文件
	var rdCfg = new(xmlFile.RDConfig)
	err = xmlFile.XmlParse("config/redisConfig.xml", new(xmlFile.RDConfig), rdCfg)
	if err != nil {
		panic(err)
	}
	//fmt.Println(rdCfg.(*xmlFile.RDConfig))
	// redis 连接初始化
	redis.InitRedis(rdCfg)
}

func redisTest() {
	rd := new(redis.RDSHandle)
	err := rd.SetNX("test_2", 2)
	fmt.Println(err)
}

func main() {
	// 本地日志初始化
	log.InitializedLog4go("config/log4go.xml")

	InitServer()
	redisTest()
}
