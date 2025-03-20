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
	log *slog.Logger
}

func (l *logStruct) Config() *Config {
	return l.cfg
}

func (l *logStruct) Handler() Handler {
	return l.log.Handler()
}

func (l *logStruct) With(args ...any) Logger {
	l2 := *l
	l2.log = l.log.With(args...)

	return &l2
}

func (l *logStruct) WithGroup(name string) Logger {
	l2 := *l
	l2.log = l.log.WithGroup(name)

	return &l2
}

func (l *logStruct) ToSlog() *slog.Logger {
	return l.log
}

func (l *logStruct) Enabled(ctx context.Context, level Level) bool {
	return l.log.Enabled(ctx, level)
}

func (l *logStruct) Log(ctx context.Context, level Level, msg string, args ...any) {
	l = l.addSourceToLogByCfg()

	l.log.Log(ctx, level, msg, args...)
}

func (l *logStruct) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	l.log.LogAttrs(ctx, level, msg, attrs...)
}

func (l *logStruct) Debug(msg string, args ...any) {
	l.Log(context.Background(), LevelDebug, msg, args...)
}

func (l *logStruct) DebugContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelDebug, msg, args)
}

func (l *logStruct) Info(msg string, args ...any) {
	l.Log(context.Background(), LevelInfo, msg, args...)
}

func (l *logStruct) InfoContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelInfo, msg, args...)
}

func (l *logStruct) Warn(msg string, args ...any) {
	l.Log(context.Background(), LevelWarn, msg, args...)
}

func (l *logStruct) WarnContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelWarn, msg, args...)
}

func (l *logStruct) Error(msg string, args ...any) {
	l.Log(context.Background(), LevelError, msg, args...)
}

func (l *logStruct) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelError, msg, args...)
}

func (l *logStruct) Fatal(msg string, args ...any) {
	l.Log(context.Background(), LevelError, msg, args...)
	os.Exit(1)
}

func (l *logStruct) FatalContext(ctx context.Context, msg string, args ...any) {
	l.Log(ctx, LevelError, msg, args...)
	os.Exit(1)
}

func (l *logStruct) addSourceToLogByCfg() *logStruct {
	_, file, line, _ := runtime.Caller(3)

	if l.cfg.AddSource {
		a := slog.String(slog.SourceKey, fmt.Sprintf("%s:%d", file, line))
		l = l.With(a).(*logStruct)
	}

	return l
}
