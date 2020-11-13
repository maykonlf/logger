package logger

// Encoding specifies the logging output formatter: console or json
type Encoding uint8

func (f Encoding) value() string {
	if f.isValid() {
		return encodingMap[f]
	}

	return encodingMap[ConsoleEncoding]
}

func (f Encoding) isValid() bool {
	return f <= JSONEncoding
}

const (
	ConsoleEncoding Encoding = iota
	JSONEncoding
)

var encodingMap = map[Encoding]string{
	ConsoleEncoding: "console",
	JSONEncoding:    "json",
}
