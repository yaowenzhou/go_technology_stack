package context_test

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	cancel()
}
