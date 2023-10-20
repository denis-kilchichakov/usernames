package network

import (
	"io"
	"net/http"
)

type RESTClient interface {
	RetrieveBody(request *Request) ([]byte, error)
	RetrieveHead(url string) (*http.Response, error)
}

type DefaultRESTClient struct{}

func (c *DefaultRESTClient) RetrieveBody(request *Request) ([]byte, error) {
	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *DefaultRESTClient) RetrieveHead(url string) (*http.Response, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
