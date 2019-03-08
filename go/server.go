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

END OF DOCSTRINGS */

import (
	"fmt"
	"net/http"
	"strconv"

	Roller "./roller"
	"github.com/labstack/echo"
)

// index route applies for `/`
func index(context echo.Context) error {
	return context.String(http.StatusOK, "Amazingg wooooah")
}

func roll(context echo.Context) error {
	var dice = context.QueryParam("dice")
	var roll = strconv.Itoa(Roller.RollDice(dice))
	return context.String(
		http.StatusOK,
		fmt.Sprintf("Your dice: %s\nYour roll: %s",
			dice,
			roll))
}

func main() {
	fmt.Println("WELCOME.")
	e := echo.New()

	e.GET("/", index)
	e.GET("/roll", roll)

	e.Start(":8080")
}
