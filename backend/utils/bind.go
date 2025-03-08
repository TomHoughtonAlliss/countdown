package utils

import (
	"github.com/countdown/number-puzzle/numbers"
	"github.com/labstack/echo/v4"
)

func CreateConfig(c echo.Context) (dud numbers.Config, err error) {
	type puzzleRequest struct {
		Choices string `json:"choices,omitempty"`
		Lower   int    `json:"lowerBound,omitempty"`
		Upper   int    `json:"upperBound,omitempty"`
	}

	var config puzzleRequest
	err = c.Bind(&config)
	if err != nil {
		return dud, err
	}

	choices := config.Choices
	lower := config.Lower
	upper := config.Upper

	if len(config.Choices) == 0 {
		choices = "LSSSSS"
	}

	if config.Lower == 0 {
		lower = 100
	}

	if config.Upper == 0 {
		upper = 1000
	}

	return numbers.NewConfig(
		choices,
		lower,
		upper,
	), nil
}
