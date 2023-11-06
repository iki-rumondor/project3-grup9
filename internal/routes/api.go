package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
)

func StartServer(handler *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("/")

	category := router.Group("api/v1")
	{
		category.POST("/categories", handler.CategoryHandlers.CreateCategory)
		category.GET("/categories", handler.CategoryHandlers.GetCategorys)
		category.PATCH("/categories/:categoryid", handler.CategoryHandlers.UpdateCategory)
		category.DELETE("/categories/:categoryid", handler.CategoryHandlers.DeleteCategory)
	}

	return router
}
