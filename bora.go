package main

import (
	"log"
	"net/http"
	"os"
)

var version string

func init() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)
}

func main() {
	log.Printf("Started version %s\n", version)

	http.HandleFunc("/", InfoHandler)
	http.HandleFunc("/incoming", IncomingHandler)
	http.ListenAndServe(":8020", nil)
}
