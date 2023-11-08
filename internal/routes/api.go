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

	admin := router.Group("").Use(middleware.IsValidJWT(), middleware.IsAdmin())
	{
		admin.POST("categories", handler.CategoryHandler.CreateCategory)
		admin.PATCH("categories/:id", handler.CategoryHandler.UpdateCategory)
		admin.DELETE("categories/:id", handler.CategoryHandler.DeleteCategory)
	}

	users := router.Group("").Use(middleware.IsValidJWT())
	{
		users.GET("categories", handler.CategoryHandler.GetCategories)
		users.PUT("update-account", middleware.SetUserID(), handler.AuthHandler.UpdateUser)
		users.DELETE("delete-account", middleware.SetUserID(), handler.AuthHandler.DeleteUser)
	}

	// categories := router.Group("categories").Use(middleware.IsValidJWT())
	{
		// categories.POST("/", middleware.IsAdmin(), handler.CategoryHandler.CreateCategory)
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
