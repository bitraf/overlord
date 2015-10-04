package model

import "net/http"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NotFoundError(e string) *Error {
	return &Error{http.StatusNotFound, e}
}

func BadRequestError(e string) *Error {
	return &Error{http.StatusBadRequest, e}
}

func InternalServerError(e string) *Error {
	return &Error{http.StatusInternalServerError, e}
}
