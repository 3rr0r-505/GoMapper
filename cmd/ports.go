package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func nmapPorts() []int {
	// Read the file from embedded assets
	content, err := assets.ReadFile("assets/nmap1000.txt")
	if err != nil {
		log.Fatalf("Error reading embedded file: %v", err)
	}

	portsArray := []int{} // Slice to store the ports

	// Read the file line by line
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		// Split the line by commas to get individual ports and ranges
		ports := strings.Split(line, ",")

		// Process each port or range
		for _, port := range ports {
			port = strings.TrimSpace(port) // Remove spaces
			if port == "" {
				continue
			}

			// Check if the port is a range (contains '-')
			if strings.Contains(port, "-") {
				rangeParts := strings.Split(port, "-")
				start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
				if err != nil {
					fmt.Println("Error parsing start of range:", err)
					continue
				}
				end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
				if err != nil {
					fmt.Println("Error parsing end of range:", err)
					continue
				}

				// Append all ports in the range to the slice
				for i := start; i <= end; i++ {
					portsArray = append(portsArray, i)
				}
			} else {
				// It's a single port, so append it to the slice
				portNumber, err := strconv.Atoi(port)
				if err != nil {
					fmt.Println("Error parsing port:", err)
					continue
				}
				portsArray = append(portsArray, portNumber)
			}
		}
	}

	return portsArray
}
