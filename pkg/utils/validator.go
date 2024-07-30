package utils

import (
	"fmt"
	"regexp"
	"time"
)

// isBirthDayValid checks if the birthday is in the correct format
func IsBirthDayValid(birthDay string) bool {
	regex := `^\d{4}-\d{2}-\d{2}$`
	match, _ := regexp.MatchString(regex, birthDay)
	if !match {
		return false
	}
	// Parse the date to ensure it is valid
	_, err := time.Parse("2006-01-02", birthDay)
	return err == nil
}

func IsNumberValid(number string) bool {
	regex := `^\d+$`
	match, _ := regexp.MatchString(regex, number)
	fmt.Println(match)
	return match
}
