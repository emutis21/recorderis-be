package swagger

import (
	"recorderis/cmd/services/memory/models"
	"recorderis/internals/utils"
)

// Auth endpoints
// --------------

// RefreshToken godoc
// @Summary      Refresh token
// @Description  Get a new access token using a refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        X-Refresh-Token  header    string  true  "Refresh Token"
// @Success      200              {object}  TokenResponse
// @Failure      401              {object}  ErrorResponse
// @Router       /auth/refresh [post]
func RefreshToken() {}

// LoginUser godoc
// @Summary      User login
// @Description  Authenticates a user and returns a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginRequest  true  "Login credentials"
// @Success      200      {object}  TokenResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /auth/login [post]
func LoginUser() {}

// RegisterUser godoc
// @Summary      Register a new user
// @Description  Creates a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterRequest  true  "Registration data"
// @Success      201      {object}  TokenResponse
// @Failure      400      {object}  ErrorResponse
// @Router       /auth/register [post]
func RegisterUser() {}

// LogoutUser godoc
// @Summary      Logout
// @Description  Invalidates user token and ends the session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  SuccessResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/auth/logout [post]
func LogoutUser() {}

// User endpoints
// -------------

// GetUsers godoc
// @Summary      List users
// @Description  Gets a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}   UserResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/users [get]
func GetUsers() {}

// GetUserById godoc
// @Summary      Get user by ID
// @Description  Returns information for a specific user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /users/{id} [get]
func GetUserById() {}

// GetUserProfile godoc
// @Summary      Get current user profile
// @Description  Returns information for the authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  UserResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /secure/users/me [get]
func GetUserProfile() {}

// CreateUser godoc
// @Summary      Create user
// @Description  Creates a new user in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateUserRequest  true  "User data"
// @Success      201      {object}  UserResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/users [post]
func CreateUser() {}

// UpdateUser godoc
// @Summary      Update user
// @Description  Updates information for an existing user
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int               true  "User ID"
// @Param        request  body      UpdateUserRequest  true  "Data to update"
// @Success      200      {object}  UserResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /secure/users/{id} [put]
func UpdateUser() {}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Permanently removes a user from the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "User ID"
// @Success      204
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/users/{id} [delete]
func DeleteUser() {}

// Memory endpoints
// --------------

// GetMemories godoc
// @Summary      List memories
// @Description  Gets all memories for the authenticated user
// @Tags         memories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}   MemoryResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/memories [get]
func GetMemories() {}

// GetMemoryById godoc
// @Summary      Get memory by ID
// @Description  Returns information for a specific memory
// @Tags         memories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Memory ID"
// @Success      200  {object}  MemoryResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/memories/{id} [get]
func GetMemoryById() {}

// CreateMemory godoc
// @Summary      Create memory
// @Description  Creates a new memory for the authenticated user
// @Tags         memories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateMemoryRequest  true  "Memory data"
// @Success      201      {object}  MemoryResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/memories [post]
func CreateMemory() {}

// UpdateMemory godoc
// @Summary      Update memory
// @Description  Updates information for an existing memory
// @Tags         memories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string               true  "Memory ID"
// @Param        request  body      UpdateMemoryRequest  true  "Data to update"
// @Success      200      {object}  MemoryResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /secure/memories/{id} [put]
func UpdateMemory() {}

// DeleteMemory godoc
// @Summary      Delete memory
// @Description  Permanently removes a memory
// @Tags         memories
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Memory ID"
// @Success      204
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/memories/{id} [delete]
func DeleteMemory() {}

type RegisterRequest struct {
	Username    string `json:"username" example:"johndoe"`
	DisplayName string `json:"display_name" example:"John Doe"`
	Email       string `json:"email" example:"john@example.com"`
	Password    string `json:"password" example:"securepassword"`
	DeviceType  string `json:"device_type" example:"web"`
}

type LoginRequest struct {
	Email      string `json:"email" example:"john@example.com"`
	Password   string `json:"password" example:"securepassword"`
	RememberMe bool   `json:"remember_me" example:"true"`
	DeviceType string `json:"device_type" example:"web"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	TokenType    string `json:"token_type" example:"Bearer"`
	ExpiresIn    int64  `json:"expires_in" example:"900"`
	RefreshToken string `json:"refresh_token,omitempty" example:"abc123def456"`
}

type UserResponse struct {
	ID          string `json:"id" example:"123"`
	Username    string `json:"username" example:"johndoe"`
	DisplayName string `json:"display_name" example:"John Doe"`
	Email       string `json:"email" example:"john@example.com"`
	Role        string `json:"role" example:"user"`
}

type CreateUserRequest struct {
	Username    string `json:"username" example:"newuser"`
	DisplayName string `json:"display_name" example:"New User"`
	Email       string `json:"email" example:"new@example.com"`
	Password    string `json:"password" example:"securepassword"`
	Role        string `json:"role" example:"user"`
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name,omitempty" example:"Updated Name"`
	Email       string `json:"email,omitempty" example:"updated@example.com"`
	Password    string `json:"password,omitempty" example:"newsecurepassword"`
}

type SuccessResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operación completada con éxito"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Error message"`
	Error   struct {
		Type    string `json:"type" example:"VALIDATION"`
		Message string `json:"message" example:"Invalid input"`
		Detail  string `json:"detail" example:"Email is required"`
	} `json:"error"`
}

type MemoryResponse struct {
	ID        string `json:"id" example:"abc123def456"`
	Title     string `json:"title" example:"Summer Vacation 2024"`
	Date      string `json:"date" example:"2024-07-15T00:00:00Z"`
	IsPublic  bool   `json:"is_public" example:"false"`
	CreatedAt string `json:"created_at" example:"2024-04-01T10:30:00Z"`
	UpdatedAt string `json:"updated_at" example:"2024-04-02T15:45:00Z"`
}

type CreateMemoryRequest struct {
	Title        string                            `json:"title" binding:"required"`
	Date         utils.JSONTime                    `json:"date" binding:"required"`
	IsPublic     bool                              `json:"is_public"`
	Descriptions []models.CreateDescriptionRequest `json:"descriptions,omitempty"`
}

type UpdateMemoryRequest struct {
	Title    string `json:"title,omitempty" example:"Updated Vacation Title"`
	Date     string `json:"date,omitempty" example:"2024-07-20"`
	IsPublic bool   `json:"is_public,omitempty" example:"true"`
}
