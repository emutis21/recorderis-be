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
	descriptionRoutes := memoryRoutes.Group(constants.IDParam + constants.DescriptionsPath)

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

	descriptionRoutes.GET("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")

		descriptions, err := memoryAdapter.GetDescriptions(c.Request.Context(), memoryID)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(descriptions, utils.MsgRetrieved)
	})

	descriptionRoutes.POST("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")
		var req models.CreateDescriptionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		description, err := memoryAdapter.CreateDescription(c.Request.Context(), memoryID, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.Created(description, utils.MsgCreated)
	})

	descriptionRoutes.GET(constants.DescriptionIDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")
		descriptionID := c.Param("description_id")

		description, err := memoryAdapter.GetDescriptionByID(c.Request.Context(), memoryID, descriptionID)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(description, utils.MsgRetrieved)
	})

descriptionRoutes.PUT(constants.DescriptionIDParam, func(c *gin.Context) {
    h := utils.NewHandler(c)
    memoryID := c.Param("id")
    descriptionID := c.Param("description_id")
		var req models.UpdateDescriptionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		description, err := memoryAdapter.UpdateDescription(c.Request.Context(), memoryID, descriptionID, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(description, utils.MsgUpdated)
	})

	descriptionRoutes.DELETE(constants.DescriptionIDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		memoryID := c.Param("id")
		descriptionID := c.Param("description_id")

		err := memoryAdapter.DeleteDescription(c.Request.Context(), memoryID, descriptionID)
		if err != nil {
			h.Error(err)
			return
		}

		h.NoContent()
	})
}
