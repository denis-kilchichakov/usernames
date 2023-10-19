package stackexchange

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

// check error on request
func TestStackExchangeErrorOnRequest(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	mockClient := &network.TestRESTClient{
		Body:  nil,
		Error: expectedErr,
	}

	stackExchange := StackExchange{}

	result, err := stackExchange.Check("someUser", "stackoverflow", mockClient)

	assert.Equal(t, expectedErr, err)
	assert.False(t, result)
	request := mockClient.Requests[0]
	expectedURL := "https://api.stackexchange.com/2.3/users?inname=someUser&order=desc&site=stackoverflow&sort=reputation"
	expectedRequest := network.NewRequest("GET", expectedURL, nil)
	assert.Len(t, mockClient.Requests, 1)
	assert.Equal(t, expectedRequest, request)
}

func TestStackExchangeErrorOnJson(t *testing.T) {
	mockClient := &network.TestRESTClient{
		Body:  []byte(`malformed json`),
		Error: nil,
	}
	stackExchange := StackExchange{}

	result, err := stackExchange.Check("someUser", "stackoverflow", mockClient)

	assert.Error(t, err)
	assert.False(t, result)
	request := mockClient.Requests[0]
	expectedURL := "https://api.stackexchange.com/2.3/users?inname=someUser&order=desc&site=stackoverflow&sort=reputation"
	expectedRequest := network.NewRequest("GET", expectedURL, nil)
	assert.Len(t, mockClient.Requests, 1)
	assert.Equal(t, expectedRequest, request)
}

func TestStackExchangeFound(t *testing.T) {
	mockClient := &network.TestRESTClient{
		Body: []byte(`{
			"items": [
				{
					"display_name": "AwesomeUser"
				},
				{
					"display_name": "someUser"
				},
				{
					"display_name": "someUser3"
				}
			]
		}`),
		Error: nil,
	}
	stackExchange := StackExchange{}

	result, err := stackExchange.Check("someUser", "stackoverflow", mockClient)

	assert.NoError(t, err)
	assert.True(t, result)
	request := mockClient.Requests[0]
	expectedURL := "https://api.stackexchange.com/2.3/users?inname=someUser&order=desc&site=stackoverflow&sort=reputation"
	expectedRequest := network.NewRequest("GET", expectedURL, nil)
	assert.Len(t, mockClient.Requests, 1)
	assert.Equal(t, expectedRequest, request)
}

func TestStackExchangeNotFound(t *testing.T) {
	mockClient := &network.TestRESTClient{
		Body: []byte(`{
			"items": [
				{
					"display_name": "AwesomeUser"
				},
				{
					"display_name": "someUser3"
				}
			]
		}`),
		Error: nil,
	}
	stackExchange := StackExchange{}

	result, err := stackExchange.Check("someUser", "stackoverflow", mockClient)

	assert.NoError(t, err)
	assert.False(t, result)
	request := mockClient.Requests[0]
	expectedURL := "https://api.stackexchange.com/2.3/users?inname=someUser&order=desc&site=stackoverflow&sort=reputation"
	expectedRequest := network.NewRequest("GET", expectedURL, nil)
	assert.Len(t, mockClient.Requests, 1)
	assert.Equal(t, expectedRequest, request)
}

func TestStackExchangeUnexpectedResponseBody(t *testing.T) {
	mockClient := &network.TestRESTClient{
		Body:  []byte(`{"not_items": [{"display_name": "AwesomeUser"}]}`),
		Error: nil,
	}
	stackExchange := StackExchange{}

	result, err := stackExchange.Check("someUser", "stackoverflow", mockClient)

	assert.ErrorContains(t, err, "unexpected response")
	assert.False(t, result)
}
