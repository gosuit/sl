package sl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithLogger(t *testing.T) {
	ctx := context.Background()
	logger := New(&Config{})
	ctxWithLogger := ContextWithLogger(ctx, logger)

	// Extract logger from new context
	extractedLogger, ok := ctxWithLogger.Value(ctxLogger{}).(Logger)
	if !ok || extractedLogger != logger {
		t.Errorf("Logger was not properly added to context")
	}
}

func TestLoggerFromContext_WithLogger(t *testing.T) {
	ctx := context.Background()
	logger := New(&Config{})
	ctxWithLogger := ContextWithLogger(ctx, logger)

	// Extract logger using loggerFromContext
	extractedLogger := loggerFromContext(ctxWithLogger)
	if extractedLogger != logger {
		t.Errorf("Did not retrieve correct logger from context")
	}
}

func TestLoggerFromContext_NoLogger(t *testing.T) {
	ctx := context.Background()

	assert.Equal(t, Default(), L(ctx))
}
