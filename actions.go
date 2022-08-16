package sebastion

import (
	"errors"
	"fmt"
)

var (
	ErrTypeMismatch      = errors.New("a type mismatch error has occured, this should not be possible, please raise an issue")
	ErrNilInputReference = errors.New("nil InputReference, an actions input is missing a reference")
)

type Action interface {
	Details() (name, description string)
	Inputs() []Input
	Run() error
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

func StringInput(v *string) InputValue {
	return InputReference[string]{v}
}
func IntInput(v *int) InputValue {
	return InputReference[int]{v}
}
func BoolInput(v *bool) InputValue {
	return InputReference[bool]{v}
}
