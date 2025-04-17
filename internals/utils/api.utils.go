package utils

import (
	"recorderis/internals/errors"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

type ErrorInfo struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

const (
	// Éxito
	MsgCreated   = "Resource created successfully"
	MsgRetrieved = "Data retrieved successfully"
	MsgUpdated   = "Update successful"
	MsgDeleted   = "Delete successful"

	// Auth
	MsgRegistered = "User registered successfully"
	MsgLoggedIn   = "Login successful"
	MsgLoggedOut  = "Logout successful"

	// Errores
	MsgInvalidInput = "Invalid input"
	MsgInvalidID    = "Invalid ID"
)

type ResponseHandler struct {
	C *gin.Context
}

func NewHandler(c *gin.Context) *ResponseHandler {
	return &ResponseHandler{C: c}
}

func (h *ResponseHandler) Success(code int, data interface{}, message string) {
	h.C.JSON(code, ApiResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

func (h *ResponseHandler) Error(err error) {
	appErr, ok := err.(*errors.AppError)
	if !ok {
		h.C.JSON(500, ApiResponse{
			Success: false,
			Message: "Error interno del servidor",
			Error: &ErrorInfo{
				Type:    "INTERNAL_ERROR",
				Message: err.Error(),
			},
		})
		return
	}

	h.C.JSON(appErr.Code, ApiResponse{
		Success: false,
		Message: appErr.Message,
		Error: &ErrorInfo{
			Type:    string(appErr.Type),
			Message: appErr.Message,
			Detail:  appErr.Detail,
		},
	})
}

func (h *ResponseHandler) OK(data interface{}, message string) {
	h.Success(200, data, message)
}

func (h *ResponseHandler) Created(data interface{}, message string) {
	h.Success(201, data, message)
}

func (h *ResponseHandler) Accepted(data interface{}, message string) {
	h.Success(202, data, message)
}

func (h *ResponseHandler) NoContent() {
	h.C.Status(204)
}
