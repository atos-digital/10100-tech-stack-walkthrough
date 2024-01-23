# Getting Started with templ

We use [templ](https://templ.guide/) to generate HTML templates from Go code.

Install templ with the following command:

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

To add templ to your `go.mod` file, run the following command:

```bash
go mod tidy
```

## Using templ in this codebase

templ components in files with the `.templ` extension can be treated and edited as normal HTML. Once you are happy with your changes, run the following command to generate Go templates for rendering HTML from your `.templ` files:

```bash
templ generate
```

# Code Explanation for 3-templ

A few structural changes have been made since the previous example.

1. The `index.html` file has been removed from the `ui` directory.
2. The `ui` directory now contains two `templ` files: `index.templ` and `body.templ`. When the `templ generate` command is run, these files are used to generate the `index_templ.go` and `body_templ.go` files. These files contain Go code that can be used to render HTML.
   - To see this in action, delete the `index_templ.go` and `body_templ.go` files and run `templ generate` again.

## templ

`body.templ` simply contains the same HTML for the `<div>` and `<button>` elements from our previous examples.

The `index.templ` file contains:

```html
templ Index(page templ.Component) {
<!DOCTYPE html>
<html lang="en">
  <head>
    <link href="/assets/output.css" rel="stylesheet" />
    <script
      src="https://unpkg.com/htmx.org@1.9.10"
      integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC"
      crossorigin="anonymous"
    ></script>
  </head>
  <body>
    @page
  </body>
</html>
}
```

The `Index` templ component takes a `page` argument of type `templ.Component` and renders it within the `<body>` element. This lets us to easily swap out the content of the `<body>` element with different HTML.

## Go

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
	ui.Index(ui.Body()).Render(context.Background(), w)
}
```

The `handleIndex` function now uses the `ui.Index` templ component, which takes the `ui.Body` component as an argument, to render the HTML for the page.
