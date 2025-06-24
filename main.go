package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server Running")
	fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome Hello")
	})

	http.ListenAndServe(":80", nil)
}
