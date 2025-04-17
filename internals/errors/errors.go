package errors

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrorType string

const (
	ErrNotFound     ErrorType = "NOT_FOUND"
	ErrValidation   ErrorType = "VALIDATION"
	ErrDatabase     ErrorType = "DATABASE"
	ErrUnauthorized ErrorType = "UNAUTHORIZED"
	ErrForbidden    ErrorType = "FORBIDDEN"
	ErrBadRequest   ErrorType = "BAD_REQUEST"
	ErrInternal     ErrorType = "INTERNAL"
	ErrConflict     ErrorType = "CONFLICT"
)

type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Detail  string    `json:"detail,omitempty"`
	Code    int       `json:"-"`
	Err     error     `json:"-"`
}

var errorStatusMap = map[ErrorType]int{
	ErrValidation:   400,
	ErrBadRequest:   400,
	ErrUnauthorized: 401,
	ErrForbidden:    403,
	ErrNotFound:     404,
	ErrConflict:     409,
	ErrDatabase:     500,
	ErrInternal:     500,
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Type, e.Message, e.Err)
	}

	return fmt.Sprintf("[%s] %s", e.Type, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrNotFound,
		Message: message,
		Code:    404,
		Err:     err,
	}
}

func NewValidationError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrValidation,
		Message: message,
		Code:    400,
		Err:     err,
	}
}

func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrUnauthorized,
		Message: message,
		Code:    401,
		Err:     err,
	}
}

func NewForbiddenError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrForbidden,
		Message: message,
		Code:    403,
		Err:     err,
	}
}

func NewBadRequestError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrBadRequest,
		Message: message,
		Code:    400,
		Err:     err,
	}
}

func NewConflictError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrConflict,
		Message: message,
		Code:    409,
		Err:     err,
	}
}

func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Type:    ErrInternal,
		Message: message,
		Code:    500,
		Err:     err,
	}
}

func AsAppError(err error) *AppError {
	appErr, ok := err.(*AppError)
	if ok {
		return appErr
	}

	return NewInternalError("Unexpected error", err)
}

var errorStatusCodes = map[ErrorType]int{
	ErrNotFound:     404,
	ErrValidation:   400,
	ErrDatabase:     500,
	ErrUnauthorized: 401,
	ErrBadRequest:   400,
}

func NewError(errType ErrorType, message string, err error) *AppError {
	return &AppError{
		Type:    errType,
		Message: message,
		Code:    errorStatusMap[errType],
		Err:     err,
	}
}

func NewErrorResponse(err error, message string) map[string]interface{} {
	appErr, ok := err.(*AppError)
	if !ok {
		return map[string]interface{}{
			"success": false,
			"message": message,
			"error":   "INTERNAL_ERROR",
		}
	}

	return map[string]interface{}{
		"success": false,
		"message": appErr.Message,
		"error":   string(appErr.Type),
	}
}

func HandleError(c *gin.Context, err error) {
	appErr := AsAppError(err)

	c.JSON(appErr.Code, gin.H{
		"success": false,
		"message": appErr.Message,
		"error": gin.H{
			"type":   string(appErr.Type),
			"detail": appErr.Detail,
		},
	})
}
