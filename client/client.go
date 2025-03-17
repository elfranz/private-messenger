package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("This is a test message from the client"))
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
}
