package gitlab

import (
	"encoding/json"
	"fmt"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceGitlab struct{}

func CreateService() contract.ServiceChecker {
	return &serviceGitlab{}
}

func (s *serviceGitlab) Name() string {
	return "gitlab"
}

func (s *serviceGitlab) Tags() []string {
	return []string{"it", "vcs", "ci/cd"}
}

func (s *serviceGitlab) Check(username string, client network.RESTClient) (bool, error) {
	req := network.NewRequest("GET", "https://gitlab.com/api/v4/users?username="+username, nil)
	respBody, err := client.RetrieveBody(req)
	if err != nil {
		return false, err
	}

	var result []map[string]interface{}
	json.Unmarshal(respBody, &result)

	if len(result) == 0 {
		return false, nil
	}

	login, ok := result[0]["username"]
	if !ok || login != username {
		return false, fmt.Errorf("unexpected response from gitlab: %s", respBody)
	}

	return true, nil
}
