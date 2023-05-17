package services

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var gitlabUsername = "someGitlabUsername"
var gitlabExpectedRequest = network.NewRequest("GET", "https://gitlab.com/api/v4/users?username="+gitlabUsername, nil)

func TestServiceGitlabName(t *testing.T) {
	s := serviceGitlab{}
	assert.Equal(t, "gitlab", s.name())
}

func TestServiceGitlabCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := serviceGitlab{}
	_, err := s.check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[]`),
	}
	s := serviceGitlab{}
	exists, err := s.check(gitlabUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckDifferentLogin(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"username":"someotherusername"}]`),
	}
	s := serviceGitlab{}
	exists, err := s.check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckFormatChanged(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"notlogin":"` + gitlabUsername + `"}]`),
	}
	s := serviceGitlab{}
	exists, err := s.check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGithlabCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"username":"` + gitlabUsername + `"}]`),
	}
	s := serviceGitlab{}
	exists, err := s.check(gitlabUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}
