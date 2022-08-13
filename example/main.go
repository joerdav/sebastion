package main

import (
	"fmt"
	"log"
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
	p := sebastion.New(Panic{}, CatSomething{}, CancelPolicy{})
	log.Println(p.Run())
}

type CancelPolicy struct{}

func (CancelPolicy) Details() (string, string) { return "Cancel Policy", "" }
func (CancelPolicy) Inputs() []sebastion.Input {
	return []sebastion.Input{
		{
			Name:        "Policy ID",
			Description: "ID of the policy",
			Type:        sebastion.InputTypeString,
		},
		{
			Name:        "reason",
			Description: "why",
			Type:        sebastion.InputTypeString,
		},
	}
}
func (CancelPolicy) Run(iv sebastion.InputValues) error {
	fmt.Print(iv.GetString(0))
	fmt.Print(iv.GetString(1))
	return nil
}

type CatSomething struct{}

func (CatSomething) Details() (string, string) { return "Cat", "Cat Something" }
func (CatSomething) Inputs() []sebastion.Input {
	return []sebastion.Input{
		{
			Name:        "Text",
			Description: "Some text to be cat-ed.",
			Type:        sebastion.InputTypeString,
		},
	}
}
func (CatSomething) Run(iv sebastion.InputValues) error {
	fmt.Println(iv.GetString(0))
	return nil
}

type Panic struct{}

func (Panic) Details() (string, string) { return "Panic", "" }
func (Panic) Inputs() []sebastion.Input {
	return []sebastion.Input{}
}
func (Panic) Run(iv sebastion.InputValues) error {
	panic("panic")
}
