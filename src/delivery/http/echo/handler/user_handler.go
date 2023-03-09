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

type userHandler struct {
	addUserUseCase              domain.AddUserUseCase
	updateUserAvatarUseCase     domain.UpdateUserAvatarUseCase
	deleteUserAvatarUseCase     domain.DeleteUserAvatarUseCase
	getUserByIDUseCase          domain.GetUserByIDUseCase
	getUserDoctorByIDUseCase    domain.GetUserDoctorByIDUseCase
	getAuthenticatedUserUseCase domain.GetAuthenticatedUserUseCase
}

func NewUserHandler(
	addUserUseCase domain.AddUserUseCase,
	updateUserAvatarUseCase domain.UpdateUserAvatarUseCase,
	deleteUserAvatarUseCase domain.DeleteUserAvatarUseCase,
	getUserByIDUseCase domain.GetUserByIDUseCase,
	getUserDoctorByIDUseCase domain.GetUserDoctorByIDUseCase,
	getAuthenticatedUserUseCase domain.GetAuthenticatedUserUseCase,
) domain.UserHandler {
	return &userHandler{
		addUserUseCase:              addUserUseCase,
		updateUserAvatarUseCase:     updateUserAvatarUseCase,
		deleteUserAvatarUseCase:     deleteUserAvatarUseCase,
		getUserByIDUseCase:          getUserByIDUseCase,
		getUserDoctorByIDUseCase:    getUserDoctorByIDUseCase,
		getAuthenticatedUserUseCase: getAuthenticatedUserUseCase,
	}
}

func (h *userHandler) PostUserHandler(c echo.Context) error {
	payload := entity.User{}
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

	addedUser, code, err := h.addUserUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(map[string]interface{}{
		"addedUser": addedUser,
	}))
}

func (h *userHandler) PutUserAvatarHandler(c echo.Context) error {
	payload := entity.UpdateAvatarPayload{}
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

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	payload.Avatar = avatar

	user, code, err := h.updateUserAvatarUseCase.Execute(payload, authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(map[string]interface{}{
		"user": user,
	}))
}

func (h *userHandler) DeleteUserAvatarHandler(c echo.Context) error {
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

	if code, err := h.deleteUserAvatarUseCase.Execute(payload, authorizationHeader); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *userHandler) GetUserByIDHandler(c echo.Context) error {
	payload := entity.UserIDPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	user, code, err := h.getUserDoctorByIDUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(map[string]interface{}{
		"user": user,
	}))
}

func (h *userHandler) GetAuthenticatedUserHandler(c echo.Context) error {
	requestAuthorizationHeader := c.Request().Header["Authorization"][0]
	authorizationHeader := entity.AuthorizationHeader{
		AccessToken: strings.Split(requestAuthorizationHeader, " ")[1],
	}

	user, code, err := h.getAuthenticatedUserUseCase.Execute(authorizationHeader)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(map[string]interface{}{
		"user": user,
	}))
}
