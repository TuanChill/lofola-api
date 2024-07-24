package helpers

import (
	"net/http"
	"strings"
)

// check identifier is email or username
func CheckIdentifier(identifier string) string {
	if strings.Contains(identifier, "@") {
		return "email"
	}
	return "username"
}

// set header and value of this for response
func SetHeaderResponse(w http.ResponseWriter, header string, value string) {
	w.Header().Set(header, value)
}
