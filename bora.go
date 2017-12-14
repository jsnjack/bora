package main

import (
	"fmt"
	"net/http"
)

var version string

// InfoHandler returns version
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bora %s", version)
}

func main() {
	http.HandleFunc("/", InfoHandler)
	http.ListenAndServe(":8020", nil)
}
