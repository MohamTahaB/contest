package pythonworker

import (
	"contest/utils"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type WorkerServer struct {
	router *gin.Engine
}

// Returns a pointer to a WorkerServer instance, or an error if encountered.
func Init() (*WorkerServer, error) {
	if err := buildSandbox(); err != nil {
		return nil, fmt.Errorf("error initiating Python WorkerServer: %w", err)
	}

	return &WorkerServer{
		router: router(),
	}, nil
}

func (ws *WorkerServer) Run() {
	ws.router.Run(":8081")
}

func router() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	router.POST("/run", func(ctx *gin.Context) {
		var submission utils.ProblemSubmission
		if err := ctx.BindJSON(&submission); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("error parsing submission: %v", err),
			})
			return
		}

		output, err := execSubmission(&submission)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("error executing submission: %v", err),
			})
			return
		}

		if output == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error executing submission: cannot dereference nil pointer",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"output": *output,
		})
	})
	return router
}

func execSubmission(ps *utils.ProblemSubmission) (*utils.ExecutionOutput, error) {
	commandArgs := []string{
		"run",
		"-i",
		"--rm",
		"--cpus=0.5",
		fmt.Sprintf("--memory=%dm", ps.MemoryLimit),
		"python-sandbox",
		"sh",
		"-c",
		fmt.Sprintf("prlimit --cpu=%d -- python -", ps.TimeLimit),
	}

	cmd := exec.Command("docker", commandArgs...)

	output := utils.ExecutionOutput{}

	cmd.Stdout = &output.Stdout
	cmd.Stderr = &output.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	io.WriteString(stdin, fmt.Sprintf("%s", ps.Submission))
	stdin.Close()

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	output.ExitCode = cmd.ProcessState.ExitCode()

	return &output, nil
}

// Builds the sandbox image, returns an error if encountered.
func buildSandbox() error {
	cmd := exec.Command("docker", "build", "-t", "python-sandbox", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error building python sandbox docker image: %w", err)
	}
	return nil
}
