//go:build sanity
// +build sanity

package gitlab

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/network"
	"github.com/stretchr/testify/assert"
)

func TestServiceGitlabUsernameExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGitlab{}
	exists, err := s.Check("abryp", &c)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestServiceGitlabUsernameNotExist(t *testing.T) {
	c := network.DefaultRESTClient{}
	s := serviceGitlab{}
	exists, err := s.Check("pyrba", &c)
	assert.NoError(t, err)
	assert.False(t, exists)
}
