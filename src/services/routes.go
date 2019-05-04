package services

import (
	"go-go-go/src/services/test"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupEngine() *gin.Engine {
	engine := gin.New()
	engine.RedirectTrailingSlash = true
	//gin.SetMode(gin.ReleaseMode)
	engine.Use(QueryMonitorMiddleware)
	if true {
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
	apiGroupLv1 := engine.Group("/lv1")
	apiGroupLv2 := apiGroupLv1.Group("/lv2")
	includeRoutes(apiGroupLv2, test.MonitorRoutes)
}

func includeRoutes(group *gin.RouterGroup, routes map[string]map[string]gin.HandlersChain) {
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			group.Handle(method, url, handlerChain...)
		}
	}
}
