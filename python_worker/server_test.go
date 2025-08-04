package pythonworker

import (
	"contest/utils"
	"testing"
)

func TestInit(t *testing.T) {
	s, err := Init()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if s == nil {
		t.Fatal("error initiating a server: expected a non nil Server instance, got nil")
	}
}

func TestExecSubmission(t *testing.T) {
	out, err := execSubmission(&utils.ProblemSubmission{
		Id:          0,
		UserID:      0,
		ProblemID:   0,
		TimeLimit:   1,
		MemoryLimit: 128,
		Submission:  `print("hello world")`,
	})
	if err != nil {
		t.Fatalf("%v", err)
	}

	if out.ExitCode != 0 {
		t.Fatalf("error executing submission: expected exit code 0, found %d", out.ExitCode)
	}

	if out.Stdout.String() != "hello world\n" {
		t.Fatalf("error executing submission: found %s", out.Stdout.String())
	}
}
