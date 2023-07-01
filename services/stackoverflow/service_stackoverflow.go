package stackoverflow

import (
	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
	"github.com/denis-kilchichakov/usernames/services/stackexchange"
)

type serviceStackoverflow struct {
	api stackexchange.StackExchangeAPI
}

func NewService() contract.ServiceChecker {
	return &serviceStackoverflow{
		api: &stackexchange.StackExchange{},
	}
}

func (s *serviceStackoverflow) Name() string {
	return "stackoverflow"
}

func (s *serviceStackoverflow) Tags() []string {
	return []string{"it", "q&a", "social"}
}

func (s *serviceStackoverflow) Check(username string, client network.RESTClient) (bool, error) {
	return s.api.Check(username, "stackoverflow", client)
}
