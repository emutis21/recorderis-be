package routes

import (
	"recorderis/cmd/middleware"
	"recorderis/cmd/services/api/models"
	api_ports "recorderis/cmd/services/api/ports/drivers"
	auth_ports "recorderis/cmd/services/auth/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/internals/errors"
	"recorderis/internals/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userAdapter api_ports.ForUser, authAdapter auth_ports.ForAuth, authMiddleware *middleware.AuthMiddleware) {

	publicUserRoutes := router.Group(constants.APIPathV1 + constants.UsersPath)
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

	secureUserRoutes := router.Group(constants.APIPathV1 + constants.SecurePath + constants.UsersPath)
	secureUserRoutes.Use(authMiddleware.RequireAuth())
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
				h.Error(errors.NewUnauthorizedError("User not authenticated", nil))
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

			h.OK(userResponse, "User profile retrieved successfully")
		})
	}
}
