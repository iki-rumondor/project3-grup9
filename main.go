package main

import (
	"log"
	"os"

	"github.com/iki-rumondor/project3-grup9/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/project3-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project3-grup9/internal/application"
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"github.com/iki-rumondor/project3-grup9/internal/repository"
	"github.com/iki-rumondor/project3-grup9/internal/routes"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewPostgresDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	migration(gormDB)

	authRepo := repository.NewAuthRepository(gormDB)
	authService := application.NewAuthService(authRepo)
	authHandler := customHTTP.NewAuthHandler(authService)

	categoryRepo := repository.NewCategoryRepository(gormDB)
	categoryService := application.NewCategoryService(categoryRepo)
	categoryHandler := customHTTP.NewCategoryHandler(categoryService)

	taskRepo := repository.NewTaskRepository(gormDB)
	taskService := application.NewTaskService(taskRepo)
	taskHandler := customHTTP.NewTaskHandler(taskService)

	handlers := customHTTP.Handlers{
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
		TaskHandler:     taskHandler,
	}

	var PORT = envPortOr("3000")
	routes.StartServer(&handlers).Run(PORT)
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func migration(db *gorm.DB) {
	db.Migrator().DropTable(&domain.User{})
	db.Migrator().CreateTable(&domain.User{})
	db.Migrator().DropTable(&domain.Category{})
	db.Migrator().CreateTable(&domain.Category{})
	db.Migrator().DropTable(&domain.Task{})
	db.Migrator().CreateTable(&domain.Task{})
	db.Create(&domain.User{
		FullName: "Administrator",
		Email: "admin@admin.com",
		Password: "123456",
		Role: "admin",
	})
}
