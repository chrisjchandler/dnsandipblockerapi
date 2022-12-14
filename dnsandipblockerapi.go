package main

import (
	"fmt"
	"net/http"
	//"strings"
)

// BlockList is a list of IP addresses or DNS names that should be blocked by iptables
var BlockList []string

func main() {
	http.HandleFunc("/block", handleBlock)
	http.ListenAndServe(":8080", nil)
}

func handleBlock(w http.ResponseWriter, r *http.Request) {
	// Get the IP address or DNS name to block from the request
	address := r.URL.Query().Get("address")

	// Add the address to the block list
	BlockList = append(BlockList, address)

	// Add the address to the iptables DROP rule
	addToIptables(address)

	fmt.Fprintf(w, "Successfully added %s to the block list", address)
}

func addToIptables(address string) {
	// TODO: Add the address to the iptables DROP rule using the "iptables" command
	// You will need to run this command as the root user or using sudo
	// Here is an example of how you could do this using the exec package:
	//
	// cmd := exec.Command("iptables", "-A", "INPUT", "-s", address, "-j", "DROP")
	// err := cmd.Run()
	// if err != nil {
	//   fmt.Fprintf(w, "Error adding %s to iptables: %s", address, err)
	// }
}

