package handler

import (
	"log"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/util"
	"github.com/labstack/echo/v4"
)

type authenticationHandler struct {
	staffLoginUseCase           domain.StaffLoginUseCase
	staffLogoutUseCase          domain.StaffLogoutUseCase
	updateAuthenticationUseCase domain.UpdateAuthenticationUseCase
}

func NewAuthenticationHandler(
	staffLoginUseCase domain.StaffLoginUseCase,
	staffLogoutUseCase domain.StaffLogoutUseCase,
	updateAuthenticatinoUseCase domain.UpdateAuthenticationUseCase,
) domain.AuthenticationHandler {
	return &authenticationHandler{
		staffLoginUseCase:           staffLoginUseCase,
		staffLogoutUseCase:          staffLogoutUseCase,
		updateAuthenticationUseCase: updateAuthenticatinoUseCase,
	}
}

func (h *authenticationHandler) PostStaffLoginHandler(c echo.Context) error {
	payload := entity.LoginPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	newLogin, code, err := h.staffLoginUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(newLogin))
}

func (h *authenticationHandler) PostStaffLogoutHandler(c echo.Context) error {
	payload := entity.RefreshTokenPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if code, err := h.staffLogoutUseCase.Execute(payload); err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponse())
}

func (h *authenticationHandler) PutAuthenticationHandler(c echo.Context) error {
	payload := entity.RefreshTokenPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(util.ClientErrorResponse(http.StatusBadRequest, err.Error()))
	}

	newAccessToken, code, err := h.updateAuthenticationUseCase.Execute(payload)
	if err != nil {
		if code != http.StatusInternalServerError {
			return c.JSON(util.ClientErrorResponse(code, err.Error()))
		}

		log.Fatal(err)
		return c.JSON(util.ServerErrorResponse())
	}

	return c.JSON(util.SuccessResponseWithData(newAccessToken))
}
