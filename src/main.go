package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LogUrlMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL.Path)
		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}

func main() {
	router := mux.NewRouter()

	bookRouter := router.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// Static assets folder
	const staticDir = "/static/"
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("src"+staticDir))))

	// Generic root request handler
	router.Use(LogUrlMiddleware)
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":80", router)
}
