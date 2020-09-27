package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, greetings("Code.education Rocks!"))
}

func greetings(s string) string{
	return "<b>" + s + "</b>"
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}