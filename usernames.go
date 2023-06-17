package usernames

import (
	"github.com/denis-kilchichakov/usernames/network"
	"github.com/denis-kilchichakov/usernames/services"
)

type CheckResult struct {
	Service string
	Err     error
	Found   bool
}

func SupportedServices() []string {
	return services.GetSupportedServiceNames()
}

func SupportedTags() []string {
	return services.GetSupportedServiceTags()
}

func ServicesByTag(tag string) []string {
	return services.GetSupportedServiceNamesByTag(tag)
}

func CheckAll(username string, parallelism int) []CheckResult {
	return checkInternal(SupportedServices(), username, parallelism)
}

func Check(services []string, username string, parallelism int) []CheckResult {
	return checkInternal(services, username, parallelism)
}

func CheckExcluding(services []string, username string, parallelism int) []CheckResult {
	excludedServices := make(map[string]bool)
	for _, service := range services {
		excludedServices[service] = true
	}
	allServices := SupportedServices()
	services = make([]string, 0)
	for _, service := range allServices {
		if !excludedServices[service] {
			services = append(services, service)
		}
	}
	return checkInternal(services, username, parallelism)
}

func CheckByTags(tags []string, username string, parallelism int) []CheckResult {
	services := make([]string, 0)
	for _, tag := range tags {
		services = append(services, ServicesByTag(tag)...)
	}
	return checkInternal(services, username, parallelism)
}

type checkTask struct {
	service  string
	username string
}

func checkInternal(services []string, username string, parallelism int) []CheckResult {
	client := &network.DefaultRESTClient{}
	tasks := make(chan checkTask, parallelism)
	results := make(chan CheckResult, len(services))

	for w := 0; w < parallelism; w++ {
		go worker(tasks, results, client)
	}

	for _, service := range services {
		tasks <- checkTask{service, username}
	}
	close(tasks)

	result := make([]CheckResult, 0, len(services))
	for i := 0; i < len(services); i++ {
		result = append(result, <-results)
	}

	return result
}

func worker(tasks chan checkTask, results chan CheckResult, client network.RESTClient) {
	for task := range tasks {
		found, err := services.Check(task.service, task.username, client)
		results <- CheckResult{task.service, err, found}
	}
}
