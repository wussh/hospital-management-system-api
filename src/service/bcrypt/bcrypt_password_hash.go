package bcrypt

import (
	"net/http"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordHash struct {
}

func NewBcryptPasswordHash() application.PasswordHash {
	return &bcryptPasswordHash{}
}

func (p *bcryptPasswordHash) Hash(password string) (string, int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return string(hashedPassword), http.StatusOK, nil
}

func (p *bcryptPasswordHash) ComparePassword(plain string, encrypted string) (int, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain)); err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}
