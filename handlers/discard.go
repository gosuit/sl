package handlers

import (
	"context"
	"log/slog"
)

type Discard struct{}

// NewDiscard returns discard handler that ignore all call.
// It can be usefull for tests or other tasks.
func NewDiscard() *Discard {
	return &Discard{}
}

func (h *Discard) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (h *Discard) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

func (h *Discard) WithGroup(_ string) slog.Handler {
	return h
}

func (h *Discard) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
