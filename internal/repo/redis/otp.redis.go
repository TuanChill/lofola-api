package redis

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tuanchill/lofola-api/configs/common/constants"
)

func SaveOtpKey(c *gin.Context, rdb *redis.Client, email string, otp string) error {
	err := rdb.Set(c, fmt.Sprintf("%s_%s", constants.OTPKey, email), otp, constants.ExpiresOTP*time.Second).Err()
	if err != nil {
		return fmt.Errorf("could not save OTP: %s", err)
	}
	return nil
}

func GetOtpKey(c *gin.Context, rdb *redis.Client, email string) (string, error) {
	otp, err := rdb.Get(c, fmt.Sprintf("%s_%s", constants.OTPKey, email)).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", fmt.Errorf("could not retrieve OTP: %s", err)
	}

	return otp, nil
}

func DeleteOtpKey(c *gin.Context, rdb *redis.Client, email string) error {
	err := rdb.Del(c, fmt.Sprintf("%s_%s", constants.OTPKey, email)).Err()
	if err != nil {
		return fmt.Errorf("could not delete OTP: %s", err)
	}
	return nil
}
