package main

import (
	"net/http"

	"github.com/tuanchill/lofola-api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	http.ListenAndServe(":8080", r)
}
