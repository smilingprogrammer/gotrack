package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var logfile = "data/log.json"

func StopTask() {

	stateF, err := os.Open(logfile)

	if err != nil {

		fmt.Println("No active task")
		return
	}
	var task TaskTracker

	json.NewDecoder(stateF).Decode(&task)

	stateF.Close()

	task.EndTime = time.Now()
	task.Duration = task.EndTime.Sub(task.StartTime).Round(time.Second).String()

	var log []TaskTracker

	if f, err := os.ReadFile(logfile); err == nil {

		json.Unmarshal(f, &log)
	}

	log = append(log, task)

	f, _ := os.Create(logfile)
	json.NewEncoder(f).Encode(log)
	os.Remove(statefile)

	fmt.Printf("Stopped task: %v (%v) \n", task.Name, task.Duration)
}
