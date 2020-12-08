package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 获取上一级目录
func GetParentDir(dir string) string {
	return substr(dir, 0, strings.LastIndex(dir, "/"))
}

// 获取上n级目录, num=1: 上一级目录
func GetParentDirM(dir string, num int) (string, error) {
	length := strings.LastIndex(dir, "/")
	if num >= length {
		return "", errors.New("超出目录")
	}
	needIndex := length - (num - 1)
	path := substr(dir, 0, needIndex)
	return path, nil
}

// 获取当前目录
func GetCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
