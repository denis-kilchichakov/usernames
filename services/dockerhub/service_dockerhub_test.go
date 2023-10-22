package dockerhub

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var dockerhubUsername = "someusername"
var dockerhubExpectedUrl = "https://hub.docker.com/v2/users/" + dockerhubUsername + "/"

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceDockerhub)
	assert.True(t, ok)
}

func TestServiceDockerhubName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "dockerhub", s.Name())
}

func TestServiceDockerhubTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"it", "registry", "docker"}, s.Tags())
}

func TestServiceDockerhubCheckErrorOnHead(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error:        expectedErr,
		HeadResponse: &http.Response{},
	}
	s := NewService()
	_, err := s.Check(dockerhubUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, dockerhubExpectedUrl)
}

func TestServiceGithubCheckFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		HeadResponse: &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body: &network.MockReadCloser{
				ReadBody:   []byte("Hi"),
				CloseError: nil,
			},
		},
	}
	s := NewService()
	exists, err := s.Check(dockerhubUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, dockerhubExpectedUrl)
}

func TestServiceGithubCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		HeadResponse: &http.Response{
			Status:     "404 Not found",
			StatusCode: 404,
			Body: &network.MockReadCloser{
				ReadBody:   []byte("Bye"),
				CloseError: nil,
			},
		},
	}
	s := NewService()
	exists, err := s.Check(dockerhubUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, dockerhubExpectedUrl)
}

func TestServiceGithubCheckUnexpectedResponse(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		HeadResponse: &http.Response{
			Status:     "403 Forbidden",
			StatusCode: 403,
			Body: &network.MockReadCloser{
				ReadBody:   []byte("Wat"),
				CloseError: nil,
			},
		},
	}
	s := NewService()
	exists, err := s.Check(dockerhubUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, dockerhubExpectedUrl)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedUrl string) {
	assert.Equal(t, 1, len(testClient.HeadRequests))
	assert.Equal(t, expectedUrl, *testClient.HeadRequests[0])
}
