package main

import "time"

func createPipeline(proj, branch string) []executer {
	return []executer{
		newExceptionStep(
			"go fmt",
			"gofmt",
			"Gofmt: SUCCESS",
			proj,
			[]string{"-l", "."},
		),
		newExceptionStep(
			"go linting",
			"golangci-lint",
			"Golint: SUCCESS",
			proj,
			[]string{"run", "."},
		),
		newExceptionStep(
			"go cyclo",
			"gocyclo",
			"Gocyclo: SUCCESS",
			proj,
			[]string{"-over", "10", "."},
		),
		newStep(
			"go test",
			"go",
			"Go Test: SUCCESS",
			proj,
			[]string{"test", "-v"},
		),
		newStep(
			"go build",
			"go",
			"Go Build: SUCCESS",
			proj,
			[]string{"build", ".", "errors"},
		),
		newTimeoutStep(
			"git push",
			"git",
			"Git Push: SUCCESS",
			proj,
			[]string{"push", "origin", branch},
			10*time.Second,
		),
	}
}
