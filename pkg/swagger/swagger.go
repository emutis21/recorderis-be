// filepath: pkg/swagger/swagger.go
package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "recorderis/docs"
)

type Routes struct{}

func NewRoutes() *Routes {
	return &Routes{}
}

func (r *Routes) Register(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.InstanceName("swagger"),
	))
}
