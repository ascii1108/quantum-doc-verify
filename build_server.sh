#!/bin/bash

echo "Building Quantum-Doc-Verify API Server..."

# Create necessary directories
mkdir -p bin
mkdir -p uploads
mkdir -p keys
mkdir -p cmd/server
mkdir -p pkg/logger

# Check if source files exist, if not create them
if [ ! -f pkg/logger/logger.go ]; then
    echo "Creating logger.go..."
    cat > pkg/logger/logger.go << 'EOF'
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
EOF
else
    # If file exists, make sure it has the Warn method
    if ! grep -q "func (l \*Logger) Warn" pkg/logger/logger.go; then
        echo "Adding Warn method to logger.go..."
        sed -i '' '/func (l \*Logger) Error/a\\
// Warn logs a warning message\\
func (l *Logger) Warn(msg string, keyvals ...interface{}) {\\
    l.log("WRN", msg, keyvals...)\\
}\\
' pkg/logger/logger.go
    fi
fi

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Build the server
echo "Building server..."
go build -o bin/server cmd/server/main.go

if [ $? -eq 0 ]; then
    echo "Server built successfully!"
    echo "Run with: ./bin/server --contract=0x364BecF1D9c4D0538929Bd0490AB9C444A2614eE"
else
    echo "Build failed!"
fi