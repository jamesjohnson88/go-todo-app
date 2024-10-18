package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo-app/handlers"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TodoItem Routes
	tiBase := "/api/todo-items"
	e.GET(tiBase, handlers.GetAllTodoItems)
	e.GET(tiBase+"/:id", handlers.GetTodoItem)
	e.POST(tiBase, handlers.CreateTodoItem)
	e.PUT(tiBase+"/:id", handlers.UpdateTodoItem)
	e.PATCH(tiBase+"/resolve/:id", handlers.ResolveTodoItem)
	e.DELETE(tiBase+"/:id", handlers.DeleteTodoItem)

	e.Logger.Fatal(e.Start(":1323"))
}
