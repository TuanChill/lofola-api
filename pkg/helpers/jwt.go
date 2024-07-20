package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
)

func JwtProvider(payload models.PayloadToken) (string, string, error) {
	accessToken, err := generateAccessToken(payload)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateRefreshToken(payload)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(payload models.PayloadToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       payload.ID,
		"email":    payload.Email,
		"userName": payload.UserName,
		"exp":      time.Now().Add(constants.ExpiresAccessToken).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(global.Config.Security.AccessTokenSecret.SecretKey))
	return tokenStr, err
}

func generateRefreshToken(payload models.PayloadToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       payload.ID,
		"email":    payload.Email,
		"userName": payload.UserName,
		"exp":      time.Now().Add(constants.ExpiresRefreshToken).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(global.Config.Security.RefreshTokenSecret.SecretKey))
	return tokenStr, err
}

func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
