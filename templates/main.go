package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	template := `
	<!DOCTYPE HTML>
	<html lang="en>
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	filename := "index.html"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("error crating file: ", err)
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(template))
}