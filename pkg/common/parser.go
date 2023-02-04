package common

type Parser interface {
	Parse(body []byte) (bool, error)
}
