package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// NewServeMux
// type handler int
// type otherHandler int

// func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	// Do something
// }

// func (o otherHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	// Do something
// }

// func main() {
// 	var h handler
// 	var o otherHandler

// 	mux := http.NewServeMux()
// 	// This can handle an exact match for /dog
// 	// This can also handle: /dog/other/things/here
// 	mux.Handle("/dog/", h)
// 	// This can only handle /cat
// 	mux.Handle("/cat", o)

// 	http.ListenAndServe(":8080", mux)
// }

// Default ServeMux
// func firstHandleFunc(writer http.ResponseWriter, request *http.Request) {
// 	// Do something
// }

// func otherHandleFunc(writer http.ResponseWriter, request *http.Request) {
// 	// Do something
// }

// func main() {
// 	http.HandleFunc("/dog/", firstHandleFunc)
// 	http.HandleFunc("/cat", otherHandleFunc)

// 	http.ListenAndServe(":8080", nil)
// }

// Exercise #1, #2, #3
// var templateContainer *template.Template

// func index(writer http.ResponseWriter, request *http.Request) {
// 	io.WriteString(writer, "In index")
// }
// func dog(writer http.ResponseWriter, request *http.Request) {
// 	io.WriteString(writer, "Im a dog")
// }
// func me(writer http.ResponseWriter, request *http.Request) {
// 	err := templateContainer.ExecuteTemplate(writer, "tpl.gohtml", "Tony")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }
// func init() {
// 	templateContainer = template.Must(template.New("").ParseGlob("templates/*"))
// }
// func main() {
// 	// http.HandleFunc("/", index)
// 	// http.HandleFunc("/dog/", dog)
// 	// http.HandleFunc("/me/", me)
// 	http.Handle("/", http.HandlerFunc(index))
// 	http.Handle("/dog/", http.HandlerFunc(dog))
// 	http.Handle("/me/", http.HandlerFunc(me))
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// Creates new goroutine per connection
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
		if ln == "" {
			// when ln is empty, header is done
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}

	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(conn)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(conn)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleIndex(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApply(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
	</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handleDefault(conn net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>"default"</h1>
		</body>
		</html>
	`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
