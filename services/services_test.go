package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterServiceAddedTwoSuccessfully(t *testing.T) {
	s1 := &MockServiceChecker{
		_name: "name1",
		_tags: []string{"tag1", "tag2"},
	}
	s2 := &MockServiceChecker{
		_name: "name2",
		_tags: []string{"tag2", "tag3"},
	}

	before := len(servicesByName)

	err := registerService(s1)
	assert.NoError(t, err)
	err = registerService(s2)
	assert.NoError(t, err)

	assert.Equal(t, before+2, len(servicesByName))
	assert.Equal(t, s1, servicesByName[s1.name()])
	assert.Equal(t, s2, servicesByName[s2.name()])
	assert.Contains(t, servicesByTag["tag1"], s1)
	assert.Contains(t, servicesByTag["tag2"], s1)
	assert.Contains(t, servicesByTag["tag2"], s2)
	assert.Contains(t, servicesByTag["tag3"], s2)
}

func TestRegisterSameServiceTwiceNoError(t *testing.T) {
	s1 := &MockServiceChecker{
		_name: "name3",
	}

	before := len(servicesByName)

	err := registerService(s1)
	assert.NoError(t, err)
	assert.Equal(t, before+1, len(servicesByName))
	err = registerService(s1)
	assert.NoError(t, err)
	assert.Equal(t, before+1, len(servicesByName))
}

func TestRegisterServicesWithSameNameFail(t *testing.T) {
	s1 := &MockServiceChecker{
		_name: "name4",
	}
	s2 := &MockServiceChecker{
		_name: "name4",
	}

	before := len(servicesByName)

	err := registerService(s1)
	assert.NoError(t, err)
	assert.Equal(t, before+1, len(servicesByName))
	err = registerService(s2)
	assert.Error(t, err)
	assert.Equal(t, before+1, len(servicesByName))
}

func TestGetSupportedServiceNames(t *testing.T) {
	s1 := &MockServiceChecker{
		_name: "name5",
	}
	s2 := &MockServiceChecker{
		_name: "name6",
	}

	err := registerService(s1)
	assert.NoError(t, err)
	err = registerService(s2)
	assert.NoError(t, err)

	names := GetSupportedServiceNames()
	assert.Contains(t, names, s1.name())
	assert.Contains(t, names, s2.name())
}

func TestGetSupportedServiceTags(t *testing.T) {
	s1 := &MockServiceChecker{
		_name: "name7",
		_tags: []string{"TestGetSupportedServiceTags1", "TestGetSupportedServiceTags2"},
	}
	s2 := &MockServiceChecker{
		_name: "name8",
		_tags: []string{"TestGetSupportedServiceTags2", "TestGetSupportedServiceTags3"},
	}

	err := registerService(s1)
	assert.NoError(t, err)
	err = registerService(s2)
	assert.NoError(t, err)

	tags := GetSupportedServiceTags()
	assert.Contains(t, tags, "TestGetSupportedServiceTags1")
	assert.Contains(t, tags, "TestGetSupportedServiceTags2")
	assert.Contains(t, tags, "TestGetSupportedServiceTags3")
}

func TestGetSupportedServiceTagsByNonExistentTag(t *testing.T) {
	names := GetSupportedServiceNamesByTag("non-existent-tag")
	assert.Empty(t, names)
}
