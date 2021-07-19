package dbTool

import (
	"bookReadingNote/infra/db"
	"bookReadingNote/infra/tool/file/xmlFile"
	"bookReadingNote/infra/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	test       = "test"
	test2      = "test2"
	testChange = "testChange"
)

type Test struct {
	Id      int       `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;comment:'Id(唯一)'"`
	Index   int       `json:"index" gorm:"column:index;unique_index;comment:'唯一索引'"`
	TestStr string    `json:"test_str" gorm:"column:test_str"`
	Time    time.Time `json:"time" gorm:"column:time;comment:'记录时间'"`
}

/*
	测试前需要进行初始化(包括db client初始化、数据库test的创建)
*/
func testInit(a *assert.Assertions) {
	var (
		err       error
		dbTestCfg *xmlFile.MysqlConfig
	)
	dbTestCfg = &xmlFile.MysqlConfig{
		DbName:     "test",
		DbUser:     "root",
		DbPwd:      "123456",
		DbUrl:      "tcp(192.168.1.89:3306)",
		DbMaxConn:  2000,
		DbMaxIdle:  200,
		DbLogModel: false,
	}
	db.InitMysql(dbTestCfg)
	db.GetDS().AutoMigrate(&Test{})

	// 清楚过期数据
	err = DataTool{DbOne: db.GetDS()}.Delete(&Test{}, map[string]interface{}{"test_str": test})
	a.Equal(err, nil, "db Delete function error, err:[%v]", err)
	err = DataTool{DbOne: db.GetDS()}.Delete(&Test{}, map[string]interface{}{"test_str": test2})
	a.Equal(err, nil, "db Delete function error, err:[%v]", err)
	err = DataTool{DbOne: db.GetDS()}.Delete(&Test{}, map[string]interface{}{"test_str": testChange})
	a.Equal(err, nil, "db Delete function error, err:[%v]", err)
}

func TestDataTool_Create(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
	)
	newTest.Index = 1
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	// 重复插入(unique数据冲突模拟)
	err = dt.Create(&newTest)
	a.NotEqual(nil, err, "db Create function error, err:[%v]", err)
}

func TestDataTool_CreateOrm(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err      error
		newTest  Test
		newTest2 Test
		newTest3 Test
	)
	newTest.Index = 1
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	newTest2.Index = 2
	newTest2.TestStr = test2
	newTest2.Time = utils.GetCurTimeUtc()

	tx := dt.DbOne.Begin()
	err = dt.CreateOrm(tx, &newTest)
	if err != nil {
		tx.Rollback()
		a.Fail("db CreateOrm function error, err:[%v] ", err)
		return
	}
	err = dt.CreateOrm(tx, &newTest2)
	if err != nil {
		tx.Rollback()
		a.Fail("db CreateOrm function error, err:[%v] ", err)
		return
	}
	tx.Commit()

	// 模拟orm事务需要回滚的情况
	tx = dt.DbOne.Begin()
	newTest3.Index = 1
	newTest3.TestStr = test2
	newTest3.Time = utils.GetCurTimeUtc()
	err = dt.CreateOrm(tx, &newTest3)
	if err != nil {
		tx.Rollback()
		return
	}
	a.NotEqual(nil, err, "db CreateOrm function error, err:[%v] ", err)
}

func TestDataTool_Save(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Save(&newTest)
	a.Equal(nil, err, "db Save function error, err:[%v]", err)
	newTest.TestStr = testChange
	err = dt.Save(&newTest)
	a.Equal(nil, err, "db Save function error, err:[%v]", err)
}

func TestDataTool_SaveOrm(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err      error
		newTest  Test
		newTest2 Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	tx := dt.DbOne.Begin()
	err = dt.SaveOrm(tx, &newTest)
	if err != nil {
		tx.Rollback()
		a.Fail("db SaveOrm function error, err:[%v]", err)
		return
	}
	tx.Commit()

	// 模拟事务需要回滚的情况
	tx = dt.DbOne.Begin()
	newTest2.Index = 3
	newTest2.TestStr = testChange
	newTest2.Time = utils.GetCurTimeUtc()
	err = dt.SaveOrm(tx, &newTest2)
	if err != nil {
		tx.Rollback()
	}
	a.NotEqual(nil, err, "db SaveOrm function error, err:[%v]", err)
}

func TestDataTool_Delete(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.Delete(&Test{}, map[string]interface{}{"test_str": test})
	a.Equal(nil, err, "db Delete function error, err:[%v]", err)
}

func TestDataTool_DeleteOrm(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	tx := dt.DbOne.Begin()
	err = dt.CreateOrm(tx, &newTest)
	if err != nil {
		tx.Rollback()
		a.Fail("db Create function error, err:[%v]", err)
		return
	}
	err = dt.DeleteOrm(tx, &Test{}, map[string]interface{}{"test_str": test})
	if err != nil {
		tx.Rollback()
		a.Fail("db Delete function error, err:[%v]", err)
		return
	}
	tx.Commit()
}

func TestDataTool_DeleteSqlIn(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err      error
		newTest  Test
		newTest2 Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	newTest2.Index = 4
	newTest2.TestStr = "test2"
	newTest2.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.Create(&newTest2)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.DeleteSqlIn(&Test{}, "test_str in (?)", []string{test, "test2"})
	a.Equal(nil, err, "db DeleteSqlIn function error, err:[%v]", err)
}

