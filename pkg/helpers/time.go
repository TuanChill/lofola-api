package helpers

import "time"

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func GetCurrentYear() string {
	return time.Now().Format("2006")
}

func GetCurrentMonth() string {
	return time.Now().Format("01")
}

func GetCurrentDay() string {
	return time.Now().Format("02")
}

func GetCurrentHour() string {
	return time.Now().Format("15")
}
