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

// Description endpoints
// --------------

// GetDescriptions godoc
// @Summary      List descriptions
// @Description  Gets all descriptions for a specific memory
// @Tags         descriptions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Memory ID"
// @Success      200  {array}   DescriptionResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/memories/{id}/descriptions [get]
func GetDescriptions() {}

// GetDescriptionByID godoc
// @Summary      Get description by ID
// @Description  Returns information for a specific description
// @Tags         descriptions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      string  true  "Memory ID"
// @Param        description_id path      string  true  "Description ID"
// @Success      200            {object}  DescriptionResponse
// @Failure      401            {object}  ErrorResponse
// @Failure      404            {object}  ErrorResponse
// @Router       /secure/memories/{id}/descriptions/{description_id} [get]
func GetDescriptionByID() {}

// CreateDescription godoc
// @Summary      Create description
// @Description  Creates a new description for a specific memory
// @Tags         descriptions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                   true  "Memory ID"
// @Param        request  body      CreateDescriptionRequest  true  "Description data"
// @Success      201      {object}  DescriptionResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /secure/memories/{id}/descriptions [post]
func CreateDescription() {}

// UpdateDescription godoc
// @Summary      Update description
// @Description  Updates information for an existing description
// @Tags         descriptions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      string                   true  "Memory ID"
// @Param        description_id path      string                   true  "Description ID"
// @Param        request        body      UpdateDescriptionRequest  true  "Data to update"
// @Success      200            {object}  DescriptionResponse
// @Failure      400            {object}  ErrorResponse
// @Failure      401            {object}  ErrorResponse
// @Failure      404            {object}  ErrorResponse
// @Router       /secure/memories/{id}/descriptions/{description_id} [put]
func UpdateDescription() {}

// DeleteDescription godoc
// @Summary      Delete description
// @Description  Permanently removes a description
// @Tags         descriptions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id             path      string  true  "Memory ID"
// @Param        description_id path      string  true  "Description ID"
// @Success      204
// @Failure      401            {object}  ErrorResponse
// @Failure      404            {object}  ErrorResponse
// @Router       /secure/memories/{id}/descriptions/{description_id} [delete]
func DeleteDescription() {}

// Location endpoints
// --------------

// GetLocations godoc
// @Summary      List locations
// @Description  Gets all locations
// @Tags         locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}   LocationResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/locations [get]
func GetLocations() {}

// GetLocationById godoc
// @Summary      Get location by ID
// @Description  Returns information for a specific location
// @Tags         locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Location ID"
// @Success      200  {object}  LocationResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/locations/{id} [get]
func GetLocationById() {}

// CreateLocation godoc
// @Summary      Create location
// @Description  Creates a new location
// @Tags         locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateLocationRequest  true  "Location data"
// @Success      201      {object}  LocationResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /secure/locations [post]
func CreateLocation() {}

// UpdateLocation godoc
// @Summary      Update location
// @Description  Updates information for an existing location
// @Tags         locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string                 true  "Location ID"
// @Param        request  body      UpdateLocationRequest  true  "Data to update"
// @Success      200      {object}  LocationResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /secure/locations/{id} [put]
func UpdateLocation() {}

// DeleteLocation godoc
// @Summary      Delete location
// @Description  Permanently removes a location
// @Tags         locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Location ID"
// @Success      204
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/locations/{id} [delete]
func DeleteLocation() {}

// Memory-Location relationship endpoints
// ---------------------------------------

// GetLocationsByMemoryID godoc
// @Summary      List locations for a memory
// @Description  Gets all locations associated with a specific memory
// @Tags         memory-locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      string  true  "Memory ID"
// @Success      200  {array}   LocationResponse
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /secure/memories/{id}/locations [get]
func GetLocationsByMemoryID() {}

// AssociateMemoryWithLocation godoc
// @Summary      Associate memory with location
// @Description  Associates a memory with an existing location
// @Tags         memory-locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path      string  true  "Memory ID"
// @Param        location_id path      string  true  "Location ID"
// @Success      201         "Location associated with memory"
// @Failure      401         {object}  ErrorResponse
// @Failure      404         {object}  ErrorResponse
// @Router       /secure/memories/{id}/locations/{location_id} [post]
func AssociateMemoryWithLocation() {}

// DisassociateMemoryFromLocation godoc
// @Summary      Remove location from memory
// @Description  Removes the association between a memory and a location
// @Tags         memory-locations
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id          path      string  true  "Memory ID"
// @Param        location_id path      string  true  "Location ID"
// @Success      204
// @Failure      401         {object}  ErrorResponse
// @Failure      404         {object}  ErrorResponse
// @Router       /secure/memories/{id}/locations/{location_id} [delete]
func DisassociateMemoryFromLocation() {}

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

type DescriptionResponse struct {
	ID            string `json:"id" example:"123"`
	DescriptionID string `json:"description_id" example:"abc-xyz-123"`
	Text          string `json:"text" example:"This was our first day at the beach, we had a great time."`
	Index         int    `json:"index" example:"0"`
	Version       int    `json:"version" example:"1"`
}

type CreateDescriptionRequest struct {
	Text  string `json:"text" binding:"required" example:"First day at the beach"`
	Index int    `json:"index" binding:"required" example:"0"`
}

type UpdateDescriptionRequest struct {
	Text  string `json:"text,omitempty" example:"Updated description text"`
	Index *int   `json:"index,omitempty" example:"1"`
}

type LocationResponse struct {
	ID         string  `json:"id" example:"abc123def456"`
	LocationID string  `json:"location_id" example:"xyz789"`
	Location   string  `json:"location" example:"Playa del Carmen"`
	Longitude  float64 `json:"longitude" example:"-87.0739"`
	Latitude   float64 `json:"latitude" example:"20.6296"`
	City       string  `json:"city" example:"Playa del Carmen"`
	Country    string  `json:"country" example:"Mexico"`
}

type CreateLocationRequest struct {
	Location  string  `json:"location" binding:"required" example:"Playa del Carmen"`
	Longitude float64 `json:"longitude" binding:"required" example:"-87.0739"`
	Latitude  float64 `json:"latitude" binding:"required" example:"20.6296"`
	City      string  `json:"city" binding:"required" example:"Playa del Carmen"`
	Country   string  `json:"country" binding:"required" example:"Mexico"`
}

type UpdateLocationRequest struct {
	Location  string   `json:"location,omitempty" example:"Updated Beach Name"`
	Longitude *float64 `json:"longitude,omitempty" example:"-87.0740"`
	Latitude  *float64 `json:"latitude,omitempty" example:"20.6297"`
	City      string   `json:"city,omitempty" example:"Playa del Carmen"`
	Country   string   `json:"country,omitempty" example:"Mexico"`
}
