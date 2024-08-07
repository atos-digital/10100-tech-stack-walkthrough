# Quickstart

Install and run Tailwind with the following commands:

```bash
npm install -g tailwindcss
```

Install and run templ with the following commands:

```bash
go install github.com/a-h/templ/cmd/templ@latest
go mod tidy
```

Install and run Air with the following commands:

```bash
go install github.com/air-verse/air@latest
air
```

# Getting Started with Air

We use Air to automatically rebuild our Tailwind CSS, templ components, and Go server whenever we make changes to our code.

## Using Air in this codebase

In this codebase, Air has been fully set up and configured. To run Air in this codebase, install it with the following command:

```bash
go install github.com/air-verse/air@latest
```

Next, run the following command. Air will run continuously and automatically rebuild your code whenever you make changes to it. Make some changes to the templ components and reloading your browser to see Air in action.

```bash
air
```

## Setting up Air in your own codebase

If you want to set up Air from scratch in your own codebase, follow the steps below.

Install Air with the following command:

```bash
go install github.com/air-verse/air@latest
```

Initialize the `.air.toml` configuration file with the default settings by running the following command:

```bash
air init
```

See our code exaplanation [below](#air) for a list of changes we made to the default configuration file. For a full list of configuration options for Air, see the [`air_example.toml`](https://github.com/cosmtrek/air/blob/master/air_example.toml) file.

Run Air with the following command. Air will run continuously and automatically rebuild your code whenever you make changes to it.

```bash
air
```

# Code Explanation for 4-air

The only structural change from the previous example is the addition of the `.air.toml` file in the root directory.

## Air

### `.air.toml`

We have made a few changes to the default configuration file generated by Air. In the code snippets below, the default configuration is at the top and our changes are at the bottom, seperated by a line.

```toml
cmd = "go build -o ./tmp/main ."
---
cmd = "npx tailwindcss -i ./input.css -o ./assets/output.css && templ generate && go build -o ./tmp/main ."
```

`npx tailwindcss -i ./input.css -o ./assets/output.css` is the command to build our Tailwind CSS. Note that we have removed the `--watch` flag, since Air will automatically rebuild our CSS whenever we make changes to it.

`templ generate` is the command to build our templ components.

`go build -o ./tmp/main .` is the command to build our Go server.

```toml
exclude_regex = ["_test.go"]
---
exclude_regex = ["_test.go", ".*_templ.go"]
```

This tells Air to ignore changes to files ending in `_templ.go` in addition to the default files ending in `_test.go`.

```toml
include_ext = ["go", "tpl", "tmpl", "html"]
---
include_ext = ["go", "tpl", "tmpl", "templ", "html", "css"]
```

This tells to Air to watch for changes to `.templ` and `.css` files in addition to the default extensions.

## Go

### `main.go`

```go
"github.com/atos-digital/10100-tech-stack-walkthrough/3-templ/ui"
---
"github.com/atos-digital/10100-tech-stack-walkthrough/4-air/ui"
```

The only change in this file is the import path for the `ui` package.
