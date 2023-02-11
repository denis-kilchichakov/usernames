package usernames

type CheckResult struct {
	Service string
	err     error
	found   bool
}

func SupportedServices() []string {
	return []string{"github"}
}

func CheckExcluding(services []string, parallelism int) []CheckResult {
	return []CheckResult{}
}

func CheckIncluding(services []string, parallelism int) []CheckResult {
	return []CheckResult{}
}
