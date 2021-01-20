package main

import (
	"backend-app/config"
	"backend-app/handler"
	"backend-app/repository"
	"backend-app/server"
	"backend-app/service"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if os.Getenv("APP_ENV") != "prod" {
		_ = godotenv.Load("../../.env")
	}

	Repository := repository.Repository{}.ConnectToDatabase()
	TodoService := service.TodoService{TodoRepository: Repository}
	Todo := handler.TodoHandler{TodoService: TodoService}

	conf, _ := config.New(".config", os.Getenv("APP_ENV"))
	conf.Print()
	servers, _ := server.New(&conf)

	servers.E.POST("/addTodoListItem/:title", Todo.AddTodoItem)
	servers.E.POST("/updateItemCompleted/:id/:completed", Todo.UpdateTodoItemCompleted)
	servers.E.POST("/deleteTodoItem/:id", Todo.DeleteTodoItem)
	servers.E.GET("/getTodoList", Todo.GetTodoList)

	_ = servers.Start()
}
