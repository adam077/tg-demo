package main

import (
	"go-go-go/src/data"
	"go-go-go/src/scheduler"
	"go-go-go/src/services"
)

func main() {
	if data.GetConfig(data.Scheduler) != "" {
		scheduler.Run()
	}
	engine := services.SetupEngine()
	engine.Run("0.0.0.0:8080")
}
