package sl

import (
	"context"
	"log/slog"
	"os"
)

type logStruct struct {
	cfg *Config
	log *slog.Logger
}

func (l *logStruct) Config() *Config {
	return l.cfg
}

func (l *logStruct) Handler() Handler {
	return l.log.Handler()
}

func (l *logStruct) With(args ...any) Logger {
	l.log = l.log.With(args...)

	return l
}

func (l *logStruct) WithGroup(name string) Logger {
	l.log = l.log.WithGroup(name)

	return l
}

func (l *logStruct) ToSlog() *slog.Logger {
	return l.log
}

func (l *logStruct) Enabled(ctx context.Context, level Level) bool {
	return l.log.Enabled(ctx, level)
}

func (l *logStruct) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.log.Log(ctx, level, msg, args...)
}

func (l *logStruct) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	l.log.LogAttrs(ctx, level, msg, attrs...)
}

func (l *logStruct) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *logStruct) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log.DebugContext(ctx, msg, args...)
}

func (l *logStruct) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *logStruct) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log.InfoContext(ctx, msg, args...)
}

func (l *logStruct) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *logStruct) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log.WarnContext(ctx, msg, args...)
}

func (l *logStruct) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l *logStruct) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
}

func (l *logStruct) Fatal(msg string, args ...any) {
	l.log.Error(msg, args...)
	os.Exit(1)
}

func (l *logStruct) FatalContext(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
