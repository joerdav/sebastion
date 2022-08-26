package ui

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joerdav/sebastion"
)

// defaultTUIHandlers outlines the default input list of TUIInputHandler. It can be overriden with AppendHandlers:
//	t := sebastion.TUI(&exampleAction{})
//	t.AppendHandlers(customStringHandler{}, myStructHander{})
var defaultTUIHandlers = []TUIInputHandler{
	stringHandler{},
	boolHandler{},
	intHandler{},
	float32Handler{},
	float64Handler{},
	multiStringHandler{},
}

// TUIRunner is an Text User Interface runner for sebastion Actions.
type TUIRunner struct {
	Actions       []sebastion.Action
	In            io.Reader
	Out           io.Writer
	inputHandlers []TUIInputHandler
}

// TUI creates a new TUIRunner with the provided actions.
//	sebastion.TUI(action1{}, action2{}).Run()
func TUI(ac ...sebastion.Action) TUIRunner {
	return TUIRunner{
		Actions: ac,
		In:      os.Stdin,
		Out:     os.Stdout,
	}
}

// AppendHandlers allows you to add custom code to retrieve inputs for an action.
// There are some handlers provided by default that use https://github.com/AlecAivazis/survey to collect inputs.
func (t *TUIRunner) AppendHandlers(h ...TUIInputHandler) {
	t.inputHandlers = append(t.inputHandlers, h...)
}

// TUIInputHandler provides an interface for retrieving action inputs via the command line.
type TUIInputHandler interface {
	// CanHandle returns true if it can handle the given Input.
	CanHandle(sebastion.Input) bool
	// Get should retrieve the input and call sebastion.Input.Value.Set if it is successful.
	Get(sebastion.Input) error
}

func actionString(a sebastion.Action) string {
	n := a.Details()
	if n.Description != "" {
		return fmt.Sprintf("%v - %v", n.Name, n.Description)
	}
	return fmt.Sprintf("%v", n.Name)
}

func (p TUIRunner) getAction() (sebastion.Action, error) {
	var options []string
	actions := make(map[string]sebastion.Action)
	for _, a := range p.Actions {
		as := actionString(a)
		options = append(options, as)
		actions[as] = a
	}
	chosen := ""
	prompt := &survey.Select{
		Message: "Choose an action:",
		Options: options,
	}
	err := survey.AskOne(prompt, &chosen)
	return actions[chosen], err
}

func (p TUIRunner) getInputs(a sebastion.Action) error {
	is := a.Inputs()
	if len(is) == 0 {
		fmt.Fprintln(p.Out, "No inputs required.")
		return nil
	}
	fmt.Fprintln(p.Out)
	for _, it := range is {
		var handler TUIInputHandler
		for _, h := range defaultTUIHandlers {
			if h.CanHandle(it) {
				handler = h
			}
		}
		for _, h := range p.inputHandlers {
			if h.CanHandle(it) {
				handler = h
			}
		}
		if defaultTUIHandlers == nil {
			return fmt.Errorf("unsupported input type [%+v]", it)
		}
		err := handler.Get(it)
		if err != nil {
			return err
		}
		fmt.Fprintln(p.Out)
	}
	return nil
}

func (p TUIRunner) Run() error {
	a, err := p.getAction()
	if err != nil {
		return err
	}
	err = p.getInputs(a)
	if err != nil {
		return err
	}
	details := a.Details()
	fmt.Fprintf(p.Out, "You are about to run the task \"%s\" with the following values:\n\n", details.Name)
	for i := range a.Inputs() {
		fmt.Fprintf(p.Out, "%s: %v\n\n", a.Inputs()[i].Name, a.Inputs()[i].Value)
	}
	inp := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Run %s?", details.Name),
	}
	err = survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	if !inp {
		return errors.New("exited")
	}
	ctx := sebastion.NewContext(context.Background())
	ctx.Logger.SetOutput(p.Out)
	ctx.Logger.SetFlags(0)
	return a.Run(ctx)
}
