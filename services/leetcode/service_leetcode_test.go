package leetcode

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

var username = "someLeetcodeUsername"
var expectedRequestBody = fmt.Sprintf(`{"query":"query userPublicProfile($username: String!) {  matchedUser(username: $username) { username }}","variables":{"username":"%s"},"operationName":"userPublicProfile"}`, username)
var expectedRequest = network.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewReader([]byte(expectedRequestBody)))

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceLeetcode)
	assert.True(t, ok)
}

func TestServiceLeetcodeName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "leetcode", s.Name())
}

func TestServiceLeetcodeTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"it", "coding", "contests"}, s.Tags())
}

func TestServiceLeetcodeCheckErrorOnBody(t *testing.T) {
	expectedErr := fmt.Errorf("Some error")
	testClient := network.TestRESTClient{
		Error: expectedErr,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(username, &testClient)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assertRequest(t, testClient, expectedRequest)
}

func TestServiceLeetcodeCheckErrorOnUnmarhsalling(t *testing.T) {
	testClient := network.TestRESTClient{
		Error: nil,
		Body:  nil,
	}
	s := NewService()
	_, err := s.Check(username, &testClient)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "unexpected end of JSON input")
	assertRequest(t, testClient, expectedRequest)
}

func TestServiceLeetcodeCheckNotFound(t *testing.T) {
	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"data":{"matchedUser":null}}`),
	}
	s := NewService()
	exists, err := s.Check(username, &testClient)
	assert.NoError(t, err)
	assert.False(t, exists)
	assertRequest(t, testClient, expectedRequest)
}

func TestServiceLeetcodeCheckFound(t *testing.T) {
	testClient := network.TestRESTClient{
		Error: nil,
		Body:  []byte(`{"data":{"matchedUser":{"username":"` + username + `"}}}`),
	}
	s := NewService()
	exists, err := s.Check(username, &testClient)
	assert.NoError(t, err)
	assert.True(t, exists)
	assertRequest(t, testClient, expectedRequest)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedRequest *network.Request) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
