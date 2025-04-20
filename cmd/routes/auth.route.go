package routes

import (
	"recorderis/cmd/middleware"
	auth_models "recorderis/cmd/services/auth/models"
	auth_ports "recorderis/cmd/services/auth/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/internals/errors"
	"recorderis/internals/utils"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, authAdapter auth_ports.ForAuth, authMiddleware *middleware.AuthMiddleware) {

	authRoutes := router.Group(constants.APIPathV1 + constants.AuthPath)
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
				h.Error(errors.NewUnauthorizedError("Missing refresh token", nil))
				return
			}

			tokenResponse, err := authAdapter.RefreshToken(c.Request.Context(), refreshToken)
			if err != nil {
				h.Error(err)
				return
			}

			h.OK(tokenResponse, "Token refreshed successfully")
		})
	}

	secureRoutes := router.Group(constants.APIPathV1 + constants.SecurePath)
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
}
