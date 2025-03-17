package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %v", inputFilePath, err)
	}

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("========================================")

	for {
		bytes := make([]byte, 8)
		n, err := f.Read(bytes)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Printf("error: %s\n", err.Error())
		}

		str := string(bytes[:n])
		fmt.Printf("read: %s\n", str)
	}
}
