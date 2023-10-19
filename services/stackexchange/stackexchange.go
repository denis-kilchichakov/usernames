package stackexchange

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/denis-kilchichakov/usernames/network"
)

type StackExchangeAPI interface {
	Check(username string, site string, client network.RESTClient) (bool, error)
}

type StackExchange struct{}

func (s *StackExchange) Check(username string, site string, client network.RESTClient) (bool, error) {
	apiURL := "https://api.stackexchange.com/2.3/users"
	params := url.Values{}
	params.Set("order", "desc")
	params.Set("sort", "reputation")
	params.Set("inname", username)
	params.Set("site", site)
	apiURL += "?" + params.Encode()

	req := network.NewRequest("GET", apiURL, nil)
	respBody, err := client.RetrieveBody(req)
	if err != nil {
		return false, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return false, err
	}

	if users, ok := data["items"].([]interface{}); ok {
		for _, user := range users {
			userData := user.(map[string]interface{})
			displayName := userData["display_name"].(string)
			if displayName == username {
				return true, nil
			}
		}

		return false, nil
	} else {
		return false, fmt.Errorf("unexpected response from stackexchange, field 'items' is missing")
	}
}
