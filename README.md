# SL

SL Logger is a lightweight logging library built on top of Go's log/slog package. It provides a flexible and configurable logging interface that allows developers to easily integrate logging into their applications. With support for different log levels, output formats, and handlers, SL Logger is designed to fit a variety of use cases.

## Installation

```zsh
go get github.com/gosuit/sl
```

## Features

• Configurable Logging Levels: Set the logging level to control the verbosity of the logs (e.g., debug, info, warn, error).

• Multiple Output Options: Log to standard output, files, or discard logs entirely.

• Structured Logging: Supports structured logging with attributes for better log analysis.

• Contextual Logging: The logger can be placed and removed from the context to transfer the logger to different parts of the application.

• Custom Handlers: Easily switch between different logging handlers (e.g., dev, pretty, discard, default).

## Usage

### Configuration

You can configure the logger using the Config struct. Here’s an example configuration:

```golang
cfg := &sl.Config{
    Level:     "info",
    AddSource: true,
    Writer:    "stderr",
    Type:      "default",
}
```

### Creating a Logger

```golang
logger := sl.New(cfg)
```

### Logging Messages

```golang
logger.Debug("Debug message", "key", "value")
logger.Info("Info message", "key", "value")
logger.Warn("Warning message", "key", "value")
logger.Error("Error message", "key", "value")
logger.Fatal("Fatal message", "key", "value")
```

### Context

```golang
ctx := context.Background()
ctx = sl.ContextWithLogger(logger)

l := sl.L(ctx)

l.Info("Info message", "key", "value")
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.