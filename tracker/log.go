package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func ShowLog() {

	data, err := os.ReadFile(logfile)
	if err != nil {

		fmt.Println("No logs file found")
		return
	}

	var log []TaskTracker

	json.Unmarshal(data, &log)

	today := time.Now().Format("2024-12-17")
	fmt.Printf("🗓️ Log for %v \n", today)

	for _, task := range log {

		if task.StartTime.Format("2024-12-17") == today {
			start := task.StartTime.Format("15:04")
			end := task.EndTime.Format("15:04")
			fmt.Printf("- %v → %v → %v → %v", start, end, task.Name, task.Duration)
		}
	}
}
