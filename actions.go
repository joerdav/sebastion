package sebastion

import (
	"errors"
	"fmt"
)

var (
	ErrTypeMismatch = errors.New("a type mismatch error has occured, this should not be possible, please raise an issue")
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
	ptr *T
}

func (si InputReference[T]) Set(v any) error {
	if s, ok := v.(T); ok {
		*si.ptr = s
		return nil
	}
	return ErrTypeMismatch
}
func (si InputReference[T]) String() string {
	return fmt.Sprint(*si.ptr)
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
