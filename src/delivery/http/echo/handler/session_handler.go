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

type sessionHandler struct {
	addSessionUseCase      domain.AddSessionUseCase
	getSessionsUseCase     domain.GetSessionsUseCase
	completeSessionUseCase domain.CompleteSessionUseCase
	cancelSessionUseCase   domain.CancelSessionUseCase
	activateSessionUseCase domain.ActivateSessionUseCase
}

func NewSessionHandler(
	addSessionUseCase domain.AddSessionUseCase,
	getSessionsUseCase domain.GetSessionsUseCase,
	completeSessionUseCase domain.CompleteSessionUseCase,
	cancelSessionUseCase domain.CancelSessionUseCase,
	activateSessionUseCase domain.ActivateSessionUseCase,
) domain.SessionHandler {
	return &sessionHandler{
		addSessionUseCase:      addSessionUseCase,
		getSessionsUseCase:     getSessionsUseCase,
		completeSessionUseCase: completeSessionUseCase,
		cancelSessionUseCase:   cancelSessionUseCase,
		activateSessionUseCase: activateSessionUseCase,
	}
}

func (h *sessionHandler) PostSessionHandler(c echo.Context) error {
	payload := entity.Session{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	session, code, err := h.addSessionUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(session))
}

func (h *sessionHandler) GetSessionsHandler(c echo.Context) error {
	payload := entity.GetSessionParams{}
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

	sessions, code, err := h.getSessionsUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(sessions))
}

func (h *sessionHandler) PostCompleteSessionHandler(c echo.Context) error {
	sessionIDPayload := entity.SessionIDPayload{}
	if err := (&echo.DefaultBinder{}).BindPathParams(c, &sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	medicalRecordPayload := entity.MedicalRecord{}
	if err := c.Bind(&medicalRecordPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(medicalRecordPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	if code, err := h.completeSessionUseCase.Execute(
		sessionIDPayload,
		medicalRecordPayload,
		authorizationHeader,
	); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *sessionHandler) PostCancelSessionHandler(c echo.Context) error {
	sessionIDPayload := entity.SessionIDPayload{}
	if err := c.Bind(&sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	if code, err := h.cancelSessionUseCase.Execute(
		sessionIDPayload,
		authorizationHeader,
	); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *sessionHandler) PostActivateSessionHandler(c echo.Context) error {
	sessionIDPayload := entity.SessionIDPayload{}
	if err := c.Bind(&sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(sessionIDPayload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	if code, err := h.activateSessionUseCase.Execute(
		sessionIDPayload,
		authorizationHeader,
	); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}
