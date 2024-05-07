package problems

import (
	"net/url"

	"github.com/moogar0880/problems"
)

type UserProblem struct {
	problems.DefaultProblem
	Id int
}

func (cp *UserProblem) ProblemType() (*url.URL, error) {
	u, err := url.Parse(cp.Type)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (cp *UserProblem) ProblemTitle() string {
	return cp.Title
}
