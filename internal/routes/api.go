package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
)

func StartServer(handler *customHTTP.Handlers) *gin.Engine{
	router := gin.Default()

	router.GET("/")

	api := router.Group("api/v1")
	{
		api.GET("/endpoints")
		api.POST("/endpoints")
		api.GET("/endpoints:id")
		api.PUT("/endpoints/:id")
		api.DELETE("/endpoints/:id")
	}

	return router
}