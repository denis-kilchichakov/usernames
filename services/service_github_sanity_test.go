//go:build sanity
// +build sanity

package services

import (
	"testing"

	"github.com/denis_kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceGithubUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGithub{}
	exists, err := s.check("denis-kilchichakov", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceGithubUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGithub{}
	exists, err := s.check("kenis-dilchichakov", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
