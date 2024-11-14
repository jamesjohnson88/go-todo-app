package handlers

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"go-todo-app/data"
	"go-todo-app/models"
	"net/http"
)

func GetAllTodoItems(c echo.Context) error {
	items, err := data.GetTodoItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, items)
}

func GetTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	item, err := data.GetTodoItemById(itemId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if item == nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, item)
}

func CreateTodoItem(c echo.Context) error {
	ti := new(models.TodoItem)
	if err := c.Bind(ti); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	item, err := data.CreateTodoItem(ti)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	location := c.Request().Host + c.Echo().URI(GetTodoItem, item.Id)
	c.Response().Header().Set("Location", location)

	return c.JSON(http.StatusCreated, item)
}

func UpdateTodoItem(c echo.Context) error {
	id := c.Param("id")
	ti := new(models.TodoItem)
	if err := c.Bind(ti); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	updatedItem, err := data.UpdateTodoItem(id, ti)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, updatedItem)
}

func CompleteTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	itemCompleted, err := data.CompleteTodoItem(itemId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if itemCompleted {
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusNotFound)
}

// todo
func DeleteTodoItem(c echo.Context) error {
	itemId := c.Param("id")
	return c.String(http.StatusOK, "Deleted Todo Item: "+itemId)
}
