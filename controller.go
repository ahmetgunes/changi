package changi

import (
	"net/http"
	"net/url"
)

//Control requests via connector

type AsyncRequest struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Tag     string            `json:"tag"`
	Id      string            `json:"id"`
	Timeout float32           `json:"timeout"`
	Body    []string          `json:"body"`
}

func (asyncRequest *AsyncRequest) toHttpRequest() request {
	httpReq := http.Request{
		URL:    url.URL{},
		Method: asyncRequest.Method
	}
}

type asyncResponse struct {
	Headers map[string]string `json:"headers"`
	Body    []byte            `json:"body"`
}

func fromHttpResponse(resp *http.Response) asyncResponse {
	return asyncResponse{}
}

func start(asyncRequest AsyncRequest) {
	//Open goroutines for each request
	//Start a ticker
	//Start the processes with the controller which waits for responses from the requesters and determines timeouts
}

func controller(response chan http.Response, ticker chan bool) {

}
