package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetCurTimeStamp(t *testing.T) {
	a := assert.New(t)
	timestamp := time.Now().Unix()
	a.LessOrEqual(timestamp, GetCurTimeStamp(false), "GetCurTimeStamp error")
	a.GreaterOrEqual(timestamp+1, GetCurTimeStamp(false), "GetCurTimeStamp error")

	a.LessOrEqual(timestamp*1000, GetCurTimeStamp(true), "GetCurTimeStamp error")
	a.GreaterOrEqual(timestamp*1000+1000, GetCurTimeStamp(true), "GetCurTimeStamp error")
}

func TestGetCurTimeStr(t *testing.T) {
	a := assert.New(t)

	a.Equal(time.Now().UTC().Format(baseFormat), GetCurTimeStr(), "GetCurTimeStr error")
}

func TestGetCurTimeUtc(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	a.GreaterOrEqual(time.Second*1, GetCurTimeUtc().Sub(timeNow), "GetCurTimeUtc error")
	a.LessOrEqual(time.Second*0, GetCurTimeUtc().Sub(timeNow), "GetCurTimeUtc error")
}

func TestSetUtcTime(t *testing.T) {
	a := assert.New(t)

	timeF := SetUtcTime(TimeFormat{
		Year:    2021,
		Month:   3,
		Day:     30,
		Hour:    10,
		Minute:  21,
		Seconds: 0,
	})
	date := time.Date(2021, 3, 30, 10, 21, 0, 0, time.UTC)
	a.Equal(date, timeF, "SetUtcTime")
}

func TestGetCurDate(t *testing.T) {
	a := assert.New(t)

	var (
		monthStr string
		dayStr   string
	)
	timeNow := time.Now().UTC()
	if timeNow.Month() < 10 {
		monthStr = fmt.Sprintf("0%d", timeNow.Month())
	} else {
		monthStr = fmt.Sprintf("%d", timeNow.Month())
	}
	if timeNow.Day() < 10 {
		dayStr = fmt.Sprintf("0%d", timeNow.Day())
	} else {
		dayStr = fmt.Sprintf("%d", timeNow.Day())
	}
	except := fmt.Sprintf("%d-%s-%s", timeNow.Year(), monthStr, dayStr)
	a.Equal(except, GetCurDate(), "GetCurDate error")
}

func TestGetAnotherTime(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	otherTime := GetAnotherTime(timeNow, TimeFormat{
		Year: 1,
	})
	a.Equal(1, otherTime.Year()-timeNow.Year(), "GetAnotherTime error")

	otherTime = GetAnotherTime(timeNow, TimeFormat{
		Month: 1,
	})
	if timeNow.Month() == 12 {
		a.Equal(1, otherTime.Year()-timeNow.Year(), "GetAnotherTime error")
		a.GreaterOrEqual(time.Month(2), otherTime.Month(), "GetAnotherTime error")
		a.LessOrEqual(time.Month(1), otherTime.Month(), "GetAnotherTime error")
	} else {
		a.LessOrEqual(time.Month(1), otherTime.Month()-timeNow.Month(), "GetAnotherTime error")
		a.GreaterOrEqual(time.Month(2), otherTime.Month()-timeNow.Month(), "GetAnotherTime error")
	}
}

func TestGetTimeSub(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	otherT := GetAnotherTime(timeNow, TimeFormat{
		Day: 10,
	})
	a.LessOrEqual(time.Hour*24*10-time.Second*1, GetTimeSub(timeNow, otherT), "GetAnotherTime error")
	a.GreaterOrEqual(time.Hour*24*101+time.Second*1, GetTimeSub(timeNow, otherT), "GetAnotherTime error")
}

func TestGetCurTimeSub(t *testing.T) {
	a := assert.New(t)
	otherT := GetAnotherTime(time.Now().UTC(), TimeFormat{
		Day: 10,
	})
	a.LessOrEqual(time.Hour*24*10-time.Second*1, GetCurTimeSub(otherT), "GetCurTimeSub error")
	a.GreaterOrEqual(time.Hour*24*10+time.Second*1, GetCurTimeSub(otherT), "GetCurTimeSub error")
}

func TestStrToDateTime(t *testing.T) {
	a := assert.New(t)

	strT := "2020-01-02 14:00:01"
	dateT, err := StrToDateTime(strT)
	exceptT := time.Date(2020, 1, 2, 14, 0, 1, 0, time.UTC)
	a.Equal(nil, err, "StrToDateTime error")
	a.Equal(exceptT, dateT, "StrToDateTime error")

	strT = "2020-01-0214:00:01"
	_, err = StrToDateTime(strT)
	a.NotEqual(nil, err, "StrToDateTime error")
}

func TestDateTimeToStr(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	a.Equal(timeNow.Format("2006-01-02 15:04:05"), DateTimeToStr(timeNow), "DateTimeToStr error")
}

func TestDateTimeToTimestamp(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	a.Equal(int(GetCurTimeStamp(false)), DateTimeToTimestamp(timeNow, false), "DateTimeToTimestamp error")
	a.Equal(int(GetCurTimeStamp(true)), DateTimeToTimestamp(timeNow, true), "DateTimeToTimestamp error")
}

func TestStrTimeToTimestamp(t *testing.T) {
	a := assert.New(t)

	timeNow := time.Now().UTC()
	nowStr := timeNow.Format("2006-01-02 15:04:05")
	timestamp, _ := StrTimeToTimestamp(nowStr, false)
	a.Equal(int(GetCurTimeStamp(false)), timestamp, "DateTimeToTimestamp error")
	timestamp, _ = StrTimeToTimestamp(nowStr, true)
	a.Equal(int(GetCurTimeStamp(true)), timestamp, "DateTimeToTimestamp error")

	nowStr = timeNow.Format("2006-01-0215:04:05")
	_, err := StrTimeToTimestamp(nowStr, false)
	a.NotEqual(nil, err, "DateTimeToTimestamp error")
}

func TestTimestampToDatetime(t *testing.T) {
	a := assert.New(t)

	exceptT := time.Date(2021, 3, 10, 2, 12, 1, 0, time.UTC)
	timestamp := DateTimeToTimestamp(exceptT, false)
	timestampMs := DateTimeToTimestamp(exceptT, true)

	a.Equal(exceptT, TimestampToDatetime(timestamp, false), "TimestampToDatetime error")
	a.Equal(exceptT, TimestampToDatetime(timestampMs, true), "TimestampToDatetime error")
}

func TestTimestampToStrTime(t *testing.T) {
	a := assert.New(t)

	exceptT := time.Date(2021, 3, 10, 2, 12, 1, 0, time.UTC)
	timestamp := DateTimeToTimestamp(exceptT, false)

	a.Equal("2021-03-10 02:12:01", TimestampToStrTime(timestamp), "TimestampToStrTime error")
}
