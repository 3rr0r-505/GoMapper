package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func nmapPorts() []int {
	// Open the nmap1000.txt file
	file, err := os.Open("cmd/assets/nmap1000.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var portsArray []int // Slice to store the ports

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by commas to get individual ports and ranges
		ports := strings.Split(line, ",")

		// Process each port or range
		for _, port := range ports {
			// Check if the port is a range (contains a '-')
			if strings.Contains(port, "-") {
				// It's a range, so split by the dash and process
				rangeParts := strings.Split(port, "-")
				start, err := strconv.Atoi(rangeParts[0])
				if err != nil {
					fmt.Println("Error parsing start of range:", err)
					continue
				}
				end, err := strconv.Atoi(rangeParts[1])
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

	// Check for errors encountered during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return portsArray

}
