package response

import (
	"net/http"
	"todo-app/code"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, v)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := code.GetCode(err).ToHTTPStatus()

	msg := err.Error()
	render.Status(r, statusCode)
	render.JSON(w, r, errorResponse{
		Code:    statusCode,
		Message: msg,
	})
}
