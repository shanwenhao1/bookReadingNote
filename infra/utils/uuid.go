package utils

import (
	"errors"
	"github.com/satori/go.uuid"
)

func NewUuid() string {
	return uuid.NewV4().String()
}

func GetMulUuid(num int) ([]string, error) {
	var (
		uuidS []string
	)
	if num < 1 {
		return uuidS, errors.New("wrong parameter")
	}
	for i := 0; i < num; i++ {
		uid := NewUuid()
		uuidS = append(uuidS, uid)
	}
	return uuidS, nil
}
