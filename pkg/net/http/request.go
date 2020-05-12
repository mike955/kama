package http

import "net/http"

type Request struct {
	*http.Request
	Params map[string]interface{}
}


func createRequest(raw *http.Request) *Request {
	return &Request{
		raw,
		make(map[string]interface{}),
	}
}
