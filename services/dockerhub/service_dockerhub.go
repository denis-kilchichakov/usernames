package dockerhub

import (
	"fmt"
	"net/http"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceDockerhub struct{}

func NewService() contract.ServiceChecker {
	return &serviceDockerhub{}
}

func (s *serviceDockerhub) Name() string {
	return "dockerhub"
}

func (s *serviceDockerhub) Tags() []string {
	return []string{"it", "registry", "docker"}
}

func (s *serviceDockerhub) Check(username string, client network.RESTClient) (bool, error) {
	url := "https://hub.docker.com/v2/users/" + username + "/"

	resp, err := client.RetrieveHead(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
