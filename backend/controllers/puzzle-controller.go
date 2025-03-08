package controllers

import (
	"fmt"

	"github.com/countdown/number-puzzle/numbers"
	"github.com/countdown/utils"
	"github.com/labstack/echo/v4"
)

func GetPuzzle(c echo.Context) error {
	config, err := utils.CreateConfig(c)
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}

	p, err := numbers.NewPuzzle(config)
	if err != nil {
		return fmt.Errorf("failed to create puzzle: %w", err)
	}

	res := map[string]any{
		"target":  p.Target,
		"numbers": p.Numbers,
	}

	return c.JSON(200, res)
}
