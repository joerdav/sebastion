package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joerdav/sebastion/examples"
	sebastionui "github.com/joerdav/sebastion/ui"
)

func main() {
	w, err := sebastionui.Web(&examples.Panic{}, &examples.EchoSomething{}, &examples.Spam{})
	if err != nil {
		fmt.Printf("An error occored: %v", err)
		os.Exit(1)
	}
	http.ListenAndServe("127.0.0.1:2020", w)
}
