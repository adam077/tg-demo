package test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tg-demo/src/data"
	"tg-demo/src/single-cache"
	"tg-demo/src/utils"
)

type User struct {
	Name string `json:"user"`
	PW   string `json:"password"`
}

func Login(c *gin.Context) {
	body, _ := c.GetRawData()
	var userBody User
	json.Unmarshal(body, &userBody)

	user := data.GetUsers(userBody.Name)
	if len(user) != 1 {
		utils.ErrorResp(c, 40000, "用户名错误")
		return
	}
	if user[0].Password != userBody.PW {
		utils.ErrorResp(c, 40000, "密码错误")
		return
	}

	token, err := utils.GetToken(user[0].ID)
	if err != nil {
		// todo log
		utils.ErrorResp(c, 40000, "请稍后")
		return
	}

	single_cache.Set(token, user[0].ID, 24*60*60)
	utils.SuccessResp(c, "", token)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Token")
	single_cache.Delete(token)
	utils.SuccessResp(c, "", nil)
}
