package middleware

import (
	"net/http"
	"os"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &jwt.Claims{},
		SigningKey:              []byte(os.Getenv("ACCESS_TOKEN_KEY")),
		ErrorHandlerWithContext: JWTErrorChecker,
	})
}

func JWTErrorChecker(err error, c echo.Context) error {
	return c.JSON(util.ClientErrorResponse(http.StatusUnauthorized, "unauthorized"))
}
