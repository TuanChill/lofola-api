package helpers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tuanchill/lofola-api/configs/common/constants"
	"github.com/tuanchill/lofola-api/global"
	"github.com/tuanchill/lofola-api/internal/models"
)

func JwtProvider(payload models.PayloadToken) (string, string, error) {
	accessToken, err := GenerateAccessToken(payload)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(payload)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func GenerateAccessToken(payload models.PayloadToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       strconv.Itoa(payload.ID),
		"email":    payload.Email,
		"userName": payload.UserName,
		"exp":      time.Now().Add(time.Duration(constants.ExpiresAccessToken) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenStr, err := token.SignedString([]byte(global.Config.Security.AccessTokenSecret.SecretKey))
	return tokenStr, err
}

func GenerateRefreshToken(payload models.PayloadToken) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       strconv.Itoa(payload.ID),
		"email":    payload.Email,
		"userName": payload.UserName,
		"exp":      time.Now().Add(time.Duration(constants.ExpiresRefreshToken) * time.Second).Unix(), // Expiration time
		"iat":      time.Now().Unix(),                                                                 // Issued at
	})

	tokenStr, err := token.SignedString([]byte(global.Config.Security.RefreshTokenSecret.SecretKey))
	return tokenStr, err
}

func VerifyToken(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// GetTokenFromHeader return access Token and Refresh Token
func GetTokenFromHeader(c *gin.Context) (string, string) {
	accToken := c.GetHeader("Authorization")
	if accToken != "" {
		accToken = accToken[len("Bearer "):]
	}

	refToken := c.GetHeader("RefreshToken")

	return accToken, refToken
}

// GetPayload return payload from token
func GetPayload(c *gin.Context) models.PayloadToken {
	tokenStr, _ := GetTokenFromHeader(c)

	tkn, _ := VerifyToken(tokenStr, global.Config.Security.AccessTokenSecret.SecretKey)
	claims, _ := tkn.Claims.(jwt.MapClaims)

	idStr, _ := claims["id"].(string)
	id, _ := strconv.Atoi(idStr) // Convert the ID from string to int

	payload := models.PayloadToken{
		ID:       id,
		Email:    claims["email"].(string),
		UserName: claims["userName"].(string),
	}

	return payload
}

func ExtractToken(token *jwt.Token) models.PayloadToken {
	claims, _ := token.Claims.(jwt.MapClaims)

	idStr, _ := claims["id"].(string)
	id, _ := strconv.Atoi(idStr) // Convert the ID from string to int

	payload := models.PayloadToken{
		ID:       id,
		Email:    claims["email"].(string),
		UserName: claims["userName"].(string),
	}

	return payload
}

func FormatBearToken(token string) string {
	return fmt.Sprintf("Bearer %s", token)
}
