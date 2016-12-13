package main

import "os"
import "fmt"
import "sync"

const (
	errPrefix  string = "[ERROR]"
	warnPrefix string = "[WARNING]"
)

// Logger keeps track of basic information about where to log messages, and how many have been seen.
type Logger struct {
	key          *sync.Mutex
	output       *os.File
	errorCount   uint
	warningCount uint
}

// NewDefaultLogger creates a new instance of the Logger type that writes messages to os.Stdout
func NewDefaultLogger() *Logger {
	var retval Logger
	retval.output = os.Stdout
	retval.key = new(sync.Mutex)
	return &retval
}

// NewFileLogger creates a new instance of the Logger type that writes messages to the specified file "target"
func NewFileLogger(target *os.File) *Logger {
	retval := NewDefaultLogger()
	retval.output = target
	return retval
}

// LogError Writes a formatted error message and does any related book keeping.
func (subject *Logger) LogError(message string) error {
	subject.key.Lock()
	_, err := fmt.Fprintf(subject.output, "%s %s\n", errPrefix, message)
	subject.errorCount++
	subject.key.Unlock()
	return err
}

// LogWarning Writes a formatted warning message and does any related book keeping.
func (subject *Logger) LogWarning(message string) error {
	subject.key.Lock()
	_, err := fmt.Fprintf(subject.output, "%s %s\n", warnPrefix, message)
	subject.warningCount++
	subject.key.Unlock()
	return err
}

// GetWarningCount returns the number of Warnings that have been logged using Logger "subject".
func (subject *Logger) GetWarningCount() uint {
	subject.key.Lock()
	retval := subject.warningCount
	subject.key.Unlock()
	return retval
}

// GetErrorCount returns the number of Errors that have been logged using Logger "subject".
func (subject *Logger) GetErrorCount() uint {
	subject.key.Lock()
	retval := subject.errorCount
	subject.key.Unlock()
	return retval
}
