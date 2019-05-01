package utils

import (
	"github.com/rs/zerolog/log"
	"runtime"
	"time"
)

func CommonRecover() {
	if err := recover(); err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)

		if err, ok := err.(error); ok {
			log.Error().Str("error", err.(error).Error()).Str("stack", string(buf[:n])).Msg("goroutine unexpected panic.")
		} else {
			log.Error().Str("stack", string(buf[:n])).Msg("goroutine unexpected error when recover.")
		}
	}
}

func TimeToDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func KeyDiff(left, right map[string]interface{}) bool {
	if len(left) != len(right) {
		return false
	}
	for key := range left {
		if _, ok := right[key]; !ok {
			return false
		}
	}
	return true
}

var ShangHaiLocation, _ = time.LoadLocation("Asia/Shanghai")

func GetLocalTimeFromShanghaiString(raw string) time.Time {
	result, _ := time.ParseInLocation(TIME_SHANGHAI_FORMAT, raw, ShangHaiLocation)
	return result.Local()
}

func GetShanghaiTimeString(raw time.Time) string {
	return raw.In(ShangHaiLocation).Format(TIME_SHANGHAI_FORMAT)
}

const (
	TIME_SHANGHAI_FORMAT = "2006-01-02T15:04:05+08:00"
	DATE_FORMAT          = "2006-01-02"
	TIME_DEFAULT_FORMAT  = "2006-01-02 15:04:05"
)

func GetTimeDateString(t time.Time) string {
	// 时间转化为日期字符串
	return t.Format(DATE_FORMAT)
}

func GetNowTime() (int64, string, int, int) {
	now := time.Now()
	today := GetTimeDateString(now)
	hour := now.Hour()
	min := (now.Minute() / 5) * 5
	return now.Unix(), today, hour, min
}
