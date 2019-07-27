package core

import (
	"encoding/json"
	"net/http"

	"github.com/juju/errors"
)

type response struct {
	Status  int
	Message string
}

// CheckForRequestBody checkss the request r for a
//   request body, returns an error if the request body is nil
func CheckForRequestBody(r *http.Request) (err error) {
	if r.Body == nil {
		err = errors.NotValidf("request body empty")
	}
	return err
}

func WriteJSONResponse(w http.ResponseWriter, s interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func WriteStatusOKResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusOK)
}

func WriteStatusOKResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusOK, message)
}

func WriteInternalServerErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusInternalServerError)
}

func WriteInternalServerErrorResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusInternalServerError, message)
}

func WriteBadRequestErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusBadRequest)
}

func WriteBadRequestErrorResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusBadRequest, message)
}

func WriteNotFoundErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusNotFound)
}

func WriteNotFoundErrorResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusNotFound, message)
}

func WriteForbiddenErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusForbidden)
}

func WriteForbiddenErrorResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusForbidden, message)
}

func WriteUnauthorizedErrorResponse(w http.ResponseWriter) {
	writeResponse(w, http.StatusUnauthorized)
}

func WriteUnauthorizedErrorResponseWithMessage(w http.ResponseWriter, message string) {
	writeResponseWithMessage(w, http.StatusUnauthorized, message)
}

func writeResponse(w http.ResponseWriter, errorCode int) {
	writeResponseWithMessage(w, errorCode, http.StatusText(errorCode))
}

func writeResponseWithMessage(w http.ResponseWriter, errorCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(
		response{
			errorCode,
			message,
		})
}
