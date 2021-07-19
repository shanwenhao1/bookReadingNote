//go:generate mockgen -destination mock_dbTool/mock_db_base.go bookReadingNote/infra/db/dbTool DataBase
package dbTool

import (
	"bookReadingNote/infra/log"
	"bookReadingNote/infra/utils"
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	RENOTFOUND string = "record not found"
)

type PageModel struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type DataBase interface {
	GetOrm() *gorm.DB
	Commit(tx *gorm.DB)
	Rollback(tx *gorm.DB)
	Create(model interface{}) error
	CreateOrm(tx *gorm.DB, model interface{}) error
	Save(model interface{}) error
	SaveOrm(tx *gorm.DB, model interface{}) error
	Delete(model interface{}, condition map[string]interface{}) error
	DeleteOrm(tx *gorm.DB, model interface{}, condition map[string]interface{}) error
	DeleteSqlIn(model interface{}, sqlField string, sqlIn interface{}) error
	DeleteSqlInOrm(tx *gorm.DB, model interface{}, sqlField string, sqlIn interface{}) error
	Count(count interface{}, model interface{}, sqlField string, sqlArgs ...interface{}) error
	Get(model interface{}, conditions map[string]interface{}) error
	GetWithPage(model interface{}, condition map[string]interface{}, page PageModel, order ...string) error
	GetSqlIn(model interface{}, sqlField string, sqlIn interface{}) error
	GetSqlOr(model interface{}, filter ...map[string]interface{}) error
	GetSqlBetween(model interface{}, condition map[string]interface{}, sqlField string, before interface{}, after interface{}) error
}

// 数据库通用方法
type DataTool struct {
	DbOne *gorm.DB
}

func (data DataTool) GetOrm() *gorm.DB {
	return data.DbOne.Begin()
}

func (data DataTool) Rollback(tx *gorm.DB) {
	tx.Rollback()
}

func (data DataTool) Commit(tx *gorm.DB) {
	tx.Commit()
}

// 通用创建方法
func (data DataTool) Create(model interface{}) error {
	var err error
	err = data.DbOne.Create(model).Error
	return err
}

// 通用事务创建方法
func (data DataTool) CreateOrm(tx *gorm.DB, model interface{}) error {
	var err error
	err = tx.Create(model).Error
	return err
}

// 通用保存方法(不存在则会创建, 存在则会更新记录)
func (data DataTool) Save(model interface{}) error {
	var err error
	err = data.DbOne.Save(model).Error
	return err
}

// 通用事务保存方法
func (data DataTool) SaveOrm(tx *gorm.DB, model interface{}) error {
	var err error
	err = tx.Save(model).Error
	return err
}

// 通用删除方法
func (data DataTool) Delete(model interface{}, condition map[string]interface{}) error {
	var err error
	err = data.DbOne.Model(model).Delete(model, condition).Error
	return err
}

// 通用事务删除方法
func (data DataTool) DeleteOrm(tx *gorm.DB, model interface{}, condition map[string]interface{}) error {
	var err error
	err = tx.Model(model).Delete(model, condition).Error
	return err
}

/*
	sql in 删除
*/
func (data DataTool) DeleteSqlIn(model interface{}, sqlField string, sqlIn interface{}) error {
	var (
		err error
	)
	// 示例"name in (?)", []string{"jinZhu", "jinZhu 2"}
	err = data.DbOne.Where(sqlField, sqlIn).Delete(model).Error
	return err
}

func (data DataTool) DeleteSqlInOrm(tx *gorm.DB, model interface{}, sqlField string, sqlIn interface{}) error {
	var (
		err error
	)
	// 示例"name in (?)", []string{"jinZhu", "jinZhu 2"}
	err = tx.Where(sqlField, sqlIn).Delete(model).Error
	return err
}

// 获取记录总数
func (data DataTool) Count(count interface{}, model interface{}, sqlField string, sqlArgs ...interface{}) error {
	var err error
	err = data.DbOne.Model(model).Where(sqlField, sqlArgs).Count(count).Error
	return err
}

// 通用查询方法
func (data DataTool) Get(model interface{}, condition map[string]interface{}) error {
	var err error
	err = data.DbOne.Find(model, condition).Error
	return err
}

// 通用按页查询方法(支持order排序). 当不需要排序的时候, order不传即可(order可支持多个排序)
func (data DataTool) GetWithPage(model interface{}, condition map[string]interface{}, page PageModel, order ...string) error {
	var (
		err      error
		orderStr string
		strTool  utils.StrUtil
	)
	if len(order) == 0 {
		err = data.DbOne.Offset(page.Page*page.PageSize-page.PageSize).Limit(page.PageSize).Find(model, condition).Error
	} else {
		for _, _order := range order {
			if strTool.CompareEqual(orderStr, "") {
				orderStr = _order
				continue
			}
			orderStr = fmt.Sprintf("%s, %s", orderStr, _order)
		}
		err = data.DbOne.Order(orderStr).Offset(page.Page*page.PageSize-page.PageSize).Limit(page.PageSize).Find(model, condition).Error
	}
	return err
}

// 通用简单sql - in查询, sqlIn为[]interface形式的数据
func (data DataTool) GetSqlIn(model interface{}, sqlField string, sqlIn interface{}) error {
	var (
		err error
	)
	// 示例"name in (?)", []string{"jinZhu", "jinZhu 2"}
	err = data.DbOne.Where(sqlField, sqlIn).Find(model).Error
	return err
}

// 通用简单sql查询(TODO 测试语法是否正确)
func (data DataTool) GetSql(model interface{}, sqlField string, sqlArgs ...interface{}) error {
	var (
		err error
	)
	// 示例"name in (?)", []string{"jinZhu", "jinZhu 2"}
	err = data.DbOne.Where(sqlField, sqlArgs).Find(model).Error
	return err
}

// 通用sql--or查询
func (data DataTool) GetSqlOr(model interface{}, filter ...map[string]interface{}) error {
	var (
		err error
		db  *gorm.DB
	)
	// 示例map[string]interface{}{"name": "jin zhu"}, map[string]interface{}{"name": "jin zhu2"}
	// err = data.DbOne.Where(filter[0]).Or(filter[1]).Find(model).Error
	for ps, _filter := range filter {
		if ps == 0 {
			db = data.DbOne.Where(_filter)
		} else {
			db = db.Or(_filter)
		}
	}
	err = db.Find(model).Error
	return err
}

// 通用sql-between查询
func (data DataTool) GetSqlBetween(model interface{}, condition map[string]interface{}, sqlField string, before interface{}, after interface{}) error {
	var (
		err error
	)
	err = data.DbOne.Where(sqlField, before, after).Find(model, condition).Error
	return err
}

//判断是否为空
func IsNull(err error) bool {
	var strTool utils.StrUtil
	if strTool.CompareEqual(err.Error(), RENOTFOUND) {
		return true
	}
	return false
}

// 断言为空, 如果不是则记录错误日志. 返回判定结果
func AssertNullWithLog(err error, logTag string) bool {
	if err.Error() == RENOTFOUND {
		return true
	} else {
		log.Tag(log.ERROR, logTag, "数据库错误: [%v]", err)
	}
	return false
}

// 判断是否为空, 如果是则记录错误日志. 返回判定结果
func AssertNotNullWithLog(err error, logTag string, msg string) bool {
	if err.Error() == RENOTFOUND {
		log.Tag(log.ERROR, logTag, msg)
		return true
	}
	return false
}
