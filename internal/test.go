package internal

import (
	"github.com/ahmetgunes/changi/internal/storage"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"path/filepath"
)

func Test() {
	filePath, _ := filepath.Abs("github.com/ahmetgunes/changi/test/data.json")
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	data := string(file)
	storage.Connect("127.0.0.1:11211")
	_ = storage.Storage.Set(&memcache.Item{Key: "request_1", Value: []byte(data)})
	requests, _ := storage.FetchRequest("request_1")
	for i, request := range requests {
		request.Id = ksuid.New().String()
		requests[i] = request
	}
	start(requests)
}
