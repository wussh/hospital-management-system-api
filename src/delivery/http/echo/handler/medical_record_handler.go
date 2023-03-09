package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type medicalRecordHandler struct {
	addMedicalRecordUseCase domain.AddMedicalRecordUseCase
}

func NewMedicalRecordHandler(addMedicalRecordUseCase domain.AddMedicalRecordUseCase) domain.MedicalRecordHandler {
	return &medicalRecordHandler{
		addMedicalRecordUseCase: addMedicalRecordUseCase,
	}
}

func (h *medicalRecordHandler) PostMedicalRecordHandler(c echo.Context) error {
	payload := entity.MedicalRecord{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.addMedicalRecordUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}
