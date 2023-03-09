package postgres

import (
	"fmt"
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/domain"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"gorm.io/gorm"
)

type authenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) domain.AuthenticationRepository {
	return &authenticationRepository{
		db: db,
	}
}

func (r *authenticationRepository) AddRefreshToken(payload entity.Authentication) (int, error) {
	result := r.db.Create(&payload)

	if result.RowsAffected == 0 {
		return http.StatusInternalServerError, result.Error
	}

	return http.StatusOK, nil
}

func (r *authenticationRepository) VerifyRefreshTokenExistence(payload entity.RefreshTokenPayload) (int, error) {
	result := r.db.First(&entity.Authentication{}, "token = ?", payload.RefreshToken)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("refresh token not found")
	}

	return http.StatusOK, nil
}

func (r *authenticationRepository) DeleteRefreshToken(payload entity.RefreshTokenPayload) (int, error) {
	result := r.db.Delete(&entity.Authentication{}, "token = ?", payload.RefreshToken)

	if result.RowsAffected == 0 {
		return http.StatusNotFound, fmt.Errorf("refresh token not found")
	}
	return http.StatusOK, nil
}
