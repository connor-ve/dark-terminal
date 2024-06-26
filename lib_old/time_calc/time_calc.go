package time_calc

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// TimeData Struct to hold previous time
type TimeData struct {
	LastRun time.Time `json:"last_run"`
}

func AfkTime() int64 {
	var newTime int64 = 0
	filename := "time.json"
	var previousTime time.Time

	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return 0
		}
		var timeData TimeData
		err = json.Unmarshal(data, &timeData)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return 0
		}
		previousTime = timeData.LastRun
	} else if !os.IsNotExist(err) {
		fmt.Println("Error checking file:", err)
		return 0
	}

	currentTime := time.Now()

	if !previousTime.IsZero() {
		duration := currentTime.Sub(previousTime)
		newTime = int64(duration.Seconds())
		if newTime >= 10000 {
			newTime = 10000
		}
	}

	timeData := TimeData{LastRun: currentTime}
	data, err := json.Marshal(timeData)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return 0
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return 0
	}
	return newTime
}
