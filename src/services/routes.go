package services

import (
	"net/http"

	"go-go-go/src/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type SpecialRouterGroup struct {
	*gin.RouterGroup
}

func (group *SpecialRouterGroup) includeWithGroup(groupName string, routes map[string]map[string]gin.HandlersChain) {
	subGroup := group.RouterGroup.Group(groupName)
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			subGroup.Handle(method, url, handlerChain...)
		}
	}
}

func (group *SpecialRouterGroup) includeRoutes(routes map[string]map[string]gin.HandlersChain) {
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			group.RouterGroup.Handle(method, url, handlerChain...)
		}
	}
}

func SetupEngine() *gin.Engine {
	engine := gin.New()
	engine.RedirectTrailingSlash = true

	if !config.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine.Use(TimeCostMiddleware)

	if config.Config.EnableCORS {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowCredentials = true
		corsConfig.AddAllowMethods(http.MethodPatch)
		corsConfig.AddAllowMethods(http.MethodDelete)
		corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Cache-Control", "Pragma")
		corsConfig.AllowOriginFunc = func(origin string) bool {
			return true
		}
		engine.Use(cors.New(corsConfig))
	}

	registerRouters(engine)
	return engine
}

func registerRouters(engine *gin.Engine) {
	apiGroupLv1 := &SpecialRouterGroup{engine.Group("/lv1")}
	registerMonitorRoutes(apiGroupLv1)
}

func registerMonitorRoutes(apiGroup *SpecialRouterGroup) {
	apiGroupLv2 := &SpecialRouterGroup{apiGroup.Group("/lv2")}
	apiGroupLv2.includeRoutes(MonitorRoutes)
}

// ---------------------------------------------------------------

var MonitorRoutes = map[string]map[string]gin.HandlersChain{
	"": {
		"GET": gin.HandlersChain{Haha1},
	},
	"hah": {
		"GET": gin.HandlersChain{Haha2},
	},
}
