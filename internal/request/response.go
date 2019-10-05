package request

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Resp *http.Response
	Id   string
}

type AsyncResponse struct {
	Headers     map[string][]string `json:"headers"`
	Body        string              `json:"body"`
	ElapsedTime float32             `json:"elapsed_time"`
	Id          string              `json:"id"`
	Status      int                 `json:"status"`
}

func FromHttpResponse(resp Response) AsyncResponse {
	defer resp.Resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Resp.Body)
	if err != nil {
		//@TODO: Do not panic maybe?
		panic(err)
	}
	return AsyncResponse{
		Headers: resp.Resp.Header,
		Status:  resp.Resp.StatusCode,
		Body:    string(body),
		Id:      resp.Id}
}
