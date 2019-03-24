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
}
