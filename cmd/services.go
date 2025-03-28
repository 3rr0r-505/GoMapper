package cmd

import (
	"bufio"
	"net"
	"regexp"
	"strings"
	"time"
)

// ScanResult struct to store scan result details
type ScanResult struct {
	Port    int
	Service string
	Version string
}

// Identify service from banner
func identifyServiceFromBanner(banner string) string {
	banner = strings.ToLower(banner)

	for keyword, service := range ServiceSignatures {
		if strings.Contains(banner, keyword) {
			return service
		}
	}
	return "unknown"
}

// Extract a proper version string from the banner

// Extract a proper version string from the banner
func extractDetailedVersion(banner string) string {
	// Keep everything after the first space or slash (for cases like "Apache/2.4.7 (Ubuntu)")
	re := regexp.MustCompile(`([a-zA-Z0-9]+(?:[/\s-]\d+(\.\d+)*).*)`)
	match := re.FindString(banner)
	if match != "" {
		return match
	}
	return "unknown"
}

// Get banner from a service running on a port
func getBanner(conn net.Conn, port int) {
	conn.SetDeadline(time.Now().Add(3 * time.Second)) // Set timeout
	defer conn.Close()

	reader := bufio.NewReader(conn)
	var banner string

	// Try each probe
	for _, probe := range Probes {
		_, err := conn.Write([]byte(probe))
		if err != nil {
			continue
		}

		// Read response
		var lines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			line = strings.TrimSpace(line)
			lines = append(lines, line)
		}

		// If HTTP response, extract Server header
		// Extract HTTP Server banner properly
		for _, line := range lines {
			if strings.HasPrefix(strings.ToLower(line), "server:") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) > 1 {
					banner = strings.TrimSpace(parts[1])
				}
				break
			}
		}

		// If no Server header, check first line for HTTP version and banner
		if banner == "" && len(lines) > 0 {
			banner = lines[0] // Some servers include version in first line
		}

		// If no Server header, join all lines
		if banner == "" && len(lines) > 0 {
			banner = strings.Join(lines, " | ")
		}

		if banner != "" {
			break
		}
	}

	// DEBUG Banners
	println("DEBUG:: Port:", port, "Banner:", banner)

	// Identify service and version
	service := identifyServiceFromBanner(banner)
	version := extractDetailedVersion(banner)

	mu.Lock()
	results = append(results, ScanResult{Port: port, Service: service, Version: version})
	mu.Unlock()
}

// Get all scan results
func getSummary() []ScanResult {
	mu.Lock()
	defer mu.Unlock()
	return results
}
