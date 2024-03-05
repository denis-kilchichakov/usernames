package reddit

import (
	"fmt"
	"net/http"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceReddit struct{}

func NewService() contract.ServiceChecker {
	return &serviceReddit{}
}

func (s *serviceReddit) Name() string {
	return "reddit"
}

func (s *serviceReddit) Tags() []string {
	return []string{"social", "entertainment", "q&a"}
}

func (s *serviceReddit) Check(username string, client network.RESTClient) (bool, error) {
	url := fmt.Sprintf("https://www.reddit.com/user/%s/about.json", username)
	req := network.NewRequest("GET", url, nil)
	req.SetHeader("User-Agent", "usernames")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return false, fmt.Errorf("failed with status code: %d", resp.StatusCode)
	}
}
