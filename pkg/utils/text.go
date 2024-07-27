package utils

import "fmt"

func FormatKeyRedis(constant string, value string) string {
	return fmt.Sprintf("%s_%s", constant, value)
}
