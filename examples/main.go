package main

import (
	"errors"
	"fmt"

	"github.com/maykonlf/logger"
)

func main() {
	l := logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
	defer l.Sync()

	var levels = []logger.Level{
		logger.DebugLevel,
		logger.InfoLevel,
		logger.WarnLevel,
		logger.ErrorLevel,
		logger.DPanicLevel,
		logger.PanicLevel,
		logger.FatalLevel,
	}

	for _, level := range levels {
		fmt.Printf("\n\n----------<logging level: %v>----------\n", level)
		log := logger.NewLoggerWithLevel(logger.JSONEncoding, level)
		log.Debug("debug log", "key", "value")
		log.Info("info log", "key", "value")
		log.Warn("warn log", "key", "value")
		log.Error("error log", "err", errors.New("meu erro!!"))
		log.DPanic("d-panic log", "key", "value")
		// log.Panic("panic log", "key", "value")
		// log.Fatal("fatal log", "key", "value")
		log.Sync()
	}
}
