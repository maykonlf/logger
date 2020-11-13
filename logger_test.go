package logger_test

import (
	"testing"

	"github.com/maykonlf/logger"
)

func BenchmarkLogger(b *testing.B) {
	l := logger.NewLoggerWithLevel(logger.JSONEncoding, logger.InfoLevel)

	for i := 0; i < b.N; i++ {
		l.Info("my info msg")
	}
}
