package network

import (
	"net/http"
	"net/http/httptest"
)

type TestRESTClient struct {
	Body     []byte
	Error    error
	Requests []*Request
}

func (c *TestRESTClient) RetrieveBody(request *Request) ([]byte, error) {
	c.Requests = append(c.Requests, request)
	return c.Body, c.Error
}

func MockServer(path string, body []byte) (url string, finalizer func()) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() == path {
			rw.Write(body)
		} else {
			rw.WriteHeader(404)
		}
	}))

	url = server.URL
	finalizer = func() {
		server.Close()
	}

	return
}