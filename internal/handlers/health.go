package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

func Health(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{"status": "ok"})
}
