package sebastion

import (
	"errors"
	"testing"
)

func TestMultiStringSelect(t *testing.T) {
	t.Run("given String is called, the selected string is returned", func(t *testing.T) {
		str := "mystring"
		ms := MultiStringSelect{Ptr: &str}
		if actual := ms.String(); actual != str {
			t.Fatalf("ms.String() = %s, expected: %s", actual, str)
		}
	})
	t.Run("given DefaultString is called, the default is returned", func(t *testing.T) {
		ms := MultiStringSelect{Props: InputProps[string]{Default: "hello"}}
		if actual := ms.DefaultString(); actual != "hello" {
			t.Fatalf("ms.String() = %s, expected: %s", actual, "hello")
		}
	})
	t.Run("given a wrong type, return an error", func(t *testing.T) {
		str := ""
		ms := MultiStringSelect{Ptr: &str}
		err := ms.Set(1)
		if !errors.Is(err, ErrTypeMismatch) {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if *ms.Ptr != str {
			t.Errorf("pointer was updated, want=%v got=%v", str, *ms.Ptr)
		}
	})
	t.Run("given an invalid input, return an error", func(t *testing.T) {
		str := ""
		ms := MultiStringSelect{Ptr: &str, Options: []string{""}, Props: InputProps[string]{
			Validator: func(s string) error { return errors.New("some error") },
		}}
		err := ms.Set("")
		if err == nil {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if *ms.Ptr != str {
			t.Errorf("pointer was updated, want=%v got=%v", str, *ms.Ptr)
		}
	})
	t.Run("given a nil input, return an error", func(t *testing.T) {
		ms := MultiStringSelect{}
		err := ms.Set("mystring")
		if !errors.Is(err, ErrNilInputReference) {
			t.Errorf("expected error. want=%v got=%v", ErrNilInputReference, err)
		}
	})
	t.Run("given a non-option, return an error", func(t *testing.T) {
		str := "old"
		ms := MultiStringSelect{Ptr: &str}
		err := ms.Set("new")
		if !errors.Is(err, ErrSelectionNotAnOption) {
			t.Errorf("expected error. want=%v got=%v", ErrSelectionNotAnOption, err)
		}
		if *ms.Ptr != str {
			t.Errorf("pointer was updated, want=%v got=%v", str, *ms.Ptr)
		}
	})
	t.Run("given a correct type, set the pointer", func(t *testing.T) {
		str := "old"
		ms := MultiStringSelect{Ptr: &str, Options: []string{"new"}}
		err := ms.Set("new")
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if *ms.Ptr != "new" {
			t.Errorf("pointer was not updated, want=%v got=%v", "new", *ms.Ptr)
		}
	})
	t.Run("given a correct type and passes validation, set the pointer", func(t *testing.T) {
		str := "old"
		ms := MultiStringSelect{Ptr: &str, Options: []string{"new"}, Props: InputProps[string]{Validator: func(s string) error { return nil }}}
		err := ms.Set("new")
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if *ms.Ptr != "new" {
			t.Errorf("pointer was not updated, want=%v got=%v", "new", *ms.Ptr)
		}
	})
}

func TestInputReference(t *testing.T) {
	t.Run("given a wrong type, return an error", func(t *testing.T) {
		number := 1
		ir := InputReference[int]{Ptr: &number}
		err := ir.Set("mystring")
		if !errors.Is(err, ErrTypeMismatch) {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if *ir.Ptr != 1 {
			t.Errorf("pointer was updated, want=%v got=%v", 1, *ir.Ptr)
		}
	})
	t.Run("given String is called, the selected string is returned", func(t *testing.T) {
		str := "mystring"
		ms := InputReference[string]{Ptr: &str}
		if actual := ms.String(); actual != str {
			t.Fatalf("ms.String() = %s, expected: %s", actual, str)
		}
	})
	t.Run("given DefaultString is called, the default is returned", func(t *testing.T) {
		ms := InputReference[string]{Props: InputProps[string]{Default: "hello"}}
		if actual := ms.DefaultString(); actual != "hello" {
			t.Fatalf("ms.String() = %s, expected: %s", actual, "hello")
		}
	})
	t.Run("given a nil input, return an error", func(t *testing.T) {
		ir := InputReference[int]{}
		err := ir.Set("mystring")
		if !errors.Is(err, ErrNilInputReference) {
			t.Errorf("expected error. want=%v got=%v", ErrNilInputReference, err)
		}
	})
	t.Run("given an invalid type, return an error", func(t *testing.T) {
		number := 1
		ir := InputReference[int]{Ptr: &number, Props: InputProps[int]{
			Validator: func(i int) error { return errors.New("some error") },
		}}
		err := ir.Set(2)
		if err == nil {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if *ir.Ptr != 1 {
			t.Errorf("pointer was updated, want=%v got=%v", 1, *ir.Ptr)
		}
	})
	t.Run("given a correct type, set the pointer", func(t *testing.T) {
		number := 1
		ir := InputReference[int]{Ptr: &number}
		err := ir.Set(2)
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if *ir.Ptr != 2 {
			t.Errorf("pointer was not updated, want=%v got=%v", 2, *ir.Ptr)
		}
	})
	t.Run("given a correct type with validation, set the pointer", func(t *testing.T) {
		number := 1
		ir := InputReference[int]{Ptr: &number, Props: InputProps[int]{
			Validator: func(i int) error { return nil },
		}}
		err := ir.Set(2)
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if *ir.Ptr != 2 {
			t.Errorf("pointer was not updated, want=%v got=%v", 2, *ir.Ptr)
		}
	})
}

func TestRequiredValidate(t *testing.T) {
	t.Run("given empty string, return error", func(t *testing.T) {
		err := Required("")
		if err == nil {
			t.Fatalf("expected err, got nil")
		}
	})
	t.Run("given non-empty string, return no error", func(t *testing.T) {
		err := Required("no error")
		if err != nil {
			t.Fatalf("expected no err, got %v", err)
		}
	})
	t.Run("given empty int, return error", func(t *testing.T) {
		err := Required(0)
		if err == nil {
			t.Fatalf("expected err, got nil")
		}
	})
	t.Run("given non-empty int, return no error", func(t *testing.T) {
		err := Required(1)
		if err != nil {
			t.Fatalf("expected no err, got %v", err)
		}
	})
}

func TestValidators(t *testing.T) {
	t.Run("given no validators, return no error", func(t *testing.T) {
		err := Validators[string]()("")
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})
	t.Run("given no validators fail, return no error", func(t *testing.T) {
		err := Validators(emptyValidator[string], emptyValidator[string])("")
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}
	})
	t.Run("given one validator fails, return an error", func(t *testing.T) {
		err := Validators(func(s string) error {
			return errors.New("some error")
		}, func(s string) error {
			return nil
		})("")
		if err == nil {
			t.Fatalf("expected err, got %v", err)
		}
	})
	t.Run("given all validators fails, return an error", func(t *testing.T) {
		err := Validators(func(s string) error {
			return errors.New("some error")
		}, func(s string) error {
			return errors.New("some error")
		})("")
		if err == nil {
			t.Fatalf("expected err, got %v", err)
		}
	})
}
