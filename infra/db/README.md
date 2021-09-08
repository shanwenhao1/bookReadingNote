# mysql 使用说明

- 使用前请参考[dbConfig.xml](../../config/dbConfig.xml)配置文件, 使用[mysql.go](mysql.go)
中的`InitMysql`函数进行初始化数据库连接
```go
import (
	"bookReadingNote/infra/db"
	"bookReadingNote/infra/log"
	"bookReadingNote/infra/tool/file/xmlFile"
)


	// 初始化mysql
	var sqlCfg = new(xmlFile.MysqlConfig)
	err := xmlFile.XmlParse("config/dbConfig.xml", new(xmlFile.MysqlConfig), sqlCfg)
	if err != nil {
		log.Tag(log.ERROR, log.InitSer, "读取数据库配置文件异常:[%v]", err)
		panic(err.Error())
	}
	db.InitMysql(sqlCfg)
```
- 使用`GetDS().AutoMigrate`同步数据结构
```go
import (
	"bookReadingNote/infra/db"
	"bookReadingNote/infra/log"
	"bookReadingNote/infra/tool/file/xmlFile"
)

db.GetDS().AutoMigrate({})
```

[Code](dbTool/db_base.go)