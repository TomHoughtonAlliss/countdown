package main

import (
	"github.com/countdown/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.POST("/v1/puzzle", controllers.GetPuzzle)

	e.Logger.Fatal(e.Start(":1323"))
}
