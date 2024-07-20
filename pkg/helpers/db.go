package helpers

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GetTimeNow() time.Time {
	return time.Now()
}

func GenerateOTP(length int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	digits := "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = digits[rand.Intn(len(digits))]
	}
	return string(otp)
}

func FormatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	minutes := int(duration.Minutes())
	return fmt.Sprintf("%d minutes", minutes)
}
