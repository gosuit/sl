package sl

import (
	"context"
	"log/slog"
	"time"
)

const (
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelDebug = slog.LevelDebug

	PrettyLogger  = "pretty"
	DevLogger     = "dev"
	DiscardLogger = "discard"
	DefaultLogger = "default"

	StdOut  = "stdout"
	FileOut = "file"
)

type (
	Attr           = slog.Attr
	Level          = slog.Level
	Handler        = slog.Handler
	Value          = slog.Value
	HandlerOptions = slog.HandlerOptions
	LogValuer      = slog.LogValuer
)

var (
	NewTextHandler = slog.NewTextHandler
	NewJSONHandler = slog.NewJSONHandler

	StringAttr   = slog.String
	BoolAttr     = slog.Bool
	Float64Attr  = slog.Float64
	AnyAttr      = slog.Any
	DurationAttr = slog.Duration
	IntAttr      = slog.Int
	Int64Attr    = slog.Int64
	Uint64Attr   = slog.Uint64

	GroupValue = slog.GroupValue
	Group      = slog.Group
)

// L returns Logger from context
// if context doesn't have logger, returns Default.
func L(ctx context.Context) Logger {
	return loggerFromContext(ctx)
}

func Float32Attr(key string, val float32) Attr {
	return Float64Attr(key, float64(val))
}

func UInt32Attr(key string, val uint32) Attr {
	return IntAttr(key, int(val))
}

func Int32Attr(key string, val int32) Attr {
	return IntAttr(key, int(val))
}

func TimeAttr(key string, time time.Time) Attr {
	return StringAttr(key, time.String())
}

func ErrAttr(err error) Attr {
	return StringAttr("error", err.Error())
}

func Default() Logger {
	cfg := &Config{
		Level:      "info",
		AddSource:  false,
		IsJSON:     false,
		Writer:     StdOut,
		SetDefault: true,
		Type:       DefaultLogger,
	}

	return &logStruct{
		log: slog.Default(),
		cfg: cfg,
	}
}

func SetDefault(log Logger) {
	slog.SetDefault(log.ToSlog())
}
