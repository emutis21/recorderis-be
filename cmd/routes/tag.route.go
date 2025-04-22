package routes

import (
	"recorderis/cmd/middleware"
	"recorderis/cmd/services/tags/models"
	tag_ports "recorderis/cmd/services/tags/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/internals/errors"
	"recorderis/internals/utils"

	"github.com/gin-gonic/gin"
)

func SetupTagRoutes(router *gin.Engine, tagAdapter tag_ports.ForTag, authMiddleware *middleware.AuthMiddleware) {
	tagRoutes := router.Group(constants.APIPathV1 + constants.SecurePath + constants.TagsPath)
	tagRoutes.Use(authMiddleware.RequireAuth())

	// GET /api/v1/secure/tags
	tagRoutes.GET("", func(c *gin.Context) {
		h := utils.NewHandler(c)

		tags, err := tagAdapter.GetTags(c.Request.Context())
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(tags, utils.MsgRetrieved)
	})

	// POST /api/v1/secure/tags
	tagRoutes.POST("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		var req models.CreateTagRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		tag, err := tagAdapter.CreateTag(c.Request.Context(), &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.Created(tag, utils.MsgCreated)
	})

	// GET /api/v1/secure/tags/:id
	tagRoutes.GET(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		tagID := c.Param("id")

		tag, err := tagAdapter.GetTagByID(c.Request.Context(), tagID)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(tag, utils.MsgRetrieved)
	})

	// PUT /api/v1/secure/tags/:id
	tagRoutes.PUT(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		tagID := c.Param("id")
		var req models.UpdateTagRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		tag, err := tagAdapter.UpdateTag(c.Request.Context(), tagID, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(tag, utils.MsgUpdated)
	})

	// DELETE /api/v1/secure/tags/:id
	tagRoutes.DELETE(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		tagID := c.Param("id")

		err := tagAdapter.DeleteTag(c.Request.Context(), tagID)
		if err != nil {
			h.Error(err)
			return
		}

		h.NoContent()
	})

}
