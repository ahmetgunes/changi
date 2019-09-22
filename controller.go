package changi

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetgunes/changi/request"
	"github.com/bradfitz/gomemcache/memcache"
	"sync"
	"time"
)

var mandatoryIds []string
var responses []request.AsyncResponse

func start(requests []*request.AsyncRequest) {
	var wg sync.WaitGroup
	//Open goroutines for each request
	//Start a ticker
	//Start the processes with the controller which waits for responses from the requesters and determines timeouts
	responseChan := make(chan request.Response)
	progressChan := make(chan string)

	count := 0
	for i, request := range requests {
		fmt.Println("Starting the request on", request.Id, request.Tag, i)
		wg.Add(i)
		if request.Mandatory {
			mandatoryIds = append(mandatoryIds, request.Id)
		}
		go makeRequest(request.ToHttpRequest(), responseChan, progressChan, &wg)
		count = i
	}
	wg.Add(count + 1)
	go controller(responseChan, &wg, count)
	wg.Wait()
	defer close(responseChan)
	defer close(progressChan)
}

func controller(response chan request.Response, wg *sync.WaitGroup, count int) bool {
	var ticker = time.NewTicker(1 * time.Millisecond)
	var tickCount = 10000
	defer wg.Done()
	defer ticker.Stop()
	for {
		select {
		case resp := <-response:
			removeIfMandatory(resp.Id)
			x, _ := json.Marshal(request.FromHttpResponse(resp))
			responses = append(responses, x)
			_ = storage.Set(&memcache.Item{Key: "response_" + resp.Id, Value: x})
			if len(mandatoryIds) == 0 {
				return true
			}
			fmt.Println("Response for:", resp.Id)
		case <-ticker.C:
			tickCount--
			if tickCount == 0 {
				fmt.Println("Ticker has reached to zero")
				if len(mandatoryIds) == 0 {
					fmt.Println("Ending requester since the ticker is off finally")
					return true
				}
			}
		}
	}
}

func removeIfMandatory(reqId string) {
	for pos, id := range mandatoryIds {
		if id == reqId {
			lastPos := len(mandatoryIds) - 1
			mandatoryIds[pos] = mandatoryIds[lastPos]
			mandatoryIds = mandatoryIds[:lastPos]
			break
		}
	}
}
