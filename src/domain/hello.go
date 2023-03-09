package domain

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/labstack/echo/v4"
)

type HelloHandler interface {
	GetHelloHandler(c echo.Context) error
}

type HelloUseCase interface {
	Execute() entity.Hello
}
