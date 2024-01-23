# Getting Started with Tailwind

We use [Tailwind](https://tailwindcss.com/) for styling HTML components.

Since Tailwind is a dev dependency, it is not necessary to have it installed to run the code in this codebase. However, we strongly recommend setting up Tailwind using [the instructions below](#using-tailwind-in-this-codebase) so you can experiment with it.

## Using Tailwind in this codebase

In this codebase, Tailwind has been fully set up and configured. To use Tailwind in this codebase, install Tailwind with the following command:

```bash
npm install -g tailwindcss
```

Next, run the following command. The `--watch` flag means Tailwind will watch for changes to your CSS and rebuild it automatically.

```bash
npx tailwindcss -i ./input.css -o ./src/output.css --watch
```

You can now head over to [`index.html`](./src/index.html) and start experimenting with Tailwind classes. Find a full list of utility classes in the [Tailwind documentation](https://tailwindcss.com/docs).

## Setting up Tailwind in your own codebase

If you want to set up Tailwind from scratch in your own codebase, follow the steps below.

Install Tailwind with the following command:

```bash
npm install -g tailwindcss
```

Create a `tailwind.config.js` file in your root directory.

```bash
npx tailwindcss init
```

Create a directory to hold your HTML, CSS, and JS files, and add the path to this directory under the `content` property in your `tailwind.config.js` file. Your files can have different extensions.

In this codebase we use `src`, but you can use any name you want - just make sure to update the `content` property in your `tailwind.config.js` file.

```js
module.exports = {
  content: ["./src/**/*.{html,js}"],
};
```

Create an `input.css` file in your root directory and add the following code. If your code editor supports it, set the language mode to `Tailwind CSS` for correct syntax highlighting.

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

To build your CSS using Tailwind, run the following command. In this command, Tailwind will take the `input.css` file as input and generate an `output.css` file under the `src` directory. Replace the file paths as necessary.

```bash
npx tailwindcss -i ./input.css -o ./src/output.css --watch
```

Add the path of your compiled CSS file to the <head> of your HTML file.

```html
<html>
  <head>
    <link href="./output.css" rel="stylesheet" />
  </head>
</html>
```

# Code Explanation for 2-tailwind

A few structural changes have been made since the previous example.

1. The HTML code has been moved out of `main.go` and into `src/index.html`.
2. The root directory now contains an `input.css` file, which Tailwind uses to generate the `src/output.css` file.

## HTML & Tailwind

Let's have a look at the Tailwind additions to our HTML code in `src/index.html`.

```html
<div id="counter" class="font-bold text-lg text-red-600">Counter: 0</div>

<button
  hx-get="/hx/add-one"
  hx-target="#counter"
  hx-swap="outerHTML"
  class="rounded-full bg-blue-600 text-white px-2 py-1"
>
  Add 1
</button>

<button
  hx-get="/hx/reset"
  hx-target="#counter"
  hx-swap="outerHTML"
  class="rounded-full bg-blue-600 text-white px-2 py-1"
>
  Reset
</button>
```

A number of Tailwind utility classes have been added to the HTML elements to style them. Find a full list of utility classes in the [Tailwind documentation](https://tailwindcss.com/docs).

- `font-bold` makes the text bold.
- `text-lg` makes the text large.
- `text-red-600` makes the text red.
- `rounded-full` makes the buttons round.
- `bg-blue-600` makes the buttons blue.
- `text-white` makes the text white.
- `px-2` adds horizontal padding.
- `py-1` adds vertical padding.

We recommend comparing the rendered HTML in this directory with the HTML in [`1-htmx`](../1-htmx/) to see how Tailwind classes have been used to style the HTML elements. You may notice that Tailwind strips out default styling of HTML elements - for example, the default 8px margin for the `<body>` element has been removed. Read more about Tailwind's base styles [here](https://tailwindcss.com/docs/preflight).

## Go

We have also made some changes to our Go code in `main.go`.

`http.HandleFunc("/", handleIndex)` and the `handleIndex` HTTP handler have been removed and replaced with:

```go
fs := http.FileServer(http.Dir("./src"))
http.Handle("/", fs)
```

`fs := http.FileServer(http.Dir("./src"))` creates a HTTP handler that serves files from the `src` directory and assigns it to the `fs` variable.

`http.Handle("/", fs)` registers the `fs` HTTP handler to be called when the root URL ("/") is accessed. When a request is made to the root, the server responds by serving `index.html` and `output.css` from the `src` directory.

Note that it is generally good practice to put assets such as CSS and JS files in a different directory and serve them separately from HTML files. In this example we have put all our files in the `src` directory for simplicity.
