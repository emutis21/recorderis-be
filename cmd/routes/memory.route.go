package routes

import (
	"context"
	"recorderis/cmd/middleware"
	"recorderis/cmd/services/memory/models"
	memory_ports "recorderis/cmd/services/memory/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/internals/errors"
	"recorderis/internals/utils"

	"github.com/gin-gonic/gin"
)

func SetupMemoryRoutes(router *gin.Engine, memoryAdapter memory_ports.ForMemory, authMiddleware *middleware.AuthMiddleware) {
	memoryRoutes := router.Group(constants.APIPathV1 + constants.SecurePath + constants.MemoriesPath)
	memoryRoutes.Use(authMiddleware.RequireAuth())

	memoryRoutes.GET("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		userID, exists := c.Get("userID")
		if !exists {
			h.Error(errors.NewUnauthorizedError("User not authenticated", nil))
			return
		}

		memories, err := memoryAdapter.GetMemories(c.Request.Context(), userID.(string))
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(memories, utils.MsgRetrieved)
	})

	memoryRoutes.POST("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		var req models.CreateMemoryRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			h.Error(errors.NewUnauthorizedError("User not authenticated", nil))
			return
		}

		ctx := context.WithValue(c.Request.Context(), constants.UserIDKey, userID)

		memory, err := memoryAdapter.CreateMemory(ctx, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.Created(memory, utils.MsgCreated)
	})

	memoryRoutes.GET(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")

		memory, err := memoryAdapter.GetMemoryByID(c.Request.Context(), memoryID)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(memory, utils.MsgRetrieved)
	})

	memoryRoutes.PUT(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")
		var req models.UpdateMemoryRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		memory, err := memoryAdapter.UpdateMemory(c.Request.Context(), memoryID, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(memory, utils.MsgUpdated)
	})

	memoryRoutes.DELETE(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")

		err := memoryAdapter.DeleteMemory(c.Request.Context(), memoryID)
		if err != nil {
			h.Error(err)
			return
		}

		h.NoContent()
	})
}
