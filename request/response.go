package request

import "net/http"

type Response struct {
	Resp *http.Response
	Id   string
}

type AsyncResponse struct {
	Headers     map[string]string `json:"headers"`
	Body        []byte            `json:"body"`
	ElapsedTime float32           `json:"elapsed_time"`
}

func fromHttpResponse(resp *http.Response) AsyncResponse {
	return AsyncResponse{}
}
