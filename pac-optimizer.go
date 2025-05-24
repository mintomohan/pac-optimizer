package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	version     = "1.0.0"
	versionText = "PAC File Optimizer\nVersion %s\n"
	helpText    = `Usage: pac-optimizer [options] <input-file> <output-file>

Options:
  --version     Show version information
  --help        Show this help message

Arguments:
  <input-file>  Path to the PAC file to optimize
  <output-file> Path where the optimized PAC file will be saved

Description:
  PAC File Optimizer removes comments, empty spaces, and blank lines from a PAC file.
`
    logo = `
  _____        _____    ____        _   _           _              
 |  __ \ /\   / ____|  / __ \      | | (_)         (_)             
 | |__) /  \ | |      | |  | |_ __ | |_ _ _ __ ___  _ _______ _ __ 
 |  ___/ /\ \| |      | |  | | '_ \| __| | '_ ` + "`" + ` _ \| |_  / _ \ '__|
 | |  / ____ \ |____  | |__| | |_) | |_| | | | | | | |/ /  __/ |   
 |_| /_/    \_\_____|  \____/| .__/ \__|_|_| |_| |_|_/___\___|_|   
                             | |                                   
                             |_|                                   
`
)

func main() {
    fmt.Println(logo)
		
	// Define flags
	showVersion := flag.Bool("version", false, "Show version information")
	showHelp := flag.Bool("help", false, "Show help message")

	// Parse flags
	flag.Parse()

	// Handle version flag
	if *showVersion {
		fmt.Printf(versionText, version)
		return
	}

	// Handle help flag
	if *showHelp {
		fmt.Println(helpText)
		return
	}

	// Check for input and output file arguments
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Error: Input and output file paths are required")
		fmt.Println(helpText)
		os.Exit(1)
	}

	inputFile := args[0]
	outputFile := args[1]

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Printf("Error: Input file '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	// Read input file
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	// Apply regex transformations
	optimizedContent := optimizePAC(string(content))

	// Write optimized content to output file
	err = ioutil.WriteFile(outputFile, []byte(optimizedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("PAC file optimized successfully: %s → %s\n", inputFile, outputFile)
}

func optimizePAC(content string) string {
	// First, detect the type of line ending used in the file
	var lineEnding string
	if strings.Contains(content, "\r\n") {
		lineEnding = "\r\n" // Windows
	} else if strings.Contains(content, "\r") {
		lineEnding = "\r" // Old Mac
	} else {
		lineEnding = "\n" // Unix/Linux
	}

	// Split the content into lines while preserving the original line endings
	lines := strings.Split(content, lineEnding)
	var optimizedLines []string

	// Process each line individually
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		
		// 1. Remove comments, but be careful with URLs
		// Look for "//" that is not part of a URL (not preceded by ":" or "/")
		commentStart := -1
		for j := 0; j < len(line)-1; j++ {
			if line[j] == '/' && line[j+1] == '/' {
				// Check if this is likely part of a URL (preceded by ":" or "/")
				if j > 0 && (line[j-1] == ':' || line[j-1] == '/') {
					// This is likely part of a URL, continue searching
					continue
				}
				// This is likely a comment
				commentStart = j
				break
			}
		}
		
		// Remove the comment if found
		if commentStart >= 0 {
			line = line[:commentStart]
		}
		
		// 2. Check if the line is empty or contains only whitespace
		if strings.TrimSpace(line) == "" {
			// Line is empty or contains only whitespace, remove it completely
			line = ""
		}
		// Note: we don't trim trailing spaces from lines with content
		
		// Add the processed line to our result (even if empty)
		optimizedLines = append(optimizedLines, line)
	}

	// Join the lines back together with the original line ending
	content = strings.Join(optimizedLines, lineEnding)
	
	// 3. Replace multiple consecutive empty lines with a single empty line
	if lineEnding == "\r\n" {
		// For Windows (CRLF)
		emptyLinesRegex := regexp.MustCompile(`(\r\n\r\n)(\r\n)+`)
		content = emptyLinesRegex.ReplaceAllString(content, "\r\n\r\n")
	} else if lineEnding == "\r" {
		// For old Mac (CR)
		emptyLinesRegex := regexp.MustCompile(`(\r\r)(\r)+`)
		content = emptyLinesRegex.ReplaceAllString(content, "\r\r")
	} else {
		// For Unix/Linux (LF)
		emptyLinesRegex := regexp.MustCompile(`(\n\n)(\n)+`)
		content = emptyLinesRegex.ReplaceAllString(content, "\n\n")
	}

	return content
}