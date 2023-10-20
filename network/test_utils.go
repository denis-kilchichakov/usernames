package network

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
)

type TestRESTClient struct {
	Body         []byte
	Error        error
	Requests     []*Request
	HeadRequests []*string
	HeadResponse *http.Response
}

func (c *TestRESTClient) RetrieveBody(request *Request) ([]byte, error) {
	c.Requests = append(c.Requests, request)
	return c.Body, c.Error
}

func (c *TestRESTClient) RetrieveHead(url string) (*http.Response, error) {
	c.HeadRequests = append(c.HeadRequests, &url)
	return c.HeadResponse, c.Error
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

type MockReader struct{}

func (m *MockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mock read error")
}

type MockReadCloser struct {
	ReadBody    []byte
	readCounter int
	CloseError  error
}

func (mrc *MockReadCloser) Read(p []byte) (n int, err error) {
	if mrc.readCounter >= len(mrc.ReadBody) {
		return 0, io.EOF
	}

	bytesToCopy := copy(p, mrc.ReadBody[mrc.readCounter:])
	mrc.readCounter += bytesToCopy

	return bytesToCopy, nil
}

func (mrc *MockReadCloser) Close() error {
	return mrc.CloseError
}
