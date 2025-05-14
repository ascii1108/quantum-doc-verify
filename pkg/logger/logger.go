package logger

import (
    "fmt"
    "os"
    "time"
)

// Logger provides structured logging
type Logger struct {
    component string
}

// New creates a new logger for a component
func New(component string) *Logger {
    return &Logger{
        component: component,
    }
}

// Info logs an informational message
func (l *Logger) Info(msg string, keyvals ...interface{}) {
    l.log("INF", msg, keyvals...)
}

// Error logs an error message
func (l *Logger) Error(msg string, keyvals ...interface{}) {
    l.log("ERR", msg, keyvals...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, keyvals ...interface{}) {
    l.log("WRN", msg, keyvals...)
}

// Fatal logs a fatal error message and exits
func (l *Logger) Fatal(msg string, keyvals ...interface{}) {
    l.log("FAT", msg, keyvals...)
    os.Exit(1)
}

// log formats and prints a log message
func (l *Logger) log(level, msg string, keyvals ...interface{}) {
    timestamp := time.Now().Format("15:04:05")
    fmt.Printf("%s %s %s: %s", timestamp, level, l.component, msg)
    
    for i := 0; i < len(keyvals); i += 2 {
        key := keyvals[i]
        var value interface{} = "missing"
        if i+1 < len(keyvals) {
            value = keyvals[i+1]
        }
        fmt.Printf(" %v=%v", key, value)
    }
    fmt.Println()
}