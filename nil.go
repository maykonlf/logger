package logger

// Nil returns a dumb logger instance which does nothing.
// Useful for a safe logger initialization.
func Nil() Logger {
	return &nilLogger{}
}

type nilLogger struct {}

func (n *nilLogger) Debug(string, ...interface{}) {}

func (n *nilLogger) Info(string, ...interface{}) {}

func (n *nilLogger) Warn(string, ...interface{}) {}

func (n *nilLogger) Error(string, ...interface{}) {}

func (n *nilLogger) DPanic(string, ...interface{}) {}

func (n *nilLogger) Panic(string, ...interface{}) {}

func (n *nilLogger) Fatal(string, ...interface{}) {}

func (n *nilLogger) Sync() {}
