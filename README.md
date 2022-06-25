# sLogger

sLogger, or (simple)Logger is a small and simple logger with no external dependencies. sLogger provides a basic logger API with
useful configuration options such as log level, times stamping and color output.

## Usage
Simply call the `GetLogger` function and provide a `LoggerOptions` struct to configure sLogger:

```go
  logger := sLogger.GetLogger(sLogger.LoggerOptions{
    LogLevel:            sLogger.LogLevelInfo,
    WithTimeStamps:      true,
    WithColorizedOutput: true,
  })

  logger.Debug("Here's a debug message...")
  logger.Info("Just some info.")
  logger.Warn("I'm warning you!")
  logger.Error("An error occurred!")
  logger.Fatal("Fatality")
```

## Log level
sLogger supports 7 different log levels with increasing severity:
```go
  LogLevelVerbose   // lowest severity
  LogLevelDebug
  LogLevelInfo
  LogLevelWarn
  LogLevelError
  LogLevelFatal     // highest severity
  LogLevelSilent
```

Setting the log level to `LogLevelWarn` will, for example, suppress all `LogLevelVerbose, LogLevelDebug, LogLevelInfo` logs.

## Contributing
Please feel free to fork, extend and improve sLogger. Any contributions are appreciated.

