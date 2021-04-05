package api

import (
	"github.com/KenLu0929/flowControlTask/pkg/components/auth"

	"github.com/gin-gonic/gin"
)

func LoginHandler(api *gin.RouterGroup) {
	loginApi := api.Group("/login")
	{
		loginApi.POST("/", auth.Login)
	}
}
