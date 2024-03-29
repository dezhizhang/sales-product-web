package router

import (
	"github.com/gin-gonic/gin"
	"sales-product-web/api"
)

func ProductRouter(Router gin.IRoutes) {
	{
		Router.GET("/banner/list", api.GetBannerList)
		Router.POST("/banner/add", api.CreateBanner)
		Router.PUT("/banner/update", api.UpdateBanner)
		Router.DELETE("/banner/delete", api.DeleteBanner)
		//Router.POST("/manager/login", api.UserLoginIn)
		//Router.DELETE("/manager/:id", api.DeleteUser)
	}
	//userRouter := Router.Group("/user").Use(middleware.Cors())
	//{
	//	userRouter.POST("/add", api.CreateUser)
	//	userRouter.GET("/list", api.GetUserList)
	//	userRouter.POST("/login", api.UserLoginIn)
	//	userRouter.DELETE("/delete/:id", api.DeleteUser)
	//}

}
