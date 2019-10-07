package internal

import (
	"encoding/json"
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal/request"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/segmentio/ksuid"
	"sync"
	"time"
)

var mandatoryIds []string
var responses []request.AsyncResponse

func Start(requests []*request.AsyncRequest) {
	var wg sync.WaitGroup
	//Open goroutines for each request
	//Start a ticker
	//Start the processes with the controller which waits for responses from the requesters and determines timeouts
	responseChan := make(chan request.Response)
	progressChan := make(chan string)

	count := 0
	for i, request := range requests {
		changi.Log.Info("Starting the request on", request.Id, request.Tag, i)

		request.Id = ksuid.New().String()
		requests[i] = request
		if request.Mandatory {
			mandatoryIds = append(mandatoryIds, request.Id)
		}
		count = i
		wg.Add(i)
		go makeRequest(request.ToHttpRequest(), responseChan, progressChan, &wg)
	}
	wg.Add(count + 1)
	go controller(responseChan, &wg)
	wg.Wait()
	defer close(responseChan)
	defer close(progressChan)
}

func controller(response chan request.Response, wg *sync.WaitGroup) bool {
	var ticker = time.NewTicker(1 * time.Millisecond)
	var tickCount = 10000
	defer wg.Done()
	defer ticker.Stop()
	for {
		select {
		case resp := <-response:
			removeIfMandatory(resp.Id)
			marshalledResponse, _ := json.Marshal(request.FromHttpResponse(resp))
			_ = Storage.Set(&memcache.Item{Key: "response_" + resp.Id, Value: marshalledResponse})
			if len(mandatoryIds) == 0 {
				return true
			}
			changi.Log.Info("Gotten response for:", resp.Id, resp.Resp.StatusCode)
		case <-ticker.C:
			tickCount--
			if tickCount == 0 {
				changi.Log.Info("Ticker has reached to zero")
				if len(mandatoryIds) == 0 {
					changi.Log.Info("Ending requester since the ticker is off finally")
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
