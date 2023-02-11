package network

import (
	"io/ioutil"
	"net/http"
)

type RESTClient interface {
	RetrieveBody(request *Request) ([]byte, error)
}

type DefaultRESTClient struct{}

func (c *DefaultRESTClient) RetrieveBody(request *Request) ([]byte, error) {
	// create a new request using http

	req, err := http.NewRequest(request.method, request.url, request.body)
	if err != nil {
		return nil, err
	}

	// add an user-agent header to the request
	// req.Header.Add("User-Agent", "UserNames-App")

	// send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
