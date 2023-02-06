package test

import (
	"net/http"
	"net/http/httptest"
)

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
