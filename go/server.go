package main

/* FILE DOCSTRINGS

Major help came from these sources:
  - Originally was going to use webassembly (*1)
  - A good portion was learned from youtube (*2)
  - Echo was a major resource in this go project
  - Pieces were gathered from my classmates (*3)
  - And of course Dani, Exercisms, and Go devs!!

*1. https://tutorialedge.net/golang/go-webassembly-tutorial/
*2. https://www.youtube.com/watch?v=_pww3NJuWnk
*3. https://github.com/fchikwekwe/FaithBot
*4. https://echo.labstack.com/guide/templates
*5. https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html

END OF DOCSTRINGS */

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// index route applies for `/`
func index(context echo.Context) error {
	return context.String(http.StatusOK, "Amazingg wooooah")
}

// roll function will route to `/roll`
func roll(context echo.Context) error {
	// var dice = context.QueryParam("dice")
	// var roll = Roller.RollDice(dice)
	// var textRoll = strconv.Itoa(roll)

	return context.HTML(http.StatusOK, `<h1>amazing</h1>`)
}

func main() {
	fmt.Println("WELCOME.")

	// e = echo server; borrowed code (*3)
	e := echo.New()

	e.GET("/", index)
	e.GET("/roll", roll)

	e.Start(":8080")
}
