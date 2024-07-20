package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Args[1]
	fmt.Println("Parsing JSON ...", path)
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Errorf("error reading file: %v", err)
		os.Exit(0)
	}

	p := NewParser(string(input))

	output, err := p.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Parsed output: %v\n", output)
}
