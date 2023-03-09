package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type clinicHandler struct {
	addClinicUseCase        domain.AddClinicUseCase
	getClinicsUseCase       domain.GetClinicsUseCase
	getClinicByIDUseCase    domain.GetClinicByIDUseCase
	updateClinicByIDUseCase domain.UpdateClinicByIDUseCase
	deleteClinicByIDUseCase domain.DeleteClinicByIDUseCase
}

func NewClinicHandler(
	addClinicUseCase domain.AddClinicUseCase,
	getClinicsUseCase domain.GetClinicsUseCase,
	getClinicByIDUseCase domain.GetClinicByIDUseCase,
	updateClinicByIDUseCase domain.UpdateClinicByIDUseCase,
	deleteClinicByIDUseCase domain.DeleteClinicByIDUseCase,
) domain.ClinicHandler {
	return &clinicHandler{
		addClinicUseCase:        addClinicUseCase,
		getClinicsUseCase:       getClinicsUseCase,
		getClinicByIDUseCase:    getClinicByIDUseCase,
		updateClinicByIDUseCase: updateClinicByIDUseCase,
		deleteClinicByIDUseCase: deleteClinicByIDUseCase,
	}
}

func (h *clinicHandler) PostClinicHandler(c echo.Context) error {
	payload := entity.Clinic{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	code, err := h.addClinicUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *clinicHandler) GetClinicsHandler(c echo.Context) error {
	clinics, code, err := h.getClinicsUseCase.Execute()
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(clinics))
}

func (h *clinicHandler) GetClinicByIDHandler(c echo.Context) error {
	payload := entity.ClinicIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	clinic, code, err := h.getClinicByIDUseCase.Execute(payload.ID)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(clinic))
}

func (h *clinicHandler) PutClinicByIDHandler(c echo.Context) error {
	payload := entity.UpdateClinicPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	code, err := h.updateClinicByIDUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *clinicHandler) DeleteClinicByIDHandler(c echo.Context) error {
	payload := entity.ClinicIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	code, err := h.deleteClinicByIDUseCase.Execute(payload.ID, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}
