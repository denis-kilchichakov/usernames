package github

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var githubUsername = "someusername"
var githubExpectedRequest = network.NewRequest("GET", "https://api.github.com/users/"+githubUsername, nil)

func TestServiceGithubName(t *testing.T) {
	s := serviceGithub{}
	assert.Equal(t, "github", s.Name())
}

func TestServiceGithubCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := serviceGithub{}
	_, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	network.AssertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"message":"Not Found"}`),
	}
	s := serviceGithub{}
	exists, err := s.Check(githubUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	network.AssertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckDifferentLogin(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"someotherusername"}`),
	}
	s := serviceGithub{}
	exists, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	network.AssertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckFormatChanged(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"notlogin":"` + githubUsername + `"}`),
	}
	s := serviceGithub{}
	exists, err := s.Check(githubUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	network.AssertRequest(t, testClient, githubExpectedRequest)
}

func TestServiceGithubCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"login":"` + githubUsername + `"}`),
	}
	s := serviceGithub{}
	exists, err := s.Check(githubUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	network.AssertRequest(t, testClient, githubExpectedRequest)
}