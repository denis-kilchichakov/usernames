# usernames

Library that checks if given username exists on popular internet services.

# Installation
```
go get github.com/denis-kilchichakov/usernames@v0.1.1
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
	parallelism := 1
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
```

## Tests

To launch sanity checks (tests with pairs of existing and non-existing usernames on supported services, to validate the algorithm):

```go
go test ./... -tags=sanity
```

## Supported services:
* GitHub

## How to add more services
Inside package `services`:
1. Implement `serviceChecker` interface
1. Call `registerService()`

## Services planned to be added in future:
* GitLab
* DockerHub
* LeetCode
* CodeWars
* ...
* you name it
