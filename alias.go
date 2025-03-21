package sl

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

const (
	// Logging levels
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelDebug = slog.LevelDebug

	// Logger types
	PrettyLogger  = "pretty"
	DevLogger     = "dev"
	DiscardLogger = "discard"
	DefaultLogger = "default"

	// Output streams
	StdErr  = "stderr"
	FileOut = "file"
)

type (
	// Logging attributes
	Attr           = slog.Attr
	Level          = slog.Level
	Handler        = slog.Handler
	Value          = slog.Value
	HandlerOptions = slog.HandlerOptions
	LogValuer      = slog.LogValuer
)

var (
	// Create a text logging handler
	NewTextHandler = slog.NewTextHandler

	// Create a JSON logging handler
	NewJSONHandler = slog.NewJSONHandler

	// Predefined attributes for logging
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

// Float32Attr creates an attribute with a float32 value.
func Float32Attr(key string, val float32) Attr {
	return Float64Attr(key, float64(val))
}

// UInt32Attr creates an attribute for an unsigned integer (uint32).
func UInt32Attr(key string, val uint32) Attr {
	return IntAttr(key, int(val))
}

// Int32Attr creates an attribute for a signed integer (int32).
func Int32Attr(key string, val int32) Attr {
	return IntAttr(key, int(val))
}

// TimeAttr creates an attribute for time.
func TimeAttr(key string, time time.Time) Attr {
	return StringAttr(key, time.String())
}

// ErrAttr creates an attribute for an error.
func ErrAttr(err error) Attr {
	return StringAttr("error", err.Error())
}

// Default returns logger that setted as default.
func Default() Logger {
	mu.Lock()
	cfg := *defaultCfg
	mu.Unlock()

	return &logStruct{
		Logger: slog.Default(),
		cfg:    &cfg,
	}
}

// SetDefault sets the given logger as the default logger.
func SetDefault(log Logger) {
	slog.SetDefault(log.ToSlog())

	mu.Lock()
	cfg := *log.Config()
	defaultCfg = &cfg
	mu.Unlock()
}

var (
	mu         sync.Mutex
	defaultCfg = &Config{
		Level:      "info",
		AddSource:  false,
		IsJSON:     false,
		Writer:     StdErr,
		SetDefault: true,
		Type:       DefaultLogger,
	}
)
