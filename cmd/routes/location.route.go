package routes

import (
	"recorderis/cmd/middleware"
	"recorderis/cmd/services/location/models"
	location_ports "recorderis/cmd/services/location/ports/drivers"
	"recorderis/internals/constants"
	"recorderis/internals/errors"
	"recorderis/internals/utils"

	"github.com/gin-gonic/gin"
)

func SetupLocationRoutes(router *gin.Engine, locationAdapter location_ports.ForLocation, authMiddleware *middleware.AuthMiddleware) {
	locationRoutes := router.Group(constants.APIPathV1 + constants.SecurePath + constants.LocationsPath)
	locationRoutes.Use(authMiddleware.RequireAuth())

	// GET /api/v1/secure/locations
	locationRoutes.GET("", func(c *gin.Context) {
		h := utils.NewHandler(c)

		locations, err := locationAdapter.GetLocations(c.Request.Context())
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(locations, utils.MsgRetrieved)
	})

	// POST /api/v1/secure/locations
	locationRoutes.POST("", func(c *gin.Context) {
		h := utils.NewHandler(c)
		var req models.CreateLocationRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		location, err := locationAdapter.CreateLocation(c.Request.Context(), &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.Created(location, utils.MsgCreated)
	})

	// GET /api/v1/secure/locations/:id
	locationRoutes.GET(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		locationID := c.Param("id")

		location, err := locationAdapter.GetLocationByID(c.Request.Context(), locationID)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(location, utils.MsgRetrieved)
	})

	// PUT /api/v1/secure/locations/:id
	locationRoutes.PUT(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		locationID := c.Param("id")
		var req models.UpdateLocationRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			h.Error(errors.NewValidationError(utils.MsgInvalidInput, err))
			return
		}

		location, err := locationAdapter.UpdateLocation(c.Request.Context(), locationID, &req)
		if err != nil {
			h.Error(err)
			return
		}

		h.OK(location, utils.MsgUpdated)
	})

	// DELETE /api/v1/secure/locations/:id
	locationRoutes.DELETE(constants.IDParam, func(c *gin.Context) {
		h := utils.NewHandler(c)
		locationID := c.Param("id")

		err := locationAdapter.DeleteLocation(c.Request.Context(), locationID)
		if err != nil {
			h.Error(err)
			return
		}

		h.NoContent()
	})
}
