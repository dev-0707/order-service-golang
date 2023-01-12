package problems

import "github.com/moogar0880/problems"

// HTTP/1.1 403 Forbidden
// Content-Type: application/problem+json
// Content-Language: en

// {
//     "type": "https://example.com/probs/out-of-credit",
//     "title": "You do not have enough credit.",
//     "detail": "Your current balance is 30, but that costs 50.",
//     "instance": "/account/12345/msgs/abc",
//     "balance": 30,
//     "accounts": ["/account/12345",
//                  "/account/67890"]
//    }

func NewProblem(status int, instance string, problemType string, details string) *problems.DefaultProblem {
	p := problems.NewStatusProblem(status)
	p.Instance = instance
	if problemType != "" {
		p.Type = problemType
	}
	if details != "" {
		p.Detail = details
	}
	if instance != "" {
		p.Instance = instance
	}
	return p
}
