package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo-app/data"
	"go-todo-app/handlers"
	"log"
	"os"
)

func main() {
	// Init Env Vars
	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	// Create PGX Connection Pool
	pool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	data.SetDbPool(pool)
	defer pool.Close()

	// Init Echo & Middleware
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
