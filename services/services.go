package services

import (
	"fmt"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
	"github.com/denis-kilchichakov/usernames/services/github"
	"github.com/denis-kilchichakov/usernames/services/gitlab"
	"github.com/denis-kilchichakov/usernames/services/instagram"
	"github.com/denis-kilchichakov/usernames/services/leetcode"
)

var servicesByName map[string]contract.ServiceChecker = make(map[string]contract.ServiceChecker)
var servicesByTag map[string][]contract.ServiceChecker = map[string][]contract.ServiceChecker{}

func init() {
	registerService(github.NewService())
	registerService(gitlab.NewService())
	registerService(leetcode.NewService())
	registerService(instagram.NewService())
}

func registerService(service contract.ServiceChecker) error {
	if s, ok := servicesByName[service.Name()]; ok {
		if s != service {
			return fmt.Errorf("service with name %s already exists", service.Name())
		}
	}

	servicesByName[service.Name()] = service

	for _, tag := range service.Tags() {
		servicesByTag[tag] = append(servicesByTag[tag], service)
	}

	return nil
}

func GetSupportedServiceNames() []string {
	names := make([]string, 0, len(servicesByName))
	for name := range servicesByName {
		names = append(names, name)
	}
	return names
}

func GetSupportedServiceTags() []string {
	tags := make([]string, 0, len(servicesByTag))
	for tag := range servicesByTag {
		tags = append(tags, tag)
	}
	return tags
}

func GetSupportedServiceNamesByTag(tag string) []string {
	if services, ok := servicesByTag[tag]; ok {
		names := make([]string, 0, len(services))
		for _, service := range services {
			names = append(names, service.Name())
		}
		return names
	}
	return []string{}
}

func Check(service string, username string, client network.RESTClient) (bool, error) {
	if s, ok := servicesByName[service]; ok {
		return s.Check(username, client)
	}
	return false, fmt.Errorf("service %s is not supported", service)
}
