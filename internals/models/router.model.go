package models

import (
	"recorderis/internals/constants"
)

type Routes struct {
	AuthenticationRoutes []Route
	UserRoutes           []Route
	MemoriesRoutes       []Route
	TagsRoutes           []Route
	LocationRoutes       []Route
}

type Route struct {
	Method       string
	Path         string
	RequiresAuth bool
	Handler      func()
	Middleware   []func()
	Description  string
}

var RoutesInstance = Routes{
	AuthenticationRoutes: []Route{
		{Method: "POST", Path: constants.APIPathV1 + constants.AuthPath + constants.RegisterPath, RequiresAuth: false, Description: "Register a new user"},
		{Method: "POST", Path: constants.APIPathV1 + constants.AuthPath + constants.LoginPath, RequiresAuth: false, Description: "Login user"},
		{Method: "POST", Path: constants.APIPathV1 + constants.AuthPath + constants.RefreshPath, RequiresAuth: false, Description: "Refresh token"},
		{Method: "POST", Path: constants.APIPathV1 + constants.SecurePath + constants.AuthPath + constants.LogoutPath, RequiresAuth: true, Description: "Logout user"},
	},
	UserRoutes: []Route{
		{Method: "GET", Path: constants.APIPathV1 + constants.UsersPath + "/:id", RequiresAuth: false, Description: "Get user by ID"},
		{Method: "GET", Path: constants.APIPathV1 + constants.SecurePath + constants.UsersPath, RequiresAuth: true, Description: "List all users"},
		{Method: "POST", Path: constants.APIPathV1 + constants.SecurePath + constants.UsersPath, RequiresAuth: true, Description: "Create user"},
		{Method: "PUT", Path: constants.APIPathV1 + constants.SecurePath + constants.UsersPath + "/:id", RequiresAuth: true, Description: "Update user"},
		{Method: "DELETE", Path: constants.APIPathV1 + constants.SecurePath + constants.UsersPath + "/:id", RequiresAuth: true, Description: "Delete user"},
		{Method: "GET", Path: constants.APIPathV1 + constants.SecurePath + constants.UsersPath + constants.MePath, RequiresAuth: true, Description: "Get current user profile"},
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
