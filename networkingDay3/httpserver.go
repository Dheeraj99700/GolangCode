package main

import (
	"fmt"
	"net/http"
)

// handler struct
type handler struct{}

// ServeHTTP method
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, %s!", r.URL.Path[1:])
}

func main() {
	// create server
	http.Handle("/", &handler{})
	http.ListenAndServe(":8082", nil)
}
