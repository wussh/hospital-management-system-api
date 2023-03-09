package handler

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type helloHandler struct {
	helloUseCase domain.HelloUseCase
}

func NewHelloHandler(helloUseCase domain.HelloUseCase) domain.HelloHandler {
	return &helloHandler{
		helloUseCase: helloUseCase,
	}
}

func (h *helloHandler) GetHelloHandler(c echo.Context) error {
	data := h.helloUseCase.Execute()

	return c.JSON(util.SuccessResponseWithData(data))
}
