package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/denis-kilchichakov/usernames/contract"
	"github.com/denis-kilchichakov/usernames/network"
)

type serviceLeetcode struct{}

type graphQLRequest struct {
	Query         string            `json:"query"`
	Variables     map[string]string `json:"variables"`
	OperationName string            `json:"operationName"`
}

type response struct {
	Data struct {
		MatchedUser struct {
			Username string `json:"username"`
		} `json:"matchedUser"`
	} `json:"data"`
}

func NewService() contract.ServiceChecker {
	return &serviceLeetcode{}
}

func (s *serviceLeetcode) Name() string {
	return "leetcode"
}

func (s *serviceLeetcode) Tags() []string {
	return []string{"it", "coding", "contests"}
}

func (s *serviceLeetcode) Check(username string, client network.RESTClient) (bool, error) {
	payload := &graphQLRequest{
		Query: "query userPublicProfile($username: String!) {  matchedUser(username: $username) { username }}",
		Variables: map[string]string{
			"username": username,
		},
		OperationName: "userPublicProfile",
	}
	payloadBytes, _ := json.Marshal(payload)

	req := network.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewReader(payloadBytes))
	respBody, err := client.RetrieveBody(req)
	if err != nil {
		return false, err
	}

	var result response
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return false, err
	}

	if result.Data.MatchedUser.Username == username {
		return true, nil
	} else {
		return false, nil
	}
}
