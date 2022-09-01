package webui

import (
	"github.com/a-h/templ"
	"github.com/joerdav/sebastion"
)

// WebInputHandler provides an interface to allow collection of inputs from the WebRunner.
type WebInputHandler interface {
	// CanHandle is used to decide what handler should be used for which input type.
	// If no handlers can process an input an error will occur, custom handlers take presidence over built in sebastion handlers.
	CanHandle(sebastion.Input) bool
	// Template should return a templ.Component that will render the form input.
	Template(sebastion.Input) templ.Component
	// Set takes in the form post value and sets the sebastion value.
	Set(sebastion.Input, string) error
}
