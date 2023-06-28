//go:build sanity
// +build sanity

package instagram

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceInstagramUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("zuck", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceInstagramUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("reallynonexistentinstauser", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
