# usernames

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
