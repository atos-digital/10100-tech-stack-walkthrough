package main

import (
	"log"
	"net/http"
)

var index = `
<!DOCTYPE html>
<html lang="en">
  <body>
  	<div id="counter">Counter: 0</div>

	<button>
		Add 1
	</button>

	<button>
		Reset!
	</button>
  </body>
</html>
`

func main() {
	http.HandleFunc("/", handleIndex)

	log.Print("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(index))
}
