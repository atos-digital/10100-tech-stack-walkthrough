package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter = 0

var counterHTML = `<div id="counter">Counter: %d</div>`

var index = `
<!DOCTYPE html>
<html lang="en">
  <script src="https://unpkg.com/htmx.org@1.9.10" integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC" crossorigin="anonymous"></script>
  <body>
  	<div id="counter">Counter: 0</div>

	<button hx-get="/hx/add-one" hx-target="#counter" hx-swap="outerHTML">
		Add 1
	</button>

	<button hx-get="/hx/reset" hx-target="#counter" hx-swap="outerHTML">
		Reset
	</button>
  </body>
</html>
`

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/hx/add-one", handleAddOne)
	http.HandleFunc("/hx/reset", handleReset)

	log.Print("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(index))
}

func handleAddOne(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, counterHTML, counter)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	counter = 0
	fmt.Fprintf(w, counterHTML, counter)
}
