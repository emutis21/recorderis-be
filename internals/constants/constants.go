package constants

// API prefixes and versioning
const (
	APIPrefix = "/api"
	V1        = "/v1"
	APIPathV1 = APIPrefix + V1
)

// Resource paths
const (
	AuthPath   = "/auth"
	UsersPath  = "/users"
	SecurePath = "/secure"
)

// Route params
const (
	IDParam = "/:id"
	MePath  = "/me"
)

// Auth endpoints
const (
	RegisterPath = "/register"
	LoginPath    = "/login"
	RefreshPath  = "/refresh"
	LogoutPath   = "/logout"
)
