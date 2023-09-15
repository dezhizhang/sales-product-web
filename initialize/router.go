package initialize

import (
	"github.com/gin-gonic/gin"
	"sales-product-web/middleware"
	"sales-product-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group("/api/v1").Use(middleware.Cors())
	router.ProductRouter(ApiGroup)
	// router.CommonRouter(ApiGroup)
	return r
}
