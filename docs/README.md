# Go Project
[![Go Report Card](https://goreportcard.com/badge/github.com/noltron000/BEW-2-5_how-i-roll)](https://goreportcard.com/report/github.com/noltron000/BEW-2-5_how-i-roll)
> [Course Documentation](https://make.sc/bew2.5)

This project was made for the backend web course @ [Make School](https://make.sc/), focusing on [GoLang](https://golang.org/).

# How I Roll
This app is "How I Roll". It's a web app that can simulate dice rolls.

## `v0.0` Functionality
At the moment, there is little to no functionality. However, one can get a local build running by following these steps:
1. fork or clone this repository.
1. `cd` into project directory within a new terminal instance.
1. run `go run ./go/main.go` in terminal.
1. navigate to `localhost:8080` in a browser of your choice.


**Dice Roller**<br />
Unfortunately the dice roller has not been connected to the html yet. However the program itself still works in the command line. Follow the steps above with the following differences:
- at step 3, replace `./go/main.go` with `./go/roller/roller.go PARAMS`.
- replace `PARAMS` with a dice amount, such as *1d6* or *8d4* before running.
- do not follow the rest of the steps.

These steps will at least get you up and running for development implementations.

## Sources
1. [Wikipedia](https://en.wikipedia.org/wiki/Dice_notation)
