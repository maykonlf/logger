# logger
![lint](https://github.com/maykonlf/logger/workflows/lint/badge.svg)

Go structured logger compatible with stackdriver build with zap

## Installation
```bash
go get -u github.com/maykonlf/logger@v1.0.0
```

## Quick Start
```go
log := logger.NewLoggerWithLevel(logger.JSONEncoding, logger.DebugLevel)
defer log.Sync()

log.Debug("debug log", "key", "value")
log.Info("info log", "key", "value")
log.Warn("warn log", "key", "value")
log.Error("error log", "key", "value", "err", errors.New("meu erro!!"))
log.DPanic("d-panic log", "key", "value")
log.Panic("panic log", "key", "value")
log.Fatal("fatal log", "key", "value")
```

## Dependencies
- [Uber's zap](https://github.com/uber-go/zap)
!
