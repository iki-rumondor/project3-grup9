package routes

import (
	"github.com/gin-gonic/gin"
	customHTTP "github.com/iki-rumondor/project3-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project3-grup9/internal/middleware"
)

func StartServer(handler *customHTTP.Handlers) *gin.Engine {
	router := gin.Default()

	public := router.Group("")
	{
		public.POST("users/register", handler.AuthHandler.Register)
		public.POST("users/login", handler.AuthHandler.Login)
	}

	users := router.Group("users").Use(middleware.IsValidJWT())
	{
		users.PUT("update-account", middleware.SetUserID(), handler.AuthHandler.UpdateUser)
		users.PUT("delete-account", middleware.SetUserID(), handler.AuthHandler.DeleteUser)
	}

	categories := router.Group("categories")
	{
		categories.POST("/", handler.CategoryHandler.CreateCategory)
		categories.GET("/", handler.CategoryHandler.GetCategories)
		// categories.PATCH("/:id", handler.CategoryHandler.UpdateCategory)
		// categories.DELETE("/:id", handler.CategoryHandler.DeleteCategory)
	}

	// tasks := router.Group("tasks")
	// {
	// 	tasks.POST("/", handler.TaskHandler.CreateTask)
	// 	tasks.GET("/", handler.TaskHandler.GetTask)
	// 	tasks.PATCH("/:id", handler.TaskHandler.UpdateTask)
	// 	tasks.DELETE("/:id", handler.TaskHandler.DeleteTask)
	// }

	return router
}
