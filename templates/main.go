package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

type course struct {
	Number string
	Name   string
	Units  string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	AcaYear string
	Fall    semester
	Spring  semester
	Summer  semester
}

type hotel struct{
	Name string
	Address string
	City string
	Zip string
	Region string
}

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

	// Nested Templates
	err = templateContainer.ExecuteTemplate(os.Stdout, "nesting_templates_index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	// Hands-on Exercise #1:
	years := []year{
		year{
			AcaYear: "2020-2021",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		year{
			AcaYear: "2021-2022",
			Fall: semester{
				Term: "Fall",
				Courses: []course{
					course{"CSCI-40", "Introduction to Programming in Go", "4"},
					course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: semester{
				Term: "Spring",
				Courses: []course{
					course{"CSCI-50", "Advanced Go", "5"},
					course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}
	err = templateContainer.ExecuteTemplate(os.Stdout, "exercise1.gohtml", years)
	if err != nil {
		log.Fatalln(err)
	}

	// Hands-on Exercise #2:
	hotels := []hotel{
		hotel{
			Name: "The Marriot",
			Address: "100 Street",
			City: "San Francisco",
			Zip: "90021",
			Region: "Northern",
		},
		hotel{
			Name: "The Inn",
			Address: "200 Street",
			City: "Fresno",
			Zip: "90022",
			Region: "Central",
		},
		hotel{
			Name: "The Wyndham",
			Address: "300 Street",
			City: "Los Angeles",
			Zip: "90023",
			Region: "Southern",
		},
	}

	err = templateContainer.ExecuteTemplate(os.Stdout, "exercise2.gohtml", hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
