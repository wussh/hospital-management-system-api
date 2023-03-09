package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/hello"
	"github.com/labstack/echo/v4"
)

func HelloRoutes(e *echo.Echo) {
	helloUseCase := hello.NewHelloUseCase()
	helloHandler := handler.NewHelloHandler(helloUseCase)

	e.GET("/hello", helloHandler.GetHelloHandler)
}
