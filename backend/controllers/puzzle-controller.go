package controllers

import (
	"fmt"
	"net/http"

	"github.com/countdown/number-puzzle/numbers"
	"github.com/labstack/echo/v4"
)

func GetPuzzle(c echo.Context) error {
	type puzzleRequest struct {
		Choices string `json:"choices"`
		Lower   int    `json:"lowerBound"`
		Upper   int    `json:"upperBound"`
	}

	var config puzzleRequest
	err := c.Bind(&config)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	p, err := numbers.NewPuzzle(numbers.NewConfig(
		config.Choices,
		config.Lower,
		config.Upper,
	))
	if err != nil {
		return fmt.Errorf("failed to create puzzle: %w", err)
	}

	res := map[string]any{
		"target":  p.Target,
		"numbers": p.Numbers,
	}

	return c.JSON(200, res)
}
