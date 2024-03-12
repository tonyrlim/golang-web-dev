package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

var templateContainer *template.Template

var functionMap = template.FuncMap{
	"upperCase": strings.ToUpper,
	"firstThree": firstThree,
	"fdateMDY": monthDayYear,
	"fdbl": double,
	"fsq": square,
	"fsqrt": sqRoot,
}

func init() {
	// Use ParseFiles if you only have a few files to parse
	// PREFER to use ParseGlob to parse a whole folder of files
	templateContainer = template.Must(template.New("").Funcs(functionMap).ParseGlob("templates/*"))

	// ERROR: Cannot do this because the functions need to be there before you parse
	// templateContainer = templateContainer.Funcs(functionMap)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func double(x int) int {
	return x*2
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
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

	foods := []string{"Pizza", "Fried Chicken", "Sushi", "Tacos"}
	err = templateContainer.ExecuteTemplate(os.Stdout, "arrays.gohtml", foods)
	if err != nil {
		log.Fatalln(err)
	}

	type hometown struct {
		City string
		State string
	}
	myHomeTown:= &hometown{"Dallas", "Texas"}
	err = templateContainer.ExecuteTemplate(os.Stdout, "struct.gohtml", myHomeTown)
	if err != nil {
		log.Fatalln(err)
	}

	// Function Map
	err = templateContainer.ExecuteTemplate(os.Stdout, "functions.gohtml", myHomeTown)
	if err != nil {
		log.Fatalln(err)
	}

	// Time
	err = templateContainer.ExecuteTemplate(os.Stdout, "time.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	// Pipelines
	err = templateContainer.ExecuteTemplate(os.Stdout, "pipelines.gohtml", 5)
	if err != nil {
		log.Fatalln(err)
	}

	// Predefined Global Functions
	xs := []string{"zero", "one", "two", "three", "four", "five"}
	err = templateContainer.ExecuteTemplate(os.Stdout, "global_functions.gohtml", xs)
	if err != nil {
		log.Fatalln(err)
	}
}
