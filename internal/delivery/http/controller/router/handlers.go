package router

import (
	"github.com/insan1a/tech-tinker/internal/lib/response"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.NotFound(w)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	response.MethodNotAllowed(w)
}
