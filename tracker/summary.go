package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func ShowSummary() {

	data, err := os.ReadFile(logfile)

	if err != nil {

		fmt.Println("No log found")
		return
	}

	var log []TaskTracker
	json.Unmarshal(data, &log)

	today := time.Now().Format("2024-12-17")
	stats := make(map[string]time.Duration)

	for _, task := range log {

		if task.StartTime.Format("2024-12-17") == today {
			d, _ := time.ParseDuration(task.Duration)
			stats[task.Category] += d
		}
	}

	fmt.Println("Today's Summary")
	var total time.Duration

	for cat, dur := range stats {

		bars := strings.Repeat("█", int(dur.Minutes()/5))
		fmt.Printf("%-10s: %-20s %s\n", cat, bars, dur.String())
		total += dur
	}

	fmt.Printf("Total       :%v\n", total.String())
}
