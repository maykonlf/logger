package logger

import "go.uber.org/zap"

type Logger interface {
	// Debug logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes log severity as 'DEBUG'
	Debug(msg string, keyAndValues ...interface{})

	// Info logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes log severity as 'INFO'
	Info(msg string, keysAndValues ...interface{})

	// Warn logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes log severity as 'WARNING'
	Warn(msg string, keyAndValues ...interface{})

	// Error logs a message with some additional context.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes log severity as 'ERROR' and includes the stacktrace.
	Error(msg string, keyAndValues ...interface{})

	// DPanic logs a message with some additional context.
	// In development, the logger then panics. (See DPanicLevel for details.)
	// The variadic key-value pairs are treated as they are in With.
	// Assumes log severity as 'CRITICAL' and includes the stacktrace.
	DPanic(msg string, keysAndValues ...interface{})

	// Panic logs a message with some additional context, then panics.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes severity as 'ALERT' and includes the stacktrace.
	Panic(msg string, keysAndValues ...interface{})

	// Fatal logs a message with some additional context, then calls os.Exit.
	// The variadic key-value pairs are treated as they are in With.
	// Assumes severity as 'EMERGENCY' and includes the stacktrace.
	Fatal(msg string, keysAndValues ...interface{})

	// Sync flushes any buffered log entries.
	Sync()
}

// New returns a new logger with json encoding and default info level.
func New() Logger {
	return NewLoggerWithLevel(JSONEncoding, InfoLevel)
}

// NewLogger returns a new logger with the given encoding(see Encoding for details)
// and default info level.
func NewLogger(encoding Encoding) Logger {
	return NewLoggerWithLevel(encoding, InfoLevel)
}

// NewLoggerWithLevel returns a new logger with the given encoding(see Encoding for details)
// and level (see Level for details).
func NewLoggerWithLevel(encoding Encoding, level Level) Logger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level.toZapCoreLevel()),
		Development:      false,
		Encoding:         encoding.value(),
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
	}

	l, _ := config.Build()
	return &logger{logger: l.Sugar()}
}

type logger struct {
	logger *zap.SugaredLogger
}

func (l *logger) Debug(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *logger) Info(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *logger) Warn(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *logger) Error(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *logger) DPanic(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

func (l *logger) Panic(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

func (l *logger) Fatal(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *logger) Sync() {
	_ = l.logger.Sync()
}
