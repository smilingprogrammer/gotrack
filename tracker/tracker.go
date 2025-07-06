package tracker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TaskTracker struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	StartTime time.Time `json:"starttime"`
	EndTime   time.Time `json:"endtime"`
	Duration  string    `json:"duration"`
}

var statefile = "data/state.json"

func StartTask(name string) {

	task := TaskTracker{

		Name:      name,
		Category:  Categorize(name),
		StartTime: time.Now(),
	}

	file, _ := os.Create(statefile)
	json.NewEncoder(file).Encode(task)
	fmt.Println("Started task: ", name)
}
