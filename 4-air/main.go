package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/atos-digital/htmx-counter/3-templ/ui"
)

var counter = 0

var counterHTML = `<div id="counter" class="font-bold text-lg text-red-600">Counter: %d</div>`

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hx/add-one", handleAddOne)
	http.HandleFunc("/hx/reset", handleReset)

	log.Print("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	ui.Index(ui.Body()).Render(context.Background(), w)
}

func handleAddOne(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, counterHTML, counter)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	counter = 0
	fmt.Fprintf(w, counterHTML, counter)
}
