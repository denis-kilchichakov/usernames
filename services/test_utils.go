package services

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServiceChecker struct {
	_name  string
	_tags  []string
	checks []mock.Arguments
	mock.Mock
}

func NewMockServiceChecker(name string, tags []string) *MockServiceChecker {
	m := &MockServiceChecker{
		_name: name,
		_tags: tags,
	}
	registerService(m)
	return m
}

func (m *MockServiceChecker) name() string {
	return m._name
}

func (m *MockServiceChecker) tags() []string {
	return m._tags
}

func (m *MockServiceChecker) check(username string, client network.RESTClient) (bool, error) {
	m.checks = append(m.checks, m.Called(username, client))
	args := m.Called(username, client)
	return args.Bool(0), args.Error(1)
}

func assertRequest(t *testing.T, testClient network.TestRESTClient, expectedRequest *network.Request) {
	assert.Equal(t, 1, len(testClient.Requests))
	assert.Equal(t, expectedRequest, testClient.Requests[0])
}
