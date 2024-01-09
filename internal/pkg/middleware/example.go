package middleware

import (
	"github.com/labstack/echo/v4"
)

func Example(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := next(ctx)

		if err != nil {
			return err
		}

		return nil
	}
}
