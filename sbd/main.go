package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type wcResult struct {
	lines int
	words int
	bytes int
}

func main() {
	lFlag := flag.Bool("l", false, "count lines")
	wFlag := flag.Bool("w", false, "count words")
	cFlag := flag.Bool("c", false, "count bytes")

	flag.Parse()

	if !*lFlag && !*wFlag && !*cFlag {
		*lFlag, *wFlag, *cFlag = true, true, true
	}

	for _, filename := range flag.Args() {
		result := wcFile(filename, *lFlag, *wFlag, *cFlag)
		printResult(result, *lFlag, *wFlag, *cFlag, filename)
	}
}

func printResult(result wcResult, countLines, countWords, countBytes bool, filename string) {
	var output string

	if countLines {
		output += fmt.Sprintf("%8d", result.lines)
	}
	if countWords {
		output += fmt.Sprintf("%8d", result.words)
	}
	if countBytes {
		output += fmt.Sprintf("%8d", result.bytes)
	}
	output += fmt.Sprintf(" %s", filename)

	fmt.Println(output)
}

func wcFile(filename string, countLines, countWords, countBytes bool) wcResult {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines, words, bytes int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		if countWords {
			words += countWordsInLine(scanner.Text())
		}
		if countBytes {
			bytes += utf8.RuneCountInString(scanner.Text()) + 1 // Add 1 for the newline character
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning file: %v\n", err)
		os.Exit(1)
	}

	if !countBytes {
		fileInfo, _ := file.Stat()
		bytes = int(fileInfo.Size())
	}

	return wcResult{lines, words, bytes}
}

func countWordsInLine(line string) int {
	scanner := bufio.NewScanner(bufio.NewReader(strings.NewReader(line)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}
