package swagger

// These are dummy functions that exist only for swagger annotations
// They won't be called in actual code

// RegisterUser godoc
// @Summary      Registrar un nuevo usuario
// @Description  Crea una nueva cuenta de usuario
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      RegisterRequest  true  "Datos de registro"
// @Success      201      {object}  TokenResponse
// @Failure      400      {object}  ErrorResponse
// @Router       /api/v1/auth/register [post]
func RegisterUser() {}

// LoginUser godoc
// @Summary      Iniciar sesión de usuario
// @Description  Autentica un usuario y devuelve un token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      LoginRequest  true  "Credenciales de acceso"
// @Success      200      {object}  TokenResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /api/v1/auth/login [post]
func LoginUser() {}

// RefreshToken godoc
// @Summary      Refrescar token
// @Description  Obtiene un nuevo access token usando el refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        X-Refresh-Token  header    string  true  "Refresh Token"
// @Success      200              {object}  TokenResponse
// @Failure      401              {object}  ErrorResponse
// @Router       /api/v1/auth/refresh [post]
func RefreshToken() {}

// LogoutUser godoc
// @Summary      Cerrar sesión
// @Description  Invalida el token del usuario y cierra la sesión
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {object}  SuccessResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /api/v1/secure/auth/logout [post]
func LogoutUser() {}

// GetUsers godoc
// @Summary      Listar usuarios
// @Description  Obtiene una lista de todos los usuarios
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200      {array}   UserResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /api/v1/secure/users [get]
func GetUsers() {}

// GetUserById godoc
// @Summary      Obtener un usuario por ID
// @Description  Devuelve la información de un usuario específico
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID del usuario"
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/users/{id} [get]
func GetUserById() {}

// CreateUser godoc
// @Summary      Crear usuario
// @Description  Crea un nuevo usuario en el sistema
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      CreateUserRequest  true  "Datos del usuario"
// @Success      201      {object}  UserResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Router       /api/v1/secure/users [post]
func CreateUser() {}

// UpdateUser godoc
// @Summary      Actualizar usuario
// @Description  Actualiza la información de un usuario existente
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      int               true  "ID del usuario"
// @Param        request  body      UpdateUserRequest  true  "Datos a actualizar"
// @Success      200      {object}  UserResponse
// @Failure      400      {object}  ErrorResponse
// @Failure      401      {object}  ErrorResponse
// @Failure      404      {object}  ErrorResponse
// @Router       /api/v1/secure/users/{id} [put]
func UpdateUser() {}

// DeleteUser godoc
// @Summary      Eliminar usuario
// @Description  Elimina permanentemente un usuario del sistema
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "ID del usuario"
// @Success      204
// @Failure      401  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Router       /api/v1/secure/users/{id} [delete]
func DeleteUser() {}

// GetUserProfile godoc
// @Summary      Obtener perfil del usuario actual
// @Description  Devuelve la información del usuario autenticado
// @Tags         users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  UserResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /api/v1/secure/users/me [get]
func GetUserProfile() {}

// Define model structures solely for Swagger documentation
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
