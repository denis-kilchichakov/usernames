package services

import (
	"fmt"

	"github.com/denis_kilchichakov/usernames/network"
)

type serviceChecker interface {
	name() string
	tags() []string
	check(username string, client network.RESTClient) (bool, error)
}

var servicesByName map[string]serviceChecker = make(map[string]serviceChecker)
var servicesByTag map[string][]serviceChecker = map[string][]serviceChecker{}

func registerService(service serviceChecker) error {
	if s, ok := servicesByName[service.name()]; ok {
		if s != service {
			return fmt.Errorf("service with name %s already exists", service.name())
		}
	}

	servicesByName[service.name()] = service

	for _, tag := range service.tags() {
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
			names = append(names, service.name())
		}
		return names
	}
	return []string{}
}

func Check(service string, username string, client network.RESTClient) (bool, error) {
	if s, ok := servicesByName[service]; ok {
		return s.check(username, client)
	}
	return false, fmt.Errorf("service %s is not supported", service)
}
