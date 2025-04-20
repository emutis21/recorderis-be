package main

import (
	"recorderis/cmd/middleware"
	"recorderis/cmd/routes"
	api_ports "recorderis/cmd/services/api/ports/drivers"
	auth_drivens "recorderis/cmd/services/auth/ports/drivens"
	auth_ports "recorderis/cmd/services/auth/ports/drivers"
	memory_ports "recorderis/cmd/services/memory/ports/drivers"
	"recorderis/pkg/swagger"

	_ "recorderis/docs"

	"github.com/gin-gonic/gin"
)

func CreateRouter(
	userAdapter api_ports.ForUser,
	authAdapter auth_ports.ForAuth,
	memoryAdapter memory_ports.ForMemory,
	tokenManager auth_drivens.ForTokenManager,
) *gin.Engine {
	router := gin.Default()

	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	routes.SetupAuthRoutes(router, authAdapter, authMiddleware)
	routes.SetupUserRoutes(router, userAdapter, authAdapter, authMiddleware)
	routes.SetupMemoryRoutes(router, memoryAdapter, authMiddleware)

	swaggerRoutes := swagger.NewRoutes()
	swaggerRoutes.Register(router)

	return router
}
