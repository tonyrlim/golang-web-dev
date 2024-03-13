package main

import (
	"log"
	"net/http"
	"html/template"
)

var templateContainer *template.Template
type handler int

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	templateContainer.ExecuteTemplate(writer, "tpl.gohtml", request.Form)
}

func init() {
	templateContainer = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
