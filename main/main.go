package main

import (
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/internal"
	"github.com/ahmetgunes/changi/logger"
)

func main() {
	config := configs.Init()
	logger := logger.Init(config.LogFile)
	internal.InitTest(config, logger)
	internal.Test()
}
