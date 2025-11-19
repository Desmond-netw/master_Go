package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type ScanRequest struct {
	IP        string `json:"ip"`
	StartPort int    `json:"startPort"`
	EndPort   int    `json:"endPort"`
}

type ScanResult struct {
	Port     int    `json:"port"`
	Status   string `json:"status"`
	Protocol string `json:"protocol"`
	Error    string `json:"error,omitempty"`
}

type ScanResponse struct {
	Results []ScanResult `json:"results"`
	Summary struct {
		TotalPorts int    `json:"totalPorts"`
		OpenPorts  int    `json:"openPorts"`
		ScanTime   string `json:"scanTime"`
	} `json:"summary"`
}

func scanPort(ip string, port int, timeout time.Duration) ScanResult {
	address := fmt.Sprintf("%s:%d", ip, port)
	network := "tcp"

	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return ScanResult{
			Port:     port,
			Status:   "closed",
			Protocol: network,
			Error:    err.Error(),
		}
	}

	defer conn.Close()

	return ScanResult{
		Port:     port,
		Status:   "open",
		Protocol: network,
	}
}

func scanHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate IP address
	if net.ParseIP(req.IP) == nil {
		http.Error(w, "Invalid IP address", http.StatusBadRequest)
		return
	}

	// Validate port range
	if req.StartPort < 1 || req.EndPort > 65535 || req.StartPort > req.EndPort {
		http.Error(w, "Invalid port range. Use 1-65535", http.StatusBadRequest)
		return
	}

	// Limit port range for safety
	if req.EndPort-req.StartPort > 1000 {
		http.Error(w, "Port range too large. Maximum 1000 ports per scan", http.StatusBadRequest)
		return
	}

	startTime := time.Now()
	results := []ScanResult{}
	openCount := 0

	// Scan ports with timeout
	for port := req.StartPort; port <= req.EndPort; port++ {
		result := scanPort(req.IP, port, 2*time.Second)
		results = append(results, result)

		if result.Status == "open" {
			openCount++
		}
	}

	scanTime := time.Since(startTime)

	response := ScanResponse{
		Results: results,
	}
	response.Summary.TotalPorts = len(results)
	response.Summary.OpenPorts = openCount
	response.Summary.ScanTime = scanTime.String()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/scan", scanHandler)

	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Server starting on :8080")
	fmt.Println("Open http://localhost:8080 in your browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
