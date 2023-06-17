package gitlab

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var gitlabUsername = "someGitlabUsername"
var gitlabExpectedRequest = network.NewRequest("GET", "https://gitlab.com/api/v4/users?username="+gitlabUsername, nil)

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceGitlab)
	assert.True(t, ok)
}

func TestServiceGitlabName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "gitlab", s.Name())
}

func TestServiceGitlabTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"it", "vcs", "ci/cd"}, s.Tags())
}

func TestServiceGitlabCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[]`),
	}
	s := NewService()
	exists, err := s.Check(gitlabUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckDifferentLogin(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"username":"someotherusername"}]`),
	}
	s := NewService()
	exists, err := s.Check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckFormatChanged(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"notlogin":"` + gitlabUsername + `"}]`),
	}
	s := NewService()
	exists, err := s.Check(gitlabUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func TestServiceGitlabCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"username":"` + gitlabUsername + `"}]`),
	}
	s := NewService()
	exists, err := s.Check(gitlabUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, gitlabExpectedRequest)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedRequest *network.Request) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
