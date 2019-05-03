package main

import (
	"flag"
	"go-go-go/src/scheduler"
	"go-go-go/src/services"
)

func main() {
	flag.Parse()
	scheduler.Run()
	engine := services.SetupEngine()
	engine.Run("0.0.0.0:8080")
}
