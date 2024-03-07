package main

import (
	"log"
	"os"
	"text/template"
)

var templateContainer *template.Template

func init() {
	// Use ParseFiles if you only have a few files to parse
	// PREFER to use ParseGlob to parse a whole folder of files
	templateContainer = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// Initial Way to create an html template
	// name := os.Args[1]
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// template := `
	// <!DOCTYPE HTML>
	// <html lang="en>
	// <head>
	// <meta charset="UTF-8">
	// <title>Hello World!</title>
	// </head>
	// <body>
	// <h1>` + name + `</h1>
	// </body>
	// </html>
	// `

	// filename := "index.html"
	// file, err := os.Create(filename)
	// if err != nil {
	// 	log.Fatal("error crating file: ", err)
	// }
	// defer file.Close()
	// io.Copy(file, strings.NewReader(template))

	// Use execute if you have only one template.
	// PREFER to use ExecuteTemplate to specify which template to execute
	// Passing 42 as data to template. Can only pass one piece of data ever. Best to use an aggregate data struct
	err := templateContainer.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	err = templateContainer.ExecuteTemplate(os.Stdout, "tpl2.gohtml", "Tony")
	if err != nil {
		log.Fatalln(err)
	}
}
