package changi

import (
	"fmt"
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/logger"
	"github.com/op/go-logging"
	"github.com/tkanos/gonfig"
	"os"
	"path/filepath"
)

const configurationFilePath = "github.com/ahmetgunes/changi/configs/file/config.%s.json"


var Config configs.Configuration
var Log *logging.Logger


func Init() {
	Config = configure()
	Log = logger.Init(Config.LogFile)
}

func configure() configs.Configuration {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	configPath, _ := filepath.Abs(fmt.Sprintf(configurationFilePath, env))
	config := configs.Configuration{}
	err := gonfig.GetConf(configPath, &config)
	if err != nil {
		Log.Fatal(err)
	}

	return config
}
