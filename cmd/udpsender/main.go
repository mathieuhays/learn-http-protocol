package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const inputAddr = "localhost:42069"

func main() {
	fmt.Println("UDP sender")

	addr, err := net.ResolveUDPAddr("udp", inputAddr)
	if err != nil {
		log.Fatalf("could not resolve UDP addr %s: %s\n", inputAddr, err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("could not dial UDP addr %s: %s\n", inputAddr, err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		log.Fatal("could not create a reader for stdin")
	}

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error reading line: %s\n", err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("error writing to UDP addr: %s\n", err)
			continue
		}

	}
}
