package error

import "net/http"

type AppError struct {
	httpCode int
	Code     string
	Message  string
}

var (
	ErrInvalidCredentials = AppError{
		httpCode: http.StatusUnauthorized,
		Code:     "4010",
		Message:  "Username or password is incorrect",
	}

	ErrUserNotFound = AppError{
		httpCode: http.StatusUnauthorized,
		Code:     "4011",
		Message:  "User not found",
	}

	ErrValidationFailed = AppError{
		httpCode: http.StatusBadRequest,
		Code:     "4001",
		Message:  "Validation error occurred",
	}

	ErrInternalServer = AppError{
		httpCode: http.StatusInternalServerError,
		Code:     "5001",
		Message:  "Something went wrong",
	}

	ErrDB = AppError{
		httpCode: http.StatusInternalServerError,
		Code:     "5002",
		Message:  "Something went wrong",
	}
)
