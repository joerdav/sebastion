package main

import (
	"fmt"
	"net/http"

	"github.com/joerdav/sebastion"
	sebastionui "github.com/joerdav/sebastion/ui"
)

func main() {
	w := sebastionui.Web(Panic{}, &CatSomething{}, &Spam{})
	http.ListenAndServe(":2020", w)
}

type Spam struct {
	message string
	repeat  int
	structa struct{}
}

func (cp *Spam) Details() (string, string) { return "Spam a message", "" }
func (cp *Spam) Inputs() []sebastion.Input {
	return []sebastion.Input{
		{
			Name:        "Message",
			Description: "The message to print",
			Value:       sebastion.StringInput(&cp.message),
		},
		{
			Name:        "Repetitions",
			Description: "How many times",
			Value:       sebastion.IntInput(&cp.repeat),
		},
	}
}
func (cp *Spam) Run() error {
	for i := 0; i < cp.repeat; i++ {
		fmt.Println(cp.message)
	}
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
