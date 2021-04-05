package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/KenLu0929/flowControlTask/api"
	_ "github.com/KenLu0929/flowControlTask/config/db"
	_ "github.com/KenLu0929/flowControlTask/config/rdb"
	"github.com/KenLu0929/flowControlTask/docs"
	_ "github.com/KenLu0929/flowControlTask/init/migrations"
	"github.com/KenLu0929/flowControlTask/pkg/middware"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api")
	apiGroup.Use(middware.FlowControlByIP())
	{
		api.UserHandler(apiGroup)
		api.LoginHandler(apiGroup)
		api.PostHandler(apiGroup)
	}

	return router
}

// @contact.name KenLu
// @contact.url http://github.com/KenLu0929
// @contact.email aaa710140505@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	docs.SwaggerInfo.Title = "flow control Task Swagger"
	docs.SwaggerInfo.Description = "This is a flow contro Task"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router := SetRouter()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
