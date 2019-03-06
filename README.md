# Go Project
> [Course Documentation](https://make.sc/bew2.5)

This project was made for the backend web course @ [Make School](https://make.sc/), focusing on [GoLang](https://golang.org/).

# How I Roll
This app is "How I Roll". It's a web app that can simulate dice rolls.

## Local Setup
1. go to project root
1. `cd go; GOOS=js GOARCH=wasm go build -o main.wasm; cd ..`
1. `go get -u github.com/shurcooL/goexec`
1. `goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'`

## Sources
1. [Wikipedia](https://en.wikipedia.org/wiki/Dice_notation)
