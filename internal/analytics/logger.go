package analytics

import (
	"fmt"
	"sync"
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

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Component string `json:"component"`
	Level     string `json:"level"`
	Event     string `json:"event"`
	Message   string `json:"message"`
}

var logs []LogEntry

var mu sync.Mutex

var metrics = map[string]int{
	"pods_created": 0,
	"pods_deleted": 0,
	"errors":       0,
	"logs":         0,
}

func Log(component, level, event, message string) {
	mu.Lock()
	defer mu.Unlock()

	metrics["logs"]++

	switch event {
	case "CreatedPod":
		metrics["pods_created"]++
	case "DeletedPod":
		metrics["pods_deleted"]++
	}

	if level == "error" {
		metrics["errors"]++
	}

	entry := LogEntry{
		Timestamp: time.Now().Format("15:04:05"),
		Component: component,
		Level:     level,
		Event:     event,
		Message:   message,
	}
	logs = append(logs, entry)

	// Color-coded console log
	color := map[string]string{
		"info":    ColorCyan,
		"success": ColorGreen,
		"warn":    ColorYellow,
		"error":   ColorRed,
	}[level]

	if color == "" {
		color = ColorWhite
	}

	fmt.Printf("%s[%s][%s][%s][%s] %s%s\n", color, entry.Timestamp, component, level, event, message, ColorReset)
}

func Summary() map[string]int {
	mu.Lock()
	defer mu.Unlock()

	return map[string]int{
		"pods_created": metrics["pods_created"],
		"pods_deleted": metrics["pods_deleted"],
		"errors":       metrics["errors"],
		"logs":         metrics["logs"],
	}
}

func GetLogs() []LogEntry {
	mu.Lock()
	defer mu.Unlock()

	out := make([]LogEntry, len(logs))
	copy(out, logs)
	return out
}

func ClearLogs() {
	mu.Lock()
	defer mu.Unlock()
	logs = []LogEntry{}
}
