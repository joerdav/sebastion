package main

import (
	"fmt"
	"os"

	"github.com/joerdav/sebastion/examples"
	sebastionui "github.com/joerdav/sebastion/ui"
)

func main() {
	p := sebastionui.TUI(&examples.Panic{}, &examples.CatSomething{}, &examples.Spam{})
	if err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
