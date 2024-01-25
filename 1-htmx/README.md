# Quickstart

To run the code in this directory, execute the following command:

```bash
go run main.go
```

# Getting Started with htmx

## Using htmx in this codebase

In this codebase, htmx has been fully set up, so there's no need to install anything. Read the [code explanation](#code-explanation-for-1-htmx) below to understand how htmx works.

## Setting up htmx in your own codebase

To set up htmx in your own codebase, follow the official instructions [here](https://htmx.org/docs/#installing).

# Code Explanation for 1-htmx

```go
var counter = 0

var counterHTML = `<div id="counter">Counter: %d</div>`
```

`counter` tracks the number of times the button has been clicked.

`counterHTML` is a HTML template for the counter. It contains the format specifier `%d` which will be replaced with the value of `counter` in the handlers.

```html
<head>
  <script
    src="https://unpkg.com/htmx.org@1.9.10"
    integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC"
    crossorigin="anonymous"
  ></script>
</head>
```

Loads htmx via CDN.

```html
<div id="counter">Counter: 0</div>

<button hx-get="/hx/add-one" hx-target="#counter" hx-swap="outerHTML">
  Add 1
</button>

<button hx-get="/hx/reset" hx-target="#counter" hx-swap="outerHTML">
  Reset
</button>
```

`hx-get` specifies the URL to make a GET request to when the button is clicked.

`hx-target` specifies the HTML element that will receive the response from the server. In this case, the responses from the `/hx/add-one` and `/hx/reset` endpoints will replace the content of the `<div/>` element.

`hx-swap` specifies how the response from the server will be swapped in relative to the target element. `outerHTML` replaces the entire target element with the response.

```go
http.HandleFunc("/hx/add-one", handleAddOne)
http.HandleFunc("/hx/reset", handleReset)
```

The handler `handleAddOne` is called when the `/hx/add-one` endpoint is accessed.

The handler `handleReset` is called when the `/hx/reset` endpoint is accessed.

```go
func handleAddOne(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, counterHTML, counter)
}
```

Increments the server-side `counter` variable.

`fmt.Fprintf(w, counterHTML, counter)` writes the updated counter value to the response writer `w` using the `counterHTML` template, replacing the `%d` format specifier with the value of `counter`.

```go
func handleReset(w http.ResponseWriter, r *http.Request) {
	counter = 0
	fmt.Fprintf(w, counterHTML, counter)
}
```

Resets the server-side `counter` variable and responds with the updated counter value using the `counterHTML` template.
