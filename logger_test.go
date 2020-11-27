package logger_test

import (
	"os"
	"testing"

	"github.com/maykonlf/logger"
)

var (
	log           logger.Logger
	loggerExample logger.Logger
)

func TestMain(m *testing.M) {
	log = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
	loggerExample = logger.Nil()
	os.Exit(m.Run())
}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Info("my info msg")
	}
}

func ExampleLogger() {
	log := logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
	defer log.Sync()

	log.Debug("my log message")
	log.Debug("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.Info("my log message")
	log.Info("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.Warn("my log message")
	log.Warn("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.Error("my log message")
	log.Error("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.DPanic("my log message")
	log.DPanic("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.Panic("my log message")
	log.Panic("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})

	log.Fatal("my log message")
	log.Fatal("my log message",
		"key1", "value 1",
		"key2", 1,
		"key3", 1.2,
		"key4", true,
		"key5", map[string]interface{}{
			"key":         "value",
			"another_key": 1,
		})
}

func ExampleNew() {
	log := logger.New()
	defer log.Sync()

	log.Info("info message")
}

func ExampleNewLogger() {
	jsonLogger := logger.NewLogger(logger.JSONEncoding)
	defer jsonLogger.Sync()

	jsonLogger.Info("info message")

	consoleLogger := logger.NewLogger(logger.ConsoleEncoding)
	defer consoleLogger.Sync()

	consoleLogger.Info("info message")
}

func ExampleNewLoggerWithLevel() {
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.InfoLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.WarnLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.ErrorLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DPanicLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.PanicLevel)
	_ = logger.NewLoggerWithLevel(logger.JSONEncoding, logger.FatalLevel)

	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.DebugLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.InfoLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.WarnLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.ErrorLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.DPanicLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.PanicLevel)
	_ = logger.NewLoggerWithLevel(logger.ConsoleEncoding, logger.FatalLevel)
}

func ExampleLoggerImp_Debug() {
	loggerExample.Debug("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Info() {
	loggerExample.Info("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Warn() {
	loggerExample.Warn("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Error() {
	loggerExample.Error("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_DPanic() {
	loggerExample.DPanic("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Panic() {
	loggerExample.Panic("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Fatal() {
	loggerExample.Fatal("my log message",
		"additional_data", "some info",
		"my_int_field", 1,
		"my_float_field", 1.2,
		"my_boolean_field", true,
		"map_field", map[string]interface{}{
			"a_key": "some info",
			"b_key": 1,
			"c_key": 1.2,
			"d_key": false,
		})
}

func ExampleLoggerImp_Sync() {
	l := logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
	defer l.Sync()

	l.Info("info message")
	l.Debug("debug info")
}

func ExampleNil() {
	log := logger.Nil()

	log.Info("silent info message")
	log.Debug("silent debug message")
}
