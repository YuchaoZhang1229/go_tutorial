// pkg/logger/logger.go
package logger

import "fmt"

const Version = "1.0.0"

func Log(message string) {
	fmt.Printf("[LOG] %s\n", message)
}

func Error(message string) {
	fmt.Printf("[ERROR] %s\n", message)
}
