package instagram

import (
	"fmt"
	"strings"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceInstagram struct{}

func NewService() contract.ServiceChecker {
	return &serviceInstagram{}
}

func (s *serviceInstagram) Name() string {
	return "instagram"
}

func (s *serviceInstagram) Tags() []string {
	return []string{"social", "photo", "video"}
}

func (s *serviceInstagram) Check(username string, client network.RESTClient) (bool, error) {
	req := network.NewRequest("GET", "https://instagram.com/_u/"+username+"/", nil)
	respBody, err := client.RetrieveBody(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return false, err
	}

	if strings.Contains(string(respBody), fmt.Sprintf(`{"username":"%s"}`, username)) {
		return true, nil
	} else {
		return false, nil
	}
}
