# usernames

[![Coverage Status](https://coveralls.io/repos/github/denis-kilchichakov/usernames/badge.svg?branch=main)](https://coveralls.io/github/denis-kilchichakov/usernames?branch=main)
![Tests+Sanity](https://github.com/denis-kilchichakov/usernames/actions/workflows/goveralls.yml/badge.svg)

Library that checks if given username exists on popular internet services.

# Installation
```
go get github.com/denis-kilchichakov/usernames@v0.5.0
```

# Usage
Simple example:
```go
package main

import (
	"fmt"

	"github.com/denis-kilchichakov/usernames"
)

func main() {
	u := "username-to-check"
	parallelism := 2
	checkResults := usernames.CheckAll(u, parallelism)
	fmt.Printf("Check results for username \"%s\":\n", u)
	for _, r := range checkResults {
		if r.Err != nil {
			fmt.Printf("Error on %s: %v\n", r.Service, r.Err)
			continue
		}
		fmt.Printf("%s: %v\n", r.Service, r.Found)
	}
}
```
Output:
```
Check results for username "username-to-check":
github: false
gitlab: false
instagram: false
leetcode: false
stackoverflow: false
```

## Tests

To launch sanity checks (which do requests to actual services instead of mocks in unit tests):
```go
go test ./... -tags=sanity
```

## Working in VS Code

For sanity test files to be processed correctly by VS Code, consider adding this to `settings.json`:
```
    "gopls.buildFlags": ["-tags=sanity"],
```

## Supported services:
* GitHub
* GitLab
* LeetCode
* Instagram
* StackOverflow

## How to add more services
Inside package `services`:
1. Create subpackage for your service
1. Implement `contract.ServiceChecker` interface
1. Add `registerService()` call in `init()` (`services` package)

## Services planned to be added in future:
* DockerHub
* CodeWars
* HackerRank
* Habr
* Kaggle
* Telegram
* Reddit
* Twitter
* LinkedIn
* Facebook
* Google
* TikTok
* ...
* you name it
