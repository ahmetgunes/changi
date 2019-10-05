package internal

import (
	"github.com/ahmetgunes/changi/configs"
	"github.com/ahmetgunes/changi/internal/storage"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"path/filepath"
)

func Test(config configs.Configuration) {
	filePath, _ := filepath.Abs(config.TestDataPath)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	data := string(file)
	storage.Connect(config.MemcachedConnectionString)
	_ = storage.Storage.Set(&memcache.Item{Key: config.TestDataKey, Value: []byte(data)})
	requests, _ := storage.FetchRequest(config.TestDataKey)
	for i, request := range requests {
		request.Id = ksuid.New().String()
		requests[i] = request
	}
	start(requests)
}
