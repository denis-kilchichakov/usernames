package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveBody(t *testing.T) {
	url, finalizer := MockServer("/some/path", []byte("OK"))
	defer finalizer()
	client := DefaultRESTClient{}

	body, err := client.RetrieveBody(NewRequest("GET", url+"/some/path", nil))

	assert.NoError(t, err)
	assert.Equal(t, []byte("OK"), body)
}
