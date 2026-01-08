package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server Running")

	// TODO: Server sin libs de terceros
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.Handle("/", http.StripPrefix("/", fs))
	// http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome Hello")
	// })
	// http.ListenAndServe(":80", nil) // el nil hace referencia al parametro del enrutador principal del server

	// TODO: Server con un MANEJADOR de REQUEST de terceros
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello word")
	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars)
		title := vars["title"]
		page := vars["page"]

		fmt.Printf("Params: title: %s, page: %s\n", title, page)
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}
