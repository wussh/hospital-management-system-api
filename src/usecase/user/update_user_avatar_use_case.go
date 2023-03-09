package user

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
)

type updateUserAvatarUseCase struct {
	userRepository  domain.UserRepository
	jwtTokenManager application.TokenManager
}

func NewUpdateUserAvatarUseCase(
	userRepository domain.UserRepository,
	jwtTokenManager application.TokenManager,
) domain.UpdateUserAvatarUseCase {
	return &updateUserAvatarUseCase{
		userRepository:  userRepository,
		jwtTokenManager: jwtTokenManager,
	}
}

func (u *updateUserAvatarUseCase) Execute(
	payload entity.UpdateAvatarPayload,
	authorizationHeader entity.AuthorizationHeader,
) (entity.UpdatedAvatar, int, error) {
	decodedPayload, code, err := u.jwtTokenManager.DecodeAccessTokenPayload(authorizationHeader.AccessToken)
	if err != nil {
		return entity.UpdatedAvatar{}, code, err
	}

	user, code, err := u.userRepository.GetUserByID(decodedPayload.ID)
	if err != nil {
		return entity.UpdatedAvatar{}, code, err
	}

	if user.ID != decodedPayload.ID {
		if user.Role != "Admin" {
			return entity.UpdatedAvatar{}, http.StatusForbidden, fmt.Errorf("restricted resource")
		}
	}

	src, err := payload.Avatar.Open()
	if err != nil {
		return entity.UpdatedAvatar{}, http.StatusInternalServerError, err
	}
	defer src.Close()

	dir, err := os.Getwd()
	if err != nil {
		return entity.UpdatedAvatar{}, http.StatusInternalServerError, err
	}
	newFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), payload.Avatar.Filename)
	path := filepath.Join(dir, "storage_data", "avatar", newFilename)

	dst, err := os.Create(path)
	if err != nil {
		return entity.UpdatedAvatar{}, http.StatusInternalServerError, err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return entity.UpdatedAvatar{}, http.StatusInternalServerError, err
	}

	if code, err := u.userRepository.UpdateUserAvatar(entity.UpdateAvatarLocationPayload{
		ID:     payload.ID,
		Avatar: newFilename,
	}); err != nil {
		return entity.UpdatedAvatar{}, code, err
	}

	var accessPath string
	if os.Getenv("STORAGE") == "local" {
		accessPath = fmt.Sprintf("%s%s/avatar/%s", os.Getenv("HOST"), os.Getenv("PORT"), newFilename)
	} else {
		accessPath = newFilename
	}

	return entity.UpdatedAvatar{
		Avatar: accessPath,
	}, http.StatusOK, nil
}
