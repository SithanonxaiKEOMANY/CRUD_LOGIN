package errs

import "net/http"

type AppError struct {
	Status  int
	Message string
	Code    int
}

func (a AppError) Error() string {
	return a.Message
}

func NewError(code int, errMsg string) error {
	return AppError{
		Status:  code,
		Message: errMsg,
	}
}

func ErrorBadRequest(errorMessage string) error {
	return AppError{
		Status:  http.StatusBadRequest,
		Message: errorMessage,
	}
}
func ErrorUnprocessableEntity(errorMessage string) error {
	return AppError{
		Status:  http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}

func ErrorInternalServerError(errorMessage string) error {
	return AppError{
		Status:  http.StatusInternalServerError,
		Message: errorMessage,
	}
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func NewNotFoundError(message string) error {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "unexpected error",
	}
}

func NewValidationError(message string) error {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
