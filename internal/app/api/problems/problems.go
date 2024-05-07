package problems

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

type Problem struct {
	Type string `json:"type"`
	// Title    string `json:"title"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func NewProblem(status int, instance string, problemType string, details string) Problem {
	// p := problems.NewStatusProblem(status)

	// p.Instance = instance
	// if problemType != "" {
	// 	p.Type = problemType
	// }
	// if details != "" {
	// 	p.Detail = details
	// }
	// if instance != "" {
	// 	p.Instance = instance
	// }
	/*
		Type     string `json:"type"`
		Title    string `json:"title"`
		Status   int    `json:"status,omitempty"`
		Detail   string `json:"detail,omitempty"`
		Instance string `json:"instance,omitempty"`
	*/
	return Problem{problemType, status, details, instance}
}
