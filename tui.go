package sebastion

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
)

type TUIRunner struct {
	Actions []Action
	In      io.Reader
	Out     io.Writer
}

func TUI(ac ...Action) TUIRunner {
	return TUIRunner{
		Actions: ac,
		In:      os.Stdin,
		Out:     os.Stdout,
	}
}

func actionString(a Action) string {
	n, d := a.Details()
	if d != "" {
		return fmt.Sprintf("%v - %v", n, d)
	}
	return fmt.Sprintf("%v", n)
}

func (p TUIRunner) getAction() (Action, error) {
	var options []string
	actions := make(map[string]Action)
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

func (p TUIRunner) getStringInput(i Input) (string, error) {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	return inp, err
}
func (p TUIRunner) getBoolInput(i Input) (bool, error) {
	inp := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	return inp, err
}
func (p TUIRunner) getIntInput(i Input) (int, error) {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(func(ans interface{}) error {
		str, ok := ans.(string)
		if !ok {
			return errors.New("This response must be a number.")
		}
		_, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("This response must be a number.")
		}
		return nil
	}))
	s, err := strconv.Atoi(inp)
	if err != nil {
		return 0, errors.New("This response must be a number.")
	}
	return s, err
}

func (p TUIRunner) getInputs(a Action) error {
	is := a.Inputs()
	if len(is) == 0 {
		fmt.Fprintln(p.Out, "No inputs required.")
		return nil
	}
	fmt.Fprintln(p.Out)
	for _, it := range is {
		switch ip := it.Value.(type) {
		case InputReference[string]:
			s, err := p.getStringInput(it)
			if err != nil {
				return err
			}
			if err := ip.Set(s); err != nil {
				return err
			}
		case InputReference[bool]:
			s, err := p.getBoolInput(it)
			if err != nil {
				return err
			}
			if err := ip.Set(s); err != nil {
				return err
			}
		case InputReference[int]:
			s, err := p.getIntInput(it)
			if err != nil {
				return err
			}
			if err := ip.Set(s); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported input type [%+v]", ip)
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
	n, _ := a.Details()
	fmt.Fprintf(p.Out, "You are about to run the task \"%s\" with the following values:\n\n", n)
	for i := range a.Inputs() {
		fmt.Fprintf(p.Out, "%s: %v\n\n", a.Inputs()[i].Name, a.Inputs()[i].Value)
	}
	inp := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Run %s?", n),
	}
	err = survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	if !inp {
		return errors.New("exited")
	}
	return a.Run()
}
