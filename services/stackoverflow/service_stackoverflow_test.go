package stackoverflow

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/denis-kilchichakov/usernames/services/stackexchange"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewService(t *testing.T) {
	s := NewService()
	assert.NotNil(t, s)
	_, ok := s.(*serviceStackoverflow)
	assert.True(t, ok)
}

func TestServiceStackoverflowName(t *testing.T) {
	s := NewService()
	assert.Equal(t, "stackoverflow", s.Name())
}

func TestServiceStackoverflowTags(t *testing.T) {
	s := NewService()
	assert.Equal(t, []string{"it", "q&a", "social"}, s.Tags())
}

func TestServiceStackoverflowCheck(t *testing.T) {
	username := "some username"
	m := stackexchange.NewMockStackExchange()
	m.Mock.On("Check", username, "stackoverflow", mock.Anything).Return(true, nil)
	s := NewService()
	s.(*serviceStackoverflow).api = m

	found, err := s.Check(username, &network.TestRESTClient{})

	assert.NoError(t, err)
	assert.True(t, found)
	m.AssertExpectations(t)
}
