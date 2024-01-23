package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

var counterHTML = `<div id="counter">Counter: %d</div>`

func main() {
	fs := http.FileServer(http.Dir("./src"))
	http.Handle("/", fs)

	http.HandleFunc("/hx/add-one", handleAddOne)
	http.HandleFunc("/hx/reset", handleReset)

	log.Print("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleAddOne(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, counterHTML, counter)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	counter = 0
	fmt.Fprintf(w, counterHTML, counter)
}
