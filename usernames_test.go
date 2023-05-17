package usernames

import (
	"testing"

	"github.com/denis-kilchichakov/usernames/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/slices"
)

var mockServiceChecker1 = services.NewMockServiceChecker("usernames_tests_mock1", []string{"utmt1", "utmt2"})
var mockServiceChecker2 = services.NewMockServiceChecker("usernames_tests_mock2", []string{"utmt2", "utmt3"})
var mockServiceChecker3 = services.NewMockServiceChecker("usernames_tests_mock3", []string{"utmt3", "utmt4"})

func TestCheck(t *testing.T) {
	services := []string{"usernames_tests_mock1", "usernames_tests_mock2", "non-existent_service"}
	username := "some_username"

	mockServiceChecker1.Mock.On("Check", username, mock.Anything).Return(true, nil)
	mockServiceChecker2.Mock.On("Check", username, mock.Anything).Return(false, nil)

	results := Check(services, username, 3)

	mockServiceChecker1.AssertExpectations(t)
	mockServiceChecker2.AssertExpectations(t)
	assert.Equal(t, 3, len(results))
	assertCheckResult(t, results, "usernames_tests_mock1", true, false)
	assertCheckResult(t, results, "usernames_tests_mock2", false, false)
	assertCheckResult(t, results, "non-existent_service", false, true)
}

func TestExcluding(t *testing.T) {
	username := "some_username"
	excludedServices := SupportedServices()

	// remove usernames_tests_mock1 from excludedServices
	i := slices.IndexFunc(excludedServices, func(s string) bool { return s == "usernames_tests_mock1" })
	excludedServices = slices.Delete(excludedServices, i, i+1)

	mockServiceChecker1.Mock.On("Check", username, mock.Anything).Return(true, nil)

	results := CheckExcluding(excludedServices, username, 3)

	mockServiceChecker1.AssertExpectations(t)

	assert.Equal(t, 1, len(results))
	assertCheckResult(t, results, "usernames_tests_mock1", true, false)
}

func TestCheckByTags(t *testing.T) {
	tags := []string{"utmt1", "utmt4"}
	username := "some_username"
	mockServiceChecker1.Mock.On("Check", username, mock.Anything).Return(true, nil)
	mockServiceChecker3.Mock.On("Check", username, mock.Anything).Return(false, nil)

	results := CheckByTags(tags, username, 3)

	mockServiceChecker1.AssertExpectations(t)
	mockServiceChecker3.AssertExpectations(t)
	assert.Equal(t, 2, len(results))
	assertCheckResult(t, results, "usernames_tests_mock1", true, false)
	assertCheckResult(t, results, "usernames_tests_mock3", false, false)
}

func assertCheckResult(t *testing.T, results []CheckResult, service string, found bool, err bool) {
	i := slices.IndexFunc(results, func(r CheckResult) bool { return r.Service == service })
	assert.Equal(t, found, results[i].Found)
	if !err {
		assert.NoError(t, results[i].Err)
	} else {
		assert.Error(t, results[i].Err)
	}
}
