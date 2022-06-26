package simplelogger

import (
	"fmt"
	"time"
)

const (
	resetColor = "\033[0m"
	green      = "\033[32m"
	blue       = "\033[34m"
	yellow     = "\033[33m"
	red        = "\033[31m"
)

// Any logs of a type lower than the current one will be suppressed.
type LogLevel struct {
	slug   string
	number int
}

var (
	LogLevelVerbose = LogLevel{"", 1}
	LogLevelDebug   = LogLevel{"[debug]", 2}
	LogLevelInfo    = LogLevel{"[info]", 3}
	LogLevelWarn    = LogLevel{"[warn]", 4}
	LogLevelError   = LogLevel{"[error]", 5}
	LogLevelFatal   = LogLevel{"[fatal]", 6}
	LogLevelSilent  = LogLevel{"", 7}
)

// The options for a Logger.
type LoggerOptions struct {
	LogLevel            LogLevel
	WithTimeStamps      bool     `default:"false"`
	WithFilePath        bool     `default:"false"`
	WithColorizedOutput bool     `default:"false"`
}

// A logger which prints to the standard output.
type Logger struct {
	options LoggerOptions
}

// Creates a new Logger with the provided LoggerOptions.
func GetLogger(options LoggerOptions) *Logger {
	return &Logger{options}
}

func (l Logger) Debug(format string) {
	if l.options.LogLevel.number <= LogLevelDebug.number {
		print(l.options, LogLevelDebug, format)
	}
}

func (l Logger) Info(format string) {
	if l.options.LogLevel.number <= LogLevelInfo.number {
		print(l.options, LogLevelInfo, format)
	}
}

func (l Logger) Warn(format string) {
	if l.options.LogLevel.number <= LogLevelWarn.number {
		print(l.options, LogLevelWarn, format)
	}
}

func (l Logger) Error(format string) {
	if l.options.LogLevel.number <= LogLevelError.number {
		print(l.options, LogLevelError, format)
	}
}

func (l Logger) Fatal(format string) {
	if l.options.LogLevel.number <= LogLevelFatal.number {
		print(l.options, LogLevelFatal, format)
	}
}

func print(options LoggerOptions, logLevel LogLevel, format string) {
	var (
		timestamp, output = "", ""
	)

	if options.WithTimeStamps {
		timestamp = fmt.Sprintf(" %s:  ", getTimestamp())
	}

	output = fmt.Sprintf("%s\t%s%s\n", logLevel.slug, timestamp, format)

	if options.WithColorizedOutput {
		fmt.Printf("%s", colorize(logLevel, output))
		return
	}

	fmt.Printf("%s", output)
}

func colorize(logLevel LogLevel, format string) (colorizedFormat string) {
	switch logLevel {
	case LogLevelDebug:
		colorizedFormat = fmt.Sprintf("%s%s%s", green, format, resetColor)
	case LogLevelInfo:
		colorizedFormat = fmt.Sprintf("%s%s%s", blue, format, resetColor)
	case LogLevelWarn:
		colorizedFormat = fmt.Sprintf("%s%s%s", yellow, format, resetColor)
	case LogLevelError:
		colorizedFormat = fmt.Sprintf("%s%s%s", red, format, resetColor)
	case LogLevelFatal:
		colorizedFormat = fmt.Sprintf("%s%s%s", red, format, resetColor)
	}
	return
}

func getTimestamp() string {
	now := time.Now()
	return fmt.Sprintf("%d:%d:%d:%d", now.Hour(), now.Minute(), now.Second(), now.UnixMilli())
}
