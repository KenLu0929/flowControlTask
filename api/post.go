package api

import (
	"github.com/gin-gonic/gin"

	"github.com/KenLu0929/flowControlTask/pkg/components/post"
	"github.com/KenLu0929/flowControlTask/pkg/middware"
)

func PostHandler(api *gin.RouterGroup) {
	postApi := api.Group("/post")
	{
		authorized := postApi.Group("/")
		authorized.Use(middware.VerifyToken())
		{
			authorized.POST("/", post.CreatePost)
			authorized.GET("/", post.GetUserPost)
		}
	}
}
