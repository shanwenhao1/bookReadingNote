package db

import (
	"bookReadingNote/infra/log"
	"bookReadingNote/infra/tool/file/xmlFile"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ds *gorm.DB

func InitMysql(cfg xmlFile.XmlFile) {
	if _, ok := cfg.(*xmlFile.MysqlConfig); !ok {
		log.Tag(log.ERROR, log.InitSer, "Mysql init failed, cfg parameter error")
		panic(errors.New("Mysql init failed, cfg parameter error "))
	}
	sqlCfg := cfg.(*xmlFile.MysqlConfig)

	// sql 连接
	dbSql, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@%v/%v?charset=utf8&parseTime=True", sqlCfg.DbUser, sqlCfg.DbPwd, sqlCfg.DbUrl, sqlCfg.DbName))
	if err != nil {
		log.Tag(log.ERROR, log.InitSer, "初始化数据源异常:%v", err)
		panic(err)
	}

	dbSql.LogMode(sqlCfg.DbLogModel)
	dbSql.SingularTable(true)
	dbSql.DB().SetMaxOpenConns(sqlCfg.DbMaxConn)
	dbSql.DB().SetMaxIdleConns(sqlCfg.DbMaxIdle)

	// 自动同步表结构, 只新增字段, 不修改或删除
	// use-example: GetDS().AutoMigrate(your model{})
	//dbSql.AutoMigrate(models.Uuioe_User{})

	ds = dbSql
	log.Tag(log.INFO, log.InitSer, "数据源已初始化完成[最大打开连接数:%v,最大空闲连接数:%v]", sqlCfg.DbMaxConn, sqlCfg.DbMaxIdle)
}

func GetDS() *gorm.DB {
	return ds
}
