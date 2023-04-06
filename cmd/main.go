package main

import (
	"log"

	"github.com/domitory66/todo-app"
	"github.com/domitory66/todo-app/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)
	err := srv.Run("8000", handlers.InitRoutes())
	if err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
