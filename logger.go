package logger

import "go.uber.org/zap"

type Logger interface {
	Debug(msg string, keyAndValues ...interface{})
	Info(msg string, keysAndValues ...interface{})
	Warn(msg string, keyAndValues ...interface{})
	Error(msg string, keyAndValues ...interface{})
	DPanic(msg string, keysAndValues ...interface{})
	Panic(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
	Sync()
}

func New() Logger {
	return NewLoggerWithLevel(JSONFormat, InfoLevel)
}

func NewLogger(format Format) Logger {
	return NewLoggerWithLevel(format, InfoLevel)
}

func NewLoggerWithLevel(format Format, level Level) Logger {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level.toZapCoreLevel()),
		Development:      false,
		Encoding:         format.value(),
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
	l.logger.Warnw(msg, keysAndValues)
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
