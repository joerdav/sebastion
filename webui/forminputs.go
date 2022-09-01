package webui

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/joerdav/sebastion"
	"github.com/joerdav/sebastion/webui/templates"
)

var defaultHandlers = []WebInputHandler{
	stringInput{},
	intInput{},
	boolInput{},
	multiStringInput{},
}

type stringInput struct{}

func (stringInput) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[string])
	return ok
}
func (stringInput) Set(i sebastion.Input, f string) error {
	return i.Value.Set(f)
}
func (stringInput) Template(i sebastion.Input) templ.Component {
	return templates.StringInput(i)
}

type intInput struct{}

func (intInput) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[int])
	return ok
}
func (intInput) Set(i sebastion.Input, f string) error {
	n, err := strconv.Atoi(f)
	if err != nil {
		return err
	}
	return i.Value.Set(n)
}
func (intInput) Template(i sebastion.Input) templ.Component {
	return templates.IntInput(i)
}

type boolInput struct{}

func (boolInput) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[bool])
	return ok
}
func (boolInput) Set(i sebastion.Input, f string) error {
	return i.Value.Set(f == "true")
}
func (boolInput) Template(i sebastion.Input) templ.Component {
	return templates.BoolInput(i)
}

type multiStringInput struct{}

func (multiStringInput) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.MultiStringSelect)
	return ok
}
func (multiStringInput) Set(i sebastion.Input, f string) error {
	return i.Value.Set(f)
}
func (multiStringInput) Template(i sebastion.Input) templ.Component {
	return templates.MultiStringInput(i)
}
