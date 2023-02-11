package services

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var username = "someusername"
var expectedRequest = network.NewRequest("GET", "https://api.github.com/users/"+username, nil)

func TestServiceGithubName(t *testing.T) {
	s := serviceGithub{}
	assert.Equal(t, "github", s.name())
}

func TestServiceGithubCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := serviceGithub{}
	_, err := s.check(username, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"message":"Not Found"}`),
	}
	s := serviceGithub{}
	exists, err := s.check(username, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckDifferentLogin(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"someotherusername"}`),
	}
	s := serviceGithub{}
	exists, err := s.check(username, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckFormatChanged(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"notlogin":"` + username + `"}`),
	}
	s := serviceGithub{}
	exists, err := s.check(username, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"` + username + `"}`),
	}
	s := serviceGithub{}
	exists, err := s.check(username, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
