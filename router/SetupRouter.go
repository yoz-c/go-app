package router

import (
	"github.com/gin-gonic/gin"
	"go-app/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../view/*")
	} else {
		router.LoadHTMLGlob("view/*")
	}

	router.Static("/assetPath", "./asset")

	indexRouting := router.Group("/")
	{
    	indexRouting.GET("", handler.GetIndex)
	}

	helloRouting := router.Group("/hello")
	{
		helloRouting.GET("", handler.GetHello)

		helloRouting.DELETE("/:id", handler.DeleteHello)
	}

	userRouting := router.Group("/user")
	{
		userRouting.GET("/:uid", handler.GetUser)

	}

	fileRouting := router.Group("/file")
	{
		fileRouting.POST("", handler.UploadSingleIndex)

	}

	return router
}