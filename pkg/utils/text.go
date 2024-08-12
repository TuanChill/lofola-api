package utils

import (
	"fmt"
	"strconv"
)

func FormatKeyRedis(constant string, value string) string {
	return fmt.Sprintf("%s_%s", constant, value)
}

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
