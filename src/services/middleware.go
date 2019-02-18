package services

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

func TimeCostMiddleware(c *gin.Context) {
	tEnter := time.Now()
	defer func(t time.Time) {
		tNow := time.Now()
		log.Info().Float64("timeCost", tNow.Sub(tEnter).Seconds()).Str("path", c.Request.URL.Path).Str("query", c.Request.URL.RawQuery).Msg("api timing")
	}(tEnter)

	c.Next()
}
