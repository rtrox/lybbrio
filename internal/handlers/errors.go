package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type AppCode int

const (
	ErrRender AppCode = iota
	ErrPaginationToken
)

type AppError struct {
	Code    AppCode `json:"code"`
	Message string  `json:"message"`
}

func (e AppError) Error() string {
	return e.Message
}

type ErrResponse struct {
	Err error `json:"-"`

	HTTPStatusCode int    `json:"code"`
	StatusText     string `json:"status"`

	AppCode   int    `json:"app_code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad Request",
		ErrorText:      err.Error(),
	}
}

func ErrInternalError(err AppError) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
		AppCode:        int(err.Code),
		ErrorText:      err.Error(),
	}
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad Request",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
