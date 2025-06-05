package utils

import (
	"fmt"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[97m"
)

func LogInfo(tag, message string) {
	fmt.Printf("%s[%s] %s %s%s\n", ColorCyan, time.Now().Format("15:04:05"), tag, message, ColorReset)
}

func LogSuccess(tag, message string) {
	fmt.Printf("%s[%s] %s %s%s\n", ColorGreen, time.Now().Format("15:04:05"), tag, message, ColorReset)
}

func LogError(tag, message string) {
	fmt.Printf("%s[%s] %s %s%s\n", ColorRed, time.Now().Format("15:04:05"), tag, message, ColorReset)
}

func LogWarn(tag, message string) {
	fmt.Printf("%s[%s] %s %s%s\n", ColorYellow, time.Now().Format("15:04:05"), tag, message, ColorReset)
}
