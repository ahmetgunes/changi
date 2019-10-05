package internal

import (
	"encoding/json"
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal/request"
	"github.com/bradfitz/gomemcache/memcache"
)

var Storage *memcache.Client

func Connect(connectionString string) {
	Storage = memcache.New(connectionString)
}

func Disconnect() {
	Storage = nil
}

func FetchRequest(key string) (req []*request.AsyncRequest, status bool) {
	//Implement decode json and fetch request
	item, err := Storage.Get(key)
	if err != nil {
		changi.Log.Fatal(err)
	}

	if item == nil {
		changi.Log.Fatal("No request were found with the key:" + key)
	}

	var requests []*request.AsyncRequest
	_ = json.Unmarshal(item.Value, &requests)
	return requests, true
}

func WriteResponse(key string, asyncResponse request.AsyncResponse) bool {
	data, _ := json.Marshal(asyncResponse)
	err := Storage.Set(&memcache.Item{Key: key, Value: data})
	if err != nil {
		changi.Log.Fatal(err)
	}

	return true
}
