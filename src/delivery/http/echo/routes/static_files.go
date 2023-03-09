package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func StaticFiles(e *echo.Echo) {
	e.GET("/avatar/:avatar", func(c echo.Context) error {
		avatar := c.Param("avatar")
		return c.File(fmt.Sprintf("storage_data/avatar/%s", avatar))
	})
}
