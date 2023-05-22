package github

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var githubUsername = "someusername"
var githubExpectedRequest = network.NewRequest("GET", "https://api.github.com/users/"+githubUsername, nil)

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceGithub)
	assert.True(t, ok)
}

func TestServiceGithubName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "github", s.Name())
}

func TestServiceGithubTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"it", "social", "vcs"}, s.Tags())
}

func TestServiceGithubCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"message":"Not Found"}`),
	}
	s := NewService()
	exists, err := s.Check(githubUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckDifferentLogin(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"someotherusername"}`),
	}
	s := NewService()
	exists, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckFormatChanged(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"notlogin":"` + githubUsername + `"}`),
	}
	s := NewService()
	exists, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"` + githubUsername + `"}`),
	}
	s := NewService()
	exists, err := s.Check(githubUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, githubExpectedRequest)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedRequest *network.Request) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
