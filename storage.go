package changi

import "github.com/bradfitz/gomemcache/memcache"

var storage *memcache.Client

func connect(connectionString string) {
	storage = memcache.New(connectionString)
}

func disconnect() {
	//Implement disconnect and destroy for "storage"
}

func fetchRequest(key string) async_request {
	//Implement decode json and fetch request
}

func writeRequest(key string, asyncRequest async_request) bool {
	//Implement encode json and writing for async_request
}
