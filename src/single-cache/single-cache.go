package single_cache

import (
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

var cacheMap sync.Map

func init() {
	// 这种实现只能单实例了
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				cacheMap.Range(UpdateTokenMap)
			}
		}
	}()
}

type token struct {
	v      string
	expire bool
	t      time.Time
}

func UpdateTokenMap(k, v interface{}) bool {
	t := v.(token)
	now := time.Now()
	if t.expire && now.After(t.t) {
		log.Info().Str("key", k.(string)).Msg("out of time")
		cacheMap.Delete(k)
	}
	return true
}

func Set(key, v string, expire ...int) bool {
	t := token{
		v: v,
	}
	if len(expire) > 0 && expire[0] > 0 {
		t.expire, t.t = true, time.Now().Add(time.Duration(expire[0])*time.Second)
	}
	cacheMap.Store(key, t)
	return true
}

func Get(key string) (string, bool) {
	result, ok := cacheMap.Load(key)
	if !ok {
		return "", false
	}
	t := result.(token)
	return t.v, true
}

func Delete(key string) {
	cacheMap.Delete(key)
}
