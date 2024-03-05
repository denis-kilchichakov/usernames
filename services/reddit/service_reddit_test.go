package reddit

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var redditUsername = "someusername"

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceReddit)
	assert.True(t, ok)
}

func TestServiceRedditName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "reddit", s.Name())
}

func TestServiceRedditTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"social", "entertainment", "q&a"}, s.Tags())
}

func TestServiceRedditCheckErrorOnDo(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(redditUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error:      nil,
		Body:       nil,
		StatusCode: 404,
	}
	s := NewService()
	exists, err := s.Check(redditUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckOtherError(t *testing.T) {

	testClient := network.TestRESTClient{
		Error:      nil,
		Body:       nil,
		StatusCode: 500,
	}
	s := NewService()
	exists, err := s.Check(redditUsername, &testClient)
	assert.Error(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient)
}

func TestServiceGithubCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error:      nil,
		Body:       nil,
		StatusCode: 200,
	}
	s := NewService()
	exists, err := s.Check(redditUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient) {
	var expectedRequest = network.NewRequest("GET", "https://www.reddit.com/user/"+redditUsername+"/about.json", nil)
	expectedRequest.SetHeader("User-Agent", "usernames")

	assert.Equal(t, 1, len(testClient.DoRequests))
	assert.Equal(t, expectedRequest, testClient.DoRequests[0])
}
