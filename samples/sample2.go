package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	username := "feedвавыв2" // Replace with the username you want to check

	// Create the URL for the Instagram API endpoint
	url := "https://instagram.com/_u/" + username + "/"

	// Send the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		//fmt.Println(string(body))

		// Check if the response body contains the substring
		if strings.Contains(string(body), fmt.Sprintf(`{"username":"%s"}`, username)) {
			fmt.Println("Username exists on Instagram")
		} else {
			fmt.Println("Username does not exist on Instagram")
		}
	} else if response.StatusCode == http.StatusNotFound {
		fmt.Println("Username does not exist on Instagram")
	} else {
		fmt.Println("Error occurred. Status code:", response.StatusCode)
	}
}
