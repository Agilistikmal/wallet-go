package handler

import (
	"github.com/agilistikmal/wallet-go/helper"
	"github.com/agilistikmal/wallet-go/model"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "Application/Json")
	exception, ok := err.(NotFoundError)
	if ok {
		w.WriteHeader(http.StatusNotFound)
		response := model.WebResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Data:   exception.Error,
		}
		helper.WriteToResponse(w, response)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		response := model.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err,
		}
		helper.WriteToResponse(w, response)
	}
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{
		Error: error,
	}
}
