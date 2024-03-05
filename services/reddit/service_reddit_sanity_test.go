//go:build sanity
// +build sanity

package reddit

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceRedditUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("kn0thing", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceRedditUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("someNonExistentGuyRly", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
