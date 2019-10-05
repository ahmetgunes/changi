package configs

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"os"
	"path/filepath"
)

const configurationFilePath = "github.com/ahmetgunes/changi/configs/file/config.%s.json"

type Configuration struct {
	MemcachedConnectionString string `json:"memcached_connection_string"`
	TestDataPath              string `json:"test_data_path"`
	TestDataKey               string `json:"test_data_key"`
	LogFile                   string `json:"log_file"`
}

func Init() Configuration {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	configPath, _ := filepath.Abs(fmt.Sprintf(configurationFilePath, env))
	config := Configuration{}
	err := gonfig.GetConf(configPath, &config)
	if err != nil {
		panic(err)
	}

	return config
}
