package request

import (
	"net/http"
)

type AsyncRequest struct {
	Url       string            `json:"url"`
	Method    string            `json:"method"`
	Headers   map[string]string `json:"headers"`
	Tag       string            `json:"tag"`
	Timeout   float32           `json:"timeout"`
	Body      []string          `json:"body"`
	Mandatory bool              `json:"mandatory"`
	Id        string
}

type Request struct {
	Req     *http.Request
	Timeout float32
	Id      string
	Tag     string
}

func (asyncRequest *AsyncRequest) ToHttpRequest() Request {
	//@TODO: Log creation
	httpReq, _ := http.NewRequest(asyncRequest.Method, asyncRequest.Url, nil)
	for key, header := range asyncRequest.Headers {
		httpReq.Header.Add(key, header)
	}
	return Request{
		Req:     httpReq,
		Timeout: asyncRequest.Timeout,
		Id:      asyncRequest.Id,
		Tag:     asyncRequest.Tag}
}
