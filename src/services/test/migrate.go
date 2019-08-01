package test

import (
	"github.com/gin-gonic/gin"
	"tg-demo/src/data"
	"tg-demo/src/utils"
)

func Migrate(c *gin.Context) {
	migrate1()
	utils.SuccessResp(c, "", nil)
}

func migrate1() {
	db := data.GetDataDB("default")
	db.AutoMigrate(&data.User{})
	db.AutoMigrate(&data.Screen{})
	db.AutoMigrate(&data.Component{})
	db.AutoMigrate(&data.UserScreen{})
	// todo 唯一索引
}
