package main

import (
	"fmt"
	"net/http"
)

func initFileServer() {
	// Wire up static asset handler
	const staticPathPrefix string = "/static/"
	fs := http.FileServer(http.Dir("src/static/"))
	http.Handle(staticPathPrefix, http.StripPrefix(staticPathPrefix, fs))
}

func main() {
	// Generic catch-all request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Requested URL: %s\n", r.URL.Path)
	})

	initFileServer()

	http.ListenAndServe(":80", nil)
}
