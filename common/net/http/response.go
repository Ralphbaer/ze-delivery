package http

import (
	"encoding/json"
	"net/http"
)

// Unauthorized respond with HTTP 401 Unauthorized and payload.
func Unauthorized(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusUnauthorized, &ResponseError{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

// Forbidden respond with HTTP 403 Forbidden.
func Forbidden(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusForbidden, &ResponseError{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

// BadRequest respond with HTTP 400 BadRequest and payload.
func BadRequest(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusBadRequest, s)
}

// Created respond with HTTP 201 StatusOK and payload.
func Created(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusCreated, s)
}

// OK respond with HTTP 200 StatusOK and payload.
func OK(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusOK, s)
}

// NoContent respond with success HTTP 204 StatusNoContent.
func NoContent(w http.ResponseWriter) {
	JSONResponse(w, http.StatusNoContent, nil)
}

// Accepted respond with HTTP 202 StatusAccepted and payload.
func Accepted(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusAccepted, s)
}


// PartialContent respond with HTTP 206 PartialContent and payload.
func PartialContent(w http.ResponseWriter, s interface{}) {
	JSONResponse(w, http.StatusPartialContent, s)
}

// RangeNotSatisfiable respond with HTTP 416 RangeNotSatisfiable.
func RangeNotSatisfiable(w http.ResponseWriter) {
	JSONResponse(w, http.StatusRequestedRangeNotSatisfiable, nil)
}

// NotFound respond with HTTP 404 NotFound and payload.
func NotFound(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusNotFound, &ResponseError{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

// Conflict respond with HTTP 409 Conflict.
func Conflict(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusConflict, &ResponseError{
		Code:    http.StatusConflict,
		Message: message,
	})
}

// NotImplemented respond with HTTP 501 Conflict.
func NotImplemented(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusNotImplemented, &ResponseError{
		Code:    http.StatusNotImplemented,
		Message: message,
	})
}

// UnprocessableEntity respond with HTTP 422 UnprocessableEntity.
func UnprocessableEntity(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusUnprocessableEntity, &ResponseError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	})
}

// InternalServerError respond with HTTP 500 InternalServerError and message.
func InternalServerError(w http.ResponseWriter, message string) {
	JSONResponse(w, http.StatusInternalServerError, &ResponseError{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

// JSONResponseError respond with a ResponseError
func JSONResponseError(w http.ResponseWriter, err ResponseError) {
	JSONResponse(w, err.Code, err)
}

// JSONResponse respond with given HTTP status and given payload.
func JSONResponse(w http.ResponseWriter, status int, s interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w)
	if s != nil {
		payload, _ := json.Marshal(s)
		w.Write(payload)
	}
}
