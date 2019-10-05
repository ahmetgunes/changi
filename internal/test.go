package internal

import (
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/internal/control"
	"github.com/ahmetgunes/changi/internal/request"
	"github.com/ahmetgunes/changi/internal/storage"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/op/go-logging"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"path/filepath"
)

var log *logging.Logger
var configuration configs.Configuration

func InitTest(config configs.Configuration, logger *logging.Logger) {
	log = logger
	configuration = config
}

func Test() {
	requests := fetchRequests()
	log.Info("Starting requests")
	for i, request := range requests {
		request.Id = ksuid.New().String()
		requests[i] = request
	}
	control.Start(requests)
}

func fetchRequests() []*request.AsyncRequest {
	data := fetchTestData()
	storage.Connect(configuration.MemcachedConnectionString)
	err := storage.Storage.Set(&memcache.Item{Key: configuration.TestDataKey, Value: []byte(data)})
	if err != nil {
		log.Fatal("An error has occurred while setting test data", err)
		panic(err)
	}
	requests, status := storage.FetchRequest(configuration.TestDataKey)
	if !status {
		log.Fatal("An error has occurred while fetching test data")
		panic("An error has occurred while fetching test data")
	}

	return requests
}

func fetchTestData() string {
	filePath, _ := filepath.Abs(configuration.TestDataPath)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error while reading test data", err)
		panic(err)
	}
	return string(file)
}
