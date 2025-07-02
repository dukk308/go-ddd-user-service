package common

import "net/http"

type AppError struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode,omitempty"`
	Message    string `json:"message,omitempty"`
	Err        error  `json:"error,omitempty"`
	TraceID    string `json:"traceId,omitempty"`
}

func NewAppError(err error, code string, message string, statusCode int, traceID string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		ErrorCode:  code,
		Message:    message,
		Err:        err,
		TraceID:    traceID,
	}
}

func (e *AppError) Peel() error {
	if err, ok := e.Err.(*AppError); ok {
		return err.Peel()
	}

	return e
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Is(err error) bool {
	return e.Err == err
}

func DbUnexpectedError(err error, errorCode string, message string) *AppError {
	return NewAppError(err, errorCode, message, http.StatusInternalServerError, "")
}

func DbRecordNotFoundError(err error) *AppError {
	return NewAppError(err, "NOT_FOUND", err.Error(), http.StatusNotFound, "")
}
