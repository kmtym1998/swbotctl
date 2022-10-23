package cfg

import (
	"io"
	"time"

	"github.com/blang/semver/v4"
	"github.com/briandowns/spinner"
)

type ExecutionContext struct {
	Spinner        *spinner.Spinner
	Stderr, Stdout io.Writer
	Version        semver.Version
}

func NewExecutionContext() *ExecutionContext {
	return &ExecutionContext{}
}

func (ec *ExecutionContext) Prepare(v string) error {
	if err := ec.setVersion(v); err != nil {
		return err
	}

	ec.setupSpinner()

	return nil
}

func (ec *ExecutionContext) setVersion(v string) error {
	parsedV, err := semver.Parse(v)
	if err != nil {
		return err
	}

	ec.Version = parsedV

	return nil
}

func (ec *ExecutionContext) setupSpinner() {
	if ec.Spinner == nil {
		spnr := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
		spnr.Writer = ec.Stderr
		ec.Spinner = spnr
	}
}

// Spin stops any existing spinner and starts a new one with the given message.
func (ec *ExecutionContext) Spin(message string) {
	ec.Spinner.Stop()
	ec.Spinner.Prefix = message
	ec.Spinner.Start()
}
