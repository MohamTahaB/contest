package utils

import "bytes"

type ExecutionOutput struct {
	Stdout   bytes.Buffer `json:"stdout"`
	Stderr   bytes.Buffer `json:"stderr"`
	ExitCode int          `json:"exit_code"`
}
