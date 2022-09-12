package sebastion

import "testing"

func TestInputReference(t *testing.T) {
	t.Run("given a wrong type, return an error", func(t *testing.T) {
		number := 1
		ir := InputReference[int]{Ptr: &number}
		err := ir.Set("mystring")
		if err != ErrTypeMismatch {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if *ir.Ptr != 1 {
			t.Errorf("pointer was updated, want=%v got=%v", 1, *ir.Ptr)
		}
	})
	t.Run("given a nil input, return an error", func(t *testing.T) {
		ir := InputReference[int]{}
		err := ir.Set("mystring")
		if err != ErrNilInputReference {
			t.Errorf("expected error. want=%v got=%v", ErrNilInputReference, err)
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
}

func TestStringInput(t *testing.T) {
	t.Run("given a wrong type, return an error", func(t *testing.T) {
		s := "somestring"
		ir := InputReference[string]{Ptr: &s}
		err := ir.Set(1)
		if err != ErrTypeMismatch {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if ir.String() != s {
			t.Errorf("pointer was updated, want=%v got=%v", 1, ir.String())
		}
	})
	t.Run("given a nil input, return an error", func(t *testing.T) {
		ir := InputReference[string]{}
		err := ir.Set("mystring")
		if err != ErrNilInputReference {
			t.Errorf("expected error. want=%v got=%v", ErrNilInputReference, err)
		}
	})
	t.Run("given a correct type, set the pointer", func(t *testing.T) {
		s := "somestring"
		ir := InputReference[string]{Ptr: &s}
		err := ir.Set("newstring")
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if ir.String() != "newstring" {
			t.Errorf("pointer was not updated, want=%v got=%v", "newstring", ir.String())
		}
	})
}

func TestBoolInput(t *testing.T) {
	t.Run("given a wrong type, return an error", func(t *testing.T) {
		s := false
		ir := InputReference[bool]{Ptr: &s}
		err := ir.Set(1)
		if err != ErrTypeMismatch {
			t.Errorf("expected error. want=%v got=%v", ErrTypeMismatch, err)
		}
		if ir.String() != "false" {
			t.Errorf("pointer was updated, want=%v got=%v", false, ir.String())
		}
	})
	t.Run("given a nil input, return an error", func(t *testing.T) {
		ir := InputReference[bool]{Ptr: nil}
		err := ir.Set("mystring")
		if err != ErrNilInputReference {
			t.Errorf("expected error. want=%v got=%v", ErrNilInputReference, err)
		}
	})
	t.Run("given a correct type, set the pointer", func(t *testing.T) {
		s := false
		ir := InputReference[bool]{Ptr: &s}
		err := ir.Set(true)
		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}
		if ir.String() != "true" {
			t.Errorf("pointer was not updated, want=%v got=%v", "true", ir.String())
		}
	})
}
