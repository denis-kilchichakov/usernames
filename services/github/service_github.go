package github

import (
	"encoding/json"
	"fmt"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceGithub struct{}

func CreateService() contract.ServiceChecker {
	return &serviceGithub{}
}

func (s *serviceGithub) Name() string {
	return "github"
}

func (s *serviceGithub) Tags() []string {
	return []string{"it", "social", "vcs"}
}

func (s *serviceGithub) Check(username string, client network.RESTClient) (bool, error) {
	req := network.NewRequest("GET", "https://api.github.com/users/"+username, nil)
	respBody, err := client.RetrieveBody(req)
	if err != nil {
		return false, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(respBody), &result)

	if result["message"] == "Not Found" {
		return false, nil
	}

	login, ok := result["login"]
	if !ok || login != username {
		return false, fmt.Errorf("unexpected response from github: %s", respBody)
	}

	return true, nil
}
