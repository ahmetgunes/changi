package web

import (
	"encoding/json"
	"fmt"
	"github.com/ahmetgunes/changi"
	"github.com/ahmetgunes/changi/internal"
	"net/http"
)

type NewRequest struct {
	RequestId  string `json:"request_id"`
	MaxTimeOut int    `json:"max_time_out"`
}

func Listen() {
	http.HandleFunc("/new", new)
	http.HandleFunc("/fetch", fetch)
	_ := http.ListenAndServe(changi.Config.Port, nil)
}

func new(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		changi.Log.Error("Request method must be 'POST'")
		_, err := fmt.Fprintf(rw, "Request method must be 'POST'")
		if err != nil {
			changi.Log.Error(err)
		}
	} else {
		changi.Log.Debug("Requests are starting")
		newRequest := readBodyFromRequest(req)
		internal.NewRequest(newRequest)
	}
}

func fetch(rw http.ResponseWriter, req *http.Request) {
}

func readBodyFromRequest(req *http.Request) NewRequest {
	defer req.Body.Close()

	var body []byte
	req.Body.Read(body)
	var newRequest NewRequest
	err := json.Unmarshal(body, newRequest)
	if err != nil {
		changi.Log.Error(err)
	}

	return newRequest
}
