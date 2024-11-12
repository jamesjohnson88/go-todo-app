package data

import "go-todo-app/models"

func CreateTodoItem(item *models.TodoItem) (dbItem models.TodoItem, err error) {
	dbItem = *item
	dbItem.Id = "1"
	return dbItem, nil
}
