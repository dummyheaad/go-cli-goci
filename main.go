package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

type executer interface {
	execute() (string, error)
}

func main() {
	proj := flag.String("p", "", "Project directory")
	branch := flag.String("b", "master", "Target branch")
	flag.Parse()

	if err := run(*proj, *branch, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(proj, branch string, out io.Writer) error {
	if proj == "" {
		return fmt.Errorf("project directory is required: %w", ErrValidation)
	}

	pipeline := createPipeline(proj, branch)

	sig := make(chan os.Signal, 1)
	errCh := make(chan error)
	done := make(chan struct{})

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for _, s := range pipeline {
			msg, err := s.execute()
			if err != nil {
				errCh <- err
				return
			}

			_, err = fmt.Fprintln(out, msg)
			if err != nil {
				errCh <- err
				return
			}
		}
		close(done)
	}()

	for {
		select {
		case rec := <-sig:
			signal.Stop(sig)
			return fmt.Errorf("%s: Exiting: %w", rec, ErrSignal)
		case err := <-errCh:
			return err
		case <-done:
			return nil
		}
	}
}
