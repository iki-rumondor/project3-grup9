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

	taskRepo := repository.NewTaskRepository(gormDB)
	taskService := application.NewTaskService(taskRepo)
	taskHandler := customHTTP.NewTaskHandler(taskService)

	Handlers := customHTTP.Handlers{
		TaskHandlers: taskHandler,
	}

	var PORT = ":8080"
	routes.StartServer(&Handlers).Run(PORT)
}
