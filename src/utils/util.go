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

var TIME_SHANGHAI_FORMAT = "2006-01-02T15:04:05+08:00"

func GetLocalTimeFromShanghaiString(raw string) time.Time {
	result, _ := time.ParseInLocation(TIME_SHANGHAI_FORMAT, raw, ShangHaiLocation)
	return result.Local()
}

func GetShanghaiTimeString(raw time.Time) string {
	return raw.In(ShangHaiLocation).Format(TIME_SHANGHAI_FORMAT)
}
