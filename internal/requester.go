package internal

import (
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal/request"
	"net/http"
	"sync"
)

//@TODO: Implement methods for requests

var client = &http.Client{}

func makeRequest(req request.Request, responseChan chan request.Response, progress chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	//@TODO: Log here
	//@TODO: Calculate elapsed time and set to response
	resp, _ := client.Do(req.Req)
	defer changi.Log.Info("Ending request", req.Tag, req.Id)
	responseChan <- request.Response{Resp: resp, Id: req.Id}
}
