package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("could not open %s: %v", inputFilePath, err)
	}

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("========================================")

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)
		currentLine := ""

		for {
			bytes := make([]byte, 8)
			n, err := f.Read(bytes)
			if err != nil {
				if currentLine != "" {
					//fmt.Printf("read: %s\n", currentLine)
					ch <- currentLine
					currentLine = ""
				}

				if errors.Is(err, io.EOF) {
					break
				}

				fmt.Printf("error: %s\n", err.Error())
			}

			str := string(bytes[:n])

			parts := strings.Split(str, "\n")
			for i, part := range parts {
				if i > 0 {
					//fmt.Printf("read: %s\n", currentLine)
					ch <- currentLine
					currentLine = ""
				}

				currentLine += part
			}
		}
	}()

	return ch
}
