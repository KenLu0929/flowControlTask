package api

import (
	"github.com/KenLu0929/flowControlTask/pkg/components/user"
	"github.com/KenLu0929/flowControlTask/pkg/middware"

	"github.com/gin-gonic/gin"
)

func UserHandler(api *gin.RouterGroup) {
	userApi := api.Group("/user")
	{
		userApi.POST("/", user.CreateUser)
		authorized := userApi.Group("/")
		authorized.Use(middware.VerifyToken())
		{
			authorized.GET("/", user.GetUser)
		}
	}
}
