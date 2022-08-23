package main

import (
	"fmt"
	"os"

	"github.com/joerdav/sebastion/examples"
	sebastionui "github.com/joerdav/sebastion/ui"
)

func openInputTTY() (*os.File, error) {
	f, err := os.Open("/dev/tty")
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	p := sebastionui.TUI(&examples.Panic{}, &examples.CatSomething{}, &examples.Spam{})
	if err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
