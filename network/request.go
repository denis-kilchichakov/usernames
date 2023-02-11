package network

import "io"

type Request struct {
	method string
	url    string
	body   io.Reader
}

func NewRequest(method, url string, body io.Reader) *Request {
	return &Request{
		method: method,
		url:    url,
		body:   body,
	}
}
