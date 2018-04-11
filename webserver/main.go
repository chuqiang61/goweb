package main

import (
	"io"
	"net/http"
)

func firstpage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, " <h1>hello go</h1>")

}

func main() {
	http.HandleFunc("/", firstpage)
	http.ListenAndServe(":8000", nil)
}
