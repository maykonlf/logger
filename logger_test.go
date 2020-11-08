package logger_test

import (
	"github.com/maykonlf/logger"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	l := logger.NewLoggerWithLevel(logger.JSONFormat, logger.InfoLevel)

	for i := 0; i < b.N; i++ {
		l.Info("my info msg")
	}
}
