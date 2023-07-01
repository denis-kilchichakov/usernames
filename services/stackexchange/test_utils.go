package stackexchange

import (
	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/mock"
)

type MockStackExchange struct {
	checks []mock.Arguments
	mock.Mock
}

func NewMockStackExchange() *MockStackExchange {
	return &MockStackExchange{}
}

func (m *MockStackExchange) Check(username string, site string, client network.RESTClient) (bool, error) {
	m.checks = append(m.checks, m.Called(username, site, client))
	args := m.Called(username, site, client)
	return args.Bool(0), args.Error(1)
}
