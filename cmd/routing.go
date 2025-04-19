package main

import (
	"recorderis/cmd/middleware"
	"recorderis/cmd/services/api/models"
	ports "recorderis/cmd/services/api/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/pkg/swagger"

	auth_models "recorderis/cmd/services/auth/models"
	auth_drivens "recorderis/cmd/services/auth/ports/drivens"
	auth_ports "recorderis/cmd/services/auth/ports/drivers"

	"recorderis/internals/errors"
	"recorderis/internals/utils"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "recorderis/docs"
)

func CreateRouter(userAdapter ports.ForUser, authAdapter auth_ports.ForAuth, tokenManager auth_drivens.ForTokenManager) *gin.Engine {
	router := gin.Default()

	authMiddleware := middleware.NewAuthMiddleware(tokenManager)

	apiV1 := router.Group(constants.APIPathV1)

	authRoutes := apiV1.Group(constants.AuthPath)
	authRoutes.Use(authMiddleware.EnrichRequest())
	{
		authRoutes.POST(constants.RegisterPath, func(c *gin.Context) {
			h := utils.NewHandler(c)
			var req auth_models.RegisterRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
				return
			}

			result, err := authAdapter.Register(c.Request.Context(), &req)
			if err != nil {
				h.Error(err)
				return
			}

			h.Created(result, utils.MsgRegistered)
		})

		authRoutes.POST(constants.LoginPath, func(c *gin.Context) {
			h := utils.NewHandler(c)
			var req auth_models.LoginRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
				return
			}

			if ip, exists := c.Get("ip_address"); exists {
				req.IPAddress = ip.(string)
			}
			if ua, exists := c.Get("user_agent"); exists {
				req.UserAgent = ua.(string)
			}

			result, err := authAdapter.Login(c.Request.Context(), &req)
			if err != nil {
				h.Error(err)
				return
			}

			h.OK(result, utils.MsgLoggedIn)
		})

		authRoutes.POST(constants.RefreshPath, func(c *gin.Context) {
			h := utils.NewHandler(c)

			refreshToken := c.GetHeader("X-Refresh-Token")
			if refreshToken == "" {
				h.Error(errors.NewUnauthorizedError("Refresh token no proporcionado", nil))
				return
			}

			tokenResponse, err := authAdapter.RefreshToken(c.Request.Context(), refreshToken)
			if err != nil {
				h.Error(err)
				return
			}

			h.OK(tokenResponse, "Token refrescado exitosamente")
		})
	}

	publicUserRoutes := apiV1.Group(constants.UsersPath)
	{
		publicUserRoutes.GET(constants.IDParam, func(c *gin.Context) {
			h := utils.NewHandler(c)

			idStr := c.Param("id")
			id, convErr := strconv.Atoi(idStr)
			if convErr != nil {
				h.Error(errors.NewValidationError(utils.MsgInvalidID, convErr))
				return
			}

			user, err := userAdapter.GetUserById(id)
			if err != nil {
				h.Error(err)
				return
			}

			h.OK(user, utils.MsgRetrieved)
		})
	}

	secureRoutes := apiV1.Group(constants.SecurePath)
	secureRoutes.Use(authMiddleware.RequireAuth())
	secureAuthRoutes := secureRoutes.Group(constants.AuthPath)
	{

		secureAuthRoutes.POST(constants.LogoutPath, func(c *gin.Context) {
			h := utils.NewHandler(c)

			userID, exists := c.Get("userID")
			if !exists {
				h.Error(errors.NewUnauthorizedError("User not authenticated", nil))
				return
			}

			err := authAdapter.Logout(c.Request.Context(), userID.(string))
			if err != nil {
				h.Error(err)
				return
			}

			h.OK(nil, utils.MsgLoggedOut)
		})

	}

	secureUserRoutes := secureRoutes.Group(constants.UsersPath)
	{
		secureUserRoutes.GET("", func(c *gin.Context) {
			h := utils.NewHandler(c)

			usersData, err := userAdapter.GetUsers(c.Request.Context())
			if err != nil {
				h.Error(err)

				return
			}

			h.OK(usersData, utils.MsgRetrieved)
		})

		secureUserRoutes.POST("", func(c *gin.Context) {
			h := utils.NewHandler(c)
			var req models.CreateUserRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				validationErr := errors.NewValidationError(utils.MsgInvalidInput, err)
				h.Error(validationErr)

				return
			}

			createdUser, err := userAdapter.CreateUser(c.Request.Context(), &req)
			if err != nil {
				h.Error(err)

				return
			}

			h.Created(createdUser, utils.MsgCreated)
		})

		secureUserRoutes.PUT(constants.IDParam, func(c *gin.Context) {
			h := utils.NewHandler(c)
			var req models.UpdateUserRequest

			if err := c.ShouldBindJSON(&req); err != nil {
				validationErr := errors.NewValidationError(utils.MsgInvalidInput, err)
				h.Error(validationErr)

				return
			}

			idStr := c.Param("id")
			id, convErr := strconv.Atoi(idStr)
			if convErr != nil {
				h.Error(errors.NewValidationError(utils.MsgInvalidID, convErr))

				return
			}

			updatedUser, err := userAdapter.UpdateUser(c.Request.Context(), id, &req)
			if err != nil {
				h.Error(err)

				return
			}

			h.OK(updatedUser, utils.MsgUpdated)
		})

		secureUserRoutes.DELETE(constants.IDParam, func(c *gin.Context) {
			h := utils.NewHandler(c)

			idStr := c.Param("id")
			id, convErr := strconv.Atoi(idStr)
			if convErr != nil {
				h.Error(errors.NewValidationError(utils.MsgInvalidID, convErr))

				return
			}

			err := userAdapter.DeleteUser(c.Request.Context(), id)
			if err != nil {
				h.Error(err)

				return
			}

			h.NoContent()
		})

		secureUserRoutes.GET(constants.MePath, func(c *gin.Context) {
			h := utils.NewHandler(c)

			userID, exists := c.Get("userID")
			if !exists {
				h.Error(errors.NewUnauthorizedError("No estás autenticado", nil))
				return
			}

			user, err := authAdapter.GetUserById(c.Request.Context(), userID.(string))
			if err != nil {
				h.Error(err)
				return
			}

			userResponse := models.UserResponse{
				ID:          strconv.Itoa(user.ID),
				UserID:      user.UserID,
				Username:    user.Username,
				DisplayName: user.DisplayName,
				Email:       user.Email,
				Role:        user.Role,
			}

			h.OK(userResponse, "Perfil de usuario recuperado exitosamente")
		})
	}

	swaggerRoutes := swagger.NewRoutes()
	swaggerRoutes.Register(router)

	return router
}
