# simplelogger

simplelogger is a small logger with no external dependencies. It provides a basic logger API with
useful configuration options such as log level, time-stamping and color output.

## Installation
Run the following to install simplelogger:
```shell
$ go get -u github.com/Renni771/simplelogger
```

## Usage
Simply call the `GetLogger` function and provide a `LoggerOptions` struct to configure simplelogger:

```go
  logger := simplelogger.GetLogger(simplelogger.LoggerOptions{
    LogLevel:            simplelogger.LogLevelInfo,
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
simplelogger supports 7 different log levels with increasing severity:
```go
  LogLevelVerbose   // lowest severity
  LogLevelDebug
  LogLevelInfo
  LogLevelWarn
  LogLevelError
  LogLevelFatal     // highest severity
  LogLevelSilent
```

Setting the log level to `LogLevelWarn` will, for example, suppress all `LogLevelVerbose`, `LogLevelDebug`, `LogLevelInfo` logs.

## Contributing
Please feel free to fork, extend and improve simplelogger. Any contributions are appreciated.

