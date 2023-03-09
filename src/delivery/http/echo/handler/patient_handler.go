package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type patientHandler struct {
	addPatientUseCase        domain.AddPatientUseCase
	getPatientsUseCase       domain.GetPatientsUseCase
	getPatientByIDUseCase    domain.GetPatientByIDUseCase
	updatePatientByIDUseCase domain.UpdatePatientByIDUseCase
	deletePatientByIDUseCase domain.DeletePatientByIDUseCase
	getPatientByNIKUseCase   domain.GetPatientByNIKUseCase
}

func NewPatientHandler(
	addPatientUseCase domain.AddPatientUseCase,
	getPatientsUseCase domain.GetPatientsUseCase,
	getPatientByIDUseCase domain.GetPatientByIDUseCase,
	updatePatientByIDUseCase domain.UpdatePatientByIDUseCase,
	deletePatientByIDUseCase domain.DeletePatientByIDUseCase,
	getPatientByNIKUseCase domain.GetPatientByNIKUseCase,
) domain.PatientHandler {
	return &patientHandler{
		addPatientUseCase:        addPatientUseCase,
		getPatientsUseCase:       getPatientsUseCase,
		getPatientByIDUseCase:    getPatientByIDUseCase,
		updatePatientByIDUseCase: updatePatientByIDUseCase,
		deletePatientByIDUseCase: deletePatientByIDUseCase,
		getPatientByNIKUseCase:   getPatientByNIKUseCase,
	}
}

func (h *patientHandler) PostPatientHandler(c echo.Context) error {
	payload := entity.Patient{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	patient, code, err := h.addPatientUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(map[string]interface{}{
		"patient": patient,
	}))
}

func (h *patientHandler) GetPatientsHandler(c echo.Context) error {
	patients, code, err := h.getPatientsUseCase.Execute()
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(patients))
}

func (h *patientHandler) GetPatientByIDHandler(c echo.Context) error {
	payload := entity.PatientIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	patient, code, err := h.getPatientByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(patient))
}

func (h *patientHandler) PutPatientByIDHandler(c echo.Context) error {
	payload := entity.UpdatePatientPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.updatePatientByIDUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *patientHandler) DeletePatientByIDHandler(c echo.Context) error {
	payload := entity.PatientIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	code, err := h.deletePatientByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *patientHandler) GetPatientByNIKHandler(c echo.Context) error {
	payload := entity.PatientNIKPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	patient, code, err := h.getPatientByNIKUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(patient))
}
