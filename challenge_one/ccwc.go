package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
    // Define a boolean flag -c to count words
    countWords := flag.Bool("c", false, "Count the number of words in the file")
    flag.Parse()

	if !*countWords {
        fmt.Println("Error: go run ccwc.go -c <filename>")
        return
    }

    if len(flag.Args()) == 0 {
        fmt.Println("Please provide a filename.")
        return
    }

    // Open the file in read mode
    file, err := os.Open(flag.Args()[0]) // Get the filename from command-line arguments
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    wordCount := 0

    // Scan through the file line by line
    for scanner.Scan() {
        // Split each line into words
        words := strings.Fields(scanner.Text())
        // Increment wordCount by the number of words in the line
        wordCount += len(words)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Print the number of words if the -c flag is provided
    if *countWords {
        fmt.Println("Number of words in the file:", wordCount)
    }
}