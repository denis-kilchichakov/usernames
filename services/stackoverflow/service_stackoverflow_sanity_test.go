//go:build sanity
// +build sanity

package stackoverflow

import (
	"testing"
	"time"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceStackoverflowUsernameExist(t *testing.T) {
	time.Sleep(time.Duration(time.Second * 2)) // to avoid rate limit
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("augur", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceStackoverflowUsernameNotExist(t *testing.T) {
	time.Sleep(time.Duration(time.Second * 2)) // to avoid rate limit
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("AwSdFghjkl", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
