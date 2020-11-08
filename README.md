# logger
Go structured logger compatible with stackdriver build with zap

## Installation
```bash
go get -u github.com/maykonlf/logger
```

## Quick Start
```go
l := logger.NewLoggerWithLevel(logger.JSONFormat, logger.WarnLevel)
defer l.Sync()

l.Debug("debug log", "key", "value")
l.Info("info log", "key", "value")
l.Warn("warn log", "key", "value")
l.Error("error log", "key", "value", "err", errors.New("meu erro!!"))
l.DPanic("d-panic log", "key", "value")
l.Panic("panic log", "key", "value")
l.Fatal("fatal log", "key", "value")
```

## Dependencies
- [Uber's zap](https://github.com/uber-go/zap)
