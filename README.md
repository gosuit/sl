# Sl

Sl is a GoLang library based on log/slog. 

It provides a flexible and configurable logging interface that allows developers to easily integrate logging into their applications.

With support for different log levels, output formats, and handlers, Sl Logger is designed to fit a variety of use cases.

## Installation

```zsh
go get github.com/gosuit/sl
```

## Features

- **Configurable Logger**: flexible configuration of the logger, including: level, output, handler type, etc.
- **Structured Logging**: Supports structured logging with attributes for better log analysis.
- **Contextual Logging**: The logger can be placed and removed from the context to transfer the logger to different parts of the application.

## Usage

### Logging

```golang
package main

import "github.com/gosuit/sl"

func main() {
	// Create logger
	cfg := &sl.Config{
		Level:     "info",
		AddSource: true,
		Writer:    "stderr",
		Type:      "default",
	}

	logger := sl.New(cfg)

	// Use logger
	logger.Debug("Debug message", "key", "value")
	logger.Info("Info message", "key", "value")
	logger.Warn("Warning message", "key", "value")
	logger.Error("Error message", "key", "value")
	logger.Fatal("Fatal message", "key", "value")
}

```

### Context

```golang
package main

import (
	"context"

	"github.com/gosuit/sl"
)

func main() {
	// Create logger
	cfg := &sl.Config{
		Level:     "info",
		AddSource: true,
		Writer:    "stderr",
		Type:      "default",
	}

	logger := sl.New(cfg)

	// Bring the logger with the context
	ctx := context.Background()
	ctx = sl.ContextWithLogger(ctx, logger)

	l := sl.L(ctx)

	l.Info("Info message", "key", "value")
}

```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.