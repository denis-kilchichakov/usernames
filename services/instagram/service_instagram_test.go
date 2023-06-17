package instagram

import (
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var expectedUsername = "someGitlabUsername"
var expectedRequest = network.NewRequest("GET", "https://instagram.com/_u/"+expectedUsername+"/", nil)

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceInstagram)
	assert.True(t, ok)
}

func TestServiceInstagramName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "instagram", s.Name())
}

func TestServiceInstagramTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"social", "photo", "video"}, s.Tags())
}

func TestServiceInstagramCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(expectedUsername, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, expectedRequest)
}

func TestServiceInstagramCheckNotFound(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[]`),
	}
	s := NewService()
	exists, err := s.Check(expectedUsername, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, expectedRequest)
}

func TestServiceInstagramCheckPassed(t *testing.T) {

	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`[{"username":"` + expectedUsername + `"}]`),
	}
	s := NewService()
	exists, err := s.Check(expectedUsername, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, expectedRequest)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedRequest *network.Request) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
