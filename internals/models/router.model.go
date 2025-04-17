package models

type Routes struct {
	AuthenticationRoutes []Route
	MemoriesRoutes       []Route
	TagsRoutes           []Route
	LocationRoutes       []Route
}

type Route struct {
	Method      string
	Path        string
	Handler     func()
	Middleware  []func()
	Description string
}

var RoutesInstance = Routes{
	AuthenticationRoutes: []Route{
		{Method: "POST", Path: "/auth/register", Handler: nil, Description: "Register a new user"},
		{Method: "POST", Path: "/auth/login", Handler: nil},
		{Method: "POST", Path: "/auth/refresh", Handler: nil},
		{Method: "DELETE", Path: "/auth/logout", Handler: nil},
	},
	MemoriesRoutes: []Route{
		{Method: "GET", Path: "/memories", Handler: nil, Description: "List memories with pagination"},
		{Method: "POST", Path: "/memories", Handler: nil},
		{Method: "GET", Path: "/memories/:id", Handler: nil},
		{Method: "PATCH", Path: "/memories/:id", Handler: nil},
		{Method: "DELETE", Path: "/memories/:id", Handler: nil},
		{Method: "POST", Path: "/memories/:id/photos", Handler: nil},
		{Method: "GET", Path: "/memories/:id/photos", Handler: nil},
		{Method: "DELETE", Path: "/memories/:id/photos/:photoId", Handler: nil},
		{Method: "POST", Path: "/memories/:id/share", Handler: nil},
		{Method: "GET", Path: "/memories/:id/shares", Handler: nil},
		{Method: "DELETE", Path: "/memories/:id/shares/:userId", Handler: nil},
		// {Method: "GET", Path: "/memories/locations/search", Handler: nil, Middleware: []func(){AuthMiddleware}, Description: "Search locations for memory"},
	},
	TagsRoutes: []Route{
		{Method: "GET", Path: "/tags", Handler: nil},
		{Method: "POST", Path: "/tags", Handler: nil},
		{Method: "DELETE", Path: "/tags/:id", Handler: nil},
	},
}
