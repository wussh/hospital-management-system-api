package jwt

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/entity"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/application"
	"github.com/golang-jwt/jwt"
)

type jwtTokenManager struct {
}

func NewJWTTokenManager() application.TokenManager {
	return &jwtTokenManager{}
}

func (j *jwtTokenManager) GenerateRefreshToken(payload entity.AuthenticationPayload) (string, int, error) {
	refreshTokenAge, _ := time.ParseDuration(os.Getenv(("REFRESH_TOKEN_AGE")))
	refreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_KEY")

	token, err := j.generateToken(payload.ID, refreshTokenAge, refreshTokenSecretKey)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func (j *jwtTokenManager) GenerateAccessToken(payload entity.AuthenticationPayload) (string, int, error) {
	accessTokenAge, _ := time.ParseDuration(os.Getenv(("ACCESS_TOKEN_AGE")))
	accessTokenSecretKey := os.Getenv("ACCESS_TOKEN_KEY")

	token, err := j.generateToken(payload.ID, accessTokenAge, accessTokenSecretKey)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func (j *jwtTokenManager) VerifyRefreshToken(refreshToken string) (int, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_KEY")), nil
	})
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid refresh token")
	}

	return http.StatusOK, nil
}

func (j *jwtTokenManager) DecodeRefreshTokenPayload(refreshToken string) (
	entity.AuthenticationPayload, int, error,
) {
	decodedPayload, err := j.decodePayload(refreshToken, os.Getenv("REFRESH_TOKEN_KEY"))
	if err != nil {
		return entity.AuthenticationPayload{}, http.StatusInternalServerError, err
	}

	return decodedPayload, http.StatusOK, nil
}

func (j *jwtTokenManager) DecodeAccessTokenPayload(accessToken string) (
	entity.AuthenticationPayload, int, error,
) {
	decodedPayload, err := j.decodePayload(accessToken, os.Getenv("ACCESS_TOKEN_KEY"))
	if err != nil {
		return entity.AuthenticationPayload{}, http.StatusInternalServerError, err
	}

	return decodedPayload, http.StatusOK, nil
}

func (j *jwtTokenManager) decodePayload(token string, secretKey string) (
	entity.AuthenticationPayload, error,
) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return entity.AuthenticationPayload{}, err
	}

	authenticationPayload := entity.AuthenticationPayload{
		ID: claims["id"].(string),
	}
	return authenticationPayload, nil
}

func (j *jwtTokenManager) generateToken(id string, expirationTime time.Duration, secretKey string) (string, error) {
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
