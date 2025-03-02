package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type exceptionStep struct {
	step
}

func newExceptionStep(name, exe, message, proj string, args []string) exceptionStep {
	s := exceptionStep{}

	s.step = newStep(name, exe, message, proj, args)

	return s
}

func (s exceptionStep) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {

		if s.name == "go cyclo" {
			return "", &stepErr{
				step:  s.name,
				msg:   fmt.Sprintf("files with complexity score > 10: %s", stdout.String()),
				cause: err,
			}
		}

		if stderr.Len() > 0 {
			return "", &stepErr{
				step:  s.name,
				msg:   fmt.Sprintf(stderr.String()),
				cause: err,
			}
		}

		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	if stdout.Len() > 0 {
		if s.name == "go fmt" {
			return "", &stepErr{
				step:  s.name,
				msg:   fmt.Sprintf("invalid format: %s", stdout.String()),
				cause: nil,
			}
		}
	}

	return s.message, nil
}
