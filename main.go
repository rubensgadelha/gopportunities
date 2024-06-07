package main

import (
	"github.com/rubensgadelha/gopportunities/config"
	"github.com/rubensgadelha/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Error("config initialization error:", err)
		return
	}

	router.Initialize()
}
