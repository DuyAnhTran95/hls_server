package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// configure the songs directory name and port
	const port = 8080

	// add a handler for the song files
	http.Handle("/client/",
		http.StripPrefix("/client/", addHeaders(http.FileServer(http.Dir("client")))))

	http.Handle("/media/",
		http.StripPrefix("/media/", addHeaders(http.FileServer(http.Dir("media")))))

	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving on HTTP port: %v\n", port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
