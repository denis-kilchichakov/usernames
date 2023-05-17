package contract

import "github.com/denis-kilchichakov/usernames/network"

type ServiceChecker interface {
	Name() string
	Tags() []string
	Check(username string, client network.RESTClient) (bool, error)
}
