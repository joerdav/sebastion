package main

import (
	"fmt"
	"os"

	"github.com/joerdav/sebastion"
)

func openInputTTY() (*os.File, error) {
	f, err := os.Open("/dev/tty")
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	p := sebastion.New(Panic{}, &CatSomething{}, &CancelPolicy{})
	if err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type CancelPolicy struct {
	policyID           string
	cancellationReason string
}

func (cp *CancelPolicy) Details() (string, string) { return "Cancel Policy", "" }
func (cp *CancelPolicy) Inputs() []sebastion.Input {
	return []sebastion.Input{
		{
			Name:        "Policy ID",
			Description: "ID of the policy",
			Value:       sebastion.StringInput(&cp.policyID),
		},
		{
			Name:        "reason",
			Description: "why",
			Value:       sebastion.StringInput(&cp.cancellationReason),
		},
	}
}
func (cp *CancelPolicy) Run() error {
	fmt.Print(cp.policyID)
	fmt.Print(cp.cancellationReason)
	return nil
}

type CatSomething struct {
	text string
}

func (c *CatSomething) Details() (string, string) { return "Cat", "Cat Something" }
func (c *CatSomething) Inputs() []sebastion.Input {
	return []sebastion.Input{
		{
			Name:        "Text",
			Description: "Some text to be cat-ed.",
			Value:       sebastion.StringInput(&c.text),
		},
	}
}
func (c *CatSomething) Run() error {
	fmt.Println(c.text)
	return nil
}

type Panic struct{}

func (Panic) Details() (string, string) { return "Panic", "" }
func (Panic) Inputs() []sebastion.Input {
	return []sebastion.Input{}
}
func (Panic) Run() error {
	panic("panic")
}
