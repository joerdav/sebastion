package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joerdav/sebastion/examples"
	sebastionui "github.com/joerdav/sebastion/webui"
)

func main() {
	w, err := sebastionui.Web(sebastionui.WebConfig{Workers: 3}, &examples.Primes{}, &examples.Spam{})
	if err != nil {
		fmt.Printf("An error occored: %v", err)
		os.Exit(1)
	}
	log.Println("Serving on localhost:2020")
	http.ListenAndServe("localhost:2020", w)
}
