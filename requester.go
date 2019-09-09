package changi

import "net/http"

//@TODO: Implement methods for requests

type request struct {
	req     http.Request
	timeout float32
}



func makeRequest(req http.Request, response chan http.Response, progress chan string){
	//Implement making a request
}