package main

import (
	"fmt"
)

// LogLevel is a logging level
type LogLevel uint8

// Enum
const (
	// iota starts at 0 in each const block and increments by one
	// until it reaches the end of the block or another assignment
	DebugLevel   LogLevel = iota + 1 // so we start at 1
	WarningLevel                     // 2
	ErrorLevel                       // 3
)

// String implements the fmt.Stringer interface
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return fmt.Sprintf("unknown log level: %d", l)
}

func main() {
	fmt.Println(WarningLevel) // warning

	lvl := LogLevel(19) // lvl is of type LogLevel, initialized with 19
	fmt.Println(lvl)    // unknown log level: 19
}
