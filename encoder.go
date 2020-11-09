package logger

import "go.uber.org/zap/zapcore"

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "eventTime",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      zapcore.OmitKey,
	FunctionKey:    zapcore.OmitKey,
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    stackDriverEncodeLevel,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func stackDriverEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[level])
}
