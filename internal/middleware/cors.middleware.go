package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigins := []string{
			"http://localhost:8080",
			"http://127.0.0.1:8080",
		}

		origin := c.GetHeader("Origin")

		for _, allowedOrigin := range allowedOrigins {
			if allowedOrigin == origin {
				setCorsHeaders(c.Writer, origin)
				break
			}
		}
	}
}

func setCorsHeaders(w http.ResponseWriter, origin string) {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Request-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
