package utils

type ProblemStatus int

const (
	NotAttempted ProblemStatus = iota
	WrongAnswer
	Accepted
	TimeLimitExceed
	MemoryLimitExceed
)

type ProblemsStatus map[int64]bool
