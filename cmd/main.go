package main

import (
	"log"

	"github.com/iki-rumondor/init-golang-service/internal/adapter/database"
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
	"github.com/iki-rumondor/init-golang-service/internal/routes"
)

func main() {
	gormDB, err := database.NewMysqlDB()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	categoryRepo := repository.NewRepository(gormDB)
	categoryService := application.NewService(categoryRepo)
	categoryHandler := customHTTP.NewHandler(categoryService)

	handlers := customHTTP{
		categoryHandler: categoryHandler,
	}
	var PORT = ":8080"
	routes.StartServer(handlers).Run(PORT)
}
