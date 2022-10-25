package sebastion

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	t.Run("given NewContext is called, a non-nil logger and the given context is returned", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "", "")
		sctx := NewContext(ctx)
		if sctx.Logger == nil {
			t.Fatalf("no logger was found")
		}
		if sctx.Context() != ctx {
			t.Fatalf("expected provided context to be returned")
		}
	})
	t.Run("given Context() is called and no ctx is provided, context.Background is returned", func(t *testing.T) {
		sctx := Context{}
		if sctx.Context() != context.Background() {
			t.Fatalf("expected Background context to be returned")
		}
	})
}
