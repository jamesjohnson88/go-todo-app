package data

import "go-todo-app/models"

func CreateTodoItem(item *models.TodoItem) (id string) {
	id = "id1" + item.Title
	return // we probably want to send back the full created item as we'll use it immediately
}
