package controllers

import "github.com/labstack/echo/v4"

func GetPuzzle(c echo.Context) error {
	return c.JSON(200, "Hello Word!")
}