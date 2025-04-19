package swagger

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
