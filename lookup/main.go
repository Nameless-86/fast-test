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
	inputFile := "subdomains.txt"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening file %s: %v", inputFile, err)
	}
	defer file.Close()

	var domains []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			domains = append(domains, line)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error reading file %s: %v", inputFile, err)
		}
	}
	if len(domains) == 0 {
		log.Printf("No domains found in %s", inputFile)
		return
	}

	for _, domain := range domains {
		fmt.Printf("performing nslookup for %s...\n", domain)
		cmd := exec.Command("nslookup", domain)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error performing ns lookup for %s: %v", domain, err)
		}

		fileName := fmt.Sprintf("%s.md", domain)
		err = os.WriteFile(fileName, output, 0644)
		if err != nil {
			log.Printf("Error writing nslookup results to file %s: %v", fileName, err)
		} else {
			fmt.Printf("Nslookup result for %s saved in %s\n", domain, fileName)
		}
	}
}
