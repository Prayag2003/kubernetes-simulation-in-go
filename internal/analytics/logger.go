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

	color := ColorWhite
	switch level {
	case "info":
		color = ColorCyan
	case "success":
		color = ColorGreen
	case "warn":
		color = ColorYellow
	case "error":
		color = ColorRed
	}

	fmt.Printf("%s[%s][%s][%s][%s] %s%s\n", color, time.Now().Format("15:04:05"), component, level, event, message, ColorReset)
}

func Summary() {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("\n--- Simulation Summary ---")
	fmt.Printf("Total Pods Created : %d\n", metrics["pods_created"])
	fmt.Printf("Total Pods Deleted : %d\n", metrics["pods_deleted"])
	fmt.Printf("Errors Logged      : %d\n", metrics["errors"])
	fmt.Printf("Total Log Events   : %d\n", metrics["logs"])
	fmt.Println("---------------------------")
}
