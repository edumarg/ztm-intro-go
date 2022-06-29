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

func serversInfo(servers map[string]int) {
	totalServers := 0
	totalOnline := 0
	totalOffline := 0
	totalMaintenance := 0
	totalRetired := 0

	for _, status := range servers {
		totalServers += 1
		if status == Online {
			totalOnline += 1
		} else if status == Offline {
			totalOffline += 1
		} else if status == Maintenance {
			totalMaintenance += 1
		} else if status == Retired {
			totalRetired += 1
		} else {
			fmt.Println(-1)
		}
	}
	fmt.Println("Total servers:", totalServers)
	fmt.Println("Total servers Online:", totalOnline)
	fmt.Println("Total servers Offline:", totalOffline)
	fmt.Println("Total servers Maintenance:", totalMaintenance)
	fmt.Println("Total servers Retired:", totalRetired)

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	serversMap := make(map[string]int)

	for _, server := range servers {
		serversMap[server] = Online
	}
	fmt.Println("servers", serversMap)
	serversInfo(serversMap)

	fmt.Println("darkstar is retired")
	serversMap["darkstar"] = Retired
	serversInfo(serversMap)

	fmt.Println("aiur is Offline")
	serversMap["aiur"] = Offline
	serversInfo(serversMap)

	fmt.Println("all servers are on Maintenance")
	for server := range serversMap {
		serversMap[server] = Maintenance
	}
	serversInfo(serversMap)

}
