package handlers

import (
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
