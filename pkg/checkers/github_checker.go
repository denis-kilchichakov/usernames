package checkers

type GithubParser struct{}

func (p *GithubParser) Parse(body []byte) (bool, error) {
	return true, nil
}
