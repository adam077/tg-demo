package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var AuthRoutes = map[string]map[string]gin.HandlersChain{
	"login": {
		http.MethodPost: gin.HandlersChain{Login},
	},
	"logout": {
		http.MethodPost: gin.HandlersChain{Logout},
	},
}

var CommonRoutes = map[string]map[string]gin.HandlersChain{
	"screens": {
		// 获得该用户下的大屏
		http.MethodGet: gin.HandlersChain{GetScreens},
		// 新增大屏
		http.MethodPost: gin.HandlersChain{AddScreen},
		// 修改大屏配置
		http.MethodPatch: gin.HandlersChain{PatchScreen},
		// 删除
		http.MethodDelete: gin.HandlersChain{DeleteScreen},
	},
	"components": {
		http.MethodGet: gin.HandlersChain{GetComponents},
	},
}
