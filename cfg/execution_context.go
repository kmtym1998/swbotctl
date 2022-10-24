package cfg

import (
	"io"
	"time"

	"github.com/blang/semver/v4"
	"github.com/briandowns/spinner"
	"github.com/kmtym1998/swbotctl/switchbot"
)

type ExecutionContext struct {
	Spinner            *spinner.Spinner
	Stderr, Stdout     io.Writer
	SwitchBotAPIClient *switchbot.Switchbot
	Version            semver.Version
}

func NewExecutionContext() *ExecutionContext {
	return &ExecutionContext{}
}

func (ec *ExecutionContext) Prepare(v, token, secret string) error {
	if err := ec.setVersion(v); err != nil {
		return err
	}

	ec.setupSpinner()
	ec.setSwitchbotAPIClient(token, secret)

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

func (ec *ExecutionContext) Spin(message string) {
	ec.Spinner.Stop()
	ec.Spinner.Prefix = message
	ec.Spinner.Start()
}

func (ec *ExecutionContext) setSwitchbotAPIClient(token, secret string) {
	c := switchbot.NewClient(token, secret)
	ec.SwitchBotAPIClient = &c
}
