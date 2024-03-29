package network

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoRequest(t *testing.T) {
	url, finalizer := MockServer("/some/path", []byte("OK"))
	defer finalizer()
	client := DefaultRESTClient{}

	resp, err := client.Do(NewRequest("GET", url+"/some/path", nil))
	assert.NoError(t, err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, []byte("OK"), body)
}

func TestRetrieveBody(t *testing.T) {
	url, finalizer := MockServer("/some/path", []byte("OK"))
	defer finalizer()
	client := DefaultRESTClient{}

	body, err := client.RetrieveBody(NewRequest("GET", url+"/some/path", nil))

	assert.NoError(t, err)
	assert.Equal(t, []byte("OK"), body)
}

func TestRequestCreationFails(t *testing.T) {
	// Create a mock request with invalid parameters
	invalidRequest := NewRequest(
		"  ",
		"http://example.com",
		nil,
	)
	client := DefaultRESTClient{}

	_, err := client.RetrieveBody(invalidRequest)

	assert.ErrorContains(t, err, "net/http: invalid method \"  \"")
}

func TestRequestDoFail(t *testing.T) {
	// Create a mock request with invalid parameters
	invalidRequest := NewRequest(
		"GET",
		"http://localhost:1234",
		nil,
	)
	client := DefaultRESTClient{}

	_, err := client.RetrieveBody(invalidRequest)

	assert.ErrorContains(t, err, "connection refused")
}

func TestRetrieveBodyNil(t *testing.T) {
	url, finalizer := MockServer("/some/path", nil)
	defer finalizer()
	client := DefaultRESTClient{}

	body, err := client.RetrieveBody(NewRequest("GET", url+"/some/path", &MockReader{}))

	assert.Nil(t, body)
	assert.ErrorContains(t, err, "mock read error")
}

func TestRetrieveHeadSuccess(t *testing.T) {
	url, finalizer := MockServer("/some/path", []byte("OK"))
	defer finalizer()
	client := DefaultRESTClient{}

	resp, err := client.RetrieveHead(url + "/some/path")
	assert.NoError(t, err)

	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 200)
}

func TestRetrieveHeadNotFound(t *testing.T) {
	url, finalizer := MockServer("/some/path", []byte("OK"))
	defer finalizer()
	client := DefaultRESTClient{}

	resp, err := client.RetrieveHead(url + "/another/path")
	assert.NoError(t, err)

	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, 404)
}
