# Quickstart

To install Go, follow the official instructions for your operating system [here](https://go.dev/doc/install).

In VS Code, open your `~/.bashrc` or `~/.zshrc` file (depending on whether you are using bash or zsh) using one of the following commands:

```bash
code ~/.bashrc
code ~/.zshrc
```

Add the following to the top of your `~/.bashrc` or `~/.zshrc` file. Replace `$HOME/go/bin` with the path to your Go bin directory.

```
export PATH="$HOME/go/bin:$PATH"
```

To run the code in this directory, execute the following command:

```bash
go run main.go
```

# Getting Started with Go

## Using Go in this codebase

In this codebase, Go has been fully set up. Simply follow the instructions in the [Quickstart](#quickstart) section to install Go and run the code.

## Setting up Go in your own codebase

To initialise your own Go project, run the following command:

```bash
go mod init github.com/<username>/<repository>
```

This will create a `go.mod` file in your project directory. The `go.mod` file contains information about your project, such as its name and dependencies.

You can find a comprehensive tutorial for creating your first Go project [here](https://go.dev/doc/tutorial/getting-started).

# Code Explanation for 0-go

```go
package main
```

Indicates that the code is part of the `main` package, which is the entry point for executables in Go.

```go
var index = `
<!DOCTYPE html>
<html lang="en">
  <body>
   <div id="counter">Counter: 0</div>

    <button>Add 1</button>

    <button>Reset!</button>

  </body>
</html>
`
```

Declares a variable named `index` and assigns it a string containing a HTML code snippet. The HTML represents a web page with a counter and two buttons.

```go
func main() {
 http.HandleFunc("/", handleIndex)

 log.Print("Listening on :8080...")
 http.ListenAndServe(":8080", nil)
}
```

`http.HandleFunc("/", handleIndex)` registers a HTTP handler function `handleIndex` to be called when the root URL ("/") is accessed.

A HTTP handler can contain any kind of logic, but must always take a `http.ResponseWriter` and a `http.Request` as its arguments. The `http.ResponseWriter` is used to write the response to the client. The `http.Request` contains information about the request, such as the URL and the request headers.

`log.Print("Listening on :8080...")` logs a message indicating that the server is listening on port 8080.

`http.ListenAndServe(":8080", nil)` starts a HTTP server listening on port 8080.

```go
func handleIndex(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte(index))
}
```

`w.Write([]byte(index))` writes the contents of the `index` variable as the response to the client's request.
