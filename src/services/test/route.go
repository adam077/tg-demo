package test

import "github.com/gin-gonic/gin"

var MonitorRoutes = map[string]map[string]gin.HandlersChain{
	"get_persons": {
		"GET": gin.HandlersChain{GetPersons},
	},
	"": {
		"GET": gin.HandlersChain{Haha1},
	},
	"hah": {
		"GET": gin.HandlersChain{Haha2},
	},
	"echarts": {
		"GET": gin.HandlersChain{GetZhihuEcharts},
	},
	"set_eat_what": {
		"GET": gin.HandlersChain{EatWhat},
	},
	"eat_what": {
		"GET": gin.HandlersChain{SeeEatWhat},
	},
}
