package main

import (
	"flag"
	"fmt"
	"go-go-go/src/config"
	"go-go-go/src/scheduler"
	"go-go-go/src/services"
)

func main() {
	flag.Parse()
	if config.Config.EnableScheduler {
		scheduler.Run()
	}
	engine := services.SetupEngine()
	engine.Run(fmt.Sprintf("0.0.0.0:%d", config.Config.ServicePort))
}