func TestDataTool_DeleteSqlInOrm(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err      error
		newTest  Test
		newTest2 Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	newTest2.Index = 4
	newTest2.TestStr = "test2"
	newTest2.Time = utils.GetCurTimeUtc()

	tx := dt.DbOne.Begin()
	err = dt.CreateOrm(tx, &newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.CreateOrm(tx, &newTest2)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.DeleteSqlInOrm(tx, &Test{}, "test_str in (?)", []string{test, "test2"})
	a.Equal(nil, err, "db DeleteSqlInOrm function error, err:[%v]", err)
	tx.Commit()
}

func TestDataTool_Count(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   Test
		count   int
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.Count(&count, &testR, "test_str = ?", "test")
	a.Equal(nil, err, "db Get function error, err:[%v]", err)
	// 时间精确度不一样
	testR.Time = newTest.Time
	a.Equal("", testR.TestStr, "db Get function error, err: Record not same")
	a.Equal(1, count, "db Get function error, err: Record count error")
}

func TestDataTool_Get(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.Get(&testR, map[string]interface{}{"test_str": test})
	a.Equal(nil, err, "db Get function error, err:[%v]", err)
	// 时间精确度不一样
	testR.Time = newTest.Time
	a.Equal(newTest, testR, "db Get function error, err: Record not same")
}

func TestDataTool_GetWithPage(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   []Test
		pg      PageModel
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	pg.Page = 1
	pg.PageSize = 10
	err = dt.GetWithPage(&testR, map[string]interface{}{"test_str": test}, pg, "id")
	a.Equal(nil, err, "db Get function error, err:[%v]", err)
	a.Equal(1, len(testR), "db GetWithPage function error")
	// 时间精确度不一样
	testR[0].Time = newTest.Time
	a.Equal(newTest, testR[0], "db GetWithPage function error, err: Record not same")
}

func TestDataTool_GetSqlIn(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   []Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.GetSqlIn(&testR, "test_str in (?)", []string{test})
	a.Equal(nil, err, "db GetSqlIn function error, err:[%v]", err)
	a.Equal(1, len(testR), "db GetSqlIn function error")
	// 时间精确度不一样
	testR[0].Time = newTest.Time
	a.Equal(newTest, testR[0], "db GetSqlIn function error, err: Record not same")
}

func TestDataTool_GetSqlOr(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   Test
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	err = dt.GetSqlOr(&testR, map[string]interface{}{"test_str": ""}, map[string]interface{}{"index": 3})
	a.Equal(nil, err, "db GetSqlOr function error, err:[%v]", err)
	// 时间精确度不一样
	testR.Time = newTest.Time
	a.Equal(newTest, testR, "db GetSqlOr function error, err: Record not same")

	err = dt.GetSqlOr(&testR, map[string]interface{}{"test_str": test}, map[string]interface{}{"index": 1})
	a.Equal(nil, err, "db GetSqlOr function error, err:[%v]", err)
	// 时间精确度不一样
	testR.Time = newTest.Time
	a.Equal(newTest, testR, "db GetSqlOr function error, err: Record not same")
}

func TestDataTool_GetSqlBetween(t *testing.T) {
	a := assert.New(t)
	testInit(a)

	dt := DataTool{
		DbOne: db.GetDS(),
	}
	var (
		err     error
		newTest Test
		testR   []Test
		beginT  time.Time
		endT    time.Time
	)
	newTest.Index = 3
	newTest.TestStr = test
	newTest.Time = utils.GetCurTimeUtc()

	err = dt.Create(&newTest)
	a.Equal(nil, err, "db Create function error, err:[%v]", err)
	beginT = utils.GetAnotherTime(newTest.Time, utils.TimeFormat{Year: 0, Month: 0, Day: -1, Hour: 0, Minute: 0, Seconds: 0})
	endT = utils.GetAnotherTime(newTest.Time, utils.TimeFormat{Year: 0, Month: 0, Day: 0, Hour: 1, Minute: 0, Seconds: 0})
	err = dt.GetSqlBetween(&testR, map[string]interface{}{"test_str": test}, "time BETWEEN ? AND ?", beginT, endT)
	a.Equal(nil, err, "db GetSqlOr function error, err:[%v]", err)
	a.Equal(1, len(testR), "db GetWithPage function error")
	// 时间精确度不一样
	testR[0].Time = newTest.Time
	a.Equal(newTest, testR[0], "db GetSqlOr function error, err: Record not same")

	beginT = utils.GetAnotherTime(newTest.Time, utils.TimeFormat{Year: 0, Month: 0, Day: -1, Hour: 0, Minute: 0, Seconds: 0})
	endT = utils.GetAnotherTime(newTest.Time, utils.TimeFormat{Year: 0, Month: 0, Day: 0, Hour: -1, Minute: 0, Seconds: 0})
	err = dt.GetSqlBetween(&testR, map[string]interface{}{"test_str": test}, "time BETWEEN ? AND ?", beginT, endT)
	a.Equal(nil, err, "db GetSqlOr function error, err:[%v]", err)
	a.Equal(0, len(testR), "db GetWithPage function error")
}
