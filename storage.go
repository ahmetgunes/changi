package changi

import (
	"encoding/json"
	"github.com/ahmetgunes/changi/request"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var storage *memcache.Client

func connect(connectionString string) {
	storage = memcache.New(connectionString)
}

func disconnect() {
	storage = nil
}

func fetchRequest(key string) (req []*request.AsyncRequest, status bool) {
	//Implement decode json and fetch request
	item, err := storage.Get("request_1")
	if err != nil {
		log.Fatal(err)
		return nil, false
	}

	if item == nil {
		log.Fatal("No request were found with the key:" + key)
		return nil, false
	}

	var requests []*request.AsyncRequest
	_ = json.Unmarshal(item.Value, &requests)
	return requests, true
}

func writeResponse(key string, asyncResponse request.AsyncResponse) bool {
	data, _ := json.Marshal(asyncResponse)
	err := storage.Set(&memcache.Item{Key: key, Value: data})
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
