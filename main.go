package main

import (
	"fmt"
	"html"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><form action='/convert'><textarea type='text' name='whatver'></textarea><button>click me</button></form></html>"))
}

func convert(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := r.FormValue("whatver")
	w.Write([]byte(fmt.Sprintf("<html>%s", html.EscapeString(input))))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/convert", convert)
	http.ListenAndServe(":3001", mux)
}
