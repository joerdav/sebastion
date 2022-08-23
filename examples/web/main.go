package main

import (
	"net/http"

	"github.com/joerdav/sebastion/examples"
	sebastionui "github.com/joerdav/sebastion/ui"
)

func main() {
	w := sebastionui.Web(&examples.Panic{}, &examples.CatSomething{}, &examples.Spam{})
	http.ListenAndServe(":2020", w)
}
