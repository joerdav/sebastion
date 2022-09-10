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
	Default string
	Options []string
}

func (si MultiStringSelect) String() string {
	return fmt.Sprint(*si.Ptr)
}
func (si MultiStringSelect) DefaultString() string {
	return si.Default
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
	Ptr        *T
	Default    T
	validators []func(T) error
}

type InputOption[T any] func(*InputReference[T])

func WithValidaton[T any](vs ...func(T) error) InputOption[T] {
	return func(ir *InputReference[T]) {
		ir.validators = append(ir.validators, vs...)
	}
}
func WithDefault[T any](d T) InputOption[T] {
	return func(ir *InputReference[T]) {
		ir.Default = d
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
	for _, v := range si.validators {
		err := v(s)
		if err != nil {
			return err
		}
	}
	*si.Ptr = s
	return nil
}
func (si InputReference[T]) DefaultString() string {
	return fmt.Sprint(si.Default)
}
func (si InputReference[T]) String() string {
	return fmt.Sprint(*si.Ptr)
}
func NewStringInput(name, description string, value *string) Input {
	return Input{Name: name, Description: description, Value: StringInputValue(value)}
}
func NewIntInput(name, description string, value *int) Input {
	return Input{Name: name, Description: description, Value: IntInputValue(value)}
}
func NewBoolInput(name, description string, value *bool) Input {
	return Input{Name: name, Description: description, Value: BoolInputValue(value)}
}
func NewInput[T any](name, description string, value *T, opts ...InputOption[T]) Input {
	ir := InputReference[T]{Ptr: value}
	for _, o := range opts {
		o(&ir)
	}
	return Input{Name: name, Description: description, Value: ir}
}
func NewMultiStringInput(name, description string, value *string, options ...string) Input {
	return Input{name, description, MultiStringSelect{Ptr: value, Options: options}}
}

func StringInputValue(v *string, opts ...InputOption[string]) InputValue {
	ir := InputReference[string]{Ptr: v}
	for _, o := range opts {
		o(&ir)
	}
	return ir
}
func IntInputValue(v *int, opts ...InputOption[int]) InputValue {
	ir := InputReference[int]{Ptr: v}
	for _, o := range opts {
		o(&ir)
	}
	return ir
}
func BoolInputValue(v *bool, opts ...InputOption[bool]) InputValue {
	ir := InputReference[bool]{Ptr: v}
	for _, o := range opts {
		o(&ir)
	}
	return ir
}
