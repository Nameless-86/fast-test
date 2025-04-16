package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	inputFile := "ips.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", inputFile, err)
	}
	defer file.Close()

	var ips []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			ips = append(ips, line)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error reading file %s: %v", inputFile, err)
		}
	}

	if len(ips) == 0 {
		log.Printf("No ips found in %s", inputFile)
		return
	}

	for _, ip := range ips {
		fmt.Printf("performing NMAP scan for %s...\n", ip)
		cmd := exec.Command("nmap", "-Pn", ip)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error performing nmap scan for %s: %v", ip, err)
		}

		fileName := fmt.Sprintf("%s.md", ip)
		err = os.WriteFile(fileName, output, 0644)
		if err != nil {
			log.Printf("Error writing nmap results to file %s: %v", fileName, err)
		} else {
			fmt.Printf("NMAP results for %s saved in %s\n", ip, fileName)
		}

	}
}
