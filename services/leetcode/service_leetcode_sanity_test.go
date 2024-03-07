//go:build sanity && exclude_gh
// +build sanity,exclude_gh

package leetcode

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceLeetcodeUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("kdenis87", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceLeetcodeUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := NewService()
	exists, err := s.Check("kdenis11", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
