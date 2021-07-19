package db

/*
	此文件暂时不用
*/
//
//import "scanPen/infra/log"
//
//const (
//	RENOTFOUND string = "record not found"
//)
//
////通用全部查询
//func QueryAll(model interface{}) error {
//	db := GetDS()
//	err := db.Find(model).Error
//	return err
//}
//
//// sql语句查询
//func QuerySql(model interface{}, sql string) error {
//	db := GetDS()
//	err := db.New().Raw(sql).Select("total").Scan(model).Error
//	return err
//}
//
////通用查询方法
//func Query(model interface{}, condtions map[string]interface{}) error {
//	db := GetDS()
//	err := db.Find(model, condtions).Error
//	return err
//}
//
///*
////通用查询方法(Test数据库)
//func QueryTestRecord(model interface{}, condtions map[string]interface{}) error {
//	db := GetTDS()
//	err := db.Find(model, condtions).Error
//	return err
//}
//
//// Test数据库保存方法
//func SaveT(model interface{}) error {
//	db := GetTDS()
//	err := db.Create(model).Error
//	return err
//}
//
//// Test数据库删除方法
//func DeleteT(model interface{}, condtions map[string]interface{}) error {
//	db := GetTDS()
//	err := db.Model(model).Delete(model, condtions).Error
//	return err
//}
// */
//
////通用查询排序方法
//func QueryOrderString(model interface{}, condtions map[string]interface{}, order_str string) error {
//	db := GetDS()
//	err := db.Order(order_str).Find(model, condtions).Error
//	return err
//}
//
////通用查询排序方法带数量
//func QueryOrderStringWithCount(model interface{}, condtions map[string]interface{}, order_str string, count *int64) error {
//	db := GetDS()
//	err := db.Order(order_str).Find(model, condtions).Count(count).Error
//	return err
//}
//
//// 通用查询列表方法
//func QueryList(model interface{}, condtions map[string]interface{}, pageModel map[string]int32) error {
//	db := GetDS()
//	var err error
//	if pageModel == nil {
//		err = db.Find(model, condtions).Error
//	} else {
//		size := pageModel["pageSize"]
//		page := pageModel["currentPage"]
//		err = db.Offset(page*size-size).Limit(size).Find(model, condtions).Error
//	}
//	return err
//}
//
////通用获取数量方法
//func QueryCount(model interface{}, condtions map[string]interface{}) (int, error) {
//	db := GetDS()
//	var value int
//	err := db.Model(model).Where(condtions).Count(&value).Error
//	return value, err
//}
//
////通用更新方法
//func Update(model interface{}, condtions map[string]interface{}) error {
//	db := GetDS()
//	err := db.Model(model).Where(condtions).Update(model).Error
//	return err
//}
//
////通用保存对象方法
//func Save(model interface{}) error {
//	db := GetDS()
//	err := db.Create(model).Error
//	return err
//}
//
////通用保存对象方法
//func Delete(model interface{}, condtions map[string]interface{}) error {
//	db := GetDS()
//	err := db.Model(model).Delete(model, condtions).Error
//	return err
//}
//
////判断是否为空
//func IsNull(err error) bool {
//	if err.Error() == RENOTFOUND {
//		return true
//	}
//	return false
//}
//
//// 断言为空, 如果不是则记录错误日志. 返回判定结果
//func AssertNullWithLog(err error, logTag string) bool {
//	if err.Error() == RENOTFOUND {
//		return true
//	}else{
//		log.Tag(log.ERROR, logTag, "数据库错误: [%v]", err)
//	}
//	return false
//}
//
//// 判断是否为空, 如果是则记录错误日志. 返回判定结果
//func AssertNotNullWithLog(err error, logTag string, msg string) bool {
//	if err.Error() == RENOTFOUND {
//		log.Tag(log.ERROR, logTag, msg)
//		return true
//	}
//	return false
//}
//
//
////通用查询方法
//func QueryOrder(order string, model interface{}, condtions map[string]interface{}) error {
//	db := GetDS()
//	err := db.Order(order).Limit(1).Find(model, condtions).Error
//	return err
//}
