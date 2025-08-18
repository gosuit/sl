package sl

import (
	"context"
	"log/slog"
	"os"

	"github.com/gosuit/boot"
	"github.com/gosuit/sl/handlers"
)

// Logger is an interface that defines methods for logging at various levels.
// It provides methods for creating log entries with different contexts and attributes.
type Logger interface {
	// Handler returns the handler used for logging.
	Handler() Handler

	// With creates a new logger instance with additional context arguments.
	With(args ...any) Logger

	// WithGroup creates a new logger instance that groups log entries under the specified name.
	WithGroup(name string) Logger

	// Enabled checks if logging is enabled for a specific level in the given context
	Enabled(ctx context.Context, level Level) bool

	// Log writes a log entry with the specified level, message, and context.
	Log(ctx context.Context, level Level, msg string, args ...any)

	// LogAttrs writes a log entry with attributes, allowing for more structured logging.
	LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)

	// Debug logs a message at the debug level.
	Debug(msg string, args ...any)

	// DebugContext logs a message at the debug level with the specified context.
	DebugContext(ctx context.Context, msg string, args ...any)

	// Info logs a message at the info level.
	Info(msg string, args ...any)

	// InfoContext logs a message at the info level with the specified context.
	InfoContext(ctx context.Context, msg string, args ...any)

	// Warn logs a message at the warning level.
	Warn(msg string, args ...any)

	// WarnContext logs a message at the warning level with the specified context.
	WarnContext(ctx context.Context, msg string, args ...any)

	// Error logs a message at the error level.
	Error(msg string, args ...any)

	// ErrorContext logs a message at the error level with the specified context.
	ErrorContext(ctx context.Context, msg string, args ...any)

	// Fatal logs a message at the error level and terminates the application.
	Fatal(msg string, args ...any)

	// FatalContext logs a message at the fatal level with the specified context and terminates the application.
	FatalContext(ctx context.Context, msg string, args ...any)

	// ToSlog returns the underlying slog.Logger instance.
	ToSlog() *slog.Logger

	// Config returns config of current logger.
	Config() *Config
}

// Config holds the configuration settings for the logger.
// It includes options for logging level, output format, and other behavior.
type Config struct {
	// Level specifies the logging level (e.g., info, debug, error).
	// It can be set via YAML configuration or environment variable.
	Level string `confy:"level" yaml:"level" json:"level" toml:"level" env:"LOGGER_LEVEL" env-default:"info" validate:"oneof=info debug error warn"`

	// AddSource indicates whether to include the source of the log message (e.g., file and line number).
	// This can be controlled through YAML or an environment variable.
	AddSource bool `confy:"add_source" yaml:"add_source" json:"add_source" toml:"add_source" env:"LOGGER_ADD_SOURCE" env-default:"true"`

	// IsJSON determines if the log output should be in JSON format.
	// If true, logs will be structured as JSON; otherwise, they will be plain text.
	IsJSON bool `confy:"is_json" yaml:"is_json" json:"is_json" toml:"is_json" env:"LOGGER_IS_JSON" env-default:"true"`

	// Writer specifies where the logs should be written.
	// It can be a "file" or a "stderr".
	Writer string `confy:"writer" yaml:"writer" json:"writer" toml:"writer" env:"LOGGER_WRITER" env-default:"stderr" validate:"oneof=stderr file"`

	// OutPath is the path to the output file if logging to a file.
	OutPath string `confy:"out_path" yaml:"out_path" json:"out_path" toml:"out_path" env:"LOGGER_OUT_PATH" env-default:""`

	// SetDefault indicates whether to set default logger options.
	// This can be used to ensure that certain configurations are applied automatically.
	SetDefault bool `confy:"set_default" yaml:"set_default" json:"set_default" toml:"set_default" env:"LOGGER_SET_DEFAULT" env-default:"true"`

	// Type defines the type of logger to use.
	// It can be one of the following:
	// - "dev" or "pretty": for human-readable logs with color and formatting.
	// - "discard": to ignore all log messages.
	// - "default": for standard logging behavior.
	Type string `confy:"type" yaml:"type" json:"type" toml:"type" env:"LOGGER_TYPE" env-default:"default" validate:"oneof=dev pretty discard default"`

	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// The attribute's value has been resolved (see [Value.Resolve]).
	// If ReplaceAttr returns a zero Attr, the attribute is discarded.
	ReplaceAttr func(groups []string, a Attr) Attr
}

// New returns the Logger with the specified configuration.
func New(cfg *Config) Logger {
	handler := setupHandler(cfg)

	logger := &logStruct{
		Logger: slog.New(handler),
		cfg:    cfg,
	}

	if cfg.SetDefault {
		SetDefault(logger)
	}

	return logger
}

func Boot[T any]() any {
	return boot.Boot[T, Config](New)
}

func setupHandler(cfg *Config) Handler {
	level := setLoggerLevel(cfg.Level)

	opts := setHandlerOptions(level, cfg)

	out := setOut(cfg)

	var handler Handler

	switch cfg.Type {

	case DevLogger:
		handler = handlers.NewDevSlog(out, opts)

	case PrettyLogger:
		handler = handlers.NewPretty(out, opts)

	case DiscardLogger:
		handler = handlers.NewDiscard()

	default:
		if cfg.IsJSON {
			handler = NewJSONHandler(out, opts)
		} else {
			handler = NewTextHandler(out, opts)
		}

	}

	return handler
}

func setLoggerLevel(lvl string) Level {
	var level Level

	switch lvl {

	case "debug":
		level = -4
	case "info":
		level = 0
	case "warn":
		level = 4
	case "error":
		level = 8
	default:
		level = 0

	}

	return level
}

func setHandlerOptions(level Level, cfg *Config) *HandlerOptions {
	return &HandlerOptions{
		AddSource:   cfg.AddSource,
		Level:       level,
		ReplaceAttr: cfg.ReplaceAttr,
	}
}

func setOut(cfg *Config) *os.File {
	if cfg.Writer == FileOut {
		return getLogFile(cfg.OutPath)
	}

	return os.Stderr
}

func getLogFile(path string) *os.File {
	if path == "" {
		path = "logs"
	}

	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(path+"/all.log", os.O_CREATE|os.O_RDWR, 0o0644)
	if err != nil {
		panic(err)
	}

	return logFile
}
