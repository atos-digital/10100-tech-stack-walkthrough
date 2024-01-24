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

For a better dev experience, we recommend installing the official `templ-vscode` plugin via the VS Code Marketplace for syntax highlighting.

## Using templ in this codebase

Files with the `.templ` extension contain templ components, which utilise a mix of Go and HTML syntax. For a full list of templ syntax and features, refer to the [templ documentation](https://templ.guide/syntax-and-usage/basic-syntax/).

To generate Go templates for rendering HTML from your `.templ` files, run the following command:

```bash
templ generate
```

Remember to run this command whenever you make changes to your `.templ` files - otherwise, your changes won't show up in the HTML output.

# Code Explanation for 3-templ

A few structural changes have been made since the previous example.

1. The `index.html` file has been removed from the `ui` directory.
2. The `ui` directory now contains three `templ` files: `index.templ`, `body.templ`, and `counter_button.tmpl`. When the `templ generate` command is run, these files are used to generate the `.go` files which are used to render HTML.
   - To see this in action, delete the `.go` files and run `templ generate` again.

## templ

### `counter_button.templ`

```go
templ CounterButton(endpoint string, text string) {
	<button
		hx-get={ endpoint }
		hx-target="#counter"
		hx-swap="outerHTML"
		class="rounded-full bg-blue-600 text-white px-2 py-1"
	>
		{ text }
	</button>
}
```

As shown above, the two buttons from our previous example have been abstracted out into a templ component.

A templ component can take arguments like a normal Go function, which can be used to specify dynamic values. In this case, the `endpoint` argument specifies the htmx endpoint to be called when the button is clicked, and the `text` argument is used to specify the text to be displayed on the button.

### `body.templ`

```go
templ Body() {
	<div id="counter" class="font-bold text-lg text-red-600">Counter: 0</div>
	@CounterButton("/hx/add-one", "Add 1")
	@CounterButton("/hx/reset", "Reset")
}
```

The `Body` templ component contains the `<div>` and `<button>` elements from our previous examples. Templ components can call other templ components with `@`, and pass in different arguments to get a different HTML output. This allows templ components to be easily reused.

### `index.templ`

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

The `Index` templ component takes a `page` argument of type `templ.Component` and renders it within the `<body>` element. This lets us easily swap out the content of the `<body>` element with different HTML.

## Tailwind

### `tailwind.config.js`

```js
module.exports = {
  content: ["./ui/**/*.{templ,html,js}"],
};
```

We have updated the `content` property to include the `templ` extension.

## Go

### `main.go`

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
	ui.Index(ui.Body()).Render(context.Background(), w)
}
```

The `handleIndex` function now uses the `ui.Index` templ component, which takes the `ui.Body` component as an argument, to render the HTML for the page.
