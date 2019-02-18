package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Resp(c *gin.Context, statusCode, errorCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{"code": errorCode, "message": message, "data": data})
}

func AbortResp(c *gin.Context, statusCode, errorCode int, message string, data interface{}) {
	c.AbortWithStatusJSON(statusCode, gin.H{"code": errorCode, "message": message, "data": data})
}

func SuccessResp(c *gin.Context, message string, data interface{}) {
	Resp(c, http.StatusOK, 0, message, data)
}

func ErrorResp(c *gin.Context, statusCode, errCode int, message string) {
	AbortResp(c, statusCode, errCode, message, nil)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
