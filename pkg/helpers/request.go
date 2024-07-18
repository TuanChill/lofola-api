package helpers

import "strings"

func CheckIdentifier(identifier string) string {
	if strings.Contains(identifier, "@") {
		return "email"
	}
	return "username"
}
