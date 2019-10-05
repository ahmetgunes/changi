package changi

import (
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/logger"
	"github.com/op/go-logging"
)

var Config configs.Configuration
var Log *logging.Logger


func Init() {
	Config = configs.Init()
	Log = logger.Init(Config.LogFile)
}