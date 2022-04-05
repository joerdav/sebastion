package sebastion

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Program struct {
	Actions []Action
	In      io.Reader
	Out     io.Writer
}

func New(ac ...Action) Program {
	return Program{
		Actions: ac,
		In:      os.Stdin,
		Out:     os.Stdout,
	}
}

func (p Program) getAction() (Action, error) {
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

func (p Program) getStringInput(i Input) string {
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
func (p Program) getBoolInput(i Input) bool {
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
func (p Program) getIntInput(i Input) int {
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

func (p Program) getInputs(a Action) (InputValues, error) {
	ivs := InputValues{}
	is := a.Inputs()
	if len(is) == 0 {
		fmt.Fprintln(p.Out, "No inputs required.")
		return ivs, nil
	}
	fmt.Fprintln(p.Out)
	for i, it := range is {
		switch it.Type {
		case InputTypeString:
			ivs[i] = p.getStringInput(it)
		case InputTypeBool:
			ivs[i] = p.getBoolInput(it)
		case InputTypeInt:
			ivs[i] = p.getIntInput(it)
		}
		fmt.Fprintln(p.Out)
	}
	return ivs, nil
}

func (p Program) Run() error {
	a, err := p.getAction()
	if err != nil {
		return err
	}
	is, err := p.getInputs(a)
	if err != nil {
		return err
	}
	n, _ := a.Details()
	fmt.Fprintf(p.Out, "You are about to run the task \"%s\" with the following values:\n", n)
	for i := range a.Inputs() {
		fmt.Fprintf(p.Out, "%s: %v\n", a.Inputs()[i].Name, is[i])
	}
	fmt.Fprintf(p.Out, "Run %s? [y/N]\n", n)
	r := bufio.NewReader(p.In)
	s, _, _ := r.ReadRune()
	if s != 'y' && s != 'Y' {
		return errors.New("exited")
	}
	return a.Run(is)
}
