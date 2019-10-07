package internal

import (
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal/request"
	"github.com/bradfitz/gomemcache/memcache"
	"io/ioutil"
	"path/filepath"
)

func Test() {
	requests := fetchRequests()
	Start(requests)
}

func fetchRequests() []*request.AsyncRequest {
	data := fetchTestData()
	ConnectStorage(changi.Config.MemcachedConnectionString)
	err := Storage.Set(&memcache.Item{Key: changi.Config.TestDataKey, Value: []byte(data)})
	if err != nil {
		changi.Log.Fatal("An error has occurred while setting test data", err)
	}
	requests, status := FetchRequest(changi.Config.TestDataKey)
	if !status {
		changi.Log.Fatal("An error has occurred while fetching test data")
	}

	return requests
}

func fetchTestData() string {
	filePath, _ := filepath.Abs(changi.Config.TestDataPath)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		changi.Log.Fatal("Error while reading test data", err)
	}
	return string(file)
}
