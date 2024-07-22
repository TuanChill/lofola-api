package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
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

func GetTokenFromHeader(c *gin.Context) (string, string) {
	accToken := c.GetHeader("Authorization")
	tokenStr := accToken[len("Bearer "):]

	refToken := c.GetHeader("RefreshToken")

	return tokenStr, refToken
}

func GetPayload(c *gin.Context) models.PayloadToken {
	tokenStr, _ := GetTokenFromHeader(c)

	tkn, _ := VerifyToken(tokenStr, global.Config.Security.AccessTokenSecret.SecretKey)
	claims, _ := tkn.Claims.(jwt.MapClaims)

	payload := models.PayloadToken{
		ID:       claims["id"].(int),
		Email:    claims["email"].(string),
		UserName: claims["userName"].(string),
	}

	return payload
}
