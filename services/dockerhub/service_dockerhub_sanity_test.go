//go:build sanity
// +build sanity

package dockerhub

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceDockerhubUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("avgur", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceDockerhubUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("nonexistentusr", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
