package sl

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

type logStruct struct {
	cfg *Config
	*slog.Logger
}

func (l *logStruct) With(args ...any) Logger {
	l2 := *l
	l2.Logger = l.Logger.With(args...)

	return &l2
}

func (l *logStruct) WithGroup(name string) Logger {
	l2 := *l
	l2.Logger = l.Logger.WithGroup(name)

	return &l2
}

func (l *logStruct) Config() *Config {
	return l.cfg
}

func (l *logStruct) ToSlog() *slog.Logger {
	return l.Logger
}

func (l *logStruct) Fatal(msg string, args ...any) {
	_, file, line, _ := runtime.Caller(1)

	cfg := l.cfg
	cfg.AddSource = false

	log := New(cfg).With(
		StringAttr("fatal_source", fmt.Sprintf("%s:%d", file, line)),
	)

	log.Error(msg, args...)

	os.Exit(1)
}

func (l *logStruct) FatalContext(ctx context.Context, msg string, args ...any) {
	_, file, line, _ := runtime.Caller(1)

	cfg := l.cfg
	cfg.AddSource = false

	log := New(cfg).With(
		StringAttr("fatal_source", fmt.Sprintf("%s:%d", file, line)),
	)

	log.ErrorContext(ctx, msg, args...)

	os.Exit(1)
}
