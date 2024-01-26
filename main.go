package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for command-line arguments
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <inputfile> <outputfile>\n", os.Args[0])
		os.Exit(1)
	}

	// Assign command-line arguments to variables
	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	// Read the content of the input file
	content, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Expand environment variables in the content
	expandedContent := os.ExpandEnv(string(content))

	// Write the expanded content to the output file
	err = os.WriteFile(outputFilePath, []byte(expandedContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
