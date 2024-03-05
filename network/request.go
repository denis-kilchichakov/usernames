package network

import "io"

type Request struct {
	method  string
	url     string
	body    io.Reader
	headers map[string]string
}

func NewRequest(method, url string, body io.Reader) *Request {
	return &Request{
		method:  method,
		url:     url,
		body:    body,
		headers: make(map[string]string),
	}
}

func (r *Request) SetHeader(key, value string) {
	r.headers[key] = value
}
