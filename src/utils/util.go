package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/satori/go.uuid"
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

func GetUUID() string {
	return uuid.NewV4().String()
}

var secrets = "shabi_labike"

func GetToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     user,
		"create_time": time.Now(), // 不同的人登陆同一个账号
	})
	return token.SignedString([]byte(secrets))

}

func GetUser(tokenStr string) (string, error) {
	// 检查token是否在redis中存在，不存在则拒绝
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(secrets), nil
	})
	if err != nil {
		// token 搞不出来，这应该是哪里实现有错误
		return "", err
	}
	// 这里从jwt中反解出userId，从而可以进行下一步操作
	userId := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok = claims["user_id"].(string)
		if !ok {
			return "", errors.New("")
		}
	}
	return userId, nil
}
