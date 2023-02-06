package network

import (
	"testing"
	"uzrnames/pkg/test"

	"github.com/stretchr/testify/assert"
)

func TestGetBody(t *testing.T) {
	url, finalizer := test.MockServer("/some/path", []byte("OK"))
	defer finalizer()

	body, err := GetBody(url + "/some/path")

	assert.NoError(t, err)
	assert.Equal(t, []byte("OK"), body)
}
