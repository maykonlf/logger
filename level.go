package logger

import "go.uber.org/zap/zapcore"

// Level is the logging priority. Higher levels are more important.
type Level uint8

func (l Level) toZapCoreLevel() zapcore.Level {
	if l.isValid() {
		return levelMap[l]
	}

	return levelMap[InfoLevel]
}

func (l Level) isValid() bool {
	return l <= FatalLevel
}

const (
	// DebugLevel are typically used on development environemt or for debug purpose.
	// Is more voluminous and are usually disabled in production.
	DebugLevel Level = iota

	// InfoLevel is the default logging priority.
	InfoLevel

	// WarnLevel logs are more important than Info, but not critical
	WarnLevel

	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel

	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel

	// PanicLevel logs a message, then panics.
	PanicLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

var levelMap = map[Level]zapcore.Level{
	DebugLevel:  zapcore.DebugLevel,
	InfoLevel:   zapcore.InfoLevel,
	WarnLevel:   zapcore.WarnLevel,
	ErrorLevel:  zapcore.ErrorLevel,
	DPanicLevel: zapcore.DPanicLevel,
	PanicLevel:  zapcore.PanicLevel,
	FatalLevel:  zapcore.FatalLevel,
}
