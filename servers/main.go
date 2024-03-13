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
	// scanner := bufio.NewScanner(conn)
	// for scanner.Scan() {
	// 	line := strings.ToLower(scanner.Text())
	// 	byteSlice := []byte(line)
	// 	rotatedLine := rotate13(byteSlice)

	// 	fmt.Fprintf(conn, "%s - %s\n\n", line, rotatedLine)
	// }
	// defer conn.Close()

	// Exercise #1
	defer conn.Close()
	request(conn)
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

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			mux(conn, line)
		}
		if line == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, line string) {
	// request line
	method := strings.Fields(line)[0]
	fmt.Println("***METHOD", method)
	url := strings.Fields(line)[1]
	fmt.Println("***URL", url)

	if method == "GET" && url == "/" {
		index(conn)
	}
	if method == "GET" && url == "/about" {
		about(conn)
	}
	if method == "GET" && url == "/contact" {
		contact(conn)
	}
	if method == "GET" && url == "/apply" {
		apply(conn)
	}
	if method == "POST" && url == "/apply" {
		applyProcess(conn)
	}
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: &d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: &d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: &d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: &d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title></title>
	</head>
	<body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 204 No Content\r\n")
	fmt.Fprintf(conn, "Content-Length: &d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
