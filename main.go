package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const tcpAddr = ":42069"

func main() {
	l, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("could not open tcp connection: %s\n", err)
	}
	defer l.Close()

	fmt.Printf("Listening for TCP traffic on %s\n", tcpAddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Accepted connection from %s\n", conn.RemoteAddr())

		go func(c net.Conn) {
			fmt.Printf("Reading data from %s\n", tcpAddr)
			fmt.Println("========================================")

			for line := range getLinesChannel(c) {
				fmt.Printf("read: %s\n", line)
			}

			fmt.Printf("connection to %s is now closed\n", conn.RemoteAddr())
		}(conn)
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
