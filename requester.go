package changi

import (
	"fmt"
	"github.com/ahmetgunes/changi/request"
	"net/http"
	"sync"
)

//@TODO: Implement methods for requests

var client = &http.Client{}

func makeRequest(req request.Request, responseChan chan request.Response, progress chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	//@TODO: Log here
	//@TODO: Calculate elapsed time and set to response
	fmt.Println("Starting the request on" + req.Tag)
	resp, _ := client.Do(req.Req)
	defer resp.Body.Close()
	var respStruct = request.Response{Resp: resp, Id: req.Id}
	responseChan <- respStruct
}
