package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

const (
	key = "xxxxxxxx-6d6b-45fc-a31e-c37ec0f526ad"
)

func GenerateMd5(str string) string {
	data := []byte(str)
	hmd5 := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", hmd5)
	return md5Str
}

// 创建会话时, 验证Md5
func SessionMd5(timeStamp int) (string, string) {
	timeStr := strconv.Itoa(timeStamp)
	seSs := key + timeStr
	data := []byte(seSs)
	hmd5 := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", hmd5)
	return md5Str, seSs
}

func VerifySeSs(md5 string, timeStamp int) bool {
	timeNow := time.Now().Unix()
	rTime := int64(timeStamp)

	ret := timeNow - rTime
	if ret < -600 || ret > 600 {
		//log.LogWithTag(log.ERROR, "MD5", "Convert Time error( Server Time: %d, Client Time: %d)", timeNow, rTime)
		return false
	}

	serMd5, _ := SessionMd5(timeStamp)
	if serMd5 != md5 {
		return false
	}
	return true
}

// 本项目相关的md5验证
func AuthAccess(timestamp int) string {
	key := "xxxxxxxx-2b13-4621-91b1-c8cd42546645"
	secret := "xxxxxxxx-dc17-48d4-a2f8-8dc1208b1b54"
	str := key + secret + strconv.Itoa(timestamp)
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
