package utils

import (
	"fmt"
	"time"
)

const baseFormat = "2006-01-02 15:04:05"

type TimeFormat struct {
	Year    int
	Month   time.Month
	Day     int
	Hour    int
	Minute  int
	Seconds int
}

// 获取当前时间时间戳(UTC与本地时区结果一致, 可使用time.Now().Unix()与GetCurTimeUtc().Unix()作对比)
func GetCurTimeStamp(needMs bool) int64 {
	curTime := time.Now().Unix()
	if needMs {
		// 返回13位的时间戳
		return curTime * 1000
	}
	return curTime
}

// 获取当前时间UTC时间戳(UTC与本地时区结果一致, 可使用time.Now().Unix()与GetCurTimeUtc().Unix()作对比)
func GetCurTimeUtcStamp(needMs bool) int64 {
	curTime := time.Now().UTC().Unix()
	if needMs {
		// 返回13位的时间戳
		return curTime * 1000
	}
	return curTime
}

// 获取当前时间(str类型)
func GetCurTimeStr() string {
	curTime := GetCurTimeUtc().Format(baseFormat)
	return curTime
}

// 获取当前时间(time.Time类型)
func GetCurTimeUtc() time.Time {
	curTime := time.Now().UTC()
	return curTime
}

// 构造时间
func SetUtcTime(timeF TimeFormat) time.Time {
	var setTime time.Time
	setTime = time.Date(timeF.Year, timeF.Month, timeF.Day, timeF.Hour, timeF.Minute, timeF.Seconds, 0, time.UTC)
	return setTime
}

// 获取当前日期(str类型)
func GetCurDate() string {
	var (
		monthStr string
		dayStr   string
	)
	nowTime := GetCurTimeUtc()
	year := nowTime.Year()
	month := nowTime.Month()
	if month < 10 {
		monthStr = fmt.Sprintf("0%d", month)
	} else {
		monthStr = fmt.Sprintf("%d", month)
	}
	day := nowTime.Day()
	if day < 10 {
		dayStr = fmt.Sprintf("0%d", day)
	} else {
		dayStr = fmt.Sprintf("%d", day)
	}
	curDate := fmt.Sprintf("%d-%s-%s", year, monthStr, dayStr)
	return curDate
}

// 获取输入时间相隔一定时间的时间
func GetAnotherTime(beforeTime time.Time, timeMove TimeFormat) time.Time {
	afterTime := beforeTime.AddDate(timeMove.Year, int(timeMove.Month), timeMove.Day)
	hour := time.Duration(timeMove.Hour)
	minute := time.Duration(timeMove.Minute)
	seconds := time.Duration(timeMove.Seconds)
	date := afterTime.Add(hour*time.Hour + minute*time.Minute + seconds*time.Second)
	return date
}

// 获取两个时间的时间差
func GetTimeSub(time1 time.Time, time2 time.Time) time.Duration {
	var subD time.Duration
	subD = time2.Sub(time1)
	return subD
}

// 获取时间与当前时间的时间差
func GetCurTimeSub(timePar time.Time) time.Duration {
	var (
		subD    time.Duration
		timeNow time.Time
	)
	timeNow = GetCurTimeUtc()
	subD = GetTimeSub(timeNow, timePar)
	return subD
}

// str时间转time.Time时间
func StrToDateTime(strTime string) (time.Time, error) {
	parseStrTime, err := time.Parse(baseFormat, strTime)
	if err != nil {
		return time.Time{}, err
	}
	return parseStrTime, nil
}

// time.Time时间转str时间
func DateTimeToStr(dateTime time.Time) string {
	var timeStr string
	timeStr = dateTime.Format(baseFormat)
	return timeStr
}

func DateTimeToTimestamp(dateTime time.Time, needMs bool) int {
	var (
		t       int
		t1      time.Time
		loc     *time.Location
		timeStr string
	)
	timeStr = dateTime.Format(baseFormat)
	// UTC为国际标准, Local为本地
	loc, _ = time.LoadLocation("UTC")
	t1, _ = time.ParseInLocation(baseFormat, timeStr, loc)
	t = int(t1.Unix())
	if needMs {
		t = t * 1000
	}
	return t
}

// str时间转成标准timestamp(默认UTC)
func StrTimeToTimestamp(timeStr string, needMs bool) (int, error) {
	var (
		err error
		t   int
		t1  time.Time
		loc *time.Location
	)
	// UTC为国际标准, Local为本地
	loc, _ = time.LoadLocation("UTC")
	t1, err = time.ParseInLocation(baseFormat, timeStr, loc)
	if err != nil {
		return t, err
	}
	t = int(t1.Unix())
	if needMs {
		t = t * 1000
	}
	return t, nil
}

// 时间戳转化为datetime时间(返回UTC时间戳)
func TimestampToDatetime(timestamp int, isMs bool) time.Time {
	var (
		t1 int64
		t  time.Time
	)
	if isMs {
		timestamp = timestamp / 1000
	}
	t1 = int64(timestamp)
	t = time.Unix(t1, 0)
	return t.UTC()
}

// 标准timestamp时间转为str时间(UTC格式)
func TimestampToStrTime(timestamp int) string {
	var (
		timeStr string
	)
	timeStr = DateTimeToStr(TimestampToDatetime(timestamp, false))
	return timeStr
}
