package main

import (
	"log"

	"github.com/domitory66/todo-app"
	"github.com/domitory66/todo-app/pkg/handler"
	"github.com/domitory66/todo-app/pkg/repository"
	"github.com/domitory66/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
