package logger

import "go.uber.org/zap/zapcore"

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
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
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
