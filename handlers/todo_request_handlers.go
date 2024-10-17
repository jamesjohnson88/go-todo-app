package handlers

import (
	"github.com/labstack/echo/v4"
	"go-todo-app/models"
	"net/http"
)

func GetAllTodoItems(c echo.Context) error {
	return c.String(http.StatusOK, "Get All Todo Items")
}

func GetTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	return c.String(http.StatusOK, "Get Todo Item By Id: "+itemId)
}

func CreateTodoItem(c echo.Context) error {
	ti := new(models.TodoItem)
	if err := c.Bind(ti); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "Created: "+ti.Title)
}

func UpdateTodoItem(c echo.Context) error {
	ti := new(models.TodoItem)
	if err := c.Bind(ti); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "Updated: "+c.Param("id"))
}

func ResolveTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	return c.String(http.StatusOK, "Resolved Todo Item: "+itemId)
}

func DeleteTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	return c.String(http.StatusOK, "Deleted Todo Item: "+itemId)
}
