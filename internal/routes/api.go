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
		users.GET("tasks", handler.TaskHandler.GetTasks)

		users.POST("tasks", middleware.SetUserID(), handler.TaskHandler.CreateTask)

		users.PUT("users/update-account", middleware.SetUserID(), handler.AuthHandler.UpdateUser)
		users.PUT("tasks/:id", middleware.SetUserID(), handler.TaskHandler.UpdateTask)

		users.PATCH("tasks/update-status/:id", middleware.SetUserID(), handler.TaskHandler.UpdateTaskStatus)
		users.PATCH("tasks/update-category/:id", middleware.SetUserID(), handler.TaskHandler.UpdateTaskCategory)

		users.DELETE("users/delete-account", middleware.SetUserID(), handler.AuthHandler.DeleteUser)
		users.DELETE("tasks/:id", middleware.SetUserID(), handler.TaskHandler.DeleteTask)
	}

	return router
}
