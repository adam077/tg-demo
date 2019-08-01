package services

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"tg-demo/src/single-cache"
	"tg-demo/src/utils"
)

func QueryMonitorMiddleware(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 40003, "message": "你的数据内容需要修改喔"})
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	defer func(body []byte) {
		if c.Writer.Status() == http.StatusOK {
			log.Info().Str("path", c.Request.URL.Path).
				Str("method", c.Request.Method).
				Str("query", c.Request.URL.RawQuery).
				Int("status", c.Writer.Status()).Msg("api record")
		} else {
			log.Error().Str("path", c.Request.URL.Path).
				Str("method", c.Request.Method).
				Str("query", c.Request.URL.RawQuery).
				Int("status", c.Writer.Status()).
				Str("body", string(body)).Msg("api record")
		}
	}(body)
	c.Next()
}

func CheckAuth(c *gin.Context) {
	token := c.GetHeader("Token")
	userId, ok := single_cache.Get(token)
	if !ok {
		utils.ErrorResp(c, 40000, "请登陆")
		return
	}
	single_cache.Set(token, userId, 24*60*60)
	c.Request.Header.Set("UserId", userId)
	c.Next()
}
