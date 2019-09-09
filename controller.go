package changi

import "net/http"

//Control requests via connector

type async_request struct {
	requests []request
	//Timeout duration in milliseconds
	timeout float32
}

func start(asyncRequest async_request) {
	//Open goroutines for each request
	//Start a ticker
	//Start the processes with the controller which waits for responses from the requesters and determines timeouts
}

func controller(response chan http.Response) {

}
