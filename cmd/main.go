package main

import (
	"log"

	"github.com/iki-rumondor/project3-grup9/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/project3-grup9/internal/adapter/http"
	"github.com/iki-rumondor/project3-grup9/internal/application"
	"github.com/iki-rumondor/project3-grup9/internal/domain"
	"github.com/iki-rumondor/project3-grup9/internal/repository"
	"github.com/iki-rumondor/project3-grup9/internal/routes"
	"gorm.io/gorm"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// go migration(gormDB)

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

	var PORT = ":8080"
	routes.StartServer(&handlers).Run(PORT)
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
