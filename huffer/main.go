package main

import (
	"fmt"
	"io"
	"os"

	"github.com/KaviiSuri/coding-challenges/huffer/huffer"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Error: File Path is Required")
		printUsage()
		os.Exit(1)
	}
	command := os.Args[1]
	if command != "encode" && command != "decode" {
		fmt.Println("Error: unrecognized command")
		printUsage()
		os.Exit(1)
	}
	path := os.Args[2]
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Printf("Error: Could not open file %v: %v\n", path, err)
		os.Exit(1)
	}

	outputPath := os.Args[3]
	outputF, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("Error: Could not open file %v: %v\n", outputPath, err)
		os.Exit(1)
	}
	defer outputF.Close()

	switch command {
	case "encode":
		data, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			os.Exit(1)
		}
		err = huffer.Encode(outputF, data)
		if err != nil {
			fmt.Printf("Error encoding file: %v", err)
			os.Exit(1)
		}
		return
	case "decode":
		data, err := huffer.Decode(f)
		if err != nil {
			fmt.Printf("Error encoding file: %v", err)
			os.Exit(1)
		}
		outputF.Write(data)
		return
	}
}

func printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("huffer encode [path] [output]")
	fmt.Println("huffer decode [path] [output]")
}
