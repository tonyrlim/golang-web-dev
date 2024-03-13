package main

import (
	"fmt"
	"net/http"
)

type whatever int

func (wh whatever) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Do whatever
	fmt.Fprintln(w, "Doing whatever in this handler")
}

func main() {
	var wh whatever
	http.ListenAndServe(":8080", wh)
}
