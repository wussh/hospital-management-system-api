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

type scheduleHandler struct {
	addScheduleUseCase           domain.AddScheduleUseCase
	getSchedulesUseCase          domain.GetSchedulesUseCase
	getScheduleByIDUseCase       domain.GetScheduleByIDUseCase
	updateScheduleByIDUseCase    domain.UpdateScheduleByIDUseCase
	deleteScheduleByIDUseCase    domain.DeleteScheduleByIDUseCase
	getScheduleByDoctorIDUseCase domain.GetSchedulesByDoctorIDUseCase
}

func NewScheduleHandler(
	addScheduleUseCase domain.AddScheduleUseCase,
	getSchedulesUseCase domain.GetSchedulesUseCase,
	getScheduleByIDUseCase domain.GetScheduleByIDUseCase,
	updateScheduleByIDUseCase domain.UpdateScheduleByIDUseCase,
	deleteScheduleByIDUseCase domain.DeleteScheduleByIDUseCase,
	getScheduleByDoctorIDUseCase domain.GetSchedulesByDoctorIDUseCase,
) domain.ScheduleHandler {
	return &scheduleHandler{
		addScheduleUseCase:           addScheduleUseCase,
		getSchedulesUseCase:          getSchedulesUseCase,
		getScheduleByIDUseCase:       getScheduleByIDUseCase,
		updateScheduleByIDUseCase:    updateScheduleByIDUseCase,
		deleteScheduleByIDUseCase:    deleteScheduleByIDUseCase,
		getScheduleByDoctorIDUseCase: getScheduleByDoctorIDUseCase,
	}
}

func (h *scheduleHandler) PostScheduleHandler(c echo.Context) error {
	payload := entity.Schedule{}
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

	code, err := h.addScheduleUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *scheduleHandler) GetSchedulesHandler(c echo.Context) error {
	Schedules, code, err := h.getSchedulesUseCase.Execute()
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(Schedules))
}

func (h *scheduleHandler) GetScheduleByIDHandler(c echo.Context) error {
	payload := entity.ScheduleIDPayload{}
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

	schedules, code, err := h.getScheduleByIDUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(schedules))
}

func (h *scheduleHandler) PutScheduleByIDHandler(c echo.Context) error {
	payload := entity.UpdateSchedulePayload{}
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

	code, err := h.updateScheduleByIDUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *scheduleHandler) DeleteScheduleByIDHandler(c echo.Context) error {
	payload := entity.ScheduleIDPayload{}
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

	code, err := h.deleteScheduleByIDUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *scheduleHandler) GetSchedulesByDoctorIDHandler(c echo.Context) error {
	payload := entity.UserIDPayload{}
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

	schedules, code, err := h.getScheduleByDoctorIDUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(schedules))
}
