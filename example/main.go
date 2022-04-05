package main

import (
	"fmt"
	"os"

	"github.com/joerdav/sebastion"
)

func main() {
	p := sebastion.New(Panic{}, CatSomething{})
	if err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
