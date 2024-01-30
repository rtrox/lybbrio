package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func statusCodeResponse(w http.ResponseWriter, r *http.Request, statusCode int) {
	render.Status(r, statusCode)
	render.DefaultResponder(w, r, render.M{"error": http.StatusText(statusCode)})
}
