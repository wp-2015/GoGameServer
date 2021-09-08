/**
 * Author :
 * CreateDate : 2020/3/2 18:24
 * Software : GoLand
 * Remark :
 **/
package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	//微秒
	_T1 = 1e3
	//毫秒
	_T2 = 1e6
)

var zoneOffsetTime = GetZoneOffsetTime()

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

//秒
func GetUnixSecond() int64 {
	return time.Now().Unix()
}

// 范围内随机 教主要求[min-max]两边都闭区间
func RandRange(min int, max int) int {
	if min == max {
		return min
	}

	r := r.Intn(max - min + 1)
	return r + min
}

//随机  教主要求[1-v]两边都闭区间
func Rand(v int) int {
	return r.Intn(v) + 1
}

func RandMicro10000() int32 {
	return int32(Rand(10000))
}

//随机数组下标，传入数组长度
func RandIndex(v int) int {
	return r.Intn(v)
}

func RandAarryInt32(array []int32) int32 {
	len :=len(array)
	if len  <= 0 {
		return 0
	}

	index := RandIndex(len)
	return array[index]
}

//毫秒
func GetUnixMillis() int64 {
	return time.Now().UnixNano() / _T2
}

//微秒
func GetUnixMicro() int64 {
	return time.Now().UnixNano() / _T1
}

//納秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

//是否在事件内
func IsInTime(beginetime, endtime time.Time) bool {
	timeNow := time.Now()
	if beginetime.Before(timeNow) && timeNow.Before(endtime) {
		return true
	}
	return false
}

//是否在一定小时和分钟内
func IsInHMTime(starthour, startminute, endhour, endminute int8) bool {
	timeNow := time.Now()
	timeRefresh := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), int(starthour), int(startminute), 0, 0, timeNow.Location())
	timerefreshend := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), int(endhour), int(endminute), 0, 0, timeNow.Location())

	return IsInTime(timeRefresh, timerefreshend)
}

func ParseInLocation(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
}

func ParseDate(date string) int64 {
	time1, error := time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
	if error == nil {
		return time1.UnixNano() / _T2
	}
	return -1
}

//获取当前时间的字符串-文件时间
func GetCurStrTime() string {
	timeUnix := time.Now().Unix()
	//01/02 03:04:05PM '06 -0700
	return time.Unix(timeUnix, 0).Format("2006-01-02_15-04-05")
}

//获取当前时间的字符串-文件时间
func GetDayStrTime(tim int64) string {
	return time.Unix(tim, 0).Format("2006_01_02")
}

//获取当前时间的字符串-文件时间
func GetDayFormatStrTime(tim int64) string {
	return time.Unix(tim, 0).Format("2006-01-02 15:04:05")
}

//获取当前时间的字符串 标准时间
func GetCurStrFormatTime() string {
	timeUnix := time.Now().Unix()
	//01/02 03:04:05PM '06 -0700
	return time.Unix(timeUnix, 0).Format("2006.1.2 15:04:05")
}

func HourTimestamp() int64 {
	now := time.Now()
	timestamp := now.Unix() - int64(now.Second()) - int64((60 * now.Minute())) - int64(3600*now.Hour())
	fmt.Println(timestamp, time.Unix(timestamp, 0), now.Unix())
	return timestamp
}

func IsTriggerTimeEvent(lastTiggerTime int64, timeInterVal int64) bool {
	timestamp := HourTimestamp()
	timeNow := time.Now().Unix()
	triggerTime := timestamp + timeInterVal
	if lastTiggerTime < triggerTime && timeNow > triggerTime {
		return true
	} else {
		return false
	}

}

//@param lastClearTime 上次清数据的时间
//@param n          清数据的时间点
func CanClearAtN(now, lastClearTime int64, n int32) bool {
	if 0 == lastClearTime {
		return true
	}
	var offsetTime = -n * 3600000
	return GetOffsetDay(now, offsetTime) != GetOffsetDay(lastClearTime, offsetTime)
}

func IsTheSame(time1, time2 int64, n int32) bool {
	var offsetTime = -n * 3600000
	return GetOffsetDay(time1, offsetTime) == GetOffsetDay(time2, offsetTime)
}

func GetOffsetDay(time1 int64, offsetTime int32) int32 {
	return int32((time1 + int64(offsetTime) + zoneOffsetTime) / 86400000)
}

//时区误差的毫秒数
func GetZoneOffsetTime() int64 {
	_, offset := time.Now().Zone()
	return int64(offset * 1000)
}

//计算间隔分钟
func CalcIntervalMinute(startMillion int64, endMillion int64) int32 {
	intervalMillion := endMillion - startMillion
	return int32(intervalMillion / 60000)
}

//在time区间内
func InStrTime(curTime int64,  strTimeStart string, strTimeEnd string) bool {
	timeStart := ParseDate(strTimeStart)
	if timeStart < 0 {
		//zlog.Errorf("InStrTime, Parse failed, timestr:%s", strTimeStart)
		return false
	}

	timeEnd := ParseDate(strTimeEnd)
	if timeEnd < 0 {
		//zlog.Errorf("InStrTime, Parse failed, timestr:%s", strTimeEnd)
		return false
	}

	if curTime > timeStart && curTime < timeEnd {
		return true
	} else {
		return false
	}
}


//转换CronTab
func TransLocalTime2Cron(strTime string) (string , error) {
	time, err :=  ParseInLocation(strTime)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("* %d %d %d %d *", time.Minute(), time.Hour(), time.Day(), time.Month()), nil
}