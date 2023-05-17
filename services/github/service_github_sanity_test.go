//go:build sanity
// +build sanity

package github

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceGithubUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGithub{}
	exists, err := s.Check("denis-kilchichakov", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceGithubUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGithub{}
	exists, err := s.Check("kenis-dilchichakov", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
