package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var templateContainer *template.Template

type handler int

func (h handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Tony-Key", "This is from Tony")
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		request.Method,
		request.URL,
		request.Form,
		request.Header,
		request.Host,
		request.ContentLength,
	}

	templateContainer.ExecuteTemplate(writer, "tpl.gohtml", data)
}

func init() {
	templateContainer = template.Must(template.New("").ParseGlob("templates/*"))
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
