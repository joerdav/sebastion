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
	// ErrSelectionNotAnOption MultiXSelect when setting a MultiSelect it must be an option.
	ErrSelectionNotAnOption = errors.New("selection not an option")
)

type ActionDetails struct {
	Name, Description string
}

// Action defines an interface for a script that takes some input and runs some code.
type Action interface {
	// Details should return the name and an optional description of the Action.
	Details() ActionDetails
	// Inputs outlines the values required to run the Action.
	Inputs() []Input
	// Run should contain the code to run the action.
	Run(ctx Context) error
}

type InputValue interface {
	fmt.Stringer
	Set(any) error
	DefaultString() string
}

type Input struct {
	Name, Description string
	Value             InputValue
}

type MultiStringSelect struct {
	Ptr     *string
	Options []string
	Props   InputProps[string]
}

func (si MultiStringSelect) String() string {
	return fmt.Sprint(*si.Ptr)
}
func (si MultiStringSelect) DefaultString() string {
	return si.Props.Default
}
func (si MultiStringSelect) Set(v any) error {
	if si.Ptr == nil {
		return ErrNilInputReference
	}
	s, ok := v.(string)
	if !ok {
		return ErrTypeMismatch
	}
	found := false
	for _, v2 := range si.Options {
		if v2 == s {
			found = true
		}
	}
	if !found {
		return ErrSelectionNotAnOption
	}
	*si.Ptr = s
	return nil
}

type InputReference[T any] struct {
	Ptr   *T
	Props InputProps[T]
}

type InputProps[T any] struct {
	Default   T
	Validator func(T) error
}

func Validators[T any](vs ...func(T) error) func(T) error {
	return func(t T) error {
		for _, v := range vs {
			err := v(t)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// Required constrains an input to non-zero valuesj
func Required[T comparable](v T) error {
	var zero T
	if v == zero {
		return errors.New("required field")
	}
	return nil
}

func (si InputReference[T]) Set(v any) error {
	if si.Ptr == nil {
		return ErrNilInputReference
	}
	s, ok := v.(T)
	if !ok {
		return ErrTypeMismatch
	}
	if si.Props.Validator == nil {
		*si.Ptr = s
		return nil
	}
	err := si.Props.Validator(s)
	if err != nil {
		return err
	}
	*si.Ptr = s
	return nil
}
func (si InputReference[T]) DefaultString() string {
	return fmt.Sprint(si.Props.Default)
}
func (si InputReference[T]) String() string {
	return fmt.Sprint(*si.Ptr)
}

func emptyValidator[T any](T) error {
	return nil
}
func NewInput[T any](name, description string, value *T, props *InputProps[T]) Input {
	ir := InputReference[T]{Ptr: value}
	if props != nil {
		ir.Props = *props
	}
	if ir.Props.Validator == nil {
		ir.Props.Validator = emptyValidator[T]
	}
	return Input{Name: name, Description: description, Value: ir}
}
func NewMultiStringInput(name, description string, value *string, options []string, props *InputProps[string]) Input {
	ir := MultiStringSelect{Ptr: value, Options: options}
	if props != nil {
		ir.Props = *props
	}
	if ir.Props.Validator == nil {
		ir.Props.Validator = func(s string) error { return nil }
	}
	return Input{name, description, ir}
}
