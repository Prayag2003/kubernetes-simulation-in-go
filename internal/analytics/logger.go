package analytics

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type LogEntry struct {
	Time    time.Time `json:"time"`
	Source  string    `json:"source"`
	Level   string    `json:"level"`
	Event   string    `json:"event"`
	Details string    `json:"details"`
}

type Collector struct {
	sync.Mutex
	entries []LogEntry

	PodsCreated int
	PodsDeleted int
	Errors      int
}

var singleton *Collector = &Collector{}

func Log(source, level, event, details string) {
	entry := LogEntry{
		Time:    time.Now(),
		Source:  source,
		Level:   level,
		Event:   event,
		Details: details,
	}

	singleton.Lock()
	defer singleton.Unlock()

	singleton.entries = append(singleton.entries, entry)

	switch event {
	case "CreatePod":
		singleton.PodsCreated++
	case "DeletePod":
		singleton.PodsDeleted++
	}
	if level == "error" {
		singleton.Errors++
	}

	fmt.Printf("[%s] %-10s %-6s %-15s %s\n",
		entry.Time.Format("15:04:05"),
		source,
		level,
		event,
		details)
}

func ExportJSON(path string) error {
	singleton.Lock()
	defer singleton.Unlock()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(singleton.entries)
}

func PrintSummary() {
	singleton.Lock()
	defer singleton.Unlock()

	fmt.Println("\n--- Simulation Summary ---")
	fmt.Printf("Total Pods Created : %d\n", singleton.PodsCreated)
	fmt.Printf("Total Pods Deleted : %d\n", singleton.PodsDeleted)
	fmt.Printf("Errors Logged      : %d\n", singleton.Errors)
	fmt.Printf("Total Log Events   : %d\n", len(singleton.entries))
	fmt.Println("---------------------------")
}
