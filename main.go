package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><form method='POST' action='/convert'><textarea type='text' name='whatver'></textarea><button>click me</button></form></html>"))
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
	err := http.ListenAndServe(":"+getListenPort(), mux)
	if err != nil {
		panic(err)
	}
}

func getListenPort() (port string) {
	port = os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}

	fmt.Printf("Listening on port %s \n", port)
	return
}
