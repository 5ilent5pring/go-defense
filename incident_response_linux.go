package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	// Create or open the output file
	outputFile, err := os.Create("incident_report.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write system information to the output file
	writeCommandOutput(outputFile, "System Information", "uname", "-a")

	// Write running processes to the output file
	writeCommandOutput(outputFile, "Running Processes", "ps", "aux")

	// Write network connections to the output file
	writeCommandOutput(outputFile, "Network Connections", "netstat", "-tuln")

	// Write user accounts to the output file
	writeCommandOutput(outputFile, "User Accounts", "cat", "/etc/passwd")

	// Write system logs to the output file
	writeCommandOutput(outputFile, "System Logs", "journalctl", "-xe")

	fmt.Println("Incident report generated successfully: incident_report.txt")
}

func writeCommandOutput(outputFile *os.File, sectionTitle string, command string, args ...string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing %s command: %v\n", command, err)
		return
	}

	fmt.Fprintf(outputFile, "=== %s ===\n%s\n\n", sectionTitle, output)
}
