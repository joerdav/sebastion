package sebastion

import (
	"errors"
	"fmt"
)

var (
	// ErrTypeMismatch should never occur, and would be an internal error with sebastion.
	ErrTypeMismatch = errors.New("a type mismatch error has occured, this should not be possible, please raise an issue")
	// ErrNilInputReference InputReference[T] types must be provided a pointer of type T that can be used.
	ErrNilInputReference = errors.New("nil InputReference, an actions input is missing a reference")
)

// Action defines an interface for a script that takes some input and runs some code.
type Action interface {
	// Details should return the name and an optional description of the Action.
	Details() (name, description string)
	// Inputs outlines the values required to run the Action.
	Inputs() []Input
	// Run should contain the code to run the action.
	Run(ctx Context) error
}

type InputValue interface {
	fmt.Stringer
	Set(any) error
}

type Input struct {
	Name, Description string
	Value             InputValue
}

type InputReference[T any] struct {
	Ptr *T
}

func (si InputReference[T]) Set(v any) error {
	if si.Ptr == nil {
		return ErrNilInputReference
	}
	if s, ok := v.(T); ok {
		*si.Ptr = s
		return nil
	}
	return ErrTypeMismatch
}
func (si InputReference[T]) String() string {
	return fmt.Sprint(*si.Ptr)
}

func NewStringInput(name, description string, value *string) Input {
	return Input{name, description, StringInputValue(value)}
}
func NewIntInput(name, description string, value *int) Input {
	return Input{name, description, IntInputValue(value)}
}
func NewBoolInput(name, description string, value *bool) Input {
	return Input{name, description, BoolInputValue(value)}
}
func NewInput[T any](name, description string, value *T) Input {
	return Input{name, description, InputReference[T]{value}}
}

func StringInputValue(v *string) InputValue {
	return InputReference[string]{v}
}
func IntInputValue(v *int) InputValue {
	return InputReference[int]{v}
}
func BoolInputValue(v *bool) InputValue {
	return InputReference[bool]{v}
}
