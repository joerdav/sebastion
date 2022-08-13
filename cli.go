package sebastion

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
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

func (p TUIRunner) getAction() (Action, error) {
	var idx int
	for idx == 0 {
		for i, a := range p.Actions {
			n, d := a.Details()
			fmt.Fprintf(p.Out, "[%d] %s\n", i+1, n)
			if d != "" {
				fmt.Fprintf(p.Out, "\t%s\n", d)
			}
		}
		fmt.Fprint(p.Out, "Select an action: ")
		r := bufio.NewReader(p.In)
		s, _, err := r.ReadRune()
		if err != nil {
			return nil, err
		}
		idx, err = strconv.Atoi(string(s))
		if err != nil {
			idx = 0
		}
	}
	return p.Actions[idx-1], nil
}

func (p TUIRunner) getStringInput(i Input) string {
	for {
		fmt.Fprintf(p.Out, "%s - %s\n", i.Name, i.Description)
		fmt.Fprint(p.Out, "Enter string: ")
		r := bufio.NewReader(p.In)
		s, err := r.ReadString('\n')
		if err != nil {
			continue
		}
		return s
	}
}
func (p TUIRunner) getBoolInput(i Input) bool {
	for {
		fmt.Fprintf(p.Out, "%s - %s\n", i.Name, i.Description)
		fmt.Fprint(p.Out, "Enter bool [t/f]: ")
		r := bufio.NewReader(p.In)
		s, _, _ := r.ReadRune()
		if s != 't' && s != 'f' {
			continue
		}
		return s == 't'
	}
}
func (p TUIRunner) getIntInput(i Input) int {
	for {
		fmt.Fprintf(p.Out, "%s - %s\n", i.Name, i.Description)
		fmt.Fprint(p.Out, "Enter number: ")
		r := bufio.NewReader(p.In)
		s, err := r.ReadString('\n')
		if err != nil {
			continue
		}
		n, err := strconv.Atoi(string(s))
		if err != nil {
			continue
		}
		return n
	}
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
			if err := ip.Set(p.getStringInput(it)); err != nil {
				return err
			}
		case InputReference[bool]:
			if err := ip.Set(p.getBoolInput(it)); err != nil {
				return err
			}
		case InputReference[int]:
			if err := ip.Set(p.getIntInput(it)); err != nil {
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
	fmt.Fprintf(p.Out, "You are about to run the task \"%s\" with the following values:\n", n)
	for i := range a.Inputs() {
		fmt.Fprintf(p.Out, "%s: %v\n", a.Inputs()[i].Name, a.Inputs()[i].Value)
	}
	fmt.Fprintf(p.Out, "Run %s? [y/N]\n", n)
	r := bufio.NewReader(p.In)
	s, _, _ := r.ReadRune()
	if s != 'y' && s != 'Y' {
		return errors.New("exited")
	}
	return a.Run()
}
