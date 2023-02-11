package services

type serviceChecker interface {
	name() string
	check(username string) (bool, error)
}
