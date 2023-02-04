package network

import (
	"io/ioutil"
	"net/http"
)

// function to request a URL and return the response body
func GetBody(url string) ([]byte, error) {
	// create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// add an user-agent header to the request
	req.Header.Add("User-Agent", "Awesome-Octocat-App")

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
