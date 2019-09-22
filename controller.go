package changi

import (
	"fmt"
	"github.com/ahmetgunes/changi/request"
	"sync"
	"time"
)

var mandatoryIds []string

func start(requests []*request.AsyncRequest) {
	var wg sync.WaitGroup
	//Open goroutines for each request
	//Start a ticker
	//Start the processes with the controller which waits for responses from the requesters and determines timeouts
	responseChan := make(chan request.Response)
	progressChan := make(chan string)

	count := 0
	for i, request := range requests {
		fmt.Println("Starting the request on" + request.Tag)
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
	defer wg.Done()
	var ticker = time.NewTicker(1 * time.Millisecond)
	var tickCount = 10000
	for {
		select {
		case resp := <-response:
			for pos, id := range mandatoryIds {
				if id == resp.Id {
					removeFromIds(pos)
				}
			}
			if len(mandatoryIds) == 0 {
				return true
			}
			fmt.Println("Response:", resp)
		case t := <-ticker.C:
			tickCount--
			fmt.Println("Remaining tick", tickCount, t.Second())
			if tickCount == 0 {
				return true
			}
		}
	}
}

func removeFromIds(i int) {
	mandatoryIds[i] = mandatoryIds[len(mandatoryIds)-1]
	mandatoryIds = mandatoryIds[:len(mandatoryIds)-1]
}
