package cmd

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:embed assets/*
var assets embed.FS

var results []ScanResult
var mu sync.Mutex

func printAsciiBanner() {
	content, err := assets.ReadFile("assets/banner.txt")
	if err != nil {
		fmt.Println("Error loading banner:", err)
		return
	}
	fmt.Println(string(content))
}

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	timeout := 10 * time.Second
	maxRetries := 3
	var conn net.Conn
	var err error

	for retries := 0; retries < maxRetries; retries++ {
		conn, err = net.DialTimeout("tcp", address, timeout)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}
	defer conn.Close()

	fmt.Printf("%d/tcp is open\n", port)
	getBanner(conn, port)
}

func parsePorts(portArg string) []int {
	var ports []int

	// Default case
	if portArg == "" || portArg == "1-1000" {
		return nmapPorts()
	}

	// For all ports
	if portArg == "*" {
		for i := 1; i <= 65535; i++ {
			ports = append(ports, i)
		}
		return ports
	}

	// For Port range
	if strings.Contains(portArg, "-") {
		parts := strings.Split(portArg, "-")
		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || start > end {
			fmt.Println("Invalid port range.")
			return nil
		}
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
		return ports
	}

	// For Specific Ports
	if strings.Contains(portArg, ",") {
		parts := strings.Split(portArg, ",")
		for _, p := range parts {
			port, err := strconv.Atoi(strings.TrimSpace(p))
			if err != nil {
				fmt.Println("Invalid port number:", p)
				return nil
			}
			ports = append(ports, port)
		}
		return ports
	}

	port, err := strconv.Atoi(portArg)
	if err != nil {
		fmt.Println("Invalid port number:", portArg)
		return nil
	}
	return []int{port}
}

func printSummary(host string, elapsed time.Duration) {
	endTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\nScan is completed on %s\n\n", endTime)
	fmt.Printf("GoMapper Scan report on %s\n", host)
	fmt.Println("PORT       SERVICE       VERSION")
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})
	for _, result := range results {
		fmt.Printf("%-10s %-13s %s\n", fmt.Sprintf("%d/tcp", result.Port), result.Service, result.Version)
	}
	fmt.Printf("\nGoMapper done: %s is scanned in %s\n", host, elapsed)
}

func StoreOutput(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Scan results saved to", filename)
		fmt.Println()
	}
}

func Execute() {
	host := flag.String("host", "127.0.0.1", "Target host or IP")
	portRange := flag.String("ports", "1-1000", "Port range or specific ports")
	outputFile := flag.String("output", "", "Store result in a txt file")

	// Custom help menu
	flag.Usage = func() {
		printAsciiBanner()
		fmt.Println("Help:")
		fmt.Println("\nOptions:")
		fmt.Println("  --host    string    Target host or IP (default: 127.0.0.1)")
		fmt.Println("  --ports   string    Port range or specific ports (default: top 1000 ports of Nmap)")
		fmt.Println("  --output  string    Store result in a txt file")
		fmt.Println("\nExamples:")
		fmt.Println("  GoMapper --host 192.168.1.1 --ports 80,443")
		fmt.Println("  GoMapper --host example.com --ports 20-25")
		fmt.Println("  GoMapper --host example.com --ports * --output \"scan.txt\"")
		os.Exit(1)
	}

	flag.Parse()

	// Capture all printed output
	var outputBuffer strings.Builder
	multiWriter := io.MultiWriter(os.Stdout, &outputBuffer)

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Use goroutine to copy the output to multiWriter
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		io.Copy(multiWriter, r)
	}()

	// Start scanning process
	printAsciiBanner()
	fmt.Println()
	startTime := time.Now()
	timestamp := startTime.Format("2006-01-02 15:04:05")
	fmt.Printf("GoMapper is running on %s at %s\n", *host, timestamp)

	parsedPorts := parsePorts(*portRange)
	if parsedPorts == nil {
		return
	}

	var scanWg sync.WaitGroup
	for _, port := range parsedPorts {
		scanWg.Add(1)
		go scanPort(*host, port, &scanWg)
	}
	scanWg.Wait()

	printSummary(*host, time.Since(startTime))

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout
	wg.Wait()

	// Store output if --output is provided
	if *outputFile != "" {
		StoreOutput(*outputFile, outputBuffer.String())
	}
}
