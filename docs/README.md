# Go Project
[![Go Report Card](https://goreportcard.com/badge/github.com/noltron000/BEW-2-5_how-i-roll)](https://goreportcard.com/report/github.com/noltron000/BEW-2-5_how-i-roll)
> [Course Documentation](https://make.sc/bew2.5)

This project was made for the backend web course @ [Make School](https://make.sc/), focusing on [GoLang](https://golang.org/).

# How I Roll
This app is "How I Roll". It's a web app that can simulate dice rolls.

<!-- ## `v0.0` Functionality
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

These steps will at least get you up and running for development implementations. -->

## `v0.1` Functionality
Version 0.1 has been released. The website is live, and its basic functionality works. However, the site does still break on unchecked errors, for example entering a zero-sided die, or entering broken text. In addition, [docsify]() has been implemented, and a [blogpost]() has been written.

### App Usage
This app is quite simple. There is a textbox, in which you enter a dice in the notation of `XdY`, or a quantity of `X` dice with `Y` sides each. When you click enter, it simulates their summed rolls randomly.

### Serve Locally for Development
If you would like to contribute, please follow these steps:
1. fork or clone this repository.
1. `cd` into project directory within a new terminal instance.
1. run `go run ./server.go` in terminal.
1. navigate to `localhost:8080` in a browser of your choice.

### Special Note
At the moment, this repo has been split to get heroku working properly. This is the main repo, and the sister repo for Heroku is located [here]().

## Sources
1. [Wikipedia](https://en.wikipedia.org/wiki/Dice_notation)
