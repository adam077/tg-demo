package main

import (
	"tg-demo/src/services"
)

func main() {
	engine := services.SetupEngine()
	engine.Run("0.0.0.0:8080")
}
