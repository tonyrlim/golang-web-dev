package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Reading and writing from a connection
	// err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	// if err != nil {
	// 	log.Println("CONN TIMEOUT")
	// }

	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println(line)
	// 	fmt.Fprintf(conn, "I heard you say: %s\n", line)
	// }
	// defer conn.Close()
	// fmt.Println("Code got here!")

	// Rot13
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		byteSlice := []byte(line)
		rotatedLine := rotate13(byteSlice)

		fmt.Fprintf(conn, "%s - %s\n\n", line, rotatedLine)
	}
	defer conn.Close()
}

func rotate13(byteSlice []byte) []byte {
	var r13 = make([]byte, len(byteSlice))
	for index, value := range byteSlice {
		if value <= 109 {
			r13[index] = value + 13
		} else {
			r13[index] = value - 13
		}
	}
	return r13
}
