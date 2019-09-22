package changi

import (
	"github.com/bradfitz/gomemcache/memcache"
)

func Test() {
	data := `[
  {
    "url": "http://www.mocky.io/v2/5d77ebfe3200006e7f92408d",
    "method": "GET",
    "headers": {
      "X-Tryout-Header": "1",
      "X-Puzzle-Header": "Hilal",
      "Accept-Encoding": "Hello"
    },
    "tag": "200 Status Code",
    "id": 1,
    "timeout":12.3
  },
  {
    "url": "http://www.mocky.io/v2/5d77eca83200006d7f924090",
    "method": "POST",
    "headers": {
      "X-Tryout-Header": "1",
      "X-Puzzle-Header": "Hilal",
      "Accept-Encoding": "Hello"
    },
    "tag": "400 Status Code",
    "id": 2,
    "timeout":12.1,
	"mandatory": true
  },
  {
    "url": "http://www.mocky.io/v2/5d77ed193200003d47924091?mock-delay=100ms",
    "method": "POST",
    "headers": {
      "X-Tryout-Header": "1",
      "X-Puzzle-Header": "Hilal",
      "Accept-Encoding": "Hello"
    },
    "tag": "400 Status Code",
    "id": 3,
    "timeout": 10,
	"mandatory": true
  }
]`

	connect("127.0.0.1:11211")
	_ = storage.Set(&memcache.Item{Key: "request_1", Value: []byte(data)})
	requests, _ := fetchRequest("request_1")
	start(requests)
}
