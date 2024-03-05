package network

import (
	"io"
	"net/http"
)

type RESTClient interface {
	Do(request *Request) (*http.Response, error)
	RetrieveBody(request *Request) ([]byte, error)
	RetrieveHead(url string) (*http.Response, error)
}

type DefaultRESTClient struct{}

func (c *DefaultRESTClient) Do(request *Request) (*http.Response, error) {
	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range request.headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *DefaultRESTClient) RetrieveBody(request *Request) ([]byte, error) {
	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
