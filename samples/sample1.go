package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const LEETCODE_API_ENDPOINT = "https://leetcode.com/graphql"

type GraphQLRequest struct {
	Query         string            `json:"query"`
	Variables     map[string]string `json:"variables"`
	OperationName string            `json:"operationName"`
}

func main1() {
	fmt.Println("Fetching daily coding challenge from LeetCode API.")

	payload := &GraphQLRequest{
		Query: "query userPublicProfile($username: String!) {  matchedUser(username: $username) { username }}",
		Variables: map[string]string{
			"username": "augur",
		},
		OperationName: "userPublicProfile",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling payload: %v", err)
		return
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", LEETCODE_API_ENDPOINT, body)
	if err != nil {
		fmt.Printf("Error creating request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v", err)
		return
	}

	var result map[string]interface{}

	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		fmt.Printf("Error unmarshalling response body: %v", err)
		return
	}

	fmt.Println(result)
}
