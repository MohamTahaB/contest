package utils

type ProblemSubmission struct {
	Id          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	ProblemID   int64  `json:"problem_id"`
	TimeLimit   int64  `json:"time_limit"`
	MemoryLimit int64  `json:"memory_limit"`
	Submission  string `json:"submission"`
}
