package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// index route applies for `/`
func index(context echo.Context) error {
	return context.String(http.StatusOK, "Amazingg wooooah")
}

func main() {
	fmt.Println("WELCOME.")
	e := echo.New()

	e.GET("/", index)

	e.Start(":8080")
}
