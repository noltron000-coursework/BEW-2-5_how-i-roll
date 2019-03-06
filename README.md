# Go Project
> [Course Documentation](https://make.sc/bew2.5)

This project was made for the backend web course @ [Make School](https://make.sc/), focusing on [GoLang](https://golang.org/).

# How I Roll
This app is "How I Roll". It's a web app that can simulate dice rolls.

## `v0.0` Functionality
At the moment, there is little to no functionality. However, one can get a local build running by following these steps:
1. fork or clone this repository.
1. `cd` into project directory within a new terminal instance.
1. run `go run ./go/main.go` in terminal.

This part is also important on this branch, for now.
1. `cd go; GOOS=js GOARCH=wasm go build -o main.wasm; cd ..`
1. `go get -u github.com/shurcooL/goexec`
1. `goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'`

These steps will at least get you up and running for development implementations.

## Sources
1. [Wikipedia](https://en.wikipedia.org/wiki/Dice_notation)
