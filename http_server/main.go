package main

import (
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static_dir"))))
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8090", nil)
}