package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SuccessResp(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": message, "data": data})
}

func ErrorResp(c *gin.Context, errCode int, message string) {
	c.AbortWithStatusJSON(400, gin.H{"code": errCode, "message": message, "data": nil})
}
