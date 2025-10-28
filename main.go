package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type APIResponse struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data []ContractInfo  `json:"data"`
}

type ContractInfo struct {
	SourceCode string `json:"sourceCode"`
}

type SourceCodeData struct {
	Language string                      `json:"language"`
	Sources  map[string]SourceFileInfo  `json:"sources"`
}

type SourceFileInfo struct {
	Content string `json:"content"`
}

func main() {
	var (
		chain    = flag.String("chain", "xlayer", "chain short name")
		address  = flag.String("address", "", "contract address (required)")
		outDir   = flag.String("out", "", "output directory (default: address value)")
	)
	flag.Parse()

	// Validate required flag
	if *address == "" {
		fmt.Fprintf(os.Stderr, "Error: --address is required\n\n")
		flag.Usage()
		os.Exit(1)
	}

	// Set output directory default
	if *outDir == "" {
		*outDir = *address
	}

	// Fetch contract source code
	url := fmt.Sprintf("https://www.oklink.com/api/v5/explorer/contract/verify-contract-info?chainShortName=%s&contractAddress=%s",
		*chain, *address)

	fmt.Printf("Fetching source code from %s...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	// Parse API response
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Fatalf("Failed to parse API response: %v", err)
	}

	if len(apiResp.Data) == 0 {
		log.Fatalf("No contract data found")
	}

	// Parse source code JSON string
	var sourceCodeData SourceCodeData
	if err := json.Unmarshal([]byte(apiResp.Data[0].SourceCode), &sourceCodeData); err != nil {
		log.Fatalf("Failed to parse source code: %v", err)
	}

	// Create output directory
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Save source files
	fmt.Printf("Saving %d source files to %s...\n", len(sourceCodeData.Sources), *outDir)
	saved := 0
	for filename, fileInfo := range sourceCodeData.Sources {
		// Create subdirectories based on forward slashes in filename
		filePath := filepath.Join(*outDir, filename)
		dir := filepath.Dir(filePath)
		
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Printf("Warning: Failed to create directory %s: %v", dir, err)
				continue
			}
		}

		// Decode escaped characters in the content
		content := fileInfo.Content
		// Replace \n with actual newlines
		content = strings.ReplaceAll(content, "\\n", "\n")
		// Replace \\ with \
		content = strings.ReplaceAll(content, "\\\\", "\\")

		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			log.Printf("Warning: Failed to save %s: %v", filename, err)
			continue
		}
		saved++
	}

	fmt.Printf("Successfully saved %d files to %s\n", saved, *outDir)
}

