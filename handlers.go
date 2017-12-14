package main

import (
	"fmt"
	"net/http"
)

// InfoHandler returns version
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bora %s", version)
}

// IncomingHandler is a handler for incoming webhooks
func IncomingHandler(w http.ResponseWriter, r *http.Request) {

}
