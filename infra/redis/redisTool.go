package redis

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RDSInterface interface {
	Set(key string, value interface{}) error
	SetNX(key string, value interface{}) error
	SetTemData(key string, value interface{}, timeOut time.Duration) error
	Get(key string, value interface{}) error
	Del(key string) error
}

type RDSHandle struct {
}

/*
	marshal data func
*/
func (handle *RDSHandle) marshalData(data interface{}) (string, error) {
	var (
		err     error
		dataB   []byte
		dataStr string
	)
	dataB, err = json.Marshal(data)
	if err != nil {
		return dataStr, err
	}
	dataStr = string(dataB)
	return dataStr, nil
}

/*
	redis store data
		before store the data, it's has been marshaled to string
		data been marshaled then store in redis

	if key is already exist, it's will be rewrite, didn't return error
	(if you want judge key is exist or not, please see SetNX)
*/
func (handle *RDSHandle) Set(key string, value interface{}) error {
	var (
		err     error
		dataStr string
	)
	dataStr, err = handle.marshalData(value)
	if err != nil {
		return err
	}
	err = GetRdS().rds.Set(ctx, key, dataStr, 0).Err()
	return err
}

/*
	redis store data, if key is exist, keep it value unchanged, else set key with value
*/
func (handle *RDSHandle) SetNX(key string, value interface{}) error {
	var (
		err     error
		dataStr string
	)
	dataStr, err = handle.marshalData(value)
	if err != nil {
		return err
	}
	err = GetRdS().rds.SetNX(ctx, key, dataStr, 0).Err()
	return err
}

/*
	redis store data and set the data timeout time.
		after timeOuT, the data will be clean
			timeOut example: 30 * time.Minute, it's means data will be overdue after 30 minutes
*/
func (handle *RDSHandle) SetTemData(key string, value interface{}, timeOut time.Duration) error {
	var (
		err     error
		dataStr string
	)
	dataStr, err = handle.marshalData(value)
	if err != nil {
		return err
	}
	err = GetRdS().rds.Set(ctx, key, dataStr, timeOut).Err()
	return err
}

/*
	redis get data by key
*/
func (handle *RDSHandle) Get(key string, value interface{}) error {
	var (
		err error
		val string
	)
	val, err = GetRdS().rds.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), value)
	return err
}

/*
	delete data by key
*/
func (handle *RDSHandle) Del(key string) error {
	var err error
	err = GetRdS().rds.Del(ctx, key).Err()
	return err
}

func IsNull(err error) bool {
	if err == redis.Nil {
		return true
	}
	return false
}
