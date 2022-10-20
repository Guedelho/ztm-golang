//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(serverStatus map[string]int) {
	fmt.Println()
	statusCounter := make(map[int]int)
	for key, value := range serverStatus {
		fmt.Println(key, value)
		switch value {
		case Online:
			statusCounter[Online] += 1
		case Offline:
			statusCounter[Offline] += 1
		case Maintenance:
			statusCounter[Maintenance] += 1
		case Retired:
			statusCounter[Retired] += 1
		}
	}
	fmt.Println(len(serverStatus), statusCounter)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}
	printServerStatus(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	printServerStatus(serverStatus)
	for key := range serverStatus {
		serverStatus[key] = Maintenance
	}
	printServerStatus(serverStatus)
}
