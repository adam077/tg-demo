package services

import (
	"net/http"
	"tg-demo/src/services/test"
	"tg-demo/src/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupEngine() *gin.Engine {
	engine := gin.New()
	engine.RedirectTrailingSlash = true
	if false {
		gin.SetMode(gin.ReleaseMode)
	}
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
	engine.Use(QueryMonitorMiddleware)
	registerRouters(engine)
	return engine
}

func registerRouters(engine *gin.Engine) {
	engine.Static("/assets", "./src/assets")
	engine.StaticFS("/assets_list", http.Dir("src/assets"))

	engine.Handle(http.MethodGet, "hi", func(context *gin.Context) {
		utils.SuccessResp(context, "hi", nil)
	})

	apiGroupLv0 := engine.Group("/migrate")
	apiGroupLv0.Handle(http.MethodPost, "", test.Migrate)

	apiGroupLv1 := engine.Group("/auth")
	includeRoutes(apiGroupLv1, test.AuthRoutes)

	apiGroupLv2 := engine.Group("/query")
	apiGroupLv2.Use(CheckAuth)
	includeRoutes(apiGroupLv2, test.CommonRoutes)
}

func includeRoutes(group *gin.RouterGroup, routes map[string]map[string]gin.HandlersChain) {
	for url, methodHandlerChain := range routes {
		for method, handlerChain := range methodHandlerChain {
			group.Handle(method, url, handlerChain...)
		}
	}
}
