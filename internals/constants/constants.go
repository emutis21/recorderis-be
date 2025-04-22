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

// Memory paths
const (
	MemoriesPath       = "/memories"
	MemoryPath         = "/memory"
	MemoryIDPath       = "/:memory_id"
	DescriptionsPath   = "/descriptions"
	DescriptionIDParam = "/:description_id"
)

// Location paths
const (
	LocationsPath       = "/locations"
	LocationPath        = "/location"
	LocationIDPath      = "/:location_id"
	LocationsByMemoryID = "/:id/locations"
	LocationIDParam     = "/:location_id"
)

// Tag paths
const (
	TagsPath     = "/tags"
	TagPath      = "/tag"
	TagIDPath    = "/:tag_id"
	TagsByMemoryID = "/:id/tags"
	TagIDParam   = "/:tag_id"
)

type ContextKey string

// Context keys
const (
	UserIDKey ContextKey = "userID"
)
