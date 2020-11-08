package logger

type Format uint8

func (f Format) value() string {
	if f.isValid() {
		return mapFormat[f]
	}

	return mapFormat[ConsoleFormat]
}

func (f Format) isValid() bool {
	return f <= JSONFormat
}

const (
	ConsoleFormat Format = iota
	JSONFormat
)

var mapFormat = map[Format]string{
	ConsoleFormat: "console",
	JSONFormat:    "json",
}
